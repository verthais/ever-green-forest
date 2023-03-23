package tests

import (
	"fmt"
	"math/rand"
	"swisscom/src/characters"
	"testing"
)

type Result struct {
	value bool
	msg   string
}

func errMsg(attr string, value int, lower int, upper int) string {
	return fmt.Sprintf(
		"%s value {%d} outside of bounds {%d, %d}", attr, value, lower, upper)
}

func assertAttr(cmp int, bounds []int) bool {
	return cmp < bounds[0] || cmp > bounds[1]
}

func assertChar(c characters.ICharacter, bounds [][]int) Result {
	if assertAttr(c.Health(), bounds[0]) {
		return Result{
			false,
			errMsg("Health", c.Health(), bounds[0][0], bounds[0][1]),
		}
	}

	c.Heal()
	if assertAttr(c.Health(), bounds[0]) {
		return Result{
			false,
			errMsg("CurrentHealth", c.Health(), bounds[0][0], bounds[0][1]),
		}
	}

	if assertAttr(c.Strength(), bounds[1]) {
		return Result{
			false,
			errMsg("Strength", c.Strength(), bounds[1][0], bounds[1][1]),
		}
	}

	if assertAttr(c.Defense(), bounds[2]) {
		return Result{
			false,
			errMsg("Defense", c.Defense(), bounds[2][0], bounds[2][1]),
		}
	}

	if assertAttr(c.Speed(), bounds[3]) {
		return Result{
			false,
			errMsg("Speed", c.Speed(), bounds[3][0], bounds[3][1]),
		}
	}

	return Result{true, ""}
}

func fuzzyTest(msg chan Result, factory func() characters.ICharacter, referenceValues [][]int) {
	execution := 100 + rand.Int()%100
	for i := 0; i < execution; i++ {
		c := factory()
		if r := assertChar(c, referenceValues); r.value != true {
			msg <- r
			break
		}
	}
	msg <- Result{true, ""}
}

func scheduleFuzzing(t *testing.T, factory func() characters.ICharacter, referenceValues [][]int) {
	msg := make(chan Result)

	execution := 100 + rand.Int()%100

	for i := 0; i < execution; i++ {
		go fuzzyTest(msg, factory, referenceValues)
	}

	for i := 0; i < execution; i++ {
		r := <-msg
		if r.value != true {
			t.Errorf(r.msg)
		}
	}
}

func TestNewBeast(t *testing.T) {
	referenceValues := [][]int{
		{60, 90},
		{60, 90},
		{40, 60},
		{40, 60},
		{25, 40},
	}

	scheduleFuzzing(t, characters.NewEnemy, referenceValues)
}

func TestNewNatelus(t *testing.T) {
	referenceValues := [][]int{
		{70, 100},
		{70, 80},
		{45, 55},
		{40, 50},
		{10, 30},
	}

	scheduleFuzzing(t, characters.NewHero, referenceValues)
}

func TestTakeDamage(t *testing.T) {
	referenceValues := [][]int{
		{70, 100},
		{70, 80},
		{45, 55},
		{40, 50},
		{10, 30},
	}

	char := characters.NewCharacter(characters.GenMonsterName(), referenceValues)

	i := char.Health()

	value := char.TakeDamage(i)

	if value != i {
		t.Errorf("Took %d - expected %d", value, i)
	}

	value = char.TakeDamage(i)

	if value != 0 {
		t.Errorf("Took %d dmg - expected %d", value, 0)
	}

	if char.Health() != 0 {
		t.Errorf("status %d - expected %d health", char.Health(), 0)
	}

	if char.Alive() {
		t.Errorf("status %t - expected %t", char.Alive(), false)
	}
}
