package dice

import (
	"math/rand"
)

// RollD12 returns a random number between 1 and 12.
func RollD12() int {
	return rand.Intn(12) + 1
}
