package characters

import "math/rand"

func GenMonsterName() string {
	// TODO: diff stat baseline based of type
	names := []string{
		"Ape",
		"Bear",
		"Beholder",
		"Centipedes",
		"Dworc",
		"Gargoyle",
		"Monkey",
		"Orc",
		"Scorpion",
		"Snake",
		"Slime",
		"Skeleton",
		"Troll",
		"Wolf",
	}

	// TODO: diff skills based of traits
	adjectives := []string{
		"Adorable",
		"Braveless",
		"Bright",
		"Clumsy",
		"Elegant",
		"Selfish",
		"Scruffy",
		"Slender",
		"Sparkling",
		"Witty",
	}

	n_idx, a_ind := rand.Intn(len(names)), rand.Intn(len(adjectives))

	return adjectives[a_ind] + " " + names[n_idx]
}

func getRandom(base int, cap int) int {
	return base + rand.Intn(cap-base)
}
