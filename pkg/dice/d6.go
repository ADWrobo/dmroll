package dice

import (
	"math/rand"
)

// RollD6 returns a random number between 1 and 6.
func RollD6() int {
	return rand.Intn(6) + 1
}
