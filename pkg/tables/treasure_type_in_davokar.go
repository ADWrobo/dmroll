// pkg/tables/treasure_type_in_davokar.go

package tables

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
)

// We'll store the table breakpoints for each region in a simple struct:
type treasureRange struct {
    min int
    max int
    typ string // "D", "C", "M", or "A"
}

var treasureMap = map[string][]treasureRange{
    "bright": {
        {1, 6,  "D"},
        {7, 12, "C"},
        {13, 18,"M"},
        {19, 20,"A"},
    },
    "wild": {
        {1, 5,  "D"},
        {6, 12, "C"},
        {13, 18,"M"},
        {19, 20,"A"},
    },
    "dark": {
        {1, 4,  "D"},
        {5, 12, "C"},
        {13, 18,"M"},
        {19, 20,"A"},
    },
}

// treasureTypeInDavokar implements the Table interface.
type treasureTypeInDavokar struct {}

func init() {
    RegisterTable(&treasureTypeInDavokar{})
}

func (t *treasureTypeInDavokar) Name() string {
    return "treasure_type_in_davokar"
}

// GetRandomEntry: Ask user which region (no waterways), roll 1d20, match the row, and return a message.
func (t *treasureTypeInDavokar) GetRandomEntry() string {
    region := PromptForDavokarRegion() // bright, wild, or dark

    // Roll 1d20
    roll, err := dice.RollDice("1d20")
    if err != nil {
        log.Println("Error rolling 1d20, defaulting to 1")
        roll = 1
    }

    // Find the correct type (D, C, M, or A) for this region & roll
    resultType := findTreasureType(region, roll)

    // Provide some explanation for each letter
    explanation := map[string]string{
        "D": "Debris (no significant value)",
        "C": "Curiosity – roll on the Curiosities table",
        "M": "Mystical Treasure – roll on Mystical Treasures table",
        "A": "Artifact – roll on Artifact table",
    }

    return fmt.Sprintf(
        "Treasure Type in Davokar result: %d\nRegion: %s\nType: %s\nMeaning: %s",
        roll, strings.Title(region), resultType, explanation[resultType],
    )
}

// GetFormatted prints the entire matrix in ASCII, for reference.
func (t *treasureTypeInDavokar) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("TREASURE TYPE IN DAVOKAR (1d20)\n")
    sb.WriteString("Region: Bright | 1–6=D, 7–12=C, 13–18=M, 19–20=A\n")
    sb.WriteString("Region: Wild   | 1–5=D, 6–12=C, 13–18=M, 19–20=A\n")
    sb.WriteString("Region: Dark   | 1–4=D, 5–12=C, 13–18=M, 19–20=A\n")
    sb.WriteString("\n")
    sb.WriteString("D = Debris, C = Curiosity, M = Mystical Treasure, A = Artifact\n")
    sb.WriteString("Curiosity => roll on Curiosities table\n")
    sb.WriteString("Mystical Treasure => roll on Mystical Treasures table\n")
    sb.WriteString("Artifact => roll on Artifact table\n")
    return sb.String()
}

// findTreasureType loops the relevant slice in treasureMap for the given region, returning the correct D/C/M/A.
func findTreasureType(region string, roll int) string {
    ranges, ok := treasureMap[region]
    if !ok {
        // if region is something unexpected, default to bright
        ranges = treasureMap["bright"]
    }
    for _, r := range ranges {
        if roll >= r.min && roll <= r.max {
            return r.typ
        }
    }
    // If it's out of range for some reason, default to "D" or handle it differently:
    return "D"
}

func (r *treasureTypeInDavokar) Category() string {
    return "Davokar"
}