package data

import (
	"github.com/hectorgimenez/d2go/pkg/data/mode"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/hectorgimenez/d2go/pkg/data/state"
)

type NPC struct {
	ID        npc.ID
	Name      string
	Positions []Position
}

type MonsterType string

type Monster struct {
	UnitID
	Name      npc.ID
	IsHovered bool
	Position  Position
	Stats     map[stat.ID]int
	Type      MonsterType
	States    state.States
	Mode      mode.NpcMode
}

type Monsters []Monster
type NPCs []NPC

func (n NPCs) FindOne(npcid npc.ID) (NPC, bool) {
	for _, np := range n {
		if np.ID == npcid {
			return np, true
		}
	}

	return NPC{}, false
}

func (m Monsters) FindOne(id npc.ID, t MonsterType) (Monster, bool) {
	for _, monster := range m {
		if monster.Name == id {
			if t == MonsterTypeNone || t == monster.Type {
				return monster, true
			}
		}
	}

	return Monster{}, false
}

func (m Monsters) Enemies(filters ...MonsterFilter) []Monster {
	monsters := make([]Monster, 0)
	for _, mo := range m {
		if !mo.IsMerc() && !mo.IsSkip() && !mo.IsGoodNPC() && !mo.IsPet() && mo.Stats[stat.Life] > 0 {
			monsters = append(monsters, mo)
		}
	}

	for _, f := range filters {
		monsters = f(monsters)
	}

	return monsters
}

type MonsterFilter func(m Monsters) []Monster

func MonsterEliteFilter() MonsterFilter {
	return func(m Monsters) []Monster {
		var filteredMonsters []Monster
		for _, mo := range m {
			if mo.IsElite() {
				filteredMonsters = append(filteredMonsters, mo)
			}
		}

		return filteredMonsters
	}
}

func MonsterAnyFilter() MonsterFilter {
	return func(m Monsters) []Monster {
		return m
	}
}

func (m Monsters) FindByID(id UnitID) (Monster, bool) {
	for _, monster := range m {
		if monster.UnitID == id {
			return monster, true
		}
	}

	return Monster{}, false
}

func (m Monster) IsImmune(resist stat.Resist) bool {
	for st, value := range m.Stats {
		// Only consider immunity threshold (>= 100).
		if value < 100 {
			continue
		}
		if resist == stat.ColdImmune && st == stat.ColdResist {
			return true
		}
		if resist == stat.FireImmune && st == stat.FireResist {
			return true
		}
		if resist == stat.LightImmune && st == stat.LightningResist {
			return true
		}
		if resist == stat.PoisonImmune && st == stat.PoisonResist {
			return true
		}
		if resist == stat.MagicImmune && st == stat.MagicResist {
			return true
		}
	}

	return false
}

func (m Monster) IsMerc() bool {
	if m.Name == npc.Guard || m.Name == npc.Act5Hireling1Hand || m.Name == npc.Act5Hireling2Hand || m.Name == npc.IronWolf || m.Name == npc.Rogue2 {
		return true
	}

	return false
}

func (m Monster) IsPet() bool {
	// Necromancer revive.
	if m.States.HasState(state.Revive) {
		return true
	}

	switch m.Name {
	case npc.DruHawk, npc.DruSpiritWolf, npc.DruFenris, npc.HeartOfWolverine,
		npc.OakSage, npc.DruBear, npc.DruPlaguePoppy, npc.VineCreature,
		npc.DruCycleOfLife, npc.ClayGolem, npc.BloodGolem, npc.IronGolem,
		npc.FireGolem, npc.NecroSkeleton, npc.NecroMage, npc.Valkyrie, npc.Decoy,
		npc.ShadowWarrior, npc.ShadowMaster:
		return true
	}

	return false
}

func (m Monster) IsGoodNPC() bool {
	switch m.Name {
	case 146, 154, 147, 150, 155, 148, 244, 210, 175, 199, 198, 177, 178, 201, 202, 200, 331, 245, 264, 255, 176,
		252, 254, 253, 297, 246, 251, 367, 521, 257, 405, 265, 520, 512, 518, 527, 515, 513, 511, 514, 266, 408, 406,
		543: // Baal Throne
		return true
	}

	return false
}

// IsPrimeEvil returns true if the monster is a Prime Evil (act boss).
func (m Monster) IsPrimeEvil() bool {
	flags, ok := npc.MonStatsFlagsForID(m.Name)
	return ok && flags.IsPrimeEvil
}

// IsUber returns true if the monster is an Uber boss.
func (m Monster) IsUber() bool {
	switch m.Name {
	case npc.Lilith, npc.UberDuriel, npc.UberIzual, npc.UberMephisto, npc.UberDiablo, npc.UberBaal:
		return true
	}

	return false
}

// IsElite returns true if the monster is an elite (champion or above).
func (m Monster) IsElite() bool {
	return m.Type == MonsterTypeChampion || m.Type == MonsterTypeMinion || m.Type == MonsterTypeUnique || m.Type == MonsterTypeSuperUnique
}

// IsSealElite returns true if the monster is an elite from a Chaos Sanctuary seal.
func (m Monster) IsSealElite() bool {
	return m.Type == MonsterTypeSuperUnique &&
		(m.Name == npc.OblivionKnight || // Lord De Seis
			m.Name == npc.VenomLord || // Infector of Souls
			m.Name == npc.StormCaster) // Grand Vizier of Chaos
}

// IsMonsterRaiser returns true if the monster can spawn new monsters.
func (m Monster) IsMonsterRaiser() bool {
	switch m.Name {
	case npc.FallenShaman, npc.CarverShaman, npc.DevilkinShaman, npc.DarkShaman, npc.WarpedShaman,
		npc.HollowOne, npc.Guardian, npc.Guardian2, npc.Unraveler, npc.Unraveler2, npc.HoradrimAncient, npc.BaalSubjectMummy, npc.HoradrimAncient2, npc.HoradrimAncient3:
		return true
	}

	return false
}

// IsSkip returns true if the monster cannot be killed as a normal enemy, for example cannot be targeted.
func (m Monster) IsSkip() bool {
	switch m.Name {
	case npc.WaterWatcherLimb, npc.WaterWatcherHead, npc.BaalTaunt, npc.Act5Combatant, npc.Act5Combatant2, npc.BarricadeTower, npc.DarkWanderer, npc.POW:
		return true
	}

	return false
}

// IsEscapingType returns true if the monster cannot be attacked when airborne or hiding in water (NpcMode 8).
func (m Monster) IsEscapingType() bool {
	switch m.Name {
	case npc.CarrionBird, npc.CarrionBird2, npc.WaterWatcherLimb, npc.RiverStalkerLimb, npc.StygianWatcherLimb,
		npc.WaterWatcherHead, npc.RiverStalkerHead, npc.StygianWatcherHead, npc.CloudStalker, npc.Sucker, npc.UndeadScavenger, npc.FoulCrow:
		return true
	}

	return false
}

// IsUndead returns true if the monster is undead.
func (m Monster) IsUndead() bool {
	flags, ok := npc.MonStatsFlagsForID(m.Name)
	return ok && (flags.IsLUndead || flags.IsHUndead)
}

// IsDemon returns true if the monster is a demon.
func (m Monster) IsDemon() bool {
	flags, ok := npc.MonStatsFlagsForID(m.Name)
	return ok && flags.IsDemon
}

// IsBeast returns true if the monster is a beast.
func (m Monster) IsBeast() bool {
	flags, ok := npc.MonStatsFlagsForID(m.Name)
	return ok && !flags.IsLUndead && !flags.IsHUndead && !flags.IsDemon
}

// IsUndeadOrDemon returns true if the monster is undead or a demon.
func (m Monster) IsUndeadOrDemon() bool {
	flags, ok := npc.MonStatsFlagsForID(m.Name)
	return ok && (flags.IsLUndead || flags.IsHUndead || flags.IsDemon)
}
