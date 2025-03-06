package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// BuildingEntry represents a single row in the Buildings table.
type BuildingEntry struct {
    Name        string
    Cost 		string
}

var buildingsTable = []BuildingEntry{
    {"Croft", "20 thaler"},
    {"Farm", "150 thaler"},
    {"Watchtower (wood)", "200 thaler"},
    {"Watchtower (stone)", "500 thaler"},
    {"Estate", "1000 thaler"},
    {"Fort (wood)", "600 thaler"},
    {"Fort (stone)", "2000 thaler"},
    {"Keep", "5000 thaler"},
    {"Castle", "10000 thaler"},
}

type buildings struct{}

func init() {
    table_registry.RegisterTable(&buildings{})
}

func (b *buildings) Name() string {
    return "buildings"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (b *buildings) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p buildings' to view it."
}

// GetFormatted returns an ASCII-formatted table of Buildings.
func (b *buildings) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("BUILDINGS (PG 173)\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %s\n", "Name", "Cost"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range buildingsTable {
        sb.WriteString(fmt.Sprintf("%-20s | %s\n", row.Name, row.Cost))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (b *buildings) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (b *buildings) SubCategory() string {
    return "Equipment and Services"
}
