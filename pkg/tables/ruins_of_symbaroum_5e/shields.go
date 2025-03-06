package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// ShieldEntry represents a single row in the Shields table.
type ShieldEntry struct {
    Name       string
    Cost       string
    ACModifier string
    Weight     string
    Properties string
}

var shieldsTable = []ShieldEntry{
    {"Buckler", "4 thaler", "+1", "2 lb.", "-"},
    {"Shield", "10 thaler", "+2", "6 lb.", "-"},
}

type shields struct{}

func init() {
    table_registry.RegisterTable(&shields{})
}

func (s *shields) Name() string {
    return "shields"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (s *shields) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p shields' to view it."
}

// GetFormatted returns an ASCII-formatted table of Shields.
func (s *shields) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("SHIELDS\n")
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-15s | %-10s | %-10s | %-10s | %s\n", "Name", "Cost", "AC Modifier", "Weight", "Properties"))
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")

    for _, row := range shieldsTable {
        sb.WriteString(fmt.Sprintf("%-15s | %-10s | %-10s | %-10s | %s\n", row.Name, row.Cost, row.ACModifier, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (s *shields) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (s *shields) SubCategory() string {
    return "Armor"
}
