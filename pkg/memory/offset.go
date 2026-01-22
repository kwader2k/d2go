package memory

type Offset struct {
	GameData                    uintptr
	UnitTable                   uintptr
	UI                          uintptr
	Hover                       uintptr
	Expansion                   uintptr
	RosterOffset                uintptr
	PanelManagerContainerOffset uintptr
	WidgetStatesOffset          uintptr
	WaypointTableOffset         uintptr
	FPS                         uintptr
	KeyBindingsOffset           uintptr
	KeyBindingsSkillsOffset     uintptr
	QuestInfo                   uintptr
	TZ                          uintptr
	Quests                      uintptr
	Ping                        uintptr
	LegacyGraphics              uintptr
	CharData                    uintptr
}

func calculateOffsets(process *Process) Offset {
	memory, err := process.getProcessMemory()
	if err != nil || len(memory) == 0 {
		// Return empty offsets - will be recalculated when game is ready
		return Offset{}
	}

	// UnitTable
	unitTableOffset := uintptr(0x1D95AF0)

	// UI
	uiOffsetPtr := uintptr(0x1DA57DA)

	// Hover
	hoverOffset := uintptr(0x1CE8400)

	// Expansion
	expOffset := uintptr(0x1CE78D0)

	// Party members offset
	rosterOffset := uintptr(0x1DABD60)

	// PanelManagerContainer
	panelManagerContainerOffset := uintptr(0x1D00968)

	// WidgetStates
	WidgetStatesOffset := uintptr(0x1DCDD40)

	// Waypoints
	WaypointTableOffset := uintptr(0x1C468B0)

	// FPS
	fpsOffset := uintptr(0x1C46894)

	// KeyBindings Skills
	keyBindingsSkillsOffset := uintptr(0x1CE8510)

	// QuestInfo
	questInfoOffset := uintptr(0)

	// Terror Zones
	tzOffset := uintptr(0x248D4C8)

	// Ping
	pingOffset := uintptr(0)

	// LegacyGraphics
	legacyGfxOffset := uintptr(0x1DB263E)

	// CharData
	charDataOffset := uintptr(0x1CED5E8)

	return Offset{
		UnitTable:                   unitTableOffset,
		UI:                          uiOffsetPtr,
		Hover:                       hoverOffset,
		Expansion:                   expOffset,
		RosterOffset:                rosterOffset,
		PanelManagerContainerOffset: panelManagerContainerOffset,
		WidgetStatesOffset:          WidgetStatesOffset,
		WaypointTableOffset:         WaypointTableOffset,
		FPS:                         fpsOffset,
		KeyBindingsSkillsOffset:     keyBindingsSkillsOffset,
		QuestInfo:                   questInfoOffset,
		TZ:                          tzOffset,
		Ping:                        pingOffset,
		LegacyGraphics:              legacyGfxOffset,
		CharData:                    charDataOffset,
	}
}
