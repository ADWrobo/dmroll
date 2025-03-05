// It's kind of silly to create a d7 for the days of the week table but whatever...
package dice

import (
	"math/rand"
)

// RollD8 returns a random number between 1 and 8.
func RollD7() int {
	return rand.Intn(7) + 1
}
