package dice

import (
    "math/rand"
)

// RollD4 returns a random number between 1 and 4.
func RollD4() int {
    return rand.Intn(4) + 1
}
