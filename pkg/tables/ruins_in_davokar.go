package tables

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
)

// RuinEntry represents one “row” in the Ruins in Davokar table.
type RuinEntry struct {
    MinRoll        int
    MaxRoll        int
    Ruin           string
    InvestigationDC string
    MaximumFinds   string
}

var ruinsTable = []RuinEntry{
    {1, 7,  "None",                               "--",      "--"},
    {8, 10, "Completely crumbled or already ransacked", "--", "--"},
    {11, 12, "Small and badly damaged",            "13",     "1d4+2"},
    {13, 14, "Small and dilapidated",             "11",     "1d6+2"},
    {15, 16, "Small and well preserved",          "10",     "1d8+2"},
    {17, 18, "Medium and badly damaged",          "17",     "2d6+2"},
    {19, 19, "Medium and dilapidated",            "15",     "2d8+2"},
    {20, 20, "Medium and well preserved",         "14",     "2d10+2"},
    {21, 21, "Grand and badly damaged",           "21",     "3d8+2"},
    {22, 22, "Grand and dilapidated",             "19",     "3d10+2"},
    // 23+ => open-ended
    {23, 0,  "Grand and well preserved",          "18",     "3d12+2"},
}

type ruinsInDavokar struct{}

func init() {
    RegisterTable(&ruinsInDavokar{})
}

func (r *ruinsInDavokar) Name() string {
    return "ruins_in_davokar"
}

// GetRandomEntry: prompt user for a location mod, roll 1d20, add mod, find the row.
func (r *ruinsInDavokar) GetRandomEntry() string {
    mod := PromptForDavokarModifier() // <-- from davokar_modifiers.go
    fmt.Println() // optional blank line

    baseRoll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling dice for Ruins in Davokar: %v\n", err)
        baseRoll = 1
    }
    finalRoll := baseRoll + mod
    entry := findRuinEntry(finalRoll)

    return fmt.Sprintf(
        "Ruins in Davokar result: %d (1d20+%d)\nRuin: %s\nInvestigation DC: %s\nMaximum Finds: %s",
        finalRoll, mod, entry.Ruin, entry.InvestigationDC, entry.MaximumFinds,
    )
}

// GetFormatted prints the entire table in ASCII.
func (r *ruinsInDavokar) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("RUINS IN DAVOKAR (1d20 + Mods)\n")
    sb.WriteString("----------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-10s | %-40s | %-15s | %s\n",
        "Roll Range", "Ruin", "Investigation DC", "Max Finds"))
    sb.WriteString("----------------------------------------------------------------------------\n")

    for _, row := range ruinsTable {
        rangeLabel := ""
        if row.MaxRoll == 0 {
            // e.g. 23+
            rangeLabel = fmt.Sprintf("%d+", row.MinRoll)
        } else if row.MinRoll == row.MaxRoll {
            rangeLabel = fmt.Sprintf("%d", row.MinRoll)
        } else {
            rangeLabel = fmt.Sprintf("%d-%d", row.MinRoll, row.MaxRoll)
        }
        sb.WriteString(fmt.Sprintf("%-10s | %-40s | %-15s | %s\n",
            rangeLabel, row.Ruin, row.InvestigationDC, row.MaximumFinds))
    }
    return sb.String()
}

// findRuinEntry picks the correct row based on final roll.
func findRuinEntry(roll int) RuinEntry {
    for _, row := range ruinsTable {
        // If MaxRoll < MinRoll, treat as open-ended
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
    // Fallback in case we got an extreme roll or something unexpected
    return RuinEntry{
        MinRoll: roll, MaxRoll: roll,
        Ruin: "???",
        InvestigationDC: "--",
        MaximumFinds: "--",
    }
}

func (r *ruinsInDavokar) Category() string {
    return "Davokar"
}