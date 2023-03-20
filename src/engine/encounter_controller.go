package game

import (
	"fmt"
	"swisscom/src/characters"
)

type IEncounterController interface {
	EncounterLogic(h, e characters.ICharacter)
}

func NewEncounterController() IEncounterController {
	return &EncounterController{
		sr: characters.NewSkillRegistry(),
	}
}

type EncounterController struct {
	sr characters.ISkillRegistry
}

func Sort(lhs, rhs characters.ICharacter) (characters.ICharacter, characters.ICharacter) {
	if lhs.Speed() < rhs.Speed() {
		return lhs, rhs
	}

	if lhs.Speed() > rhs.Speed() {
		return rhs, lhs
	}

	if lhs.Luck() < rhs.Luck() {
		return lhs, rhs
	}

	if lhs.Luck() > rhs.Luck() {
		return rhs, lhs
	}

	return lhs, rhs
}

func (ec *EncounterController) action(attacker, defender characters.ICharacter) {
	if !defender.Lucky() {
		attack := ec.sr.GetSkill(*attacker.GetOffensiveSkill())
		attack.Effect(attacker, defender, ec.sr)
	} else {
		fmt.Printf("%s got lucky!\n", defender.Name())
	}
}

func (ec *EncounterController) EncounterLogic(h, e characters.ICharacter) {
	defender, attacker := Sort(h, e)

	ec.action(attacker, defender)
	ec.action(defender, attacker)
}
