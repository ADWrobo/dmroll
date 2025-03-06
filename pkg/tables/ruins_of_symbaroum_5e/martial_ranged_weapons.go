package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// MartialRangedWeaponEntry represents a single row in the Martial Ranged Weapons table.
type MartialRangedWeaponEntry struct {
    Name       string
    Cost       string
    Damage     string
    Weight     string
    Properties string
}

var martialRangedWeaponsTable = []MartialRangedWeaponEntry{
    {"Arbalest", "50 thaler", "1d10 piercing", "18 lb.", "Ammunition (range 100/400), deep impact, heavy, loading, two-handed"},
    {"Blowpipe", "10 thaler", "1 piercing", "1 lb.", "Ammunition (range 25/100), loading"},
    {"Bolas", "5 thaler", "1 bludgeoning", "2 lb.", "Finesse, restraining, thrown (range 30/90)"},
    {"Composite Bow", "30 thaler", "1d8 piercing", "3 lb.", "Ammunition (range 80/320), deep impact, two-handed"},
    {"Crossbow, hand", "75 thaler", "1d6 piercing", "3 lb.", "Ammunition (range 30/120), light, loading"},
    {"Crossbow, heavy", "25 thaler", "1d10 piercing", "18 lb.", "Ammunition (range 100/400), heavy, loading, two-handed"},
    {"Crossbow, repeating", "100 thaler", "1d8 piercing", "20 lb.", "Ammunition (range 80/320), two-handed"},
    {"Longbow", "20 thaler", "1d8 piercing", "2 lb.", "Ammunition (range 150/600), heavy, two-handed"},
    {"Net", "1 thaler", "—", "3 lb.", "Ensnaring, special, thrown (range 5/15)"},
    {"Throwing Wing", "10 thaler", "1d6 bludgeoning", "1 lb.", "Returning, thrown (range 30/90)"},
}

type martialRangedWeapons struct{}

func init() {
    table_registry.RegisterTable(&martialRangedWeapons{})
}

func (m *martialRangedWeapons) Name() string {
    return "martial_ranged_weapons"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (m *martialRangedWeapons) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p martial_ranged_weapons' to view it."
}

// GetFormatted returns an ASCII-formatted table of martial ranged weapons.
func (m *martialRangedWeapons) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("MARTIAL RANGED WEAPONS (PG 165)\n")
    sb.WriteString("-----------------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %-12s | %-15s | %-6s | %s\n", "Name", "Cost", "Damage", "Wt.", "Properties"))
    sb.WriteString("-----------------------------------------------------------------------------------------------------------\n")

    for _, row := range martialRangedWeaponsTable {
        sb.WriteString(fmt.Sprintf("%-20s | %-12s | %-15s | %-6s | %s\n", row.Name, row.Cost, row.Damage, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (m *martialRangedWeapons) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (m *martialRangedWeapons) SubCategory() string {
    return "Weapons"
}
