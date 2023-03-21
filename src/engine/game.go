package game

import (
	"fmt"
	"swisscom/src/characters"
	"swisscom/src/engine/dice"
)

func NewGame() Game {
	dice.Init()
	return Game{
		MaxLoops: 20,
		ec:       NewEncounterController(),
		cf:       characters.NewCharacterFactory,
	}
}

type Game struct {
	MaxLoops int
	ec       IEncounterController
	cf       func(characters.CharacterType) func() characters.ICharacter
}

func (g *Game) Run() {
	hero := g.cf(characters.HERO)()
	enemy := g.cf(characters.BEAST)()
	loop := 0

	for ; loop < g.MaxLoops; loop++ {

		g.ec.EncounterLogic(hero, enemy)

		if !hero.Alive() || !enemy.Alive() {
			break
		}
	}

	if hero.Alive() {
		fmt.Printf("Hero %s won!\n", hero.Name())
	} else if enemy.Alive() {
		fmt.Printf("Beast %s won!\n", enemy.Name())
	} else {
		fmt.Printf("Hero %s survived for %d rounds\n", hero.Name(), loop)
	}
}
