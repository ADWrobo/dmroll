package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
    "dmroll/pkg/table_registry"
)

// RuinEntryway represents a single row in the Entryways to Other Levels table.
type RuinEntryway struct {
    Roll     int
    Entryway string
}

// ruinEntrywaysToOtherLevelsTable holds all entries.
var ruinEntrywaysToOtherLevelsTable = []RuinEntryway{
    {1, "Intact stone staircase: difficult to destroy (damage threshold 20, AC 17, 120 hit points)"},
    {2, "Intact wooden staircase: can be destroyed (damage threshold 10, AC 15, 60 hit points)"},
    {3, "Damaged staircase: easy to destroy (AC 13, 30 hit points)"},
    {4, "Destroyed staircase: replaced by rope ladder/bridge (AC 10, 10 hit points)"},
    {5, "Completely destroyed staircase"},
    {6, "Hole in the ground/ceiling"},
    {7, "Hole in the ground/ceiling, with a ladder which can easily be pulled up/down"},
    {8, "Roll twice, ignoring further results of 8"},
}

// ruinEntrywaysToOtherLevels implements the Table interface
type ruinEntrywaysToOtherLevels struct{}

// Register the table automatically
func init() {
    table_registry.RegisterTable(&ruinEntrywaysToOtherLevels{})
}

// Name returns the table name
func (r *ruinEntrywaysToOtherLevels) Name() string {
    return "ruin_entryways_to_other_levels"
}

// GetRandomEntry returns a random row
func (r *ruinEntrywaysToOtherLevels) GetRandomEntry() string {
    roll, err := dice.RollDice("1d8")
    if err != nil {
        log.Printf("Error rolling 1d8 for ruin entryways to other levels: %v", err)
        roll = 1
    }
    // Convert the dice result (1-8) to a zero-based index.
    index := roll - 1
    if index < 0 || index >= len(ruinEntrywaysToOtherLevelsTable) {
        index = 0
    }
    entry := ruinEntrywaysToOtherLevelsTable[index]
    return fmt.Sprintf("Entryway: %s", entry.Entryway)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinEntrywaysToOtherLevels) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("ENTRYWAYS TO OTHER LEVELS (1d8)\n")
    sb.WriteString("-------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Entryway"))
    sb.WriteString("-------------------------------------------------------------\n")

    for _, row := range ruinEntrywaysToOtherLevelsTable {
        sb.WriteString(fmt.Sprintf("%-5d | %s\n", row.Roll, row.Entryway))
    }

    return sb.String()
}

func (r *ruinEntrywaysToOtherLevels) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinEntrywaysToOtherLevels) SubCategory() string {
    return "Ruin"
}
