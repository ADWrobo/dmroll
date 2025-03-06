package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// SimpleRangedWeaponEntry represents a single row in the Simple Ranged Weapons table.
type SimpleRangedWeaponEntry struct {
    Name       string
    Cost       string
    Damage     string
    Weight     string
    Properties string
}

var simpleRangedWeaponsTable = []SimpleRangedWeaponEntry{
    {"Crossbow, light", "15 thaler", "1d8 piercing", "5 lb.", "Ammunition (range 80/320), loading, two-handed"},
    {"Dart", "5 ortegs", "1d4 piercing", "1/4 lb.", "Finesse, thrown (range 20/60)"},
    {"Horseman’s Bow", "10 thaler", "1d6 piercing", "2 lb.", "Ammunition (range 80/320), two-handed"},
    {"Sling", "1 shilling", "1d4 bludgeoning", "—", "Ammunition (range 30/120)"},
    {"Spear Sling", "5 thaler", "1d6 piercing", "2 lb.", "Ammunition (range 40/160), deep impact"},
}

type simpleRangedWeapons struct{}

func init() {
    table_registry.RegisterTable(&simpleRangedWeapons{})
}

func (s *simpleRangedWeapons) Name() string {
    return "simple_ranged_weapons"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (s *simpleRangedWeapons) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p simple_ranged_weapons' to view it."
}

// GetFormatted returns an ASCII-formatted table of simple ranged weapons.
func (s *simpleRangedWeapons) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("SIMPLE RANGED WEAPONS (PG 162)\n")
    sb.WriteString("------------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %-12s | %-15s | %-6s | %s\n", "Name", "Cost", "Damage", "Wt.", "Properties"))
    sb.WriteString("------------------------------------------------------------------------------------------------------\n")

    for _, row := range simpleRangedWeaponsTable {
        sb.WriteString(fmt.Sprintf("%-20s | %-12s | %-15s | %-6s | %s\n", row.Name, row.Cost, row.Damage, row.Weight, row.Properties))
    }

    return sb.String()
}

func (s *simpleRangedWeapons) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (s *simpleRangedWeapons) SubCategory() string {
	return "Weapons"
}
