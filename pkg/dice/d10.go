package dice

import (
	"math/rand"
)

// RollD10 returns a random number between 1 and 10.
func RollD10() int {
	return rand.Intn(10) + 1
}
