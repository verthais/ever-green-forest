package tests

import (
	"swisscom/src/characters"
	game "swisscom/src/engine"
	"testing"
)

func TestSort(t *testing.T) {
	slow := characters.NewCharacter(characters.GenMonsterName(), [][]int{
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 40},
		{0, 1},
	})
	fast := characters.NewCharacter(characters.GenMonsterName(), [][]int{
		{0, 1},
		{0, 1},
		{0, 1},
		{50, 100},
		{0, 1},
	})

	p_slow, p_fast := game.Sort(slow, fast)

	if slow != p_slow || fast != p_fast {
		t.Error("Fail sort")
	}

	p_slow, p_fast = game.Sort(fast, slow)

	if slow != p_slow || fast != p_fast {
		t.Error("Fail sort")
	}
}
