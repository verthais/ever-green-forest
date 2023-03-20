package characters

import "fmt"

type ISkillRegistry interface {
	Register(SkillType, *Skill) *SkillRegistry
	GetSkill(SkillType) *Skill
	Size() int
}

func NewSkillRegistry() ISkillRegistry {
	sr := SkillRegistry{skills: SkillMap{}}
	sr.Register(DefaultAttack, DefaultOffensiveSkill())
	sr.Register(DefaultDefense, DefaultDefensiveSkill())
	sr.Register(RapidStrike, RapidStrikeSkill())
	sr.Register(MagicShield, MagicShieldSkill())
	return &sr
}

type SkillMap map[SkillType]*Skill

type SkillRegistry struct {
	skills SkillMap
}

func (sr *SkillRegistry) Register(st SkillType, s *Skill) *SkillRegistry {
	sr.skills[st] = s
	return sr
}

func (sr SkillRegistry) GetSkill(skill SkillType) *Skill {
	return sr.skills[skill]
}

func (sr SkillRegistry) Size() int {
	return len(sr.skills)
}

func DefaultOffensiveSkill() *Skill {
	name := "Attack"
	return &Skill{Name: name, Effect: func(attacker, defender ICharacter, sr ISkillRegistry) {
		fmt.Printf("%s is using: %s\n", attacker.Name(), name)

		defense := sr.GetSkill(*defender.GetDefensiveSkill())
		defense.Effect(attacker, defender, sr)
	}}
}

func DefaultDefensiveSkill() *Skill {
	name := "Defend"
	return &Skill{Name: name, Effect: func(attacker, defender ICharacter, sr ISkillRegistry) {
		fmt.Printf("%s is using: %s\n", defender.Name(), name)
		damage := attacker.Strength() - defender.Defense()
		damage = defender.TakeDamage(damage)
		fmt.Printf(
			"%s[%d/%d] took %d damage from %s\n",
			defender.Name(), defender.Health(), defender.BaseHealth(), damage, attacker.Name())
	}}
}

func RapidStrikeSkill() *Skill {
	name := "Rapid Strike"
	return &Skill{Name: name, Effect: func(attacker, defender ICharacter, sr ISkillRegistry) {
		fmt.Printf("%s is using: %s\n", attacker.Name(), name)

		defense := sr.GetSkill(*defender.GetDefensiveSkill())
		defense.Effect(attacker, defender, sr)

		defense = sr.GetSkill(*defender.GetDefensiveSkill())
		defense.Effect(attacker, defender, sr)
	}}
}

func MagicShieldSkill() *Skill {
	name := "Magic Shield"
	return &Skill{Name: name, Effect: func(attacker, defender ICharacter, sr ISkillRegistry) {
		fmt.Printf("%s is using: %s\n", defender.Name(), name)
		damage := attacker.Strength() - defender.Defense()
		fmt.Printf(
			"%s[%d/%d] blocked %d damage from %s\n",
			defender.Name(), defender.Health(), defender.BaseHealth(), damage, attacker.Name())
	}}
}
