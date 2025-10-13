package memory

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	d2gsSendPacketPattern = "\xE8\x00\x00\x00\x00\x0F\xB6\x85\x00\x00\x00\x00\x48\x03\xF0"
	d2gsSendPacketMask    = "x????xxx????xxx"
	maxPacketSize         = 65536
)

var (
	sendPacketStubBase = []byte{
		0xF3, 0x0F, 0x1E, 0xFA,
		0x53,
		0x48, 0x89, 0xCB,
		0x48, 0x83, 0xEC, 0x20,
		0x48, 0x8B, 0x03,
		0x48, 0x8B, 0x4B, 0x08,
		0x48, 0x8B, 0x53, 0x10,
		0x45, 0x33, 0xC0,
		0xFF, 0xD0,
		0xC7, 0x43, 0x18, 0x01, 0x00, 0x00, 0x00,
		0xB8, 0x01, 0x00, 0x00, 0x00,
		0x48, 0x83, 0xC4, 0x20,
		0x5B,
		0xC3,
	}

	kernel32                       = windows.NewLazySystemDLL("kernel32.dll")
	procVirtualAllocEx             = kernel32.NewProc("VirtualAllocEx")
	procVirtualFreeEx              = kernel32.NewProc("VirtualFreeEx")
	procQueueUserAPC               = kernel32.NewProc("QueueUserAPC")
	procSuspendThread              = kernel32.NewProc("SuspendThread")
	procResumeThread               = kernel32.NewProc("ResumeThread")
	procGetExitCodeThread          = kernel32.NewProc("GetExitCodeThread")
	procGetThreadTimes             = kernel32.NewProc("GetThreadTimes")
	procVirtualProtectEx           = kernel32.NewProc("VirtualProtectEx")
	procSetProcessValidCallTargets = kernel32.NewProc("SetProcessValidCallTargets")

	d2gsCachedFn uintptr
	d2gsCacheMu  sync.RWMutex
	d2gsCachePID uint32

	metaBufPool = sync.Pool{
		New: func() interface{} {
			return new([sendPacketMetaSize]byte)
		},
	}
)

const (
	sendPacketProcessAccess = windows.PROCESS_VM_OPERATION | windows.PROCESS_VM_READ | windows.PROCESS_VM_WRITE | windows.PROCESS_QUERY_INFORMATION

	THREAD_SUSPEND_RESUME    = 0x0002
	THREAD_SET_CONTEXT       = 0x0010
	THREAD_QUERY_INFORMATION = 0x0040
	sendPacketThreadAccess   = THREAD_SET_CONTEXT | THREAD_SUSPEND_RESUME | THREAD_QUERY_INFORMATION

	sendPacketStatusOffset = 24
	sendPacketMetaSize     = 32

	threadStillActive = 259

	cfgCallTargetValid = 0x1
)

type sendPacketState struct {
	mu                  sync.Mutex
	handle              windows.Handle
	stub                uintptr
	meta                uintptr
	packet              uintptr
	packetCap           uintptr
	fn                  uintptr
	thread              windows.Handle
	threadID            uint32
	processPID          uint32
	threadLastValidated time.Time
	leakedBuffers       []uintptr
}

type cfgCallTargetInfo struct {
	Offset uintptr
	Flags  uintptr
}

func findPatternOffset(memory []byte, pattern, mask string) int {
	if len(pattern) != len(mask) {
		return -1
	}
	patternLength := len(pattern)
	limit := len(memory) - patternLength
	for i := 0; i <= limit; i++ {
		match := true
		for j := 0; j < patternLength; j++ {
			if mask[j] == 'x' && memory[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func (s *sendPacketState) ensureHandle(pid uint32) (windows.Handle, error) {
	if s.handle != 0 && s.processPID == pid {
		return s.handle, nil
	}

	if s.handle != 0 && s.processPID != pid {
		if s.packet != 0 {
			virtualFreeEx(s.handle, s.packet)
		}
		if s.meta != 0 {
			virtualFreeEx(s.handle, s.meta)
		}
		if s.stub != 0 {
			virtualFreeEx(s.handle, s.stub)
		}

		windows.CloseHandle(s.handle)
		s.handle = 0
		s.processPID = 0
		s.stub = 0
		s.meta = 0
		s.packet = 0
		s.packetCap = 0
		s.fn = 0
	}

	h, err := windows.OpenProcess(sendPacketProcessAccess, false, pid)
	if err != nil {
		return 0, fmt.Errorf("open process %d: %w", pid, err)
	}
	s.handle = h
	s.processPID = pid
	return h, nil
}

func (s *sendPacketState) ensureStub(handle windows.Handle) error {
	if s.stub != 0 {
		return nil
	}

	stubSize := uintptr(len(sendPacketStubBase))

	addr, err := virtualAllocEx(handle, stubSize, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		return fmt.Errorf("allocate remote stub: %w", err)
	}

	if err := writeRemoteMemory(handle, addr, sendPacketStubBase); err != nil {
		virtualFreeEx(handle, addr)
		return fmt.Errorf("write remote stub: %w", err)
	}

	if err := markCallTargetValid(handle, addr, stubSize); err != nil {
		virtualFreeEx(handle, addr)
		return fmt.Errorf("register call target: %w", err)
	}

	if err := virtualProtectEx(handle, addr, stubSize, windows.PAGE_EXECUTE_READ); err != nil {
		virtualFreeEx(handle, addr)
		return fmt.Errorf("set stub protection: %w", err)
	}

	s.stub = addr
	return nil
}

func (s *sendPacketState) ensureMeta(handle windows.Handle) error {
	if s.meta != 0 {
		return nil
	}
	addr, err := virtualAllocEx(handle, sendPacketMetaSize, windows.PAGE_READWRITE)
	if err != nil {
		return fmt.Errorf("allocate remote metadata: %w", err)
	}
	s.meta = addr
	return nil
}

func (s *sendPacketState) ensurePacketBuffer(handle windows.Handle, size uintptr) error {
	if size == 0 {
		return errors.New("packet size must be greater than zero")
	}

	if size > maxPacketSize {
		return fmt.Errorf("packet too large: %d bytes (max %d)", size, maxPacketSize)
	}

	if size <= s.packetCap && s.packet != 0 {
		return nil
	}

	allocSize := size
	if size < 4096 {
		allocSize = 4096
	} else {
		allocSize = (size + 4095) &^ 4095
	}

	newPacket, err := virtualAllocEx(handle, allocSize, windows.PAGE_READWRITE)
	if err != nil {
		return fmt.Errorf("allocate packet buffer of %d bytes: %w", allocSize, err)
	}

	if s.packet != 0 {
		if err := virtualFreeEx(handle, s.packet); err != nil {
			log.Printf("Warning: failed to free old packet buffer at 0x%X: %v", s.packet, err)
			s.leakedBuffers = append(s.leakedBuffers, s.packet)
		}
	}

	s.packet = newPacket
	s.packetCap = allocSize
	return nil
}

func (s *sendPacketState) ensureFunction(p *Process) (uintptr, error) {
	if s.fn != 0 {
		return s.fn, nil
	}

	fn, err := p.GetD2GSSendPacketFn()
	if err != nil {
		return 0, err
	}

	if fn == 0 {
		return 0, errors.New("D2GS_SendPacket function pointer is null")
	}

	s.fn = fn
	return fn, nil
}

func (s *sendPacketState) ensureThreadHandle(p *Process) (windows.Handle, error) {
	if s.thread != 0 && s.processPID == p.pid {
		if err := validateThreadActive(s.thread); err == nil {
			s.threadLastValidated = time.Now()
			return s.thread, nil
		}
		windows.CloseHandle(s.thread)
		s.thread = 0
		s.threadID = 0
		s.processPID = 0
	} else if s.thread != 0 && s.processPID != p.pid {
		windows.CloseHandle(s.thread)
		s.thread = 0
		s.threadID = 0
		s.processPID = 0
	}

	threadID, err := findMainThreadID(p.pid)
	if err != nil {
		return 0, err
	}

	handle, err := windows.OpenThread(sendPacketThreadAccess, false, threadID)
	if err != nil {
		return 0, fmt.Errorf("open thread %d: %w", threadID, err)
	}

	s.thread = handle
	s.threadID = threadID
	s.processPID = p.pid
	s.threadLastValidated = time.Now()
	return handle, nil
}

func validateThreadActive(thread windows.Handle) error {
	var exitCode uint32
	ret, _, err := procGetExitCodeThread.Call(
		uintptr(thread),
		uintptr(unsafe.Pointer(&exitCode)),
	)
	if ret == 0 {
		if err != nil {
			return fmt.Errorf("GetExitCodeThread: %w", err)
		}
		return errors.New("GetExitCodeThread failed")
	}

	if exitCode != threadStillActive {
		return fmt.Errorf("thread is not active (exit code: %d)", exitCode)
	}

	return nil
}

func (s *sendPacketState) dispatchAPC(thread windows.Handle, start, parameter uintptr) error {
	if thread == 0 {
		return errors.New("thread handle is zero")
	}

	if err := suspendThread(thread); err != nil {
		return err
	}

	queued := false
	defer func() {
		if !queued {
			_ = resumeThread(thread)
		}
	}()

	if err := queueUserAPC(start, thread, parameter); err != nil {
		return err
	}
	queued = true

	if err := resumeThread(thread); err != nil {
		queued = false
		return err
	}

	return nil
}

func queueUserAPC(start uintptr, thread windows.Handle, parameter uintptr) error {
	ret, _, err := procQueueUserAPC.Call(start, uintptr(thread), parameter)
	if ret == 0 {
		if err != nil {
			return fmt.Errorf("QueueUserAPC: %w", err)
		}
		return errors.New("QueueUserAPC failed")
	}
	return nil
}

func suspendThread(thread windows.Handle) error {
	ret, _, err := procSuspendThread.Call(uintptr(thread))
	if ret == ^uintptr(0) {
		if err != nil {
			return fmt.Errorf("SuspendThread: %w", err)
		}
		return errors.New("SuspendThread failed")
	}
	return nil
}

func resumeThread(thread windows.Handle) error {
	ret, _, err := procResumeThread.Call(uintptr(thread))
	if ret == ^uintptr(0) {
		if err != nil {
			return fmt.Errorf("ResumeThread: %w", err)
		}
		return errors.New("ResumeThread failed")
	}
	return nil
}

// getThreadCreationTime gets the creation time of a thread
func getThreadCreationTime(thread windows.Handle) (int64, error) {
	var creationTime, exitTime, kernelTime, userTime windows.Filetime
	ret, _, err := procGetThreadTimes.Call(
		uintptr(thread),
		uintptr(unsafe.Pointer(&creationTime)),
		uintptr(unsafe.Pointer(&exitTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)),
	)
	if ret == 0 {
		if err != nil {
			return 0, fmt.Errorf("GetThreadTimes: %w", err)
		}
		return 0, errors.New("GetThreadTimes failed")
	}

	return int64(creationTime.HighDateTime)<<32 | int64(creationTime.LowDateTime), nil
}

func findMainThreadID(pid uint32) (uint32, error) {
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPTHREAD, 0)
	if err != nil {
		return 0, fmt.Errorf("CreateToolhelp32Snapshot: %w", err)
	}
	defer windows.CloseHandle(snapshot)

	var entry windows.ThreadEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	if err := windows.Thread32First(snapshot, &entry); err != nil {
		return 0, fmt.Errorf("Thread32First: %w", err)
	}

	var mainThreadID uint32
	var earliestCreationTime int64 = 0x7FFFFFFFFFFFFFFF
	var found bool

	for {
		if entry.OwnerProcessID == pid {
			thread, err := windows.OpenThread(sendPacketThreadAccess, false, entry.ThreadID)
			if err == nil {
				creationTime, err := getThreadCreationTime(thread)
				windows.CloseHandle(thread)

				if err == nil && creationTime < earliestCreationTime {
					earliestCreationTime = creationTime
					mainThreadID = entry.ThreadID
					found = true
				}
			}
		}

		if err := windows.Thread32Next(snapshot, &entry); err != nil {
			if errors.Is(err, windows.ERROR_NO_MORE_FILES) {
				break
			}
			return 0, fmt.Errorf("Thread32Next: %w", err)
		}
	}

	if !found {
		return 0, fmt.Errorf("no threads found for process %d", pid)
	}

	return mainThreadID, nil
}

func validatePatternMatch(memory []byte, offset int) bool {
	if offset >= len(memory) || memory[offset] != 0xE8 {
		return false
	}

	if offset+5 > len(memory) {
		return false
	}

	if offset+14 < len(memory) {
		if memory[offset+5] != 0x0F || memory[offset+6] != 0xB6 {
			return false
		}
	}

	return true
}

func (p *Process) GetD2GSSendPacketFn() (uintptr, error) {
	if p == nil {
		return 0, errors.New("process is nil")
	}

	if p.handler == 0 {
		return 0, errors.New("process handle is invalid")
	}

	d2gsCacheMu.RLock()
	if d2gsCachedFn != 0 && d2gsCachePID == p.pid {
		fn := d2gsCachedFn
		d2gsCacheMu.RUnlock()
		return fn, nil
	}
	d2gsCacheMu.RUnlock()

	modules, err := GetProcessModules(p.pid)
	if err != nil {
		return 0, fmt.Errorf("enumerate modules: %w", err)
	}

	for _, module := range modules {
		name := strings.ToLower(module.ModuleName)

		if strings.Contains(name, "windows") || strings.Contains(name, "system32") {
			continue
		}

		if module.ModuleBaseSize == 0 || module.ModuleBaseSize > 100*1024*1024 {
			continue
		}

		memory := make([]byte, int(module.ModuleBaseSize))
		if err := windows.ReadProcessMemory(p.handler, module.ModuleBaseAddress, &memory[0], uintptr(module.ModuleBaseSize), nil); err != nil {
			log.Printf("Warning: failed to read module %s: %v", module.ModuleName, err)
			continue
		}

		offset := findPatternOffset(memory, d2gsSendPacketPattern, d2gsSendPacketMask)
		if offset < 0 {
			continue
		}

		if !validatePatternMatch(memory, offset) {
			log.Printf("Warning: false positive pattern match at offset 0x%X in %s", offset, module.ModuleName)
			continue
		}

		patternAddr := module.ModuleBaseAddress + uintptr(offset)
		relOffset := int32(binary.LittleEndian.Uint32(memory[offset+1 : offset+5]))
		absolute := uintptr(int64(patternAddr+5) + int64(relOffset))

		if absolute == 0 || absolute < module.ModuleBaseAddress {
			log.Printf("Warning: invalid computed address 0x%X in %s", absolute, module.ModuleName)
			continue
		}

		log.Printf("D2GS_SendPacket resolved at 0x%X in module %s", absolute, module.ModuleName)

		d2gsCacheMu.Lock()
		d2gsCachedFn = absolute
		d2gsCachePID = p.pid
		d2gsCacheMu.Unlock()

		return absolute, nil
	}

	return 0, errors.New("D2GS_SendPacket pattern not found in any module")
}

func (p *Process) SendPacket(packet []byte) (err error) {
	if p == nil {
		return errors.New("process is nil")
	}

	defer func() {
		if err != nil {
			log.Printf("SendPacket(%d bytes) error: %v", len(packet), err)
		}
	}()

	if len(packet) == 0 {
		return errors.New("packet payload cannot be empty")
	}

	if len(packet) > maxPacketSize {
		return fmt.Errorf("packet too large: %d bytes (max %d)", len(packet), maxPacketSize)
	}

	p.sendPacketMu.Lock()
	if p.sendPacket == nil {
		p.sendPacket = &sendPacketState{}
	}
	state := p.sendPacket
	state.mu.Lock()
	p.sendPacketMu.Unlock()
	defer state.mu.Unlock()

	fnAddr, err := state.ensureFunction(p)
	if err != nil {
		return fmt.Errorf("resolve D2GS_SendPacket: %w", err)
	}

	handle, err := state.ensureHandle(p.pid)
	if err != nil {
		return fmt.Errorf("open process: %w", err)
	}

	if err := state.ensureStub(handle); err != nil {
		return err
	}

	if err := state.ensureMeta(handle); err != nil {
		return err
	}

	if err := state.ensurePacketBuffer(handle, uintptr(len(packet))); err != nil {
		return err
	}

	if err := writeRemoteMemory(handle, state.packet, packet); err != nil {
		return fmt.Errorf("write remote packet: %w", err)
	}

	metaBuf := metaBufPool.Get().(*[sendPacketMetaSize]byte)
	defer func() {
		*metaBuf = [sendPacketMetaSize]byte{}
		metaBufPool.Put(metaBuf)
	}()

	binary.LittleEndian.PutUint64(metaBuf[0:], uint64(fnAddr))
	binary.LittleEndian.PutUint64(metaBuf[8:], uint64(state.packet))
	binary.LittleEndian.PutUint64(metaBuf[16:], uint64(len(packet)))
	binary.LittleEndian.PutUint32(metaBuf[sendPacketStatusOffset:], 0)

	if err := writeRemoteMemory(handle, state.meta, metaBuf[:]); err != nil {
		return fmt.Errorf("write remote metadata: %w", err)
	}

	threadHandle, err := state.ensureThreadHandle(p)
	if err != nil {
		return fmt.Errorf("resolve main thread: %w", err)
	}

	if err := state.dispatchAPC(threadHandle, state.stub, state.meta); err != nil {
		return fmt.Errorf("dispatch APC: %w", err)
	}

	return nil
}

func writeRemoteMemory(handle windows.Handle, address uintptr, data []byte) error {
	if address == 0 {
		return errors.New("attempt to write to null address")
	}
	if len(data) == 0 {
		return nil
	}
	return windows.WriteProcessMemory(handle, address, &data[0], uintptr(len(data)), nil)
}

func virtualAllocEx(handle windows.Handle, size uintptr, protect uint32) (uintptr, error) {
	if size == 0 {
		return 0, errors.New("allocation size cannot be zero")
	}

	addr, _, err := procVirtualAllocEx.Call(
		uintptr(handle),
		0,
		size,
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		uintptr(protect),
	)
	if addr == 0 {
		if err != nil {
			return 0, err
		}
		return 0, errors.New("VirtualAllocEx failed")
	}
	return addr, nil
}

func virtualFreeEx(handle windows.Handle, address uintptr) error {
	if address == 0 {
		return nil
	}

	ret, _, err := procVirtualFreeEx.Call(
		uintptr(handle),
		address,
		0,
		windows.MEM_RELEASE,
	)
	if ret == 0 {
		if err != nil {
			return err
		}
		return errors.New("VirtualFreeEx failed")
	}
	return nil
}

func virtualProtectEx(handle windows.Handle, address uintptr, size uintptr, protect uint32) error {
	if address == 0 || size == 0 {
		return errors.New("attempt to protect invalid region")
	}

	if err := procVirtualProtectEx.Find(); err != nil {
		return nil
	}

	var oldProtect uint32
	ret, _, err := procVirtualProtectEx.Call(
		uintptr(handle),
		address,
		size,
		uintptr(protect),
		uintptr(unsafe.Pointer(&oldProtect)),
	)
	if ret == 0 {
		if err != nil {
			return fmt.Errorf("VirtualProtectEx: %w", err)
		}
		return errors.New("VirtualProtectEx failed")
	}
	return nil
}

func markCallTargetValid(handle windows.Handle, address uintptr, size uintptr) error {
	if err := procSetProcessValidCallTargets.Find(); err != nil {
		return nil
	}

	if address == 0 {
		return errors.New("attempt to register null call target")
	}

	regionSize := alignUp(size, 0x1000)
	info := cfgCallTargetInfo{
		Offset: 0,
		Flags:  cfgCallTargetValid,
	}

	ret, _, err := procSetProcessValidCallTargets.Call(
		uintptr(handle),
		address,
		regionSize,
		1,
		uintptr(unsafe.Pointer(&info)),
	)
	if ret == 0 {
		if err != nil {
			return fmt.Errorf("SetProcessValidCallTargets: %w", err)
		}
		return errors.New("SetProcessValidCallTargets failed")
	}

	return nil
}

func alignUp(value, alignment uintptr) uintptr {
	if alignment == 0 {
		return value
	}
	mask := alignment - 1
	return (value + mask) &^ mask
}
