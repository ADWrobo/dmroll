package dice

import (
    "math/rand"
)

// RollD20 returns a random number between 1 and 20.
func RollD20() int {
    return rand.Intn(20) + 1
}
