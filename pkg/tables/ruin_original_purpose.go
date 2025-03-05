package tables

import (
    "fmt"
    "strings"
    "log"

    "dmroll/pkg/dice"
)

// RuinOriginalPurposeEntry represents a single row in the Ruin Original Purpose table.
type RuinOriginalPurposeEntry struct {
    Roll      int
    Function  string
    Details   [4]string
}

// ruinOriginalPurposeTable holds all entries.
var ruinOriginalPurposeTable = []RuinOriginalPurposeEntry{
    {1, "Fortification", [4]string{"Small fort", "Stone fort", "Fortress ruin", "Massive fortress (walls and moat)"}},
    {2, "Prison", [4]string{"Mighty individual", "Dangerous artifact", "Hungry monster", "Shackled deity of nature"}},
    {3, "Pyramid", [4]string{"Ziggurat", "Ziggurat", "Traditional", "1d4 Ziggurats"}},
    {4, "Labyrinth", [4]string{"Garden", "Entertainment", "Used for experiments", "Used for experiments"}},
    {5, "Mine", [4]string{"Precious metal", "Gems", "Alchemical compounds", "Crystallized Corruption"}},
    {6, "Temple", [4]string{"Ancestor worship", "Spider worship", "Serpent temple", "Hero cult"}},
    {7, "Seat of Power", [4]string{"Magician’s tower", "Priest’s estate", "Royal palace", "Imperial pleasure palace"}},
    {8, "Tomb", [4]string{"Lord/lady", "Revered mystic", "Petty king", "Member of the imperial or royal family"}},
    {9, "Exhibition Halls", [4]string{"Art & sculptures", "Antique jewelry", "Historical objects", "Monster display"}},
    {10, "Ritual Chamber", [4]string{"Funerals", "Daemon summoning", "Fleshcrafters", "Daemon exaltation"}},
}

// ruinOriginalPurpose implements the Table interface
type ruinOriginalPurpose struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinOriginalPurpose{})
}

// Name returns the table name
func (r *ruinOriginalPurpose) Name() string {
    return "ruin_original_purpose"
}

// GetRandomEntry returns a random row
func (r *ruinOriginalPurpose) GetRandomEntry() string {
    // Roll 1d10 for the main entry
    roll, err := dice.RollDice("1d10")
    if err != nil {
        log.Printf("Error rolling 1d10 for ruin original purpose: %v", err)
        roll = 1
    }
    // Convert to zero-based index
    index := roll - 1
    if index < 0 || index >= len(ruinOriginalPurposeTable) {
        index = 0
    }
    entry := ruinOriginalPurposeTable[index]

    // Roll 1d4 for selecting one of the details
    subRoll, err := dice.RollDice("1d4")
    if err != nil {
        log.Printf("Error rolling 1d4 for ruin original purpose details: %v", err)
        subRoll = 1
    }
    subIndex := subRoll - 1
    if subIndex < 0 || subIndex >= len(entry.Details) {
        subIndex = 0
    }

    return fmt.Sprintf("%s | Detail: %s", entry.Function, entry.Details[subIndex])
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinOriginalPurpose) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("RUIN'S ORIGINAL PURPOSE (1d10 + 1d4)\n")
    sb.WriteString("----------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %-15s | %-15s | %-15s | %-15s | %-15s\n",
        "Roll", "Function", "Detail 1", "Detail 2", "Detail 3", "Detail 4"))
    sb.WriteString("----------------------------------------------------------------------------\n")

    for _, row := range ruinOriginalPurposeTable {
        sb.WriteString(fmt.Sprintf("%-5d | %-15s | %-15s | %-15s | %-15s | %-15s\n",
            row.Roll, row.Function, row.Details[0], row.Details[1], row.Details[2], row.Details[3]))
    }

    return sb.String()
}

// Category returns the table category
func (r *ruinOriginalPurpose) Category() string {
    return "Ruin"
}
