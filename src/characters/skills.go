package characters

type SkillType int

const (
	DefaultAttack SkillType = iota
	DefaultDefense
	RapidStrike
	MagicShield
)

type SkillDescriptor struct {
	skillType SkillType
	chance    int
}

type Skills struct {
	all []SkillDescriptor
}

func (ss *Skills) Register(s SkillType, chance int) {
	// TODO Sort from lowest to highest probability
	ss.all = append(ss.all, SkillDescriptor{s, chance})
}

func (ss *Skills) GetSkill(luck int) *SkillType {
	for _, st := range ss.all {
		if luck < st.chance {
			return &st.skillType
		}
	}

	return nil
}

type Skill struct {
	Name   string `json:"name"`
	Effect func(attacker, defender ICharacter, sr ISkillRegistry)
}
