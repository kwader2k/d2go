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
	unitTableOffset := uintptr(0x1E9B350)

	// UI
	uiOffsetPtr := uintptr(0x1EAB042)

	// Hover
	hoverOffset := uintptr(0x1DEF000)

	// Expansion
	expOffset := uintptr(0x1DEE468)

	// Party members offset
	rosterOffset := uintptr(0x1EB1660)

	// PanelManagerContainer
	panelManagerContainerOffset := uintptr(0x1E05DC0)

	// WidgetStates
	WidgetStatesOffset := uintptr(0x1ED3678)

	// Waypoints
	WaypointTableOffset := uintptr(0x1D4D3C0)

	// FPS
	fpsOffset := uintptr(0x1D4D394)

	// KeyBindings
	keyBindingsOffset := uintptr(0x19C65B4)

	// KeyBindings Skills
	keyBindingsSkillsOffset := uintptr(0x1DEF12C)

	// QuestInfo
	questInfoOffset := uintptr(0x1EB7CD8)

	// Terror Zones
	tzOffset := uintptr(0x25A5AB0)

	// Ping
	pingOffset := uintptr(0x1DEE468)

	// LegacyGraphics
	legacyGfxOffset := uintptr(0x1EB7E7E)

	// CharData
	charDataOffset := uintptr(0x1DF25F8)

	// Selected Char Name
	selectedCharNameOffset := uintptr(0x1D44195)

	// Last Game Name
	lastGameNameOffset := uintptr(0x25EE370)

	// Last Game Password
	lastGamePasswordOffset := uintptr(0x25EE3C8)

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
