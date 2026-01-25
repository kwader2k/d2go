package npc

import (
	"github.com/hectorgimenez/d2go/pkg/data/monplace"
	"github.com/hectorgimenez/d2go/pkg/data/superunique"
)

type PresetUnit struct {
	PresetID int // Constructed by appending monstats.txt + superuniques.txt + monplace.txt to match map object npc id (d2go relies on 1.13c map data).
}

var PresetUnitBySuperUniqueID = map[superunique.ID]PresetUnit{
	0:  {PresetID: 734}, // Bishibosh
	1:  {PresetID: 735}, // Bonebreak
	2:  {PresetID: 736}, // Coldcrow
	3:  {PresetID: 737}, // Rakanishu
	4:  {PresetID: 738}, // TreeheadWoodFist
	5:  {PresetID: 739}, // Griswold
	6:  {PresetID: 740}, // TheCountess
	7:  {PresetID: 741}, // PitspawnFouldog
	8:  {PresetID: 742}, // FlamespiketheCrawlerER
	9:  {PresetID: 743}, // Boneash
	10: {PresetID: 744}, // Radament
	11: {PresetID: 745}, // BloodwitchtheWild
	12: {PresetID: 746}, // Fangskin
	13: {PresetID: 747}, // Beetleburst
	14: {PresetID: 748}, // Leatherarm
	15: {PresetID: 749}, // ColdwormtheBurrowerR
	16: {PresetID: 750}, // FireEye
	17: {PresetID: 751}, // DarkElder
	18: {PresetID: 752}, // TheSummoner
	19: {PresetID: 753}, // AncientKaatheSoullessLESS
	20: {PresetID: 754}, // TheSmith
	21: {PresetID: 755}, // WebMagetheBurning
	22: {PresetID: 756}, // WitchDoctorEndugu
	23: {PresetID: 757}, // Stormtree
	24: {PresetID: 758}, // SarinatheBattlemaidD
	25: {PresetID: 759}, // IcehawkRiftwing
	26: {PresetID: 760}, // IsmailVilehand
	27: {PresetID: 761}, // GelebFlamefinger
	28: {PresetID: 762}, // BremmSparkfist
	29: {PresetID: 763}, // ToorcIcefist
	30: {PresetID: 764}, // WyandVoidfinger
	31: {PresetID: 765}, // MafferDragonhand
	32: {PresetID: 766}, // WingedDeath
	33: {PresetID: 767}, // TheTormentor
	34: {PresetID: 768}, // Taintbreeder
	35: {PresetID: 769}, // RiftwraiththeCannibalBAL
	36: {PresetID: 770}, // InfectorofSouls
	37: {PresetID: 771}, // LordDeSeis
	38: {PresetID: 772}, // GrandVizierofChaosS
	39: {PresetID: 773}, // TheCowKing
	40: {PresetID: 774}, // Corpsefire
	41: {PresetID: 775}, // TheFeatureCreep
	42: {PresetID: 776}, // SiegeBoss
	43: {PresetID: 777}, // AncientBarbarian1
	44: {PresetID: 778}, // AncientBarbarian2
	45: {PresetID: 779}, // AncientBarbarian3
	46: {PresetID: 780}, // AxeDweller
	47: {PresetID: 781}, // BonesawBreaker
	48: {PresetID: 782}, // DacFarren
	49: {PresetID: 783}, // MegaflowRectifier
	50: {PresetID: 784}, // EyebackUnleashed
	51: {PresetID: 785}, // ThreashSocket
	52: {PresetID: 786}, // Pindleskin
	53: {PresetID: 787}, // SnapchipShatter
	54: {PresetID: 788}, // AnodizedElite
	55: {PresetID: 789}, // VinvearMolech
	56: {PresetID: 790}, // SharpToothSayer
	57: {PresetID: 791}, // MagmaTorquer
	58: {PresetID: 792}, // BlazeRipper
	59: {PresetID: 793}, // Frozenstein
	60: {PresetID: 794}, // NihlathakBoss
	61: {PresetID: 795}, // BaalSubject1
	62: {PresetID: 796}, // BaalSubject2
	63: {PresetID: 797}, // BaalSubject3
	64: {PresetID: 798}, // BaalSubject4
	65: {PresetID: 799}, // BaalSubject5
}

var PresetUnitByMonPlaceID = map[monplace.ID]PresetUnit{
	0:  {PresetID: 800}, // nothing
	1:  {PresetID: 801}, // npc_pack
	2:  {PresetID: 802}, // unique_pack
	3:  {PresetID: 803}, // champion
	4:  {PresetID: 804}, // rogue_warner
	5:  {PresetID: 805}, // bloodraven
	6:  {PresetID: 806}, // rivermonster_right
	7:  {PresetID: 807}, // rivermonster_left
	8:  {PresetID: 808}, // tightspotboss
	9:  {PresetID: 809}, // amphibian
	10: {PresetID: 810}, // tentacle_ns
	11: {PresetID: 811}, // tentacle_ew
	12: {PresetID: 812}, // fallennest
	13: {PresetID: 813}, // fetishnest
	14: {PresetID: 814}, // talkingrogue
	15: {PresetID: 815}, // talkingguard
	16: {PresetID: 816}, // dumbguard
	17: {PresetID: 817}, // fallen
	18: {PresetID: 818}, // fallenshaman
	19: {PresetID: 819}, // maggot
	20: {PresetID: 820}, // maggotegg
	21: {PresetID: 821}, // mosquitonest
	22: {PresetID: 822}, // fetish
	23: {PresetID: 823}, // fetishshaman
	24: {PresetID: 824}, // impgroup
	25: {PresetID: 825}, // imp
	26: {PresetID: 826}, // miniongroup
	27: {PresetID: 827}, // minion
	28: {PresetID: 828}, // bloodlord
	29: {PresetID: 829}, // deadminion
	30: {PresetID: 830}, // deadimp
	31: {PresetID: 831}, // deadbarb
	32: {PresetID: 832}, // reanimateddead
	33: {PresetID: 833}, // group25
	34: {PresetID: 834}, // group50
	35: {PresetID: 835}, // group75
	36: {PresetID: 836}, // group100
}

func PresetUnitForSuperUniqueID(id superunique.ID) (PresetUnit, bool) {
	presetunit, ok := PresetUnitBySuperUniqueID[id]

	return presetunit, ok
}

func PresetUnitForMonPlaceID(id monplace.ID) (PresetUnit, bool) {
	presetunit, ok := PresetUnitByMonPlaceID[id]

	return presetunit, ok
}
