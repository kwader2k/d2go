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
	unitTableOffset := uintptr(0x1E9E350)

	// UI
	uiOffsetPtr := uintptr(0x1EAE04A)

	// Hover
	hoverOffset := uintptr(0x1DF2000)

	// Expansion
	expOffset := uintptr(0x1DF1468)

	// Party members offset
	rosterOffset := uintptr(0x1EB4668)

	// PanelManagerContainer
	panelManagerContainerOffset := uintptr(0x1E08DC0)

	// WidgetStates
	WidgetStatesOffset := uintptr(0x1ED6680)

	// Waypoints
	WaypointTableOffset := uintptr(0x1D503C0)

	// FPS
	fpsOffset := uintptr(0x1D50394)

	// KeyBindings
	keyBindingsOffset := uintptr(0x19C95B4)

	// KeyBindings Skills
	keyBindingsSkillsOffset := uintptr(0x1DF2110)

	// QuestInfo
	questInfoOffset := uintptr(0x1EBACD8)

	// Terror Zones
	tzOffset := uintptr(0x25A8AF0)

	// Ping
	pingOffset := uintptr(0x1DF1468)

	// LegacyGraphics
	legacyGfxOffset := uintptr(0x1EBAF46)

	// CharData
	charDataOffset := uintptr(0x1DF55F8)

	// Selected Char Name
	selectedCharNameOffset := uintptr(0x1D47195)

	// Last Game Name
	lastGameNameOffset := uintptr(0x25F1450)

	// Last Game Password
	lastGamePasswordOffset := uintptr(0x25F14A8)

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
