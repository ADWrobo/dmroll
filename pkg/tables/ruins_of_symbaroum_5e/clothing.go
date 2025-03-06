package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// ClothingEntry represents a single row in the Clothing table.
type ClothingEntry struct {
    Name string
    Cost string
}

var clothingTable = []ClothingEntry{
    {"Rags", "5 ortegs"},
    {"Simple garb", "1 shilling"},
    {"Artisan work-clothes", "5 shillings"},
    {"Burgher’s dress-clothes", "1 thaler"},
    {"Priest’s vestments", "5 thaler or more"},
    {"Noble’s outfit", "10 thaler or more"},
    {"Boots", "1 shilling"},
    {"Cap", "2-4 ortegs"},
    {"Cloak", "2-6 ortegs"},
    {"Coat", "5-10 ortegs"},
    {"Dress", "1-10 shillings"},
    {"Gown", "1-5 ortegs"},
    {"Hat", "2-4 ortegs"},
    {"Mask", "1-5 ortegs"},
    {"Pants", "1-5 ortegs"},
    {"Robe", "1-5 shillings"},
    {"Scarf", "1-2 ortegs"},
    {"Shirt", "1-4 ortegs"},
    {"Skirt", "5-10 ortegs"},
    {"Tunic", "2-5 ortegs"},
}

type clothing struct{}

func init() {
    table_registry.RegisterTable(&clothing{})
}

func (c *clothing) Name() string {
    return "clothing"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (c *clothing) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p clothing' to view it."
}

// GetFormatted returns an ASCII-formatted table of Clothing.
func (c *clothing) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("CLOTHING (PG 174)\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %s\n", "Name", "Cost"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range clothingTable {
        sb.WriteString(fmt.Sprintf("%-20s | %s\n", row.Name, row.Cost))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (c *clothing) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (c *clothing) SubCategory() string {
    return "Equipment and Services"
}
