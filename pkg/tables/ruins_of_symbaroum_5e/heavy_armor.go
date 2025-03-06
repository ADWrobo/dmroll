package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// HeavyArmorEntry represents a single row in the Heavy Armor table.
type HeavyArmorEntry struct {
    Name       string
    Cost       string
    ArmorClass string
    Weight     string
    Properties string
}

var heavyArmorTable = []HeavyArmorEntry{
    {"Chain and Plate", "250 thaler", "15", "50 lb.", "-"},
    {"Field Armor", "500 thaler", "17", "70 lb.", "Cumbersome, Weighty (13)"},
    {"Field Armor of the Pansars", "750 thaler", "18", "70 lb.", "Noisy, Weighty (15)"},
    {"Full Plate", "500 thaler", "16", "65 lb.", "Noisy, Weighty (15)"},
}

type heavyArmor struct{}

func init() {
    table_registry.RegisterTable(&heavyArmor{})
}

func (h *heavyArmor) Name() string {
    return "heavy_armor"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (h *heavyArmor) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p heavy_armor' to view it."
}

// GetFormatted returns an ASCII-formatted table of Heavy Armor.
func (h *heavyArmor) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("HEAVY ARMOR (PG 170)\n")
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-30s | %-10s | %-10s | %-10s | %s\n", "Name", "Cost", "Armor Class", "Weight", "Properties"))
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")

    for _, row := range heavyArmorTable {
        sb.WriteString(fmt.Sprintf("%-30s | %-10s | %-10s | %-10s | %s\n", row.Name, row.Cost, row.ArmorClass, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (h *heavyArmor) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (h *heavyArmor) SubCategory() string {
    return "Armor"
}
