package nip

import (
	"testing"

	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
	"github.com/stretchr/testify/require"
)

func TestRule_Evaluate(t *testing.T) {
	type fields struct {
		RawLine    string
		Filename   string
		LineNumber int
		Enabled    bool
	}
	type args struct {
		item data.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    RuleResult
		wantErr bool
	}{
		{
			name: "Basic rule with posion dmg, ethereal is not specified as a condition so it should be ignored",
			fields: fields{
				RawLine:    "[name] == smallcharm && [quality] == magic  # (([poisonlength]*25)*[poisonmaxdam])/256 >= 123",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					ID:         618,
					Name:       "SmAlLCharM",
					Quality:    item.QualityMagic,
					Identified: true,
					Ethereal:   true,
					Stats: []stat.Data{
						{ID: stat.PoisonLength, Value: 20},
						{ID: stat.PoisonMaxDamage, Value: 100},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Complex rule with flags",
			fields: fields{
				RawLine:    "[type] == armor && [quality] <= superior && [flag] != ethereal # ([itemmaxdurabilitypercent] == 0 || [itemmaxdurabilitypercent] == 15) && ([sockets] == 0 || [sockets] == 3 || [sockets] == 4)",
				Filename:   "test.nip",
				LineNumber: 1,
				Enabled:    true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "mageplate",
					Quality:    item.QualitySuperior,
					Identified: true,
					Ethereal:   false,
					Stats: []stat.Data{
						{ID: stat.MaxDurabilityPercent, Value: 15},
						{ID: stat.NumSockets, Value: 4},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Armor with +3 Sorc skills",
			fields: fields{
				RawLine: "[type] == armor # [sorceressskills] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "mageplate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.AddClassSkills, Value: 3, Layer: 1},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Armor with +3 Glacial Spike",
			fields: fields{
				RawLine: "[type] == armor  # [skillglacialspike] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "mageplate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SingleSkill, Value: 3, Layer: 55},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Unid item matching base stats should return partial match",
			fields: fields{
				RawLine: "[type] == armor && [quality] == magic # [defense] == 200",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Identified: false,
					Name:       "mageplate",
					Quality:    item.QualityMagic,
				},
			},
			want: RuleResultPartial,
		},
		{
			name: "Basic rule without stats or maxquantity",
			fields: fields{
				RawLine: "[type] == assassinclaw && [class] == exceptional && [quality] == magic # #",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         187,
					Identified: false,
					Name:       "GreaterTalons",
					Quality:    item.QualityMagic,
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Basic rule for a white superior item with enhanceddefense",
			fields: fields{
				RawLine: "[type] == armor && [quality] == superior # [enhanceddefense] >= 15 #",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         373,
					Name:       "mageplate",
					Quality:    item.QualitySuperior,
					Stats: []stat.Data{
						{ID: stat.EnhancedDefense, Value: 15},
						{ID: stat.Defense, Value: 301},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Basic rule for a white superior item with enhanceddamage",
			fields: fields{
				RawLine: "[type] == sword # [enhanceddamage] >= 15 #",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         234,
					Name:       "colossusblade",
					Quality:    item.QualitySuperior,
					Stats: []stat.Data{
						{ID: stat.EnhancedDamage, Value: 15},
						{ID: stat.MinDamage, Value: 28},
						{ID: stat.MaxDamage, Value: 74},
						{ID: stat.TwoHandedMinDamage, Value: 66},
						{ID: stat.TwoHandedMaxDamage, Value: 132},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Magic Giant Sword with +3 Warcries and 6% mana leech",
			fields: fields{
				RawLine: "[name] == giantsword && [quality] == magic # [warcriesskilltab] >= 3 && ([coldmindam] >= 4 || [firemindam] >= 9 || [lightmaxdam] >= 40 || [poisonmindam] >= 10 || [poisondamage] >= 10 || [maxdamage] >= 10 || [mindamage] >= 5 || [strength] >= 15 || [hpregen] >= 5 || [maxhp] >= 30 || [itemknockback] >= 1 || [lifeleech] >= 6 || [manaleech] >= 6 || [itemreqpercent] == -30 || [itemskillonattack] >= 1 || [itemchargedskill] >= 1)",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         35,
					Name:       "GiantSword",
					Quality:    item.QualityMagic,
					Stats: []stat.Data{
						{ID: stat.AddSkillTab, Value: 3, Layer: 34},
						{ID: stat.ManaSteal, Value: 6},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name:    "Ensure [itemskillonstriking] returns error, not supported yet",
			fields:  fields{RawLine: "[quality] == magic # [itemskillonstriking] >= 1", Enabled: true},
			args:    args{item: data.Item{Identified: true, ID: 516, Name: "healingpotion", Quality: item.QualityMagic}},
			wantErr: true,
		},
		{
			name: "Magic Flail with itemskillonhit",
			fields: fields{
				RawLine: "[name] == flail && [quality] == magic && [flag] != ethereal # [itemskillonhit] == 48",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         21,
					Name:       "Flail",
					Quality:    item.QualityMagic,
					Ethereal:   false,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 10, Layer: 3075},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Magic Giant Sword with itemskillonhit",
			fields: fields{
				RawLine: "[name] == giantsword && [quality] == magic # [itemskillonhit] >= 1",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         35,
					Name:       "GiantSword",
					Quality:    item.QualityMagic,
					Ethereal:   false,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 10, Layer: 3075},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "Magic Flail with itemskillonhitlevel",
			fields: fields{
				RawLine: "[name] == flail && [quality] == magic && [flag] != ethereal # [itemskillonhitlevel] >= 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					Identified: true,
					ID:         21,
					Name:       "Flail",
					Quality:    item.QualityMagic,
					Ethereal:   false,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 10, Layer: 3075},
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnAttack: Fire Ball (47) Level 5 (match)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonattack] == 47 && [itemskillonattacklevel] == 5",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnAttack, Value: 10, Layer: 3013}, // (47 << 6) | 5
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnAttack: Fire Ball (47) Level 5 (match distinct level condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonattacklevel] >= 4",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnAttack, Value: 10, Layer: 3013}, // (47 << 6) | 5
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnAttack: Fire Ball (47) Level 5 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonattack] == 47 || [itemskillonattack] == 50",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnAttack, Value: 10, Layer: 3013}, // (47 << 6) | 5
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnAttack: Fire Ball (47) Level 5 (no match)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonattack] == 48",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnAttack, Value: 10, Layer: 3013},
					},
				},
			},
			want: RuleResultNoMatch,
		},
		{
			name: "SkillOnKill: Chain Lightning (53) Level 12 (match)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonkill] == 53 && [itemskillonkilllevel] == 12",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnKill, Value: 10, Layer: 3404}, // (53 << 6) | 12
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnKill: Chain Lightning (53) Level 12 (match distinct level condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonkilllevel] >= 10",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnKill, Value: 10, Layer: 3404}, // (53 << 6) | 12
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnKill: Chain Lightning (53) Level 12 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonkill] == 53 || [itemskillonkill] == 60",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnKill, Value: 10, Layer: 3404}, // (53 << 6) | 12
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnKill: Chain Lightning (53) Level 12 (no match)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonkill] == 54",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnKill, Value: 10, Layer: 3404}, // (53 << 6) | 12
					},
				},
			},
			want: RuleResultNoMatch,
		},
		{
			name: "SkillOnDeath: Shiver Armor (50) Level 3 (match)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillondeath] == 50 && [itemskillondeathlevel] == 3",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnDeath, Value: 100, Layer: 3203}, // (50 << 6) | 3
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnDeath: Shiver Armor (50) Level 3 (match distinct level condition)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillondeathlevel] >= 2",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnDeath, Value: 100, Layer: 3203}, // (50 << 6) | 3
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnDeath: Shiver Armor (50) Level 3 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillondeath] == 50 || [itemskillondeath] == 55",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnDeath, Value: 100, Layer: 3203}, // (50 << 6) | 3
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnDeath: Shiver Armor (50) Level 3 (no match)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillondeath] == 51",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnDeath, Value: 100, Layer: 3203}, // (50 << 6) | 3
					},
				},
			},
			want: RuleResultNoMatch,
		},
		{
			name: "SkillOnHit: Amplify Damage (66) Level 1 (match)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonhit] == 66 && [itemskillonhitlevel] == 1",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 5, Layer: 4225}, // (66 << 6) | 1
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnHit: Amplify Damage (66) Level 1 (match distinct level condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonhitlevel] >= 1",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 5, Layer: 4225}, // (66 << 6) | 1
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnHit: Amplify Damage (66) Level 1 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonhit] == 66 || [itemskillonhit] == 70",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 5, Layer: 4225}, // (66 << 6) | 1
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnHit: Amplify Damage (66) Level 1 (no match)",
			fields: fields{
				RawLine: "[name] == flail # [itemskillonhit] == 67",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnHit, Value: 5, Layer: 4225}, // (66 << 6) | 1
					},
				},
			},
			want: RuleResultNoMatch,
		},
		{
			name: "SkillOnLevelUp: Blizzard (59) Level 7 (match)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillonlevelup] == 59 && [itemskillonleveluplevel] >= 5",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnLevelUp, Value: 10, Layer: 3783}, // (59 << 6) | 7
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnLevelUp: Blizzard (59) Level 7 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillonlevelup] == 59 || [itemskillonlevelup] == 60",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnLevelUp, Value: 10, Layer: 3783}, // (59 << 6) | 7
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnLevelUp: Blizzard (59) Level 7 (no match)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillonlevelup] == 60",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnLevelUp, Value: 10, Layer: 3783}, // (59 << 6) | 7
					},
				},
			},
			want: RuleResultNoMatch,
		},
		{
			name: "SkillOnGetHit: Nova (48) Level 9 (match)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillongethit] == 48 && [itemskillongethitlevel] == 9",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnGetHit, Value: 10, Layer: 3081}, // (48 << 6) | 9
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnGetHit: Nova (48) Level 9 (match distinct level condition)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillongethitlevel] >= 5",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnGetHit, Value: 10, Layer: 3081}, // (48 << 6) | 9
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnGetHit: Nova (48) Level 9 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillongethit] == 48 || [itemskillongethit] == 50",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnGetHit, Value: 10, Layer: 3081}, // (48 << 6) | 9
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "SkillOnGetHit: Nova (48) Level 9 (no match)",
			fields: fields{
				RawLine: "[name] == mageplate # [itemskillongethit] == 49",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         373,
					Name:       "MagePlate",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.SkillOnGetHit, Value: 10, Layer: 3081}, // (48 << 6) | 9
					},
				},
			},
			want: RuleResultNoMatch,
		},
		{
			name: "ChargedSkill: Teleport (54) Level 1 (match)",
			fields: fields{
				RawLine: "[name] == flail # [itemchargedskill] == 54 && [itemchargedskilllevel] == 1",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.ItemChargedSkill, Value: 20, Layer: 3457}, // (54 << 6) | 1
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "ChargedSkill: Teleport (54) Level 1 (match distinct level condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemchargedskilllevel] >= 1",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.ItemChargedSkill, Value: 20, Layer: 3457}, // (54 << 6) | 1
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "ChargedSkill: Teleport (54) Level 1 (match distinct OR condition)",
			fields: fields{
				RawLine: "[name] == flail # [itemchargedskill] == 54 || [itemchargedskill] == 60",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.ItemChargedSkill, Value: 20, Layer: 3457}, // (54 << 6) | 1
					},
				},
			},
			want: RuleResultFullMatch,
		},
		{
			name: "ChargedSkill: Teleport (54) Level 1 (no match - different skill)",
			fields: fields{
				RawLine: "[name] == flail # [itemchargedskill] == 55",
				Enabled: true,
			},
			args: args{
				item: data.Item{
					ID:         21,
					Name:       "Flail",
					Identified: true,
					Stats: []stat.Data{
						{ID: stat.ItemChargedSkill, Value: 20, Layer: 3457}, // (54 << 6) | 1
					},
				},
			},
			want: RuleResultNoMatch,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewRule(tt.fields.RawLine, tt.fields.Filename, tt.fields.LineNumber)
			require.NoError(t, err)
			got, err := r.Evaluate(tt.args.item)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		rawRule    string
		filename   string
		lineNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    Rule
		wantErr bool
	}{
		{
			name: "Ensure [color] returns error, not supported yet",
			args: args{
				rawRule: "[type] == armor && [color] == 1000 && [quality] == magic",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRule(tt.args.rawRule, tt.args.filename, tt.args.lineNumber)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
			}
		})
	}
}

func BenchmarkEvaluate(b *testing.B) {
	it := data.Item{
		ID:      0,
		Name:    "Axe",
		Quality: item.QualitySuperior,
	}

	rule, err := NewRule(
		"[type] == amulet && [quality] == crafted # ([shapeshiftingskilltab] >= 2 || [elementalskilltab] >= 2 || [druidsummoningskilltab] >= 2) && [fcr] >= 10 && ([strength]+[maxhp]+[maxmana] >= 60 || [dexterity]+[maxhp]+[maxmana] >= 60 || [strength]+[dexterity]+[maxhp] >= 50 || [strength]+[dexterity]+[maxmana] >= 55)",
		"test",
		1,
	)
	require.NoError(b, err)

	for n := 0; n < b.N; n++ {
		rule.Evaluate(it)
	}
}
