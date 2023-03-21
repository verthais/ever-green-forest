package tests

import (
	"reflect"
	"swisscom/src/characters"
	"testing"
)

func TestCharacterFact(t *testing.T) {
	factory := characters.NewCharacterFactory

	if reflect.ValueOf(factory(characters.HERO)) != reflect.ValueOf(characters.NewHero) {
		t.Error("Return wrong hero factory")
	}

	if reflect.ValueOf(factory(characters.BEAST)) != reflect.ValueOf(characters.NewEnemy) {
		t.Error("Return wrong enemy factory")
	}
}
