package dice

import (
    "fmt"
    "math/rand"
    "regexp"
    "strconv"
    "strings"
)

// RollDice parses a dice notation (like "1d20" or "4d6") and rolls accordingly.
func RollDice(notation string) (int, error) {
    // Example valid strings:  "1d20", "3d6", "d10" (which means "1d10"), etc.

    // Use a regex to parse the format XdY
    re := regexp.MustCompile(`^(\d*)[dD](\d+)$`)
    match := re.FindStringSubmatch(strings.ToLower(strings.TrimSpace(notation)))
    if match == nil {
        return 0, fmt.Errorf("invalid dice notation: %s", notation)
    }

    // If the first capture (number of dice) is empty, default to 1
    numDiceStr := match[1]
    if numDiceStr == "" {
        numDiceStr = "1"
    }
    numDice, err := strconv.Atoi(numDiceStr)
    if err != nil {
        return 0, fmt.Errorf("invalid number of dice: %v", err)
    }

    // The second capture is the die size
    dieSizeStr := match[2]
    dieSize, err := strconv.Atoi(dieSizeStr)
    if err != nil {
        return 0, fmt.Errorf("invalid die size: %v", err)
    }

    // Roll all dice and sum
    total := 0
    for i := 0; i < numDice; i++ {
        roll := rand.Intn(dieSize) + 1
        total += roll
    }

    return total, nil
}
