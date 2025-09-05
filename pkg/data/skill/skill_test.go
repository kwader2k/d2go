package skill

import (
	"testing"
)

// checks that all SkillDesc keys in Desc map have a corresponding skill in Skills.
func TestSkillsHaveDescriptions(t *testing.T) {
	// Create a set of all SkillDesc values from Skills
	skillDescSet := make(map[SkillDesc]bool)
	for _, skill := range Skills {
		if skill.SkillDesc != "" {
			skillDescSet[skill.SkillDesc] = true
		}
	}

	// Check that every SkillDesc key in Desc has a corresponding skill
	for skillDescKey := range Desc {
		if !skillDescSet[skillDescKey] {
			t.Errorf("SkillDesc %q is in Desc map but no skill uses it", skillDescKey)
		}
	}
}

// checks that Skills map keys match their ID field values.
func TestSkillIDConsistency(t *testing.T) {
	for mapKey, skill := range Skills {
		if mapKey != skill.ID {
			t.Errorf("Skill map key %d does not match skill.ID %d for skill %q", mapKey, skill.ID, skill.Name)
		}
	}
}
