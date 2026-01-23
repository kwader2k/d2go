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
	SelectedCharName            uintptr
	LastGameName                uintptr
	LastGamePassword            uintptr
}

func calculateOffsets(_ *Process) Offset {
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

	// KeyBindings
	keyBindingsOffset := uintptr(0x18C2894)

	// KeyBindings Skills
	keyBindingsSkillsOffset := uintptr(0x1CE8510)

	// QuestInfo
	questInfoOffset := uintptr(0x1DB23D8)

	// Terror Zones
	tzOffset := uintptr(0x248D430)

	// Ping
	pingOffset := uintptr(0x1CE78D0)

	// LegacyGraphics
	legacyGfxOffset := uintptr(0x1DB263E)

	// CharData
	charDataOffset := uintptr(0x1CED5E8)

	// Selected Char Name
	selectedCharNameOffset := uintptr(0x1C3D694)

	// Last Game Name
	lastGameNameOffset := uintptr(0x24D5A90)

	// Last Game Password
	lastGamePasswordOffset := uintptr(0x24D5AE8)

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
		KeyBindingsOffset:           keyBindingsOffset,
		KeyBindingsSkillsOffset:     keyBindingsSkillsOffset,
		QuestInfo:                   questInfoOffset,
		TZ:                          tzOffset,
		Ping:                        pingOffset,
		LegacyGraphics:              legacyGfxOffset,
		CharData:                    charDataOffset,
		SelectedCharName:            selectedCharNameOffset,
		LastGameName:                lastGameNameOffset,
		LastGamePassword:            lastGamePasswordOffset,
	}
}
