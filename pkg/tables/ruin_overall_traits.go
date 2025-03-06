package tables

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
)

// RuinOverallTrait represents a single row in the Ruin Overall Traits table.
type RuinOverallTrait struct {
    Roll    int
    Feature string
}

// ruinOverallTraitsTable holds all entries.
var ruinOverallTraitsTable = []RuinOverallTrait{
    {1, "Corruptive. Inside the ruin whenever you gain temporary Corruption you gain 1 more point than normal."},
    {2, "Sparse with treasure. You have disadvantage on treasure value rolls."},
    {3, "Crowded. Use 1d8 instead of 1d6 to determine the number of creatures in a room."},
    {4, "Desolate. Use 1d4 instead of 1d6 to determine the number of creatures in a room."},
    {5, "Rich with treasure. You have advantage on treasure value rolls."},
    {6, "Lesser corruption. Reduce any temporary Corruption gained within the ruin by 1 point (minimum 1 point gained)."},
}

// ruinOverallTraits implements the Table interface
type ruinOverallTraits struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinOverallTraits{})
}

// Name returns the table name
func (r *ruinOverallTraits) Name() string {
    return "ruin_overall_traits"
}

// GetRandomEntry returns a random row
func (r *ruinOverallTraits) GetRandomEntry() string {
    roll, err := dice.RollDice("1d6")
    if err != nil {
        log.Printf("Error rolling 1d6 for ruin overall traits: %v", err)
        roll = 1
    }
    // Convert the dice result (1-6) to a zero-based index.
    index := roll - 1
    // Safety check in case the table doesn't exactly have 6 entries.
    if index < 0 || index >= len(ruinOverallTraitsTable) {
        index = 0
    }
    entry := ruinOverallTraitsTable[index]
    return fmt.Sprintf("Trait: %s", entry.Feature)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinOverallTraits) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("RUIN OVERALL TRAITS (1d6)\n")
    sb.WriteString("--------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Feature"))
    sb.WriteString("--------------------------------------------------------------\n")

    for _, row := range ruinOverallTraitsTable {
        sb.WriteString(fmt.Sprintf("%-5d | %s\n", row.Roll, row.Feature))
    }

    return sb.String()
}

func (r *ruinOverallTraits) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinOverallTraits) SubCategory() string {
    return "Ruin"
}
