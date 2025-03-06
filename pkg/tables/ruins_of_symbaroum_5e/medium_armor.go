package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// MediumArmorEntry represents a single row in the Medium Armor table.
type MediumArmorEntry struct {
    Name       string
    Cost       string
    ArmorClass string
    Weight     string
    Properties string
}

var mediumArmorTable = []MediumArmorEntry{
    {"Chain Shirt", "50 thaler", "13 + Dexterity modifier (max 2)", "20 lb.", "-"},
    {"Crow Armor", "5 thaler", "14 + Dexterity modifier (max 2)", "30 lb.", "Cumbersome"},
    {"Double Chain Mail", "125 thaler", "14 + Dexterity modifier (max 2)", "40 lb.", "Noisy"},
    {"Lacquered Silk Cuirass", "60 thaler", "14 + Dexterity modifier (max 2)", "18 lb.", "-"},
    {"Laminated Armor", "150 thaler", "15 + Dexterity modifier (max 2)", "40 lb.", "-"},
    {"Scale Mail", "50 thaler", "15 + Dexterity modifier (max 2)", "45 lb.", "Noisy"},
}

type mediumArmor struct{}

func init() {
    table_registry.RegisterTable(&mediumArmor{})
}

func (m *mediumArmor) Name() string {
    return "medium_armor"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (m *mediumArmor) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p medium_armor' to view it."
}

// GetFormatted returns an ASCII-formatted table of Medium Armor.
func (m *mediumArmor) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("MEDIUM ARMOR (PG 170)\n")
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-25s | %-10s | %-30s | %-10s | %s\n", "Name", "Cost", "Armor Class", "Weight", "Properties"))
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")

    for _, row := range mediumArmorTable {
        sb.WriteString(fmt.Sprintf("%-25s | %-10s | %-30s | %-10s | %s\n", row.Name, row.Cost, row.ArmorClass, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (m *mediumArmor) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (m *mediumArmor) SubCategory() string {
    return "Armor"
}
