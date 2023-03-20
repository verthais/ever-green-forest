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
	}
}

type Game struct {
	MaxLoops int
	ec       IEncounterController
}

func (g *Game) Run() {
	hero := characters.NewHero()
	enemy := characters.NewEnemy()
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
