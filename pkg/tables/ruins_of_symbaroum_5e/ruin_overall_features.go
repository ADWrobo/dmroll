package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
    "dmroll/pkg/table_registry"
)

// RuinOverallFeature represents a single row in the Ruin Overall Features table.
type RuinOverallFeature struct {
    Roll    int
    Feature string
}

// ruinOverallFeaturesTable holds all entries.
var ruinOverallFeaturesTable = []RuinOverallFeature{
    {1, "Water damaged"},
    {2, "Dusty and desolate"},
    {3, "Overgrown"},
    {4, "Untouched and ominously clean"},
    {5, "Inexplicably cold/hot"},
    {6, "Darkened (all light sources are diminished, half normal range)"},
}

// ruinOverallFeatures implements the Table interface
type ruinOverallFeatures struct{}

// Register the table automatically
func init() {
    table_registry.RegisterTable(&ruinOverallFeatures{})
}

// Name returns the table name
func (r *ruinOverallFeatures) Name() string {
    return "ruin_overall_features"
}

// GetRandomEntry returns a random row
func (r *ruinOverallFeatures) GetRandomEntry() string {
    roll, err := dice.RollDice("1d6")
    if err != nil {
        log.Printf("Error rolling 1d6 for ruin overall features: %v", err)
        roll = 1
    }
    // Adjust the 1-based roll to a 0-based index.
    index := roll - 1
    // Safety check in case the table isn't exactly 6 entries.
    if index < 0 || index >= len(ruinOverallFeaturesTable) {
        index = 0
    }
    entry := ruinOverallFeaturesTable[index]
    return fmt.Sprintf("Feature: %s", entry.Feature)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinOverallFeatures) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("RUIN OVERALL FEATURES (1d6)\n")
    sb.WriteString("-----------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Feature"))
    sb.WriteString("-----------------------------\n")

    for _, row := range ruinOverallFeaturesTable {
        sb.WriteString(fmt.Sprintf("%-5d | %s\n", row.Roll, row.Feature))
    }

    return sb.String()
}

func (r *ruinOverallFeatures) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinOverallFeatures) SubCategory() string {
    return "Ruin"
}
