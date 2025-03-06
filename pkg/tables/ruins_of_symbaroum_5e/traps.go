package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// TrapEntry represents a single row in the Traps table.
type TrapEntry struct {
    Name      string
    Difficulty string
    Damage    string
    Cost      string
}

// TrapCategory holds a category name and its corresponding trap entries.
type TrapCategory struct {
    Category string
    Entries  []TrapEntry
}

var trapsTable = []TrapCategory{
    {"Alchemical Mines", []TrapEntry{
        {"Weak", "DC 13", "2d6 fire", "10 thaler"},
        {"Moderate", "DC 15", "4d6 fire", "20 thaler"},
        {"Strong", "DC 17", "6d6 fire", "30 thaler"},
    }},
    {"Mechanical Traps", []TrapEntry{
        {"Weak (Snare)", "DC 10", "1d12 bludgeoning", "5 shillings"},
        {"Moderate (Bear Trap)", "DC 13", "2d12 piercing", "1 thaler"},
        {"Strong (Dragon Trap)", "DC 17", "3d12 bludgeoning", "2 thaler"},
    }},
}

type traps struct{}

func init() {
    table_registry.RegisterTable(&traps{})
}

func (t *traps) Name() string {
    return "traps"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (t *traps) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p traps' to view it."
}

// GetFormatted returns an ASCII-formatted table of Traps with category headers.
func (t *traps) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("TRAPS (PG 186)\n")
    sb.WriteString("================================================================\n")

    for _, category := range trapsTable { // Ensures ordered iteration
        sb.WriteString(fmt.Sprintf("\n%s\n", strings.ToUpper(category.Category)))
        sb.WriteString("----------------------------------------------------------------\n")
        sb.WriteString(fmt.Sprintf("%-25s | %-10s | %-15s | %s\n", "Name", "Difficulty", "Damage", "Cost"))
        sb.WriteString("----------------------------------------------------------------\n")

        for _, row := range category.Entries {
            sb.WriteString(fmt.Sprintf("%-25s | %-10s | %-15s | %s\n", row.Name, row.Difficulty, row.Damage, row.Cost))
        }
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (t *traps) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (t *traps) SubCategory() string {
    return "Equipment and Services"
}
