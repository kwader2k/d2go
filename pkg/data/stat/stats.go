package stat

import (
	"strconv"
	"strings"
)

type ID int16

type Data struct {
	ID    ID
	Value int
	Layer int
}

func (d Data) String() string {
	return strings.Replace(StatStringMap[int(d.ID)][d.Layer], "#", strconv.Itoa(d.Value), 1)
}

type Stats []Data

func (i Stats) FindStat(id ID, layer int) (Data, bool) {
	for _, s := range i {
		if s.ID == id && s.Layer == layer {
			return s, true
		}
	}

	return Data{}, false
}
func (s ID) String() string {
	return StringStats[s]
}

const (
	Strength ID = iota
	Energy
	Dexterity
	Vitality
	StatPoints
	SkillPoints
	Life
	MaxLife
	Mana
	MaxMana
	Stamina
	MaxStamina
	Level
	Experience
	Gold
	StashGold
	EnhancedDefense
	EnhancedDamageMin
	EnhancedDamage
	AttackRating
	ChanceToBlock
	MinDamage
	MaxDamage
	TwoHandedMinDamage
	TwoHandedMaxDamage
	DamagePercent
	ManaRecovery
	ManaRecoveryBonus
	StaminaRecoveryBonus
	LastExp
	NextExp
	Defense
	DefenseVsMissiles
	DefenseVsHth
	NormalDamageReduction
	MagicDamageReduction
	DamageReduced
	MagicResist
	MaxMagicResist
	FireResist
	MaxFireResist
	LightningResist
	MaxLightningResist
	ColdResist
	MaxColdResist
	PoisonResist
	MaxPoisonResist
	DamageAura
	FireMinDamage
	FireMaxDamage
	LightningMinDamage
	LightningMaxDamage
	MagicMinDamage
	MagicMaxDamage
	ColdMinDamage
	ColdMaxDamage
	ColdLength
	PoisonMinDamage
	PoisonMaxDamage
	PoisonLength
	LifeSteal
	LifeStealMax
	ManaSteal
	ManaStealMax
	StaminaDrainMinDamage
	StaminaDrainMaxDamage
	StunLength
	VelocityPercent
	AttackRate
	OtherAnimRate
	Quantity
	Value
	Durability
	MaxDurability
	ReplenishLife
	MaxDurabilityPercent
	MaxLifePercent
	MaxManaPercent
	AttackerTakesDamage
	GoldFind
	MagicFind
	Knockback
	TimeDuration
	AddClassSkills
	Unused84
	AddExperience
	LifeAfterEachKill
	ReducePrices
	DoubleHerbDuration
	LightRadius
	LightColor
	Requirements
	LevelRequire
	IncreasedAttackSpeed
	LevelRequirePercent
	LastBlockFrame
	FasterRunWalk
	NonClassSkill
	State
	FasterHitRecovery
	PlayerCount
	PoisonOverrideLength
	FasterBlockRate
	BypassUndead
	BypassDemons
	FasterCastRate
	BypassBeasts
	SingleSkill
	SlainMonstersRestInPeace
	CurseResistance
	PoisonLengthReduced
	NormalDamage
	HitCausesMonsterToFlee
	HitBlindsTarget
	DamageTakenGoesToMana
	IgnoreTargetsDefense
	TargetDefense
	PreventMonsterHeal
	HalfFreezeDuration
	AttackRatingPercent
	MonsterDefensePerHit
	DemonDamagePercent
	UndeadDamagePercent
	DemonAttackRating
	UndeadAttackRating
	Throwable
	FireSkills
	AllSkills
	AttackerTakesLightDamage
	IronMaidenLevel
	LifeTapLevel
	ThornsPercent
	BoneArmor
	BoneArmorMax
	FreezesTarget
	OpenWounds
	CrushingBlow
	KickDamage
	ManaAfterKill
	HealAfterDemonKill
	ExtraBlood
	DeadlyStrike
	AbsorbFirePercent
	AbsorbFire
	AbsorbLightningPercent
	AbsorbLightning
	AbsorbMagicPercent
	AbsorbMagic
	AbsorbColdPercent
	AbsorbCold
	SlowsTarget
	Aura
	Indestructible
	CannotBeFrozen
	SlowerStaminaDrain
	Reanimate
	Pierce
	MagicArrow
	ExplosiveArrow
	ThrowMinDamage
	ThrowMaxDamage
	SkillHandofAthena
	SkillStaminaPercent
	SkillPassiveStaminaPercent
	SkillConcentration
	SkillEnchant
	SkillPierce
	SkillConviction
	SkillChillingArmor
	SkillFrenzy
	SkillDecrepify
	SkillArmorPercent
	Alignment
	Target0
	Target1
	GoldLost
	ConverisonLevel
	ConverisonMaxHP
	UnitDooverlay
	AttackVsMonType
	DamageVsMonType
	Fade
	ArmorOverridePercent
	Unused183
	Unused184
	Unused185
	Unused186
	Unused187
	AddSkillTab
	Unused189
	Unused190
	Unused191
	Unused192
	Unused193
	NumSockets
	SkillOnAttack
	SkillOnKill
	SkillOnDeath
	SkillOnHit
	SkillOnLevelUp
	Unused200
	SkillOnGetHit
	Unused202
	Unused203
	ItemChargedSkill
	Unused205
	Unused206
	Unused207
	Unused208
	Unused209
	Unused210
	Unused211
	Unused213
	Unused212
	DefensePerLevel
	ArmorPercentPerLevel
	LifePerLevel
	ManaPerLevel
	MaxDamagePerLevel
	MaxDamagePercentPerLevel
	StrengthPerLevel
	DexterityPerLevel
	EnergyPerLevel
	VitalityPerLevel
	AttackRatingPerLevel
	AttackRatingPercentPerLevel
	ColdDamageMaxPerLevel
	FireDamageMaxPerLevel
	LightningDamageMaxPerLevel
	PoisonDamageMaxPerLevel
	ResistColdPerLevel
	ResistFirePerLevel
	ResistLightningPerLevel
	ResistPoisonPerLevel
	AbsorbColdPerLevel
	AbsorbFirePerLevel
	AbsorbLightningPerLevel
	AbsorbPoisonPerLevel
	ThornsPerLevel
	ExtraGoldPerLevel
	MagicFindPerLevel
	RegenStaminaPerLevel
	StaminaPerLevel
	DamageDemonPerLevel
	DamageUndeadPerLevel
	AttackRatingDemonPerLevel
	AttackRatingUndeadPerLevel
	CrushingBlowPerLevel
	OpenWoundsPerLevel
	KickDamagePerLevel
	DeadlyStrikePerLevel
	FindGemsPerLevel
	ReplenishDurability
	ReplenishQuantity
	ExtraStack
	FindItem
	SlashDamage
	SlashDamagePercent
	CrushDamage
	CrushDamagePercent
	ThrustDamage
	ThrustDamagePercent
	AbsorbSlash
	AbsorbCrush
	AbsorbThrust
	AbsorbSlashPercent
	AbsorbCrushPercent
	AbsorbThrustPercent
	ArmorByTime
	ArmorPercentByTime
	LifeByTime
	ManaByTime
	MaxDamageByTime
	MaxDamagePercentByTime
	StrengthByTime
	DexterityByTime
	EnergyByTime
	VitalityByTime
	AttackRatingByTime
	AttackRatingPercentByTime
	ColdDamageMaxByTime
	FireDamageMaxByTime
	LightningDamageMaxByTime
	PoisonDamageMaxByTime
	ResistColdByTime
	ResistFireByTime
	ResistLightningByTime
	ResistPoisonByTime
	AbsorbColdByTime
	AbsorbFireByTime
	AbsorbLightningByTime
	AbsorbPoisonByTime
	FindGoldByTime
	MagicFindByTime
	RegenStaminaByTime
	StaminaByTime
	DamageDemonByTime
	DamageUndeadByTime
	AttackRatingDemonByTime
	AttackRatingUndeadByTime
	CrushingBlowByTime
	OpenWoundsByTime
	KickDamageByTime
	DeadlyStrikeByTime
	FindGemsByTime
	PierceCold
	PierceFire
	PierceLightning
	PiercePoison
	DamageVsMonster
	DamagePercentVsMonster
	AttackRatingVsMonster
	AttackRatingPercentVsMonster
	AcVsMonster
	AcPercentVsMonster
	FireLength
	BurningMin
	BurningMax
	ProgressiveDamage
	ProgressiveSteal
	ProgressiveOther
	ProgressiveFire
	ProgressiveCold
	ProgressiveLightning
	ExtraCharges
	ProgressiveAttackRating
	PoisonCount
	DamageFrameRate
	PierceIdx
	FireSkillDamage
	LightningSkillDamage
	ColdSkillDamage
	PoisonSkillDamage
	EnemyFireResist
	EnemyLightningResist
	EnemyColdResist
	EnemyPoisonResist
	PassiveCriticalStrike
	PassiveDodge
	PassiveAvoid
	PassiveEvade
	PassiveWarmth
	PassiveMasteryMeleeAttackRating
	PassiveMasteryMeleeDamage
	PassiveMasteryMeleeCritical
	PassiveMasteryThrowAttackRating
	PassiveMasteryThrowDamage
	PassiveMasteryThrowCritical
	PassiveWeaponBlock
	SummonResist
	ModifierListSkill
	ModifierListLevel
	LastSentHPPercent
	SourceUnitType
	SourceUnitID
	ShortParam1
	QuestItemDifficulty
	PassiveMagicMastery
	PassiveMagicPierce
	SkillCooldown
	SkillMissileDamageScale
	Psychicward
	Psychicwardmax
	SkillChannelingTick
	CustomizationIndex
	MagicDamageMaxPerLevel
	PassiveDamagePierce
	HeraldTier
)

var StringStats = []string{
	"strength",
	"energy",
	"dexterity",
	"vitality",
	"statpoints",
	"skillpoints",
	"life",
	"maxlife",
	"mana",
	"maxmana",
	"stamina",
	"maxstamina",
	"level",
	"experience",
	"gold",
	"stashgold",
	"enhanceddefense",
	"enhanceddamagemin",
	"enhanceddamage",
	"attackrating",
	"chancetoblock",
	"mindamage",
	"maxdamage",
	"twohandedmindamage",
	"twohandedmaxdamage",
	"damagepercent",
	"manarecovery",
	"manarecoverybonus",
	"staminarecoverybonus",
	"lastexp",
	"nextexp",
	"defense",
	"defensevsmissiles",
	"defensevshth",
	"normaldamagereduction",
	"magicdamagereduction",
	"damagereduced",
	"magicresist",
	"maxmagicresist",
	"fireresist",
	"maxfireresist",
	"lightningresist",
	"maxlightningresist",
	"coldresist",
	"maxcoldresist",
	"poisonresist",
	"maxpoisonresist",
	"damageaura",
	"firemindamage",
	"firemaxdamage",
	"lightningmindamage",
	"lightningmaxdamage",
	"magicmindamage",
	"magicmaxdamage",
	"coldmindamage",
	"coldmaxdamage",
	"coldlength",
	"poisonmindamage",
	"poisonmaxdamage",
	"poisonlength",
	"lifesteal",
	"lifestealmax",
	"manasteal",
	"manastealmax",
	"staminadrainmindamage",
	"staminadrainmaxdamage",
	"stunlength",
	"velocitypercent",
	"attackrate",
	"otheranimrate",
	"quantity",
	"value",
	"durability",
	"maxdurability",
	"replenishlife",
	"maxdurabilitypercent",
	"maxlifepercent",
	"maxmanapercent",
	"attackertakesdamage",
	"goldfind",
	"magicfind",
	"knockback",
	"timeduration",
	"addclassskills",
	"unused84",
	"addexperience",
	"lifeaftereachkill",
	"reduceprices",
	"doubleherbduration",
	"lightradius",
	"lightcolor",
	"requirements",
	"levelrequire",
	"increasedattackspeed",
	"levelrequirepercent",
	"lastblockframe",
	"fasterrunwalk",
	"nonclassskill",
	"state",
	"fasterhitrecovery",
	"playercount",
	"poisonoverridelength",
	"fasterblockrate",
	"bypassundead",
	"bypassdemons",
	"fastercastrate",
	"bypassbeasts",
	"singleskill",
	"slainmonstersrestinpeace",
	"curseresistance",
	"poisonlengthreduced",
	"normaldamage",
	"hitcausesmonstertoflee",
	"hitblindstarget",
	"damagetakengoestomana",
	"ignoretargetsdefense",
	"targetdefense",
	"preventmonsterheal",
	"halffreezeduration",
	"attackratingpercent",
	"monsterdefenseperhit",
	"demondamagepercent",
	"undeaddamagepercent",
	"demonattackrating",
	"undeadattackrating",
	"throwable",
	"fireskills",
	"allskills",
	"attackertakeslightdamage",
	"ironmaidenlevel",
	"lifetaplevel",
	"thornspercent",
	"bonearmor",
	"bonearmormax",
	"freezestarget",
	"openwounds",
	"crushingblow",
	"kickdamage",
	"manaafterkill",
	"healafterdemonkill",
	"extrablood",
	"deadlystrike",
	"absorbfirepercent",
	"absorbfire",
	"absorblightningpercent",
	"absorblightning",
	"absorbmagicpercent",
	"absorbmagic",
	"absorbcoldpercent",
	"absorbcold",
	"slowstarget",
	"aura",
	"indestructible",
	"cannotbefrozen",
	"slowerstaminadrain",
	"reanimate",
	"pierce",
	"magicarrow",
	"explosivearrow",
	"throwmindamage",
	"throwmaxdamage",
	"skillhandofathena",
	"skillstaminapercent",
	"skillpassivestaminapercent",
	"skillconcentration",
	"skillenchant",
	"skillpierce",
	"skillconviction",
	"skillchillingarmor",
	"skillfrenzy",
	"skilldecrepify",
	"skillarmorpercent",
	"alignment",
	"target0",
	"target1",
	"goldlost",
	"converisonlevel",
	"converisonmaxhp",
	"unitdooverlay",
	"attackvsmontype",
	"damagevsmontype",
	"fade",
	"armoroverridepercent",
	"unused183",
	"unused184",
	"unused185",
	"unused186",
	"unused187",
	"addskilltab",
	"unused189",
	"unused190",
	"unused191",
	"unused192",
	"unused193",
	"numsockets",
	"skillonattack",
	"skillonkill",
	"skillondeath",
	"skillonhit",
	"skillonlevelup",
	"unused200",
	"skillongethit",
	"unused202",
	"unused203",
	"itemchargedskill",
	"unused205",
	"unused206",
	"unused207",
	"unused208",
	"unused209",
	"unused210",
	"unused211",
	"unused213",
	"unused212",
	"defenseperlevel",
	"armorpercentperlevel",
	"lifeperlevel",
	"manaperlevel",
	"maxdamageperlevel",
	"maxdamagepercentperlevel",
	"strengthperlevel",
	"dexterityperlevel",
	"energyperlevel",
	"vitalityperlevel",
	"attackratingperlevel",
	"attackratingpercentperlevel",
	"colddamagemaxperlevel",
	"firedamagemaxperlevel",
	"lightningdamagemaxperlevel",
	"poisondamagemaxperlevel",
	"resistcoldperlevel",
	"resistfireperlevel",
	"resistlightningperlevel",
	"resistpoisonperlevel",
	"absorbcoldperlevel",
	"absorbfireperlevel",
	"absorblightningperlevel",
	"absorbpoisonperlevel",
	"thornsperlevel",
	"extragoldperlevel",
	"magicfindperlevel",
	"regenstaminaperlevel",
	"staminaperlevel",
	"damagedemonperlevel",
	"damageundeadperlevel",
	"attackratingdemonperlevel",
	"attackratingundeadperlevel",
	"crushingblowperlevel",
	"openwoundsperlevel",
	"kickdamageperlevel",
	"deadlystrikeperlevel",
	"findgemsperlevel",
	"replenishdurability",
	"replenishquantity",
	"extrastack",
	"finditem",
	"slashdamage",
	"slashdamagepercent",
	"crushdamage",
	"crushdamagepercent",
	"thrustdamage",
	"thrustdamagepercent",
	"absorbslash",
	"absorbcrush",
	"absorbthrust",
	"absorbslashpercent",
	"absorbcrushpercent",
	"absorbthrustpercent",
	"armorbytime",
	"armorpercentbytime",
	"lifebytime",
	"manabytime",
	"maxdamagebytime",
	"maxdamagepercentbytime",
	"strengthbytime",
	"dexteritybytime",
	"energybytime",
	"vitalitybytime",
	"attackratingbytime",
	"attackratingpercentbytime",
	"colddamagemaxbytime",
	"firedamagemaxbytime",
	"lightningdamagemaxbytime",
	"poisondamagemaxbytime",
	"resistcoldbytime",
	"resistfirebytime",
	"resistlightningbytime",
	"resistpoisonbytime",
	"absorbcoldbytime",
	"absorbfirebytime",
	"absorblightningbytime",
	"absorbpoisonbytime",
	"findgoldbytime",
	"magicfindbytime",
	"regenstaminabytime",
	"staminabytime",
	"damagedemonbytime",
	"damageundeadbytime",
	"attackratingdemonbytime",
	"attackratingundeadbytime",
	"crushingblowbytime",
	"openwoundsbytime",
	"kickdamagebytime",
	"deadlystrikebytime",
	"findgemsbytime",
	"piercecold",
	"piercefire",
	"piercelightning",
	"piercepoison",
	"damagevsmonster",
	"damagepercentvsmonster",
	"attackratingvsmonster",
	"attackratingpercentvsmonster",
	"acvsmonster",
	"acpercentvsmonster",
	"firelength",
	"burningmin",
	"burningmax",
	"progressivedamage",
	"progressivesteal",
	"progressiveother",
	"progressivefire",
	"progressivecold",
	"progressivelightning",
	"extracharges",
	"progressiveattackrating",
	"poisoncount",
	"damageframerate",
	"pierceidx",
	"fireskilldamage",
	"lightningskilldamage",
	"coldskilldamage",
	"poisonskilldamage",
	"enemyfireresist",
	"enemylightningresist",
	"enemycoldresist",
	"enemypoisonresist",
	"passivecriticalstrike",
	"passivedodge",
	"passiveavoid",
	"passiveevade",
	"passivewarmth",
	"passivemasterymeleeattackrating",
	"passivemasterymeleedamage",
	"passivemasterymeleecritical",
	"passivemasterythrowattackrating",
	"passivemasterythrowdamage",
	"passivemasterythrowcritical",
	"passiveweaponblock",
	"summonresist",
	"modifierlistskill",
	"modifierlistlevel",
	"lastsenthppercent",
	"sourceunittype",
	"sourceunitid",
	"shortparam1",
	"questitemdifficulty",
	"passivemagicmastery",
	"passivemagicpierce",
	"skillcooldown",
	"skillmissiledamagescale",
	"psychicward",
	"psychicwardmax",
	"skillchannelingtick",
	"customizationindex",
	"magicdamagemaxperlevel",
	"passivedamagepierce",
	"heraldtier",
}

// StatStringMap maps [Stat ID] -> [layer] -> [Stat String].
// Format strings use "#" as a placeholder for the stat value.
var StatStringMap = map[int]map[int]string{
	0:  {0: "+# to Strength"},
	1:  {0: "+# to Energy"},
	2:  {0: "+# to Dexterity"},
	3:  {0: "+# to Vitality"},
	4:  {0: "Stat Points: #"},
	5:  {0: "New Skills: #"},
	6:  {0: "Hit Points: #"},
	7:  {0: "+# to Life"},
	8:  {0: "Mana: #"},
	9:  {0: "+# to Mana"},
	10: {0: "Stamina: #"},
	11: {0: "+# Maximum Stamina"},
	12: {0: "Level: #"},
	13: {0: "Experience: #"},
	14: {0: "Gold: #"},
	15: {0: "Gold in Bank: #"},
	16: {0: "+#% Enhanced Defense"},
	17: {0: "+#% Enhanced Damage (Max)"},
	18: {0: "+#% Enhanced Damage (Min)"},
	19: {0: "+# to Attack Rating"},
	20: {0: "#% Increased Chance of Blocking"},
	21: {
		0: "#-# Damage",
		1: "+# to Minimum Damage",
	},
	22: {
		0: "#",
		1: "+# to Maximum Damage",
	},
	23: {0: "#"},
	24: {0: "#"},
	25: {0: "damagepercent ????"},
	26: {0: "manarecovery ????"},
	27: {0: "Regenerate Mana #%"},
	28: {0: "Heal Stamina Plus #%"},
	29: {0: "lastexp ????"},
	30: {0: "nextexp ????"},
	31: {0: "+# Defense"},
	32: {0: "# Defense vs. Missile"},
	33: {0: "# Defense vs. Melee"},
	34: {0: "Damage Reduced by #"},
	35: {0: "Magic Damage Reduced by #"},
	36: {0: "Damage Reduced by #%"},
	37: {0: "Magic Resist +#%"},
	38: {0: "+#% to Maximum Magic Resist"},
	39: {0: "Fire Resist +#%"},
	40: {0: "+#% to Maximum Fire Resist"},
	41: {0: "Lightning Resist +#%"},
	42: {0: "+#% to Maximum Lightning Resist"},
	43: {0: "Cold Resist +#%"},
	44: {0: "+#% to Maximum Cold Resist"},
	45: {0: "Poison Resist +#%"},
	46: {0: "+#% to Maximum Poison Resist"},
	47: {0: "damageaura ????"},
	48: {0: "+# to Minimum Fire Damage"},
	49: {0: "+# to Maximum Fire Damage"},
	50: {0: "Adds #-# Lightning Damage"},
	51: {0: "+# to Maximum Lightning Damage"},
	52: {0: "Adds #-# Magic Damage"},
	53: {0: "+# to Maximum Magic Damage"},
	54: {0: "Adds #-# Cold Damage"},
	55: {0: "+# to Maximum Cold Damage"},
	56: {0: "Half Freeze Duration"},
	57: {
		0: "+# to Poison Damage",
		1: "+# to Maximum Poison Damage",
	},
	58: {0: "+# to Maximum Poison Damage"},
	59: {0: "Poison Length Reduced by #%"},
	60: {0: "#% Life stolen per hit"},
	61: {0: "lifedrainmaxdam ????"},
	62: {0: "#% Mana stolen per hit"},
	63: {0: "manadrainmaxdam ????"},
	64: {0: "stamdrainmindam ????"},
	65: {0: "stamdrainmaxdam ????"},
	66: {0: "stunlength ????"},
	67: {0: "#% Velocity"},
	68: {0: "#"},
	69: {0: "otheranimrate ????"},
	70: {0: "Quantity: #"},
	71: {0: "Value: #"},
	72: {0: "Durability: # of #"},
	73: {0: "# Max Durability"},
	74: {0: "+# Replenish Life"},
	75: {0: "Increase Maximum Durability #%"},
	76: {0: "Increase Maximum Life #%"},
	77: {0: "Increase Maximum Mana #%"},
	78: {0: "Attacker Takes Damage of #"},
	79: {0: "#% Extra Gold from Monsters"},
	80: {0: "#% Better Chance of Getting Magic Items"},
	81: {0: "+# Knockback"},
	82: {0: "itemtimeduration ????"},
	83: {
		0: "+# to Amazon Skill Levels",
		1: "+# to Sorceress Skill Levels",
		2: "+# to Necromancer Skill Levels",
		3: "+# to Paladin Skill Levels",
		4: "+# to Barbarian Skill Levels",
		5: "+# to Druid Skill Levels",
		6: "+# to Assassin Skill Levels",
		7: "+# to Warlock Skill Levels",
	},
	84: {0: "unsentparam1 ????"},
	85: {0: "+#% to Experience Gained"},
	86: {0: "+# Life after each Kill"},
	87: {0: "Reduces all Vendor Prices #%"},
	88: {0: "itemdoubleherbduration ????"},
	89: {0: "+# to Light Radius"},
	90: {0: "itemlightcolor ????"},
	91: {0: "Requirements -#%"},
	92: {0: "Required Level: #"},
	93: {0: "+#% Increased Attack Speed"},
	94: {0: "itemlevelreqpct ????"},
	95: {0: "lastblockframe ????"},
	96: {0: "+#% Faster Run/Walk"},
	97: {
		0:   "+# to [Skill]",
		6:   "+# To Magic Arrow",
		7:   "+# To Fire Arrow",
		8:   "+# To Inner Sight",
		9:   "+# To Critical Strike",
		10:  "+# To Jab",
		11:  "+# To Cold Arrow",
		12:  "+# To Multiple Shot",
		13:  "+# To Dodge",
		14:  "+# To Power Strike",
		15:  "+# To Poison Javelin",
		16:  "+# To Exploding Arrow",
		17:  "+# To Slow Missiles",
		18:  "+# To Avoid",
		19:  "+# To Impale",
		20:  "+# To Lightning Bolt",
		21:  "+# To Ice Arrow",
		22:  "+# To Guided Arrow",
		23:  "+# To Penetrate",
		24:  "+# To Charged Strike",
		25:  "+# To Plague Javelin",
		26:  "+# To Strafe",
		27:  "+# To Immolation Arrow",
		28:  "+# To Decoy",
		29:  "+# To Evade",
		30:  "+# To Fend",
		31:  "+# To Freezing Arrow",
		32:  "+# To Valkyrie",
		33:  "+# To Pierce",
		34:  "+# To Lightning Strike",
		35:  "+# To Lightning Fury",
		36:  "+# To Fire Bolt",
		37:  "+# To Warmth",
		38:  "+# To Charged Bolt",
		39:  "+# To Ice Bolt",
		40:  "+# To Frozen Armor",
		41:  "+# To Inferno",
		42:  "+# To Static Field",
		43:  "+# To Telekinesis",
		44:  "+# To Frost Nova",
		45:  "+# To Ice Blast",
		46:  "+# To Blaze",
		47:  "+# To Fire Ball",
		48:  "+# To Nova",
		49:  "+# To Lightning",
		50:  "+# To Shiver Armor",
		51:  "+# To Fire Wall",
		52:  "+# To Enchant",
		53:  "+# To Chain Lightning",
		54:  "+# To Teleport",
		55:  "+# To Glacial Spike",
		56:  "+# To Meteor",
		57:  "+# To Thunder Storm",
		58:  "+# To Energy Shield",
		59:  "+# To Blizzard",
		60:  "+# To Chilling Armor",
		61:  "+# To Fire Mastery",
		62:  "+# To Hydra",
		63:  "+# To Lightning Mastery",
		64:  "+# To Frozen Orb",
		65:  "+# To Cold Mastery",
		66:  "+# To Amplify Damage",
		67:  "+# To Teeth",
		68:  "+# To Bone Armor",
		69:  "+# To Skeleton Mastery",
		70:  "+# To Raise Skeleton",
		71:  "+# To Dim Vision",
		72:  "+# To Weaken",
		73:  "+# To Poison Dagger",
		74:  "+# To Corpse Explosion",
		75:  "+# To Clay Golem",
		76:  "+# To Iron Maiden",
		77:  "+# To Terror",
		78:  "+# To Bone Wall",
		79:  "+# To Golem Mastery",
		80:  "+# To Raise Skeletal Mage",
		81:  "+# To Confuse",
		82:  "+# To Life Tap",
		83:  "+# To Poison Explosion",
		84:  "+# To Bone Spear",
		85:  "+# To Blood Golem",
		86:  "+# To Attract",
		87:  "+# To Decrepify",
		88:  "+# To Bone Prison",
		89:  "+# To Summon Resist",
		90:  "+# To Iron Golem",
		91:  "+# To Lower Resist",
		92:  "+# To Poison Nova",
		93:  "+# To Bone Spirit",
		94:  "+# To Fire Golem",
		95:  "+# To Revive",
		96:  "+# To Sacrifice",
		97:  "+# To Smite",
		98:  "+# To Might",
		99:  "+# To Prayer",
		100: "+# To Resist Fire",
		101: "+# To Holy Bolt",
		102: "+# To Holy Fire",
		103: "+# To Thorns",
		104: "+# To Defiance",
		105: "+# To Resist Cold",
		106: "+# To Zeal",
		107: "+# To Charge",
		108: "+# To Blessed Aim",
		109: "+# To Cleansing",
		110: "+# To Resist Lightning",
		111: "+# To Vengeance",
		112: "+# To Blessed Hammer",
		113: "+# To Concentration",
		114: "+# To Holy Freeze",
		115: "+# To Vigor",
		116: "+# To Conversion",
		117: "+# To Holy Shield",
		118: "+# To Holy Shock",
		119: "+# To Sanctuary",
		120: "+# To Meditation",
		121: "+# To Fist of the Heavens",
		122: "+# To Fanaticism",
		123: "+# To Conviction",
		124: "+# To Redemption",
		125: "+# To Salvation",
		126: "+# To Bash",
		127: "+# To Sword Mastery",
		128: "+# To Axe Mastery",
		129: "+# To Mace Mastery",
		130: "+# To Howl",
		131: "+# To Find Potion",
		132: "+# To Leap",
		133: "+# To Double Swing",
		134: "+# To Pole Arm Mastery",
		135: "+# To Throwing Mastery",
		136: "+# To Spear Mastery",
		137: "+# To Taunt",
		138: "+# To Shout",
		139: "+# To Stun",
		140: "+# To Double Throw",
		141: "+# To Increased Stamina",
		142: "+# To Find Item",
		143: "+# To Leap Attack",
		144: "+# To Concentrate",
		145: "+# To Iron Skin",
		146: "+# To Battle Cry",
		147: "+# To Frenzy",
		148: "+# To Increased Speed",
		149: "+# To Battle Orders",
		150: "+# To Grim Ward",
		151: "+# To Whirlwind",
		152: "+# To Berserk",
		153: "+# To Natural Resistance",
		154: "+# To War Cry",
		155: "+# To Battle Command",
		221: "+# To Raven",
		222: "+# To Werewolf",
		223: "+# To Werewolf",
		224: "+# To Lycanthropy",
		225: "+# To Firestorm",
		226: "+# To Oak Sage",
		227: "+# To Summon Spirit Wolf",
		228: "+# To Werebear",
		229: "+# To Molten Boulder",
		230: "+# To Arctic Blast",
		231: "+# To Carrion Vine",
		232: "+# To Feral Rage",
		233: "+# To Maul",
		234: "+# To Fissure",
		235: "+# To Cyclone Armor",
		236: "+# To Heart of Wolverine",
		237: "+# To Summon Dire Wolf",
		238: "+# To Rabies",
		239: "+# To Fire Claws",
		240: "+# To Twister",
		241: "+# To Solar Creeper",
		242: "+# To Hunger",
		243: "+# To Shock Wave",
		244: "+# To Volcano",
		245: "+# To Tornado",
		246: "+# To Spirit of Barbs",
		247: "+# To Summon Grizzly",
		248: "+# To Fury",
		249: "+# To Armageddon",
		250: "+# To Hurricane",
		251: "+# To Fire Blast",
		252: "+# To Claw Mastery",
		253: "+# To Psychic Hammer",
		254: "+# To Tiger Strike",
		255: "+# To Dragon Talon",
		256: "+# To Shock Web",
		257: "+# To Blade Sentinel",
		258: "+# To Burst of Speed",
		259: "+# To Fists of Fire",
		260: "+# To Dragon Claw",
		261: "+# To Charged Bolt Sentry",
		262: "+# To Wake of Fire",
		263: "+# To Weapon Block",
		264: "+# To Cloak of Shadows",
		265: "+# To Cobra Strike",
		266: "+# To Blade Fury",
		267: "+# To Fade",
		268: "+# To Shadow Warrior",
		269: "+# To Claws of Thunder",
		270: "+# To Dragon Tail",
		271: "+# To Lightning Sentry",
		272: "+# To Wake of Inferno",
		273: "+# To Mind Blast",
		274: "+# To Blades of Ice",
		275: "+# To Dragon Flight",
		276: "+# To Death Sentry",
		277: "+# To Blade Shield",
		278: "+# To Venom",
		279: "+# To Shadow Master",
		280: "+# To Phoenix Strike",
		373: "+# To Summon Goatman",
		374: "+# To Demonic Mastery",
		375: "+# To Death Mark",
		376: "+# To Summon Tainted",
		377: "+# To Summon Defiler",
		378: "+# To Blood Oath",
		379: "+# To Engorge",
		380: "+# To Blood Boil",
		381: "+# To Consume",
		382: "+# To Bind Demon",
		383: "+# To Levitate",
		384: "+# To Eldritch Blast",
		385: "+# To Hex Bane",
		386: "+# To Hex Siphon",
		387: "+# To Psychic Ward",
		388: "+# To Echoing Strike",
		389: "+# To Hex Purge",
		390: "+# To Blade Warp",
		391: "+# To Cleave",
		392: "+# To Mirrored Blades",
		393: "+# To Sigil Lethargy",
		394: "+# To Ring of Fire",
		395: "+# To Miasma Bolt",
		396: "+# To Sigil Rancor",
		397: "+# To Enhanced Entropy",
		398: "+# To Flame Wave",
		399: "+# To Miasma Chains",
		400: "+# To Sigil Death",
		401: "+# To Apocalypse",
		402: "+# To Abyss",
	},
	98:  {0: "state ????"},
	99:  {0: "+#% Faster Hit Recovery"},
	100: {0: "monsterplayercount ????"},
	101: {0: "skillpoisonoverridelength ????"},
	102: {0: "+#% Faster Block Rate"},
	103: {0: "skillbypassundead ????"},
	104: {0: "skillbypassdemons ????"},
	105: {0: "+#% Faster Cast Rate"},
	106: {0: "skillbypassbeasts ????"},
	107: {
		0:   "+# to [Skill] ([Class] only)",
		6:   "+# to Magic Arrow (Amazon only)",
		7:   "+# to Fire Arrow (Amazon only)",
		8:   "+# to Inner Sight (Amazon only)",
		9:   "+# to Critical Strike (Amazon only)",
		10:  "+# to Jab (Amazon only)",
		11:  "+# to Cold Arrow (Amazon only)",
		12:  "+# to Multiple Shot (Amazon only)",
		13:  "+# to Dodge (Amazon only)",
		14:  "+# to Power Strike (Amazon only)",
		15:  "+# to Poison Javelin (Amazon only)",
		16:  "+# to Exploding Arrow (Amazon only)",
		17:  "+# to Slow Missiles (Amazon only)",
		18:  "+# to Avoid (Amazon only)",
		19:  "+# to Impale (Amazon only)",
		20:  "+# to Lightning Bolt (Amazon only)",
		21:  "+# to Ice Arrow (Amazon only)",
		22:  "+# to Guided Arrow (Amazon only)",
		23:  "+# to Penetrate (Amazon only)",
		24:  "+# to Charged Strike (Amazon only)",
		25:  "+# to Plague Javelin (Amazon only)",
		26:  "+# to Strafe (Amazon only)",
		27:  "+# to Immolation Arrow (Amazon only)",
		28:  "+# to Decoy (Amazon only)",
		29:  "+# to Evade (Amazon only)",
		30:  "+# to Fend (Amazon only)",
		31:  "+# to Freezing Arrow (Amazon only)",
		32:  "+# to Valkyrie (Amazon only)",
		33:  "+# to Pierce (Amazon only)",
		34:  "+# to Lightning Strike (Amazon only)",
		35:  "+# to Lightning Fury (Amazon only)",
		36:  "+# to Fire Bolt (Sorceress only)",
		37:  "+# to Warmth (Sorceress only)",
		38:  "+# to Charged Bolt (Sorceress only)",
		39:  "+# to Ice Bolt (Sorceress only)",
		40:  "+# to Frozen Armor (Sorceress only)",
		41:  "+# to Inferno (Sorceress only)",
		42:  "+# to Static Field (Sorceress only)",
		43:  "+# to Telekinesis (Sorceress only)",
		44:  "+# to Frost Nova (Sorceress only)",
		45:  "+# to Ice Blast (Sorceress only)",
		46:  "+# to Blaze (Sorceress only)",
		47:  "+# to Fire Ball (Sorceress only)",
		48:  "+# to Nova (Sorceress only)",
		49:  "+# to Lightning (Sorceress only)",
		50:  "+# to Shiver Armor (Sorceress only)",
		51:  "+# to Fire Wall (Sorceress only)",
		52:  "+# to Enchant (Sorceress only)",
		53:  "+# to Chain Lightning (Sorceress only)",
		54:  "+# to Teleport (Sorceress only)",
		55:  "+# to Glacial Spike (Sorceress only)",
		56:  "+# to Meteor (Sorceress only)",
		57:  "+# to Thunder Storm (Sorceress only)",
		58:  "+# to Energy Shield (Sorceress only)",
		59:  "+# to Blizzard (Sorceress only)",
		60:  "+# to Chilling Armor (Sorceress only)",
		61:  "+# to Fire Mastery (Sorceress only)",
		62:  "+# to Hydra (Sorceress only)",
		63:  "+# to Lightning Mastery (Sorceress only)",
		64:  "+# to Frozen Orb (Sorceress only)",
		65:  "+# to Cold Mastery (Sorceress only)",
		66:  "+# to Amplify Damage (Necromancer only)",
		67:  "+# to Teeth (Necromancer only)",
		68:  "+# to Bone Armor (Necromancer only)",
		69:  "+# to Skeleton Mastery (Necromancer only)",
		70:  "+# to Raise Skeleton (Necromancer only)",
		71:  "+# to Dim Vision (Necromancer only)",
		72:  "+# to Weaken (Necromancer only)",
		73:  "+# to Poison Dagger (Necromancer only)",
		74:  "+# to Corpse Explosion (Necromancer only)",
		75:  "+# to Clay Golem (Necromancer only)",
		76:  "+# to Iron Maiden (Necromancer only)",
		77:  "+# to Terror (Necromancer only)",
		78:  "+# to Bone Wall (Necromancer only)",
		79:  "+# to Golem Mastery (Necromancer only)",
		80:  "+# to Raise Skeletal Mage (Necromancer only)",
		81:  "+# to Confuse (Necromancer only)",
		82:  "+# to Life Tap (Necromancer only)",
		83:  "+# to Poison Explosion (Necromancer only)",
		84:  "+# to Bone Spear (Necromancer only)",
		85:  "+# to Blood Golem (Necromancer only)",
		86:  "+# to Attract (Necromancer only)",
		87:  "+# to Decrepify (Necromancer only)",
		88:  "+# to Bone Prison (Necromancer only)",
		89:  "+# to Summon Resist (Necromancer only)",
		90:  "+# to Iron Golem (Necromancer only)",
		91:  "+# to Lower Resist (Necromancer only)",
		92:  "+# to Poison Nova (Necromancer only)",
		93:  "+# to Bone Spirit (Necromancer only)",
		94:  "+# to Fire Golem (Necromancer only)",
		95:  "+# to Revive (Necromancer only)",
		96:  "+# to Sacrifice (Paladin only)",
		97:  "+# to Smite (Paladin only)",
		98:  "+# to Might (Paladin only)",
		99:  "+# to Prayer (Paladin only)",
		100: "+# to Resist Fire (Paladin only)",
		101: "+# to Holy Bolt (Paladin only)",
		102: "+# to Holy Fire (Paladin only)",
		103: "+# to Thorns (Paladin only)",
		104: "+# to Defiance (Paladin only)",
		105: "+# to Resist Cold (Paladin only)",
		106: "+# to Zeal (Paladin only)",
		107: "+# to Charge (Paladin only)",
		108: "+# to Blessed Aim (Paladin only)",
		109: "+# to Cleansing (Paladin only)",
		110: "+# to Resist Lightning (Paladin only)",
		111: "+# to Vengeance (Paladin only)",
		112: "+# to Blessed Hammer (Paladin only)",
		113: "+# to Concentration (Paladin only)",
		114: "+# to Holy Freeze (Paladin only)",
		115: "+# to Vigor (Paladin only)",
		116: "+# to Conversion (Paladin only)",
		117: "+# to Holy Shield (Paladin only)",
		118: "+# to Holy Shock (Paladin only)",
		119: "+# to Sanctuary (Paladin only)",
		120: "+# to Meditation (Paladin only)",
		121: "+# to Fist of the Heavens (Paladin only)",
		122: "+# to Fanaticism (Paladin only)",
		123: "+# to Conviction (Paladin only)",
		124: "+# to Redemption (Paladin only)",
		125: "+# to Salvation (Paladin only)",
		126: "+# to Bash (Barbarian only)",
		127: "+# to Sword Mastery (Barbarian only)",
		128: "+# to Axe Mastery (Barbarian only)",
		129: "+# to Mace Mastery (Barbarian only)",
		130: "+# to Howl (Barbarian only)",
		131: "+# to Find Potion (Barbarian only)",
		132: "+# to Leap (Barbarian only)",
		133: "+# to Double Swing (Barbarian only)",
		134: "+# to Pole Arm Mastery (Barbarian only)",
		135: "+# to Throwing Mastery (Barbarian only)",
		136: "+# to Spear Mastery (Barbarian only)",
		137: "+# to Taunt (Barbarian only)",
		138: "+# to Shout (Barbarian only)",
		139: "+# to Stun (Barbarian only)",
		140: "+# to Double Throw (Barbarian only)",
		141: "+# to Increased Stamina (Barbarian only)",
		142: "+# to Find Item (Barbarian only)",
		143: "+# to Leap Attack (Barbarian only)",
		144: "+# to Concentrate (Barbarian only)",
		145: "+# to Iron Skin (Barbarian only)",
		146: "+# to Battle Cry (Barbarian only)",
		147: "+# to Frenzy (Barbarian only)",
		148: "+# to Increased Speed (Barbarian only)",
		149: "+# to Battle Orders (Barbarian only)",
		150: "+# to Grim Ward (Barbarian only)",
		151: "+# to Whirlwind (Barbarian only)",
		152: "+# to Berserk (Barbarian only)",
		153: "+# to Natural Resistance (Barbarian only)",
		154: "+# to War Cry (Barbarian only)",
		155: "+# to Battle Command (Barbarian only)",
		221: "+# to Raven (Druid only)",
		222: "+# to Werewolf (Druid only)",
		223: "+# to Werewolf (Druid only)",
		224: "+# to Lycanthropy (Druid only)",
		225: "+# to Firestorm (Druid only)",
		226: "+# to Oak Sage (Druid only)",
		227: "+# to Summon Spirit Wolf (Druid only)",
		228: "+# to Werebear (Druid only)",
		229: "+# to Molten Boulder (Druid only)",
		230: "+# to Arctic Blast (Druid only)",
		231: "+# to Carrion Vine (Druid only)",
		232: "+# to Feral Rage (Druid only)",
		233: "+# to Maul (Druid only)",
		234: "+# to Fissure (Druid only)",
		235: "+# to Cyclone Armor (Druid only)",
		236: "+# to Heart of Wolverine (Druid only)",
		237: "+# to Summon Dire Wolf (Druid only)",
		238: "+# to Rabies (Druid only)",
		239: "+# to Fire Claws (Druid only)",
		240: "+# to Twister (Druid only)",
		241: "+# to Solar Creeper (Druid only)",
		242: "+# to Hunger (Druid only)",
		243: "+# to Shock Wave (Druid only)",
		244: "+# to Volcano (Druid only)",
		245: "+# to Tornado (Druid only)",
		246: "+# to Spirit of Barbs (Druid only)",
		247: "+# to Summon Grizzly (Druid only)",
		248: "+# to Fury (Druid only)",
		249: "+# to Armageddon (Druid only)",
		250: "+# to Hurricane (Druid only)",
		251: "+# to Fire Blast (Assassin only)",
		252: "+# to Claw Mastery (Assassin only)",
		253: "+# to Psychic Hammer (Assassin only)",
		254: "+# to Tiger Strike (Assassin only)",
		255: "+# to Dragon Talon (Assassin only)",
		256: "+# to Shock Web (Assassin only)",
		257: "+# to Blade Sentinel (Assassin only)",
		258: "+# to Burst of Speed (Assassin only)",
		259: "+# to Fists of Fire (Assassin only)",
		260: "+# to Dragon Claw (Assassin only)",
		261: "+# to Charged Bolt Sentry (Assassin only)",
		262: "+# to Wake of Fire (Assassin only)",
		263: "+# to Weapon Block (Assassin only)",
		264: "+# to Cloak of Shadows (Assassin only)",
		265: "+# to Cobra Strike (Assassin only)",
		266: "+# to Blade Fury (Assassin only)",
		267: "+# to Fade (Assassin only)",
		268: "+# to Shadow Warrior (Assassin only)",
		269: "+# to Claws of Thunder (Assassin only)",
		270: "+# to Dragon Tail (Assassin only)",
		271: "+# to Lightning Sentry (Assassin only)",
		272: "+# to Wake of Inferno (Assassin only)",
		273: "+# to Mind Blast (Assassin only)",
		274: "+# to Blades of Ice (Assassin only)",
		275: "+# to Dragon Flight (Assassin only)",
		276: "+# to Death Sentry (Assassin only)",
		277: "+# to Blade Shield (Assassin only)",
		278: "+# to Venom (Assassin only)",
		279: "+# to Shadow Master (Assassin only)",
		280: "+# to Phoenix Strike (Assassin only)",
		373: "+# to Summon Goatman (Warlock only)",
		374: "+# to Demonic Mastery (Warlock only)",
		375: "+# to Death Mark (Warlock only)",
		376: "+# to Summon Tainted (Warlock only)",
		377: "+# to Summon Defiler (Warlock only)",
		378: "+# to Blood Oath (Warlock only)",
		379: "+# to Engorge (Warlock only)",
		380: "+# to Blood Boil (Warlock only)",
		381: "+# to Consume (Warlock only)",
		382: "+# to Bind Demon (Warlock only)",
		383: "+# to Levitate (Warlock only)",
		384: "+# to Eldritch Blast (Warlock only)",
		385: "+# to Hex Bane (Warlock only)",
		386: "+# to Hex Siphon (Warlock only)",
		387: "+# to Psychic Ward (Warlock only)",
		388: "+# to Echoing Strike (Warlock only)",
		389: "+# to Hex Purge (Warlock only)",
		390: "+# to Blade Warp (Warlock only)",
		391: "+# to Cleave (Warlock only)",
		392: "+# to Mirrored Blades (Warlock only)",
		393: "+# to Sigil Lethargy (Warlock only)",
		394: "+# to Ring of Fire (Warlock only)",
		395: "+# to Miasma Bolt (Warlock only)",
		396: "+# to Sigil Rancor (Warlock only)",
		397: "+# to Enhanced Entropy (Warlock only)",
		398: "+# to Flame Wave (Warlock only)",
		399: "+# to Miasma Chains (Warlock only)",
		400: "+# to Sigil Death (Warlock only)",
		401: "+# to Apocalypse (Warlock only)",
		402: "+# to Abyss (Warlock only)",
	},
	108: {0: "+# Slain Monsters Rest in Peace"},
	109: {0: "curseresistance ????"},
	110: {0: "Poison Length Reduced by #%"},
	111: {0: "Damage +#"},
	112: {0: "Hit Causes Monster to Flee #%"},
	113: {0: "Hit Blinds Target +#"},
	114: {0: "#% Damage Taken Goes To Mana"},
	115: {0: "Ignore Target's Defense"},
	116: {0: "-#% Target Defense"},
	117: {0: "Prevent Monster Heal"},
	118: {0: "Half Freeze Duration"},
	119: {0: " #% Bonus to Attack Rating"},
	120: {0: "-# to Monster Defense Per Hit"},
	121: {0: "+#% Damage to Demons"},
	122: {0: "+#% Damage to Undead"},
	123: {0: "+# to Attack Rating against Demons"},
	124: {0: "+# to Attack Rating against Undead"},
	125: {0: "itemthrowable ????"},
	126: {0: "+# to Fire Skills"},
	127: {0: "+# to All Skills"},
	128: {0: "Attacker Takes Lightning Damage of #"},
	129: {0: "ironmaidenlevel ????"},
	130: {0: "lifetaplevel ????"},
	131: {0: "thornspercent ????"},
	132: {0: "bonearmor ????"},
	133: {0: "bonearmormax ????"},
	134: {0: "Freezes Target +#"},
	135: {0: "#% Chance of Open Wounds"},
	136: {0: "#% Chance of Crushing Blow"},
	137: {0: "+# Kick Damage"},
	138: {0: "+# to Mana after each Kill"},
	139: {0: "+# Life after each Demon Kill"},
	140: {0: "itemextrablood"},
	141: {0: "#% Deadly Strike"},
	142: {0: "+# Fire Absorb"},
	143: {0: "Fire Absorb #%"},
	144: {0: "+# Lightning Absorb"},
	145: {0: "Lightning Absorb #%"},
	146: {0: "Magic Absorb #%"},
	147: {0: "+# Magic Absorb"},
	148: {0: "Cold Absorb #%"},
	149: {0: "+# Cold Absorb"},
	150: {0: "Slows Target by #%"},
	151: {
		0:   "Level # [Skill] Aura When Equipped",
		98:  "Level # Might Aura When Equipped",
		102: "Level # Holy Fire Aura When Equipped",
		103: "Level # Thorns Aura When Equipped",
		104: "Level # Defiance Aura When Equipped",
		113: "Level # Concentration Aura When Equipped",
		114: "Level # Holy Freeze Aura When Equipped",
		115: "Level # Vigor Aura When Equipped",
		118: "Level # Holy Shock Aura When Equipped",
		119: "Level # Sanctuary Aura When Equipped",
		120: "Level # Meditation Aura When Equipped",
		122: "Level # Fanaticism Aura When Equipped",
		123: "Level # Conviction Aura When Equipped",
		124: "Level # Redemption Aura When Equipped",
	},
	152: {0: "Indestructible"},
	153: {0: "Cannot be Frozen"},
	154: {0: "#% Slower Stamina Drain"},
	155: {0: "Reanimate As: [Returned]"},
	156: {0: "Piercing Attack"},
	157: {0: "Fires Magic Arrows"},
	158: {0: "Fires Explosive Arrows or Bolts"},
	159: {0: "# To Minimum Damage"},
	160: {0: "+# Maximum Damage"},
	161: {0: "itemskillhandofathena ????"},
	162: {0: "itemskillstaminapercent ????"},
	163: {0: "itemskillpassivestaminapercent ????"},
	164: {0: "itemskillconcentration ????"},
	165: {0: "itemskillenchant ????"},
	166: {0: "itemskillpierce ????"},
	167: {0: "itemskillconviction ????"},
	168: {0: "itemskillchillingarmor ????"},
	169: {0: "itemskillfrenzy ????"},
	170: {0: "itemskilldecrepify ????"},
	171: {0: "itemskillarmorpercent ????"},
	172: {0: "alignment ????"},
	173: {0: "target0 ????"},
	174: {0: "target1 ????"},
	175: {0: "goldlost ????"},
	176: {0: "conversionlevel ????"},
	177: {0: "conversionmaxhp ????"},
	178: {0: "unitdooverlay ????"},
	179: {0: "attackvsmontype ????"},
	180: {0: "damagevsmontype ????"},
	181: {0: "fade ????"},
	182: {0: "armoroverridepercent ????"},
	183: {0: "unused183 ????"},
	184: {0: "unused184 ????"},
	185: {0: "unused185 ????"},
	186: {0: "unused186 ????"},
	187: {0: "Monster Cold Immunity is Sundered"},
	188: {
		0:  "+# to Bow and Crossbow Skills (Amazon only)",
		1:  "+# to Passive and Magic Skills (Amazon only)",
		2:  "+# to Javelin and Spears Skills (Amazon only)",
		8:  "+# to Fire Skills (Sorceress only)",
		9:  "+# to Lightning Skills (Sorceress only)",
		10: "+# to Cold Skills (Sorceress only)",
		16: "+# to Curses Skills (Necromancer only)",
		17: "+# to Poison and Bone Skills (Necromancer only)",
		18: "+# to Summoning Skills (Necromancer only)",
		24: "+# to Paladin Combat Skills (Paladin only)",
		25: "+# to Offensive Aura Skills (Paladin only)",
		26: "+# to Defensive Aura Skills (Paladin only)",
		32: "+# to Barbarian Combat Skills (Barbarian only)",
		33: "+# to Mastery Skills (Barbarian only)",
		34: "+# to War Cry Skills (Barbarian only)",
		40: "+# to Druid Summoning Skills (Druid only)",
		41: "+# to Shapeshifting Skills (Druid only)",
		42: "+# to Elemental Skills (Druid only)",
		48: "+# to Traps Skills (Assassin only)",
		49: "+# to Shadow Discipline Skills (Assassin only)",
		50: "+# to Martial Arts Skills (Assassin only)",
		56: "+# to Demon Skills (Warlock only)",
		57: "+# to Eldritch Skills (Warlock only)",
		58: "+# to Chaos Skills (Warlock only)",
	},
	189: {0: "Monster Fire Immunity is Sundered"},
	190: {0: "Monster Lightning Immunity is Sundered"},
	191: {0: "Monster Poison Immunity is Sundered"},
	192: {0: "Monster Physical Immunity is Sundered"},
	193: {0: "Monster Magic Immunity is Sundered"},
	194: {0: "Socketed (#)"},
	195: {
		1:    "#% Chance to cast level # [Skill] on attack",
		2:    "itemskillonattacklevel ????",
		3395: "#% Chance to cast level 3 Chain Lightning on attack",
	},
	196: {
		1: "#% Chance to cast level # [Skill] when you Kill an Enemy",
		2: "itemskillonkilllevel ????",
	},
	197: {
		1: "#% Chance to cast level # [Skill] when you Die",
		2: "itemskillondeathlevel ????",
	},
	198: {
		1:    "#% Chance to cast level # [Skill] on striking",
		2:    "itemskillonhitlevel ????",
		4225: "Amplify Damage on Hit",
	},
	199: {
		1: "#% Chance to cast level # [Skill] when you Level-Up",
		2: "itemskillonleveluplevel ????",
	},
	200: {0: "unused200 ????"},
	201: {
		1:    "#% Chance to cast level # [Skill] when struck",
		2:    "itemskillongethitlevel ????",
		5903: "#% Chance to cast level 15 Poison Nova when struck",
		7751: "#% Chance to cast level 7 Fist of Heavens when struck",
	},
	202: {0: "unused202 ????"},
	203: {0: "unused203 ????"},
	204: {
		1:     "itemchargedskill ????",
		2:     "itemchargedskilllevel ????",
		3461:  "Teleport (charged)",
		17795: "Venom level 3 (charged)",
	},
	205: {0: "unused204 ????"},
	206: {0: "unused205 ????"},
	207: {0: "unused206 ????"},
	208: {0: "unused207 ????"},
	209: {0: "unused208 ????"},
	210: {0: "unused209 ????"},
	211: {0: "unused210 ????"},
	212: {0: "unused211 ????"},
	213: {0: "unused212 ????"},
	214: {0: "+# Defense (Based on Character Level)"},
	215: {0: "+#% Enhanced Defense (Based on Character Level)"},
	216: {0: "+# to Life (Based on Character Level)"},
	217: {0: "+# to Mana (Based on Character Level)"},
	218: {0: "+# to Maximum Damage (Based on Character Level)"},
	219: {0: "+#% Enhanced Maximum Damage (Based on Character Level)"},
	220: {0: "+# to Strength (Based on Character Level)"},
	221: {0: "+# to Dexterity (Based on Character Level)"},
	222: {0: "+# to Energy (Based on Character Level)"},
	223: {0: "+# to Vitality (Based on Character Level)"},
	224: {0: "+# to Attack Rating (Based on Character Level)"},
	225: {0: "#% Bonus to Attack Rating (Based on Character Level)"},
	226: {0: "+# to Maximum Cold Damage (Based on Character Level)"},
	227: {0: "+# to Maximum Fire Damage (Based on Character Level)"},
	228: {0: "+# to Maximum Lightning Damage (Based on Character Level)"},
	229: {0: "+# to Maximum Poison Damage (Based on Character Level)"},
	230: {0: "Cold Resist +#% (Based on Character Level)"},
	231: {0: "Fire Resist +#% (Based on Character Level)"},
	232: {0: "Lightning Resist +#% (Based on Character Level)"},
	233: {0: "Poison Resist +#% (Based on Character Level)"},
	234: {0: "Absorbs Cold Damage (Based on Character Level)"},
	235: {0: "Absorbs Fire Damage (Based on Character Level)"},
	236: {0: "Absorbs Lightning Damage (Based on Character Level)"},
	237: {0: "Absorbs Poison Damage (Based on Character Level)"},
	238: {0: "Attacker Takes Damage of # (Based on Character Level)"},
	239: {0: "#% Extra Gold from Monsters (Based on Character Level)"},
	240: {0: "#% Better Chance of Getting Magic Items (Based on Character Level)"},
	241: {0: "Heal Stamina Plus #% (Based on Character Level)"},
	242: {0: "+# Maximum Stamina (Based on Character Level)"},
	243: {0: "+#% Damage to Demons (Based on Character Level)"},
	244: {0: "+#% Damage to Undead (Based on Character Level)"},
	245: {0: "+# to Attack Rating against Demons (Based on Character Level)"},
	246: {0: "+# to Attack Rating against Undead (Based on Character Level)"},
	247: {0: "#% Chance of Crushing Blow (Based on Character Level)"},
	248: {0: "#% Chance of Open Wounds (Based on Character Level)"},
	249: {0: "+# Kick Damage (Based on Character Level)"},
	250: {0: "#% Deadly Strike (Based on Character Level)"},
	251: {0: "itemfindgemsperlevel ????"},
	252: {0: "Repairs 1 durability in # seconds"},
	253: {0: "Replenishes quantity"},
	254: {0: "Increased Stack Size"},
	255: {0: "itemfinditem ????"},
	256: {0: "itemslashdamage ????"},
	257: {0: "itemslashdamagepercent ????"},
	258: {0: "itemcrushdamage ????"},
	259: {0: "itemcrushdamagepercent ????"},
	260: {0: "itemthrustdamage ????"},
	261: {0: "itemthrustdamagepercent ????"},
	262: {0: "itemabsorbslash ????"},
	263: {0: "itemabsorbcrush ????"},
	264: {0: "itemabsorbthrust ????"},
	265: {0: "itemabsorbslashpercent ????"},
	266: {0: "itemabsorbcrushpercent ????"},
	267: {0: "itemabsorbthrustpercent ????"},
	268: {0: "+# Defense (Increases near [Day/Dusk/Night/Dawn])"},
	269: {0: "+#% Enhanced Defense (Increases near [Day/Dusk/Night/Dawn])"},
	270: {0: "+# to Life (Increases near [Day/Dusk/Night/Dawn])"},
	271: {0: "+# to Mana (Increases near [Day/Dusk/Night/Dawn])"},
	272: {0: "+# to Maximum Damage (Increases near [Day/Dusk/Night/Dawn])"},
	273: {0: "+#% Enhanced Maximum Damage (Increases near [Day/Dusk/Night/Dawn])"},
	274: {0: "+# to Strength (Increases near [Day/Dusk/Night/Dawn])"},
	275: {0: "+# to Dexterity (Increases near [Day/Dusk/Night/Dawn])"},
	276: {0: "+# to Energy (Increases near [Day/Dusk/Night/Dawn])"},
	277: {0: "+# to Vitality (Increases near [Day/Dusk/Night/Dawn])"},
	278: {0: "+# to Attack Rating (Increases near [Day/Dusk/Night/Dawn])"},
	279: {0: "+#% Bonus to Attack Rating (Increases near [Day/Dusk/Night/Dawn])"},
	280: {0: "+# to Maximum Cold Damage (Increases near [Day/Dusk/Night/Dawn])"},
	281: {0: "+# to Maximum Fire Damage (Increases near [Day/Dusk/Night/Dawn])"},
	282: {0: "+# to Maximum Lightning Damage (Increases near [Day/Dusk/Night/Dawn])"},
	283: {0: "+# to Maximum Poison Damage (Increases near [Day/Dusk/Night/Dawn])"},
	284: {0: "Cold Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},
	285: {0: "Fire Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},
	286: {0: "Lightning Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},
	287: {0: "Poison Resist +#% (Increases near [Day/Dusk/Night/Dawn])"},
	288: {0: "Absorbs Cold Damage (Increases near [Day/Dusk/Night/Dawn])"},
	289: {0: "Absorbs Fire Damage (Increases near [Day/Dusk/Night/Dawn])"},
	290: {0: "Absorbs Lightning Damage (Increases near [Day/Dusk/Night/Dawn])"},
	291: {0: "Absorbs Poison Damage (Increases near [Day/Dusk/Night/Dawn])"},
	292: {0: "#% Extra Gold from Monsters (Increases near [Day/Dusk/Night/Dawn])"},
	293: {0: "#% Better Chance of Getting Magic Items (Increases near [Day/Dusk/Night/Dawn])"},
	294: {0: "Heal Stamina Plus #% (Increases near [Day/Dusk/Night/Dawn])"},
	295: {0: "+# Maximum Stamina (Increases near [Day/Dusk/Night/Dawn])"},
	296: {0: "+#% Damage to Demons (Increases near [Day/Dusk/Night/Dawn])"},
	297: {0: "+#% Damage to Undead (Increases near [Day/Dusk/Night/Dawn])"},
	298: {0: "+# to Attack Rating against Demons (Increases near [Day/Dusk/Night/Dawn])"},
	299: {0: "+# to Attack Rating against Undead (Increases near [Day/Dusk/Night/Dawn])"},
	300: {0: "#% Chance of Crushing Blow (Increases near [Day/Dusk/Night/Dawn])"},
	301: {0: "#% Chance of Open Wounds (Increases near [Day/Dusk/Night/Dawn])"},
	302: {0: "+# Kick Damage (Increases near [Day/Dusk/Night/Dawn])"},
	303: {0: "#% Deadly Strike (Increases near [Day/Dusk/Night/Dawn])"},
	304: {0: "itemfindgemsbytime ????"},
	305: {0: "-#% to Enemy Cold Resistance"},
	306: {0: "-#% to Enemy Fire Resistance"},
	307: {0: "-#% to Enemy Lightning Resistance"},
	308: {0: "-#% to Enemy Poison Resistance"},
	309: {0: "itemdamagevsmonster ????"},
	310: {0: "itemdamagepercentvsmonster ????"},
	311: {0: "itemtohitvsmonster ????"},
	312: {0: "itemtohitpercentvsmonster ????"},
	313: {0: "itemacvsmonster ????"},
	314: {0: "itemacpercentvsmonster ????"},
	315: {0: "firelength ????"},
	316: {0: "burningmin ????"},
	317: {0: "burningmax ????"},
	318: {0: "progressivedamage ????"},
	319: {0: "progressivesteal ????"},
	320: {0: "progressiveother ????"},
	321: {0: "progressivefire ????"},
	322: {0: "progressivecold ????"},
	323: {0: "progressivelightning ????"},
	324: {0: "itemextracharges ????"},
	325: {0: "progressivetohit ????"},
	326: {0: "poisoncount ????"},
	327: {0: "damageframerate ????"},
	328: {0: "pierceidx ????"},
	329: {0: "+#% to Fire Skill Damage"},
	330: {0: "+#% to Lightning Skill Damage"},
	331: {0: "+#% to Cold Skill Damage"},
	332: {0: "+#% to Poison Skill Damage"},
	333: {0: "-#% to Enemy Fire Resistance"},
	334: {0: "-#% to Enemy Lightning Resistance"},
	335: {0: "-#% to Enemy Cold Resistance"},
	336: {0: "-#% to Enemy Poison Resistance"},
	337: {0: "passivecriticalstrike ????"},
	338: {0: "passivedodge ????"},
	339: {0: "passiveavoid ????"},
	340: {0: "passiveevade ????"},
	341: {0: "passivewarmth ????"},
	342: {0: "passivemasterymeleeth ????"},
	343: {0: "passivemasterymeleedmg ????"},
	344: {0: "passivemasterymeleecrit ????"},
	345: {0: "passivemasterythrowth ????"},
	346: {0: "passivemasterythrowdmg ????"},
	347: {0: "passivemasterythrowcrit ????"},
	348: {0: "passiveweaponblock ????"},
	349: {0: "passivesummonresist ????"},
	350: {0: "modifierlistskill ????"},
	351: {0: "modifierlistlevel ????"},
	352: {0: "lastsenthppct ????"},
	353: {0: "sourceunittype ????"},
	354: {0: "sourceunitid ????"},
	355: {0: "shortparam1 ????"},
	356: {0: "questitemdifficulty ????"},
	357: {0: "passivemagmastery ????"},
	358: {0: "passivemagpierce ????"},
	359: {0: "skillcooldown ????"},
	360: {0: "skillmissiledamagescale ????"},
	555: {0: "All Resistances +#"},
}
