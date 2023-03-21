package characters

type CharacterType int

const (
	HERO CharacterType = iota
	BEAST
)

func NewCharacterFactory(t CharacterType) func() ICharacter {
	switch t {
	case HERO:
		return NewHero
	case BEAST:
		return NewEnemy
	default:
		return nil
	}
}
