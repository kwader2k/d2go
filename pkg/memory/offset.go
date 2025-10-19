package memory

import (
	"encoding/binary"
)

type Offset struct {
	GameData                    uintptr
	UnitTable                   uintptr
	UI                          uintptr
	Hover                       uintptr
	Expansion                   uintptr
	RosterOffset                uintptr
	PanelManagerContainerOffset uintptr
	WidgetStatesOffset          uintptr
	WaypointsOffset             uintptr
	FPS                         uintptr
	KeyBindingsOffset           uintptr
	KeyBindingsSkillsOffset     uintptr
	TZ                          uintptr
	Quests                      uintptr
	Ping                        uintptr
	LegacyGraphics              uintptr
}

func calculateOffsets(process *Process) Offset {
	// ignoring errors, always best practices
	memory, _ := process.getProcessMemory()

	// GameReader
	pattern := process.FindPattern(memory, "\x44\x88\x25\x00\x00\x00\x00\x66\x44\x89\x25\x00\x00\x00\x00", "xxx????xxxx????")
	bytes := process.ReadBytesFromMemory(pattern+0x3, 4)
	offsetInt := uintptr(binary.LittleEndian.Uint32(bytes))
	gameDataOffset := (pattern - process.moduleBaseAddressPtr) - 0x121 + offsetInt

	// UnitTable
	pattern = process.FindPattern(memory, "\x48\x03\xC7\x49\x8B\x8C\xC6", "xxxxxxx")
	bytes = process.ReadBytesFromMemory(pattern+7, 4)
	unitTableOffset := uintptr(binary.LittleEndian.Uint32(bytes))

	// UI
	pattern = process.FindPattern(memory, "\x40\x84\xed\x0f\x94\x05", "xxxxxx")
	uiOffset := process.ReadUInt(pattern+6, Uint32)
	uiOffsetPtr := (pattern - process.moduleBaseAddressPtr) + 10 + uintptr(uiOffset)

	// Hover
	pattern = process.FindPattern(memory, "\xc6\x84\xc2\x00\x00\x00\x00\x00\x48\x8b\x74", "xxx?????xxx")
	hoverOffset := process.ReadUInt(pattern+3, Uint32) - 1

	// Expansion
	pattern = process.FindPattern(memory, "\x48\x8B\x05\x00\x00\x00\x00\x48\x8B\xD9\xF3\x0F\x10\x50\x00", "xxx????xxxxxxx?")
	offsetPtr := uintptr(process.ReadUInt(pattern+3, Uint32))
	expOffset := pattern - process.moduleBaseAddressPtr + 7 + offsetPtr

	// Party members offset
	pattern = process.FindPattern(memory, "\x02\x45\x33\xD2\x4D\x8B", "xxxxxx")
	offsetPtr = uintptr(process.ReadUInt(pattern-3, Uint32))
	rosterOffset := pattern - process.moduleBaseAddressPtr + 1 + offsetPtr

	// PanelManagerContainer
	pattern = process.FindPatternByOperand(memory, "\x48\x89\x05\x00\x00\x00\x00\x48\x85\xDB\x74\x1E", "xxx????xxxxx")
	bytes = process.ReadBytesFromMemory(pattern, 8)
	panelManagerContainerOffset := (pattern - process.moduleBaseAddressPtr) // uintptr(binary.LittleEndian.Uint64(bytes))

	// WidgetStates
	pattern = process.FindPattern(memory, "\x48\x8B\x0D\x00\x00\x00\x00\x4C\x8D\x44\x24\x00\x48\x03\xC2", "xxx????xxxx?xxx")
	WidgetStatesPtr := process.ReadUInt(pattern+3, Uint32)
	WidgetStatesOffset := pattern - process.moduleBaseAddressPtr + 7 + uintptr(WidgetStatesPtr)

	// Waypoints
	pattern = process.FindPattern(memory, "\x48\x89\x05\x00\x00\x00\x00\x0F\x11\x00", "xxx????xxx")
	offsetBuffer := process.ReadUInt(pattern+3, Uint32)
	WaypointsOffset := pattern - process.moduleBaseAddressPtr + 23 + uintptr(offsetBuffer)

	// FPS
	pattern = process.FindPattern(memory, "\x8B\x1D\x00\x00\x00\x00\x48\x8D\x05\x00\x00\x00\x00\x48\x8D\x4C\x24\x40", "xx????xxx????xxxxx")
	fpsOffsetPtr := uintptr(process.ReadUInt(pattern+2, Uint32))
	fpsOffset := pattern - process.moduleBaseAddressPtr + 6 + fpsOffsetPtr

	// Keybindings
	pattern = process.FindPattern(memory, "\x48\x8D\x05\xAF\xEE", "xxxxx")
	bytes = process.ReadBytesFromMemory(pattern+3, 4)
	relativeOffset := int32(binary.LittleEndian.Uint32(bytes))
	keyBindingsOffset := pattern - process.moduleBaseAddressPtr + 7 + uintptr(relativeOffset)

	// KeyBindings Skills
	pattern = process.FindPattern(memory, "\x0F\x10\x04\x24\x48\x6B\xC8\x1C\x48\x8D\x05", "xxxxxxxxxxx")
	var keyBindingsSkillsOffset uintptr
	bytes = process.ReadBytesFromMemory(pattern+11, 4)
	relativeOffset = int32(binary.LittleEndian.Uint32(bytes))
	keyBindingsSkillsOffset = uintptr(int64(pattern) + 15 + int64(relativeOffset))

	// Terror Zones
	pattern = process.FindPattern(memory, "\x48\x89\x05\xCC\xCC\xCC\xCC\x48\x8D\x05\xCC\xCC\xCC\xCC\x48\x89\x05\xCC\xCC\xCC\xCC\x48\x8D\x05\xCC\xCC\xCC\xCC\x48\x89\x15\xCC\xCC\xCC\xCC\x48\x89\x15", "xxx????xxx????xxx????xxx????xxx????xxx")
	tzPtr := process.ReadUInt(pattern+3, Uint32)
	tzOffset := pattern - process.moduleBaseAddressPtr + 7 + uintptr(tzPtr)

	// Quest Bytes Data
	pattern = process.FindPattern(memory, "\x42\xc6\x84\x28\x00\x00\x00\x00\x00\x49\xff\xc5\x49\x83\xfd\x29", "xxxx?????xxxxxxx")
	bytes = process.ReadBytesFromMemory(pattern+4, 4)
	questOffset := uintptr(binary.LittleEndian.Uint32(bytes))
	questDataOffset := questOffset + 1

	// Ping
	pattern = process.FindPattern(memory, "\x48\x8B\x0D\xCC\xCC\xCC\xCC\x49\x2B\xC7", "xxx????xxx")
	bytes = process.ReadBytesFromMemory(pattern+3, 4)
	relativeOffset = int32(binary.LittleEndian.Uint32(bytes))
	pingOffset := pattern - process.moduleBaseAddressPtr + 7 + uintptr(relativeOffset)

	// LegacyGraphics
	pattern = process.FindPattern(memory, "\x80\x3D\x00\x00\x00\x00\x00\x48\x8D\x54\x24\x30", "xx?????xxxxx")
	legacyGfxPtr := uintptr(process.ReadUInt(pattern+2, Uint32))
	legacyGfxOffset := pattern - process.moduleBaseAddressPtr + 7 + legacyGfxPtr

	return Offset{
		GameData:                    gameDataOffset,
		UnitTable:                   unitTableOffset,
		UI:                          uiOffsetPtr,
		Hover:                       uintptr(hoverOffset),
		Expansion:                   expOffset,
		RosterOffset:                rosterOffset,
		PanelManagerContainerOffset: panelManagerContainerOffset,
		WidgetStatesOffset:          WidgetStatesOffset,
		WaypointsOffset:             WaypointsOffset,
		FPS:                         fpsOffset,
		KeyBindingsOffset:           keyBindingsOffset,
		KeyBindingsSkillsOffset:     keyBindingsSkillsOffset,
		TZ:                          tzOffset,
		Quests:                      questDataOffset,
		Ping:                        pingOffset,
		LegacyGraphics:              legacyGfxOffset,
	}
}
