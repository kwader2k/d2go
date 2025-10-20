package memory

import (
	"github.com/hectorgimenez/d2go/pkg/data/quest"
)

func (gd *GameReader) getQuests(questBytes []byte) quest.Quests {
	return quest.Quests{
		// Act 1
		quest.Act1DenOfEvil:             gd.getQuestStatus(gd.readQuestFlags(questBytes, 1)),
		quest.Act1SistersBurialGrounds:  gd.getQuestStatus(gd.readQuestFlags(questBytes, 2)),
		quest.Act1ToolsOfTheTrade:       gd.getQuestStatus(gd.readQuestFlags(questBytes, 3)),
		quest.Act1TheSearchForCain:      gd.getQuestStatus(gd.readQuestFlags(questBytes, 4)),
		quest.Act1TheForgottenTower:     gd.getQuestStatus(gd.readQuestFlags(questBytes, 5)),
		quest.Act1SistersToTheSlaughter: gd.getQuestStatus(gd.readQuestFlags(questBytes, 6)),
		// Act 2
		quest.Act2RadamentsLair:    gd.getQuestStatus(gd.readQuestFlags(questBytes, 9)),
		quest.Act2TheHoradricStaff: gd.getQuestStatus(gd.readQuestFlags(questBytes, 10)),
		quest.Act2TaintedSun:       gd.getQuestStatus(gd.readQuestFlags(questBytes, 11)),
		quest.Act2ArcaneSanctuary:  gd.getQuestStatus(gd.readQuestFlags(questBytes, 12)),
		quest.Act2TheSummoner:      gd.getQuestStatus(gd.readQuestFlags(questBytes, 13)),
		quest.Act2TheSevenTombs:    gd.getQuestStatus(gd.readQuestFlags(questBytes, 14)),
		// Act 3
		quest.Act3LamEsensTome:          gd.getQuestStatus(gd.readQuestFlags(questBytes, 17)),
		quest.Act3KhalimsWill:           gd.getQuestStatus(gd.readQuestFlags(questBytes, 18)),
		quest.Act3BladeOfTheOldReligion: gd.getQuestStatus(gd.readQuestFlags(questBytes, 19)),
		quest.Act3TheGoldenBird:         gd.getQuestStatus(gd.readQuestFlags(questBytes, 20)),
		quest.Act3TheBlackenedTemple:    gd.getQuestStatus(gd.readQuestFlags(questBytes, 21)),
		quest.Act3TheGuardian:           gd.getQuestStatus(gd.readQuestFlags(questBytes, 22)),
		// Act 4
		quest.Act4TheFallenAngel: gd.getQuestStatus(gd.readQuestFlags(questBytes, 25)),
		quest.Act4HellForge:      gd.getQuestStatus(gd.readQuestFlags(questBytes, 27)),
		quest.Act4TerrorsEnd:     gd.getQuestStatus(gd.readQuestFlags(questBytes, 26)),
		// Act 5
		quest.Act5SiegeOnHarrogath:    gd.getQuestStatus(gd.readQuestFlags(questBytes, 35)),
		quest.Act5RescueOnMountArreat: gd.getQuestStatus(gd.readQuestFlags(questBytes, 36)),
		quest.Act5PrisonOfIce:         gd.getQuestStatus(gd.readQuestFlags(questBytes, 37)),
		quest.Act5BetrayalOfHarrogath: gd.getQuestStatus(gd.readQuestFlags(questBytes, 38)),
		quest.Act5RiteOfPassage:       gd.getQuestStatus(gd.readQuestFlags(questBytes, 39)),
		quest.Act5EveOfDestruction:    gd.getQuestStatus(gd.readQuestFlags(questBytes, 40)),
	}
}

func (gd *GameReader) readQuestFlags(questBytes []byte, questIndex int) uint16 {
	byteOffset := 2 * questIndex
	if byteOffset+1 >= len(questBytes) {
		return 0
	}
	return uint16(questBytes[byteOffset]) | (uint16(questBytes[byteOffset+1]) << 8)
}

func (gd *GameReader) getQuestStatus(flags uint16) quest.Status {
	return quest.Status(flags)
}
