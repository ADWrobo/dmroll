package ruins_of_symbaroum_5e

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
    "dmroll/pkg/table_registry"
    "dmroll/pkg/tables"
)

// EncountersEntry represents one row of the Encounters in Davokar table.
type EncountersEntry struct {
    MinRoll      int
    MaxRoll      int
    Encounter    string
    PerceptionDC string
    Number       string
}

// This table covers rolls 1..8 explicitly, then "9+" is open-ended => "None"
var encountersTable = []EncountersEntry{
    {1, 1, "Hunting party",                 "15", "1d6+2"},
    {2, 2, "Missionaries",                  "10", "1d6+4"},
    {3, 3, "Treasure hunters",              "14", "1d4+6"},
    {4, 4, "Local settlement",              "10", "2d10+20"},
    {5, 5, "Rangers",                       "15", "1d4+4"},
    {6, 6, "Large expedition",              "12", "2d10+4"},
    {7, 7, "Nomadic goblins",               "16", "4d10+4"},
    {8, 8, "Diplomatic elves\\Civilized trolls\\Peaceful bestials", "20", "1d10+4"},
    // 9+ => None
    {9, 0, "None",                          "--",  "--"},
}

type encountersInDavokar struct{}

func init() {
    table_registry.RegisterTable(&encountersInDavokar{})
}

func (e *encountersInDavokar) Name() string {
    return "encounters_in_davokar"
}

// GetRandomEntry: prompts for a modifier, then rolls 1d20 + that modifier.
func (e *encountersInDavokar) GetRandomEntry() string {
    mod := tables.PromptForDavokarModifier()  // <-- from davokar_modifiers.go
    fmt.Println() // optional blank line

    baseRoll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling dice for Encounters in Davokar: %v\n", err)
        baseRoll = 1
    }
    finalRoll := baseRoll + mod
    entry := findEncounterEntry(finalRoll)

    return fmt.Sprintf(
        "Encounters in Davokar result: %d (1d20+%d)\nEncounter: %s\nPerception DC: %s\nNumber: %s",
        finalRoll, mod, entry.Encounter, entry.PerceptionDC, entry.Number,
    )
}

func (e *encountersInDavokar) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("ENCOUNTERS IN DAVOKAR (1d20 + Mods)\n")
    sb.WriteString("---------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-8s | %-40s | %-13s | %s\n",
        "Roll", "Encounter", "Perception DC", "Number"))
    sb.WriteString("---------------------------------------------------------------\n")

    for _, row := range encountersTable {
        rangeLabel := ""
        if row.MaxRoll == 0 {
            rangeLabel = fmt.Sprintf("%d+", row.MinRoll) // 9+
        } else if row.MinRoll == row.MaxRoll {
            rangeLabel = fmt.Sprintf("%d", row.MinRoll)  // single value
        } else {
            rangeLabel = fmt.Sprintf("%d-%d", row.MinRoll, row.MaxRoll)
        }
        sb.WriteString(fmt.Sprintf("%-8s | %-40s | %-13s | %s\n",
            rangeLabel, row.Encounter, row.PerceptionDC, row.Number))
    }
    return sb.String()
}

// findEncounterEntry matches finalRoll to the correct row in the table.
func findEncounterEntry(roll int) EncountersEntry {
    for _, row := range encountersTable {
        // If row.MaxRoll < row.MinRoll, treat it as open-ended "9+"
        if row.MaxRoll < row.MinRoll {
            if roll >= row.MinRoll {
                return row
            }
        } else {
            if roll >= row.MinRoll && roll <= row.MaxRoll {
                return row
            }
        }
    }
    // Fallback if out of range
    return EncountersEntry{
        MinRoll: roll, MaxRoll: roll,
        Encounter: "???",
        PerceptionDC: "--",
        Number: "--",
    }
}

func (r *encountersInDavokar) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *encountersInDavokar) SubCategory() string {
    return "Davokar"
}
