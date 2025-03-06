package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// DonningDoffingEntry represents a single row in the Donning and Doffing Armor table.
type DonningDoffingEntry struct {
    Category string
    Don      string
    Doff     string
}

var donningDoffingArmorTable = []DonningDoffingEntry{
    {"Light Armor", "1 minute", "1 minute"},
    {"Medium Armor", "5 minutes", "5 minutes"},
    {"Heavy Armor", "10 minutes", "5 minutes"},
    {"Buckler", "Object Interaction*", "Gesture*"},
    {"Shield", "Action", "Action"},
}

type donningDoffingArmor struct{}

func init() {
    table_registry.RegisterTable(&donningDoffingArmor{})
}

func (d *donningDoffingArmor) Name() string {
    return "donning_and_doffing_armor"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (d *donningDoffingArmor) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p donning_and_doffing_armor' to view it."
}

// GetFormatted returns an ASCII-formatted table of Donning and Doffing Armor.
func (d *donningDoffingArmor) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("DONNING AND DOFFING ARMOR (PG 171)\n")
    sb.WriteString("-----------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %-15s | %s\n", "Category", "Don", "Doff"))
    sb.WriteString("-----------------------------------------------------\n")

    for _, row := range donningDoffingArmorTable {
        sb.WriteString(fmt.Sprintf("%-20s | %-15s | %s\n", row.Category, row.Don, row.Doff))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (d *donningDoffingArmor) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (d *donningDoffingArmor) SubCategory() string {
    return "Armor"
}
