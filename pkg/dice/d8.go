package dice

import (
	"math/rand"
)

// RollD8 returns a random number between 1 and 8.
func RollD8() int {
	return rand.Intn(8) + 1
}
