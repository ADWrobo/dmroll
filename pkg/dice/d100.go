package dice

import (
	"math/rand"
)

// RollD100 returns a random number between 1 and 100
func RollD100() int {
	return rand.Intn(100) + 1
}
