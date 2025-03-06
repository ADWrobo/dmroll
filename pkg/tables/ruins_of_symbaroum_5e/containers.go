package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// ContainerEntry represents a single row in the Containers table.
type ContainerEntry struct {
    Name        string
    Cost		string
}

var containersTable = []ContainerEntry{
    {"Backpack", "1 thaler"},
    {"Barrel", "4 ortegs"},
    {"Basket", "2 ortegs"},
    {"Belt pouch", "5 ortegs"},
    {"Chest, small", "3 shillings"},
    {"Chest, large", "1 thaler"},
    {"Clay pitcher", "5 ortegs"},
    {"Coin purse", "3 ortegs"},
    {"Decorated box", "2-10 thaler"},
    {"Glass vial", "1 thaler"},
    {"Knapsack", "2 shillings"},
    {"Quiver", "5 shillings"},
    {"Sack", "2 ortegs"},
}

type containers struct{}

func init() {
    table_registry.RegisterTable(&containers{})
}

func (c *containers) Name() string {
    return "containers"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (c *containers) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p containers' to view it."
}

// GetFormatted returns an ASCII-formatted table of Containers.
func (c *containers) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("CONTAINERS (PG 173)\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %s\n", "Name", "Cost"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range containersTable {
        sb.WriteString(fmt.Sprintf("%-20s | %s\n", row.Name, row.Cost))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (c *containers) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (c *containers) SubCategory() string {
    return "Equipment and Services"
}
