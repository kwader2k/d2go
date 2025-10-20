package quest

type Quest int16
type Status uint16

func (s Status) Completed() bool {
	return s&StatusMaskCompleted != 0
}

func (s Status) NotStarted() bool {
	return s == 0
}

func (s Status) HasStatus(st Status) bool {
	return (s & st) == st
}

// Note that quest order isnt what is represented visually in game
const (
	Act1DenOfEvil Quest = iota
	Act1SistersBurialGrounds
	Act1ToolsOfTheTrade
	Act1TheSearchForCain
	Act1TheForgottenTower
	Act1SistersToTheSlaughter
	Act2RadamentsLair
	Act2TheHoradricStaff
	Act2TaintedSun
	Act2ArcaneSanctuary
	Act2TheSummoner
	Act2TheSevenTombs
	Act3LamEsensTome
	Act3KhalimsWill
	Act3BladeOfTheOldReligion
	Act3TheGoldenBird
	Act3TheBlackenedTemple
	Act3TheGuardian
	Act4TheFallenAngel
	Act4HellForge
	Act4TerrorsEnd
	Act5SiegeOnHarrogath
	Act5RescueOnMountArreat
	Act5PrisonOfIce
	Act5BetrayalOfHarrogath
	Act5RiteOfPassage
	Act5EveOfDestruction
)

type Quests map[Quest]Status
