package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// LightArmorEntry represents a single row in the Light Armor table.
type LightArmorEntry struct {
    Name       string
    Cost       string
    ArmorClass string
    Weight     string
    Properties string
}

var lightArmorTable = []LightArmorEntry{
    {"Blessed Robe", "10 thaler", "11 + Dexterity modifier", "5 lb.", "-"},
    {"Concealed Armor", "20 thaler", "11 + Dexterity modifier", "4 lb.", "Concealable"},
    {"Order Cloak", "10 thaler", "11 + Dexterity modifier", "5 lb.", "-"},
    {"Skald’s Cuirass", "30 thaler", "12 + Dexterity modifier", "20 lb.", "-"},
    {"Studded Leather", "20 thaler", "12 + Dexterity modifier", "15 lb.", "-"},
    {"Witch Gown", "10 thaler", "11 + Dexterity modifier", "5 lb.", "-"},
    {"Wolf Skin", "1 thaler", "12 + Dexterity modifier", "15 lb.", "Cumbersome"},
    {"Woven Silk", "50 thaler", "12 + Dexterity modifier", "6 lb.", "-"},
}

type lightArmor struct{}

func init() {
    table_registry.RegisterTable(&lightArmor{})
}

func (l *lightArmor) Name() string {
    return "light_armor"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (l *lightArmor) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p light_armor' to view it."
}

// GetFormatted returns an ASCII-formatted table of Light Armor.
func (l *lightArmor) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("LIGHT ARMOR\n")
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %-10s | %-25s | %-10s | %s\n", "Name", "Cost", "Armor Class", "Weight", "Properties"))
    sb.WriteString("-----------------------------------------------------------------------------------------------------\n")

    for _, row := range lightArmorTable {
        sb.WriteString(fmt.Sprintf("%-20s | %-10s | %-25s | %-10s | %s\n", row.Name, row.Cost, row.ArmorClass, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (l *lightArmor) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (l *lightArmor) SubCategory() string {
    return "Armor"
}
