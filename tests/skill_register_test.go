package tests

import (
	"testing"

	"swisscom/src/characters"
)

func TestSkillRegister(t *testing.T) {
	sr := characters.NewSkillRegistry()

	KNOWN_SKILLS := 4

	if sr.Size() != KNOWN_SKILLS {
		t.Error("Known skills dataset has changed!")
	}
}

func TestOffensiveSkill(t *testing.T) {
	c := characters.NewEnemy()

	offSkill := *c.GetOffensiveSkill()

	if offSkill != characters.DefaultAttack {
		t.Error("Wrong Enemy Default Attack")
	}
}

func TestDefensiveSkill(t *testing.T) {
	c := characters.NewEnemy()

	defSkill := *c.GetDefensiveSkill()

	if defSkill != characters.DefaultDefense {
		t.Error("Wrong Enemy Default Defense")
	}
}

func TestSkillRegistry(t *testing.T) {
	sr := characters.NewSkillRegistry()

	skill := sr.GetSkill(characters.DefaultAttack)
	expected := characters.DefaultOffensiveSkill()
	if skill.Name != expected.Name {
		t.Errorf("Skill %s does not match %s", skill.Name, expected.Name)
	}

	skill = sr.GetSkill(characters.DefaultDefense)
	expected = characters.DefaultDefensiveSkill()
	if skill.Name != expected.Name {
		t.Errorf("Skill %s does not match %s", skill.Name, expected.Name)
	}

	skill = sr.GetSkill(characters.RapidStrike)
	expected = characters.RapidStrikeSkill()
	if skill.Name != expected.Name {
		t.Errorf("Skill %s does not match %s", skill.Name, expected.Name)
	}

	skill = sr.GetSkill(characters.MagicShield)
	expected = characters.MagicShieldSkill()
	if skill.Name != expected.Name {
		t.Errorf("Skill %s does not match %s", skill.Name, expected.Name)
	}
}
