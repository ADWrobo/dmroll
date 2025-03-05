package tables

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
)

// EnemiesEntry represents one row in the Enemies in Davokar table.
type EnemiesEntry struct {
    MinRoll  int
    MaxRoll  int
    Examples string
    CR       string // Challenge Rating
    Number   string
}

// enemiesTable holds the ranges and data for “Enemies in Davokar”.
var enemiesTable = []EnemiesEntry{
    {1, 8,  "None",                                                   "--",         "--"},
    {9, 10, "Village Guards or Kotkas or Blight-born Humans",         "1/2 or less","1d4+2"},
    {11,12, "Fortune Hunters or Jakaars or Frost Lights",             "1 or less",  "1d6+3"},
    {13,14, "Colossi or Hunger Wolves or Stone Boars",                "2 or less",  "2d4+2"},
    {15,16, "Ferber Swarm or Rage Trolls or Killer Shrubs",           "4 or less",  "1d4+1"},
    {17,18, "Liege Trolls or Primal Blight Beasts or Ravenous Willows","10 or less","1d4"},
    {19,20, "Robbers + Robber Chief",                                 "1/2+2",      "1d4+3"},
    {21,22, "Dragouls + Necromage",                                   "1+15",       "2d4+2"},
    {23,23, "Skullbiter Crushers + Skullbiter Queen",                 "5+15",       "1d4"},
    {24,24, "Aboars + Lindworm",                                      "5+15",       "1d6"},
    {25,25, "World Serpent Tunnelers + World Serpent Wallower",       "17+26",      "1d2"},
    // There's no row for 26+, so a roll over 25 will fallback to ??? row.
}

// enemiesInDavokar implements the Table interface, so the CLI can -t enemies_in_davokar
type enemiesInDavokar struct{}

func init() {
    RegisterTable(&enemiesInDavokar{})
}

func (e *enemiesInDavokar) Name() string {
    return "enemies_in_davokar"
}

// GetRandomEntry prompts for the Davokar modifier, rolls 1d20 + mod, then returns the resulting row.
func (e *enemiesInDavokar) GetRandomEntry() string {
    // 1) Prompt user for environment-based modifier (from davokar_modifiers.go)
    mod := PromptForDavokarModifier()
    
    // optional blank line after user input
    fmt.Println()

    // 2) Roll 1d20
    baseRoll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling dice for Enemies in Davokar: %v\n", err)
        baseRoll = 1
    }
    finalRoll := baseRoll + mod

    // 3) Look up the correct row in the table
    entry := findEnemiesEntry(finalRoll)

    // 4) Return the formatted result
    return fmt.Sprintf(
        "Enemies in Davokar result: %d (1d20+%d)\nExamples: %s\nChallenge Rating: %s\nNumber: %s",
        finalRoll, mod, entry.Examples, entry.CR, entry.Number,
    )
}

// GetFormatted prints the entire Enemies in Davokar table (e.g. for `-p enemies_in_davokar`).
func (e *enemiesInDavokar) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("ENEMIES IN DAVOKAR (1d20 + Mods)\n")
    sb.WriteString("--------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-8s | %-45s | %-15s | %s\n",
        "Roll", "Examples", "Challenge Rating", "Number"))
    sb.WriteString("--------------------------------------------------------------\n")

    for _, row := range enemiesTable {
        rangeLabel := ""
        if row.MaxRoll == 0 {
            // If you had an open-ended row (e.g. 25+), you'd do something like fmt.Sprintf("%d+", row.MinRoll).
            // But here, we don’t have a 25+ row, so it doesn't apply. We'll ignore this case for now.
            rangeLabel = fmt.Sprintf("%d+", row.MinRoll)
        } else if row.MinRoll == row.MaxRoll {
            // Single value like "23"
            rangeLabel = fmt.Sprintf("%d", row.MinRoll)
        } else {
            // A range like "1-8" or "9-10"
            rangeLabel = fmt.Sprintf("%d-%d", row.MinRoll, row.MaxRoll)
        }

        sb.WriteString(fmt.Sprintf("%-8s | %-45s | %-15s | %s\n",
            rangeLabel, row.Examples, row.CR, row.Number))
    }
    return sb.String()
}

// findEnemiesEntry finds which row in enemiesTable corresponds to the final roll.
func findEnemiesEntry(roll int) EnemiesEntry {
    for _, row := range enemiesTable {
        // If row.MaxRoll=0, it might indicate an open-ended like "25+",
        // but in our table we didn’t define a 25+ row, so you can skip that logic or adapt if needed.
        if row.MaxRoll < row.MinRoll {
            // e.g. 23+ or 25+ if you had that scenario
            if roll >= row.MinRoll {
                return row
            }
        } else {
            // Normal range: roll must be >= MinRoll AND <= MaxRoll
            if roll >= row.MinRoll && roll <= row.MaxRoll {
                return row
            }
        }
    }
    // If user gets 26 or more, we have no row. Provide a fallback:
    return EnemiesEntry{
        MinRoll:  roll,
        MaxRoll:  roll,
        Examples: "???",
        CR:       "--",
        Number:   "--",
    }
}

func (r *enemiesInDavokar) Category() string {
    return "Davokar"
}