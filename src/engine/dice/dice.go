package dice

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RollK100() int {
	return rand.Intn(100)
}
