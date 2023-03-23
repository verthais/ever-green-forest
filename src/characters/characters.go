package characters

import (
	"swisscom/src/engine/dice"
)

type ICharacter interface {
	Alive() bool
	BaseHealth() int
	Defense() int
	GetOffensiveSkill() *SkillType
	GetDefensiveSkill() *SkillType
	Heal()
	Health() int
	Lucky() bool
	Luck() int
	Name() string
	Strength() int
	Speed() int
	TakeDamage(int) int
}

func NewEnemy() ICharacter {
	bounds := [][]int{
		{60, 90},
		{60, 90},
		{40, 60},
		{40, 60},
		{25, 40},
	}

	c := NewCharacter(GenMonsterName(), bounds)

	c.offensiveSkills.Register(DefaultAttack, 100)

	c.defensiveSkills.Register(DefaultDefense, 100)

	return c
}

func NewHero() ICharacter {
	bounds := [][]int{
		{70, 100},
		{70, 80},
		{45, 55},
		{40, 50},
		{10, 30},
	}

	c := NewCharacter("Natelus", bounds)

	c.offensiveSkills.Register(RapidStrike, 10)
	c.offensiveSkills.Register(DefaultAttack, 100)

	c.defensiveSkills.Register(MagicShield, 20)
	c.defensiveSkills.Register(DefaultDefense, 100)

	return c
}

func NewCharacter(name string, bounds [][]int) *Character {
	return &Character{
		name:          name,
		baseHealth:    NewAttribute("Base Health", bounds[0]),
		currentHealth: NewAttribute("Current Health", bounds[0]),
		strength:      NewAttribute("Strength", bounds[1]),
		defense:       NewAttribute("Defense", bounds[2]),
		speed:         NewAttribute("Speed", bounds[3]),
		luck:          NewAttribute("Luck", bounds[4]),
	}
}

type Character struct {
	name            string
	currentHealth   Attribute
	baseHealth      Attribute
	strength        Attribute
	defense         Attribute
	speed           Attribute
	luck            Attribute
	offensiveSkills Skills
	defensiveSkills Skills
}

func (c *Character) Lucky() bool {
	roll := dice.RollK100()

	if c.luck.Value >= roll {
		return true
	}

	return false
}

func (c Character) Name() string {
	return c.name
}

func (c Character) Health() int {
	return c.currentHealth.Value
}

func (c Character) BaseHealth() int {
	return c.baseHealth.Value
}

func (c *Character) TakeDamage(dmg int) int {
	c.currentHealth.Value -= dmg

	if c.currentHealth.Value < 0 {
		dmg = dmg + c.currentHealth.Value
		c.currentHealth.Value = 0
		return dmg
	}

	return dmg
}

func (c Character) Alive() bool {
	return c.currentHealth.Value > 0
}

func (c *Character) Heal() {
	c.currentHealth.Value = c.baseHealth.Value
}

func (c Character) Strength() int {
	return c.strength.Value
}

func (c Character) Defense() int {
	return c.defense.Value
}

func (c Character) Speed() int {
	return c.speed.Value
}

func (c Character) Luck() int {
	return c.luck.Value
}

func (c Character) GetOffensiveSkill() *SkillType {
	return c.offensiveSkills.GetSkill()
}

func (c *Character) GetDefensiveSkill() *SkillType {
	return c.defensiveSkills.GetSkill()
}
