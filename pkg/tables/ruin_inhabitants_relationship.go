package tables

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
)

// InhabitantRelationship represents a single row in the Inhabitants Relationship table.
type InhabitantRelationship struct {
    Roll    int
    Feature string
}

// ruinInhabitantsRelationshipTable holds all entries.
var ruinInhabitantsRelationshipTable = []InhabitantRelationship{
    {1, "The groups are bitter enemies."},
    {2, "The groups are unaware of each other."},
    {3, "The groups are aware of each other, but do not interact."},
    {4, "One group has been subjugated by the other."},
    {5, "The groups work together to an extent, but there is tension between them."},
    {6, "The groups work together harmoniously."},
}

// ruinInhabitantsRelationship implements the Table interface
type ruinInhabitantsRelationship struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinInhabitantsRelationship{})
}

// Name returns the table name
func (r *ruinInhabitantsRelationship) Name() string {
    return "ruin_inhabitants_relationship"
}

// GetRandomEntry returns a random row
func (r *ruinInhabitantsRelationship) GetRandomEntry() string {
    roll, err := dice.RollDice("1d6")
    if err != nil {
        log.Printf("Error rolling 1d6 for ruin inhabitants relationship: %v", err)
        roll = 1
    }
    // Convert the dice result (1-6) to a zero-based index.
    index := roll - 1
    if index < 0 || index >= len(ruinInhabitantsRelationshipTable) {
        index = 0
    }
    entry := ruinInhabitantsRelationshipTable[index]
    return fmt.Sprintf("Relationship: %s", entry.Feature)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinInhabitantsRelationship) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("INHABITANTS RELATIONSHIP (1d6)\n")
    sb.WriteString("-------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Relationship"))
    sb.WriteString("-------------------------------------------\n")

    for _, row := range ruinInhabitantsRelationshipTable {
        sb.WriteString(fmt.Sprintf("%-5d | %s\n", row.Roll, row.Feature))
    }

    return sb.String()
}

func (r *ruinInhabitantsRelationship) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinInhabitantsRelationship) SubCategory() string {
    return "Ruin"
}
