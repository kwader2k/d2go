package nip

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/item"
	"github.com/hectorgimenez/d2go/pkg/data/stat"
)

const (
	RuleResultFullMatch RuleResult = 1
	RuleResultPartial   RuleResult = 2
	RuleResultNoMatch   RuleResult = 3
)

var (
	fixedPropsRegexp = regexp.MustCompile(`(\[(type|quality|class|name|flag|color|prefix|suffix)]\s*(<=|<|>|>=|!=|==)\s*([a-zA-Z0-9]+))`)
	statsRegexp      = regexp.MustCompile(`\[(.*?)]`)
	maxQtyRegexp     = regexp.MustCompile(`(\[maxquantity]\s*(<=|<|>|>=|!=|==)\s*([0-9]+))`)
	tierRegexp       = regexp.MustCompile(`(\[tier]\s*(<=|<|>|>=|!=|==)\s*([0-9]+))`)
	mercTierRegexp   = regexp.MustCompile(`(\[merctier]\s*(<=|<|>|>=|!=|==)\s*([0-9]+))`)
)

type Rule struct {
	RawLine       string // Original line, don't use it for evaluation
	Filename      string
	LineNumber    int
	Enabled       bool
	maxQuantity   int
	tier          float64
	mercTier      float64
	stage1        *vm.Program
	stage2        *vm.Program
	requiredStats []string
}

type RuleResult int
type Rules []Rule

func (r Rules) EvaluateAll(it data.Item) (Rule, RuleResult) {
	bestMatch := RuleResultNoMatch
	bestMatchingRule := Rule{}
	for _, rule := range r {
		if rule.Enabled {
			result, err := rule.Evaluate(it)
			if err != nil {
				continue
			}
			if result == RuleResultFullMatch {
				return rule, result
			}
			if result == RuleResultPartial {
				bestMatch = result
				bestMatchingRule = rule
			}
		}
	}

	return bestMatchingRule, bestMatch
}

func (r Rules) EvaluateAllIgnoreTiers(it data.Item) (Rule, RuleResult) {
	bestMatch := RuleResultNoMatch
	bestMatchingRule := Rule{}
	for _, rule := range r {
		if rule.Enabled {
			if rule.tier > 0 || rule.mercTier > 0 {
				continue
			}
			result, err := rule.Evaluate(it)
			if err != nil {
				continue
			}
			if result == RuleResultFullMatch {
				return rule, result
			}
			if result == RuleResultPartial {
				bestMatch = result
				bestMatchingRule = rule
			}
		}
	}

	return bestMatchingRule, bestMatch
}

func (r Rules) EvaluateTiers(it data.Item, tierRulesIndexes []int) (Rule, Rule) {
	highestTierRule := Rule{}
	highestMercTierRule := Rule{}
	for _, ruleIndex := range tierRulesIndexes {
		if ruleIndex < len(r) {
			rule := r[ruleIndex]
			if rule.Enabled {
				result, err := rule.Evaluate(it)
				if err != nil {
					continue
				}
				if result == RuleResultFullMatch || result == RuleResultPartial {
					if rule.tier > highestTierRule.tier {
						highestTierRule = rule
					}
					if rule.mercTier > highestMercTierRule.mercTier {
						highestMercTierRule = rule
					}
				}
			}
		}
	}

	return highestTierRule, highestMercTierRule
}

var fixedPropsList = map[string]int{"type": 0, "quality": 0, "class": 0, "name": 0, "flag": 0, "color": 0, "prefix": 0, "suffix": 0}

func NewRule(rawRule string, filename string, lineNumber int) (Rule, error) {
	rule := sanitizeLine(rawRule)

	// Try to get the maxquantity value and purge it from the rule, we can not evaluate it
	maxQuantity := 0
	for _, prop := range maxQtyRegexp.FindAllStringSubmatch(rule, -1) {
		mxQty, err := strconv.Atoi(prop[3])
		if err != nil {
			return Rule{}, fmt.Errorf("error parsing maxquantity value %s: %w", prop[3], err)
		}
		maxQuantity = mxQty
		rule = strings.ReplaceAll(rule, prop[0], "")
	}

	// Try to get the tier value and purge it from the rule, we can not evaluate it yet
	tier := 0.0
	for _, prop := range tierRegexp.FindAllStringSubmatch(rule, -1) {
		parsedTier, err := strconv.Atoi(prop[3])
		if err != nil {
			return Rule{}, fmt.Errorf("error parsing tier value %s: %w", prop[3], err)
		}
		tier = float64(parsedTier)
		rule = strings.ReplaceAll(rule, prop[0], "")
	}

	// Try to get the merctier value and purge it from the rule, we can not evaluate it yet
	mercTier := 0.0
	for _, prop := range mercTierRegexp.FindAllStringSubmatch(rule, -1) {
		parsedMercTier, err := strconv.Atoi(prop[3])
		if err != nil {
			return Rule{}, fmt.Errorf("error parsing merctier value %s: %w", prop[3], err)
		}
		mercTier = float64(parsedMercTier)
		rule = strings.ReplaceAll(rule, prop[0], "")
	}

	// Sanitize again, just in case we messed up the rule while parsing maxquantity
	rule = sanitizeLine(rule)
	if rule == "" {
		return Rule{}, ErrEmptyRule
	}

	r := Rule{
		RawLine:     rawRule,
		Filename:    filename,
		LineNumber:  lineNumber,
		Enabled:     true,
		maxQuantity: maxQuantity,
		tier:        tier,
		mercTier:    mercTier,
	}

	parts := strings.Split(rule, "#")

	if len(parts) > 0 {
		stage1 := strings.TrimSpace(parts[0])
		if stage1 != "" {
			line, err := replaceStringPropertiesInStage1(stage1)
			if err != nil {
				return Rule{}, err
			}

			line = strings.ReplaceAll(line, "[", "")
			line = strings.ReplaceAll(line, "]", "")
			program, err := expr.Compile(line, expr.Env(fixedPropsList))
			if err != nil {
				return Rule{}, fmt.Errorf("error compiling rule stage1: %w", err)
			}
			r.stage1 = program
		}
	}

	if len(parts) > 1 {
		stage2 := strings.TrimSpace(parts[1])
		if stage2 != "" {
			// Extract stats before removing brackets for compilation
			r.requiredStats = getRequiredStatsForRule(stage2)

			statsMap := make(map[string]int)
			for _, prop := range r.requiredStats {
				statsMap[prop] = 0
			}

			// Normalize whitespace around operators in parenthesized expressions
			stage2 = normalizeParenthesizedExpressions(stage2)

			// Remove brackets for compilation
			compileReady := strings.ReplaceAll(stage2, "[", "")
			compileReady = strings.ReplaceAll(compileReady, "]", "")

			program, err := expr.Compile(compileReady, expr.Env(statsMap))
			if err != nil {
				return Rule{}, fmt.Errorf("error compiling rule stage2: %w, expression: %s", err, compileReady)
			}
			r.stage2 = program
		}
	}

	return r, nil
}

func normalizeParenthesizedExpressions(expr string) string {
	// Normalize common operators
	expr = strings.ReplaceAll(expr, "||", " || ")
	expr = strings.ReplaceAll(expr, "&&", " && ")
	expr = strings.ReplaceAll(expr, "==", " == ")
	expr = strings.ReplaceAll(expr, "!=", " != ")
	expr = strings.ReplaceAll(expr, ">=", " >= ")
	expr = strings.ReplaceAll(expr, "<=", " <= ")

	// Fix extra spaces
	expr = regexp.MustCompile(`\s+`).ReplaceAllString(expr, " ")

	// Normalize parentheses spacing
	expr = strings.ReplaceAll(expr, "( ", "(")
	expr = strings.ReplaceAll(expr, " )", ")")

	return expr
}

func (r Rule) Evaluate(it data.Item) (RuleResult, error) {
	// Stage 1: Basic properties evaluation
	stage1Props := make(map[string]int)
	for prop := range fixedPropsList {
		switch prop {
		case "type":
			stage1Props["type"] = it.Type().ID
		case "quality":
			stage1Props["quality"] = int(it.Quality)
		case "class":
			stage1Props["class"] = int(it.Desc().Tier())
		case "name":
			stage1Props["name"] = it.ID
		case "flag":
			// 0x400000 (eth) | 0x4000000 (runeword) kolbot style
			currentFlag := 0

			if it.Ethereal {
				currentFlag |= 0x400000
			}
			if it.IsRuneword {
				currentFlag |= 0x4000000
			}

			stage1Props["flag"] = currentFlag
		case "prefix":
			if it.Affixes.Rare.Prefix != 0 {
				stage1Props["prefix"] = int(it.Affixes.Rare.Prefix)
			}
			for _, prefix := range it.Affixes.Magic.Prefixes {
				if prefix != 0 {
					stage1Props["prefix"] = int(prefix)
					break
				}
			}
		case "suffix":
			if it.Affixes.Rare.Suffix != 0 {
				stage1Props["suffix"] = int(it.Affixes.Rare.Suffix)
			}
			for _, suffix := range it.Affixes.Magic.Suffixes {
				if suffix != 0 {
					stage1Props["suffix"] = int(suffix)
					break
				}
			}
		case "color":
			// TODO: Not supported yet
		}
	}

	// Check if stage1 exists before evaluating
	if r.stage1 == nil {
		return RuleResultNoMatch, fmt.Errorf("stage1 program is nil")
	}

	// Let's evaluate first stage
	stage1Result, err := expr.Run(r.stage1, stage1Props)
	if err != nil {
		return RuleResultNoMatch, fmt.Errorf("error evaluating rule stage1: %w", err)
	}

	// If stage1 does not match, we can stop here, nothing else to match
	if !stage1Result.(bool) {
		return RuleResultNoMatch, nil
	}

	// If we have no stage2 (no stat requirements), allow full match even for unidentified items
	if r.stage2 == nil {
		return RuleResultFullMatch, nil
	}

	// From here on we have stat requirements - return partial match for unidentified items
	if !it.Identified {
		return RuleResultPartial, nil
	}

	stage2Props := make(map[string]int)
	stage2 := ""
	if len(strings.Split(r.RawLine, "#")) > 1 {
		stage2 = strings.ToLower(strings.Split(r.RawLine, "#")[1])
	}

	// Special handling for skill tabs
	if strings.Contains(stage2, "[itemaddskilltab]") {
		stage2Props["itemaddskilltab"] = evaluateSkillTabSum(it)
	}

	// Special handling for class skills
	if strings.Contains(stage2, "[itemaddclassskills]") {
		stage2Props["itemaddclassskills"] = evaluateClassSkillsSum(it)
	}

	// Preprocess stage2 to see if certain stats are being compared to zero
	zeroCheckStats := make(map[string]bool)
	for _, statName := range r.requiredStats {
		if strings.Contains(stage2, "["+statName+"] == 0") ||
			strings.Contains(stage2, "["+statName+"]==0") {
			zeroCheckStats[statName] = true
		}
	}

	// Handle resist sums - check if item has any resist stats (including negative values)
	hasAnyResist := false

	// Detect if this is a rule with a resist sum expression
	isResistSum := false
	if strings.Contains(stage2, "resist") {
		isResistSum = strings.Contains(stage2, "+") || strings.Contains(stage2, "-") ||
			(strings.Contains(stage2, "(") && strings.Contains(stage2, ")"))
	}

	// Check if rule ONLY checks resist sums (no other OR conditions)
	hasNonResistConditions := false
	for _, statName := range r.requiredStats {
		if !strings.Contains(statName, "resist") {
			hasNonResistConditions = true
			break
		}
	}

	if isResistSum {
		// Check if the item has any resist stats at all (including negative values like sunders)
		for _, statName := range r.requiredStats {
			if !strings.Contains(statName, "resist") {
				continue
			}
			statData, found := statAliases[statName]
			if !found {
				continue
			}
			layer := 0
			if len(statData) > 1 {
				layer = statData[1]
			}

			// Check if stat exists at all, regardless of value (including negative)
			if _, found := it.FindStat(stat.ID(statData[0]), layer); found {
				hasAnyResist = true
				break
			}
		}
	}

	// Evaluate each required stat
	for _, statName := range r.requiredStats {
		// Skip stats we've already handled
		if statName == "itemaddskilltab" || statName == "itemaddclassskills" {
			continue
		}

		// --- Special composite stats ---
		// Enhanced Damage (ED) can be stored in different stat IDs depending on the item source.
		// In practice we want a single value that represents the visible ED% on the tooltip.
		//
		// Common cases seen in the wild:
		//  - stat 17 (item_maxdamage_percent)
		//  - stat 18 (item_mindamage_percent)
		//  - stat 25 (damagepercent) on some runewords/edge cases
		//
		// Some items store both 17 and 18 with the same value; others only populate one of them.
		// If a rule asks for [enhanceddamage] (or the internal UI shorthand [eddmg]), we should
		// consider all of them.
		switch strings.ToLower(statName) {
		case "enhanceddamage", "eddmg":
			if ed, ok := findItemEDPercent(it); ok {
				stage2Props[statName] = ed
			} else {
				stage2Props[statName] = 0
			}
			continue
		}
		statData, found := statAliases[statName]
		if !found {
			return RuleResultNoMatch, fmt.Errorf("property %s is not valid or not supported", statName)
		}

		layer := 0
		if len(statData) > 1 {
			layer = statData[1]
		}

		// Use the FindStat method which handles both Stats and BaseStats
		statFound := false
		var statValue int

		if itemStat, found := it.FindStat(stat.ID(statData[0]), layer); found {
			statValue = itemStat.Value
			statFound = true
		}

		// Special handling for stats not found
		if !statFound {
			isResistStat := strings.Contains(statName, "resist")
			// Special case: Pure resist-sum rules (like sunders) should fail early if no resists exist
			// But mixed rules (resists OR other stats) should continue to check other conditions
			if isResistStat && isResistSum && !hasAnyResist && !hasNonResistConditions {
				// This is a pure resist rule and item has no resists - can't match
				return RuleResultNoMatch, nil
			}
			// For all other cases, default missing stats to 0 and let expression evaluate
			stage2Props[statName] = 0
		} else {
			stage2Props[statName] = statValue
		}
	}

	res, err := expr.Run(r.stage2, stage2Props)
	if err != nil {
		return RuleResultNoMatch, fmt.Errorf("error evaluating rule stage2: %w", err)
	}

	// 100% rule match, we can return here
	if res.(bool) {
		return RuleResultFullMatch, nil
	}

	return RuleResultNoMatch, nil
}


func replaceStringPropertiesInStage1(stage1 string) (string, error) {
	baseProperties := fixedPropsRegexp.FindAllStringSubmatch(stage1, -1)
	for _, prop := range baseProperties {
		replaceWith := ""
		switch prop[2] {
		case "type":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", item.ItemTypes[typeAliases[prop[4]]].ID))
		case "quality":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", qualityAliases[prop[4]]))
		case "class":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", classAliases[prop[4]]))
		case "name":
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", item.GetIDByName(prop[4])))
		case "flag":
			val := 0
			switch strings.ToLower(prop[4]) {
			case "runeword":
				val = 0x4000000
			case "ethereal":
				val = 0x400000
			default:
				val = 1
			}
			replaceWith = strings.ReplaceAll(prop[0], prop[4], fmt.Sprintf("%d", val))
		case "prefix", "suffix":
			// Handle prefix/suffix IDs
			replaceWith = strings.ReplaceAll(prop[0], prop[4], prop[4])
		case "color":
			// TODO: Not supported yet
			return "", fmt.Errorf("property %s is not supported yet", prop[2])
		}

		if replaceWith != "" {
			stage1 = strings.ReplaceAll(stage1, prop[0], replaceWith)
		}
	}

	return stage1, nil
}

func getRequiredStatsForRule(line string) []string {
	statsList := make([]string, 0)
	statsFound := make(map[string]bool)

	for _, statName := range statsRegexp.FindAllStringSubmatch(line, -1) {
		if !statsFound[statName[1]] {
			statsList = append(statsList, statName[1])
			statsFound[statName[1]] = true
		}
	}
	return statsList
}

func evaluateClassSkillsSum(it data.Item) int {
	// Check all class skills stats
	totalClassSkills := 0
	maxLayer := 6 // in aliases.go the max layer for class skills is 6 (itemaddassassinskills)

	for layer := 0; layer <= maxLayer; layer++ {
		if itemStat, found := it.FindStat(stat.AddClassSkills, layer); found && itemStat.Value > 0 {
			totalClassSkills += itemStat.Value
		}
	}

	return totalClassSkills
}
func evaluateSkillTabSum(it data.Item) int {
	// Check all skill tab stats
	totalSkillTabs := 0
	maxLayer := 50 // in aliases.go the max layer for skill tabs is 50 (itemaddmartialartsskilltab)

	for layer := 0; layer <= maxLayer; layer++ {
		if itemStat, found := it.FindStat(stat.AddSkillTab, layer); found && itemStat.Value > 0 {
			totalSkillTabs += itemStat.Value
		}
	}

	return totalSkillTabs
}

func findItemEDPercent(it data.Item) (value int, ok bool) {
	// Prefer layer 0 when present.
	if ed, found := it.FindStat(stat.EnhancedDamageMin, 0); found {
		return ed.Value, true
	}
	if ed, found := it.FindStat(stat.EnhancedDamage, 0); found {
		return ed.Value, true
	}
	if ed, found := it.FindStat(stat.DamagePercent, 0); found {
		return ed.Value, true
	}

	// Some items store ED under a non-zero layer; scan all layers (stats + basestats).
	if ed, found := findStatAnyLayerInItem(it, stat.EnhancedDamageMin); found {
		return ed.Value, true
	}
	if ed, found := findStatAnyLayerInItem(it, stat.EnhancedDamage); found {
		return ed.Value, true
	}
	if ed, found := findStatAnyLayerInItem(it, stat.DamagePercent); found {
		return ed.Value, true
	}

	// Some runewords (and occasionally other items) do not expose ED% as an explicit stat entry,
	// but the resulting weapon damage stats still reflect the roll. As a fallback, try to derive
	// ED% from base vs current damage.
	if minED, maxED, exact, derived := deriveWeaponEDPercentFromDamageStats(it); derived {
		if exact {
			return minED, true
		}
		// Best-effort: return the upper bound to avoid false negatives when filtering.
		return maxED, true
	}

	return 0, false
}

// deriveWeaponEDPercentFromDamageStats tries to infer on-weapon ED% by comparing base damage
// (BaseStats/Desc) to current damage (Stats/BaseStats). This is mainly used for runewords where
// ED may not appear as an explicit stat.
func deriveWeaponEDPercentFromDamageStats(it data.Item) (min int, max int, exact bool, ok bool) {
	// Prefer 2H damage when present.
	baseMax, curMax, baseMin, curMin, has2H := getDamagePair(it, stat.TwoHandedMaxDamage, stat.TwoHandedMinDamage, true)
	if !has2H {
		baseMax, curMax, baseMin, curMin, _ = getDamagePair(it, stat.MaxDamage, stat.MinDamage, false)
	}

	if baseMax <= 0 || curMax <= 0 {
		return 0, 0, false, false
	}

	minMax, maxMax, okMax := edPercentRangeFromBaseCurrent(baseMax, curMax)
	if !okMax {
		return 0, 0, false, false
	}
	min = minMax
	max = maxMax

	// If we can also compute from min damage, intersect. If intersection is empty (common when the
	// item adds flat min/max damage), fall back to the tighter of the two ranges.
	if baseMin > 0 && curMin > 0 {
		minMin, maxMin, okMin := edPercentRangeFromBaseCurrent(baseMin, curMin)
		if okMin {
			iMin := min
			if minMin > iMin {
				iMin = minMin
			}
			iMax := max
			if maxMin < iMax {
				iMax = maxMin
			}
			if iMin <= iMax {
				min, max = iMin, iMax
			} else {
				// Pick tighter range.
				widthMax := maxMax - minMax
				widthMin := maxMin - minMin
				if widthMin < widthMax {
					min, max = minMin, maxMin
				} else {
					min, max = minMax, maxMax
				}
			}
		}
	}

	exact = min == max
	return min, max, exact, true
}

func getDamagePair(it data.Item, maxID, minID stat.ID, preferDesc2H bool) (baseMax, curMax, baseMin, curMin int, ok bool) {
	// Base: prefer BaseStats, fallback to Stats.
	if s, found := it.BaseStats.FindStat(maxID, 0); found {
		baseMax = s.Value
	} else if s, found := it.Stats.FindStat(maxID, 0); found {
		baseMax = s.Value
	}
	if s, found := it.BaseStats.FindStat(minID, 0); found {
		baseMin = s.Value
	} else if s, found := it.Stats.FindStat(minID, 0); found {
		baseMin = s.Value
	}

	// Current: prefer Stats, fallback to BaseStats.
	if s, found := it.Stats.FindStat(maxID, 0); found {
		curMax = s.Value
	} else if s, found := it.BaseStats.FindStat(maxID, 0); found {
		curMax = s.Value
	}
	if s, found := it.Stats.FindStat(minID, 0); found {
		curMin = s.Value
	} else if s, found := it.BaseStats.FindStat(minID, 0); found {
		curMin = s.Value
	}

	// If base damage is missing, fall back to the item description (template/base values).
	// This is not perfect for superior/eth bases, but it's better than failing completely.
	if baseMax <= 0 {
		d := it.Desc()
		if preferDesc2H && d.TwoHandMaxDamage > 0 {
			baseMax = d.TwoHandMaxDamage
		} else if d.MaxDamage > 0 {
			baseMax = d.MaxDamage
		}
	}
	if baseMin <= 0 {
		d := it.Desc()
		if preferDesc2H && d.TwoHandMinDamage > 0 {
			baseMin = d.TwoHandMinDamage
		} else if d.MinDamage > 0 {
			baseMin = d.MinDamage
		}
	}

	// Current max is required to do anything useful.
	ok = baseMax > 0 && curMax > 0
	return baseMax, curMax, baseMin, curMin, ok
}

// edPercentRangeFromBaseCurrent returns the possible ED% integer range that could result in the
// observed current value given a base value, using integer truncation semantics.
func edPercentRangeFromBaseCurrent(base, current int) (minED int, maxED int, ok bool) {
	if base <= 0 || current <= 0 {
		return 0, 0, false
	}

	// x = 100 + ed
	// current = floor(base * x / 100)
	// => current <= base*x/100 < current+1
	// => 100*current/base <= x < 100*(current+1)/base
	// We need integer x, so:
	//   minX = ceil(100*current/base)
	//   maxX = floor((100*(current+1)-1)/base)
	// ed = x - 100
	numerMin := int64(100) * int64(current)
	minX := int((numerMin + int64(base) - 1) / int64(base))

	numerMax := int64(100)*(int64(current)+1) - 1
	maxX := int(numerMax / int64(base))

	minED = minX - 100
	maxED = maxX - 100
	if minED < 0 {
		minED = 0
	}
	if maxED < 0 {
		maxED = 0
	}
	if minED > maxED {
		return 0, 0, false
	}
	return minED, maxED, true
}

func findStatAnyLayerInItem(it data.Item, id stat.ID) (stat.Data, bool) {
	for _, s := range it.Stats {
		if s.ID == id {
			return s, true
		}
	}
	for _, s := range it.BaseStats {
		if s.ID == id {
			return s, true
		}
	}
	return stat.Data{}, false
}

// MaxQuantity returns the maximum quantity of items that character can have, 0 means no limit
func (r Rule) MaxQuantity() int {
	return r.maxQuantity
}

func (r Rule) Tier() float64 {
	return r.tier
}

func (r Rule) MercTier() float64 {
	return r.mercTier
}
