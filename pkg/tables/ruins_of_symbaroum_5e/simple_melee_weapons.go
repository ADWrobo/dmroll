package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// SimpleMeleeWeaponEntry represents a single row in the Simple Melee Weapons table.
type SimpleMeleeWeaponEntry struct {
    Name       string
    Cost       string
    Damage     string
    Weight     string
    Properties string
}

var simpleMeleeWeaponsTable = []SimpleMeleeWeaponEntry{
    {"Club", "1 shilling", "1d4 bludgeoning", "2 lb.", "Light"},
    {"Dagger", "2 thaler", "1d4 piercing", "1 lb.", "Finesse, light, thrown (range 20/60)"},
    {"Greatclub", "2 shillings", "1d8 bludgeoning", "10 lb.", "Two-handed"},
    {"Handaxe", "5 thaler", "1d6 slashing", "2 lb.", "Light, thrown (range 20/60)"},
    {"Javelin", "5 shillings", "1d6 piercing", "2 lb.", "Thrown (range 30/120)"},
    {"Light hammer", "2 thaler", "1d4 bludgeoning", "2 lb.", "Light, thrown (range 20/60)"},
    {"Mace", "5 thaler", "1d6 bludgeoning", "4 lb.", "—"},
    {"Quarterstaff", "2 shillings", "1d8 bludgeoning", "4 lb.", "Two-handed"},
    {"Sickle", "1 thaler", "1d4 slashing", "2 lb.", "Light"},
    {"Spear", "1 thaler", "1d6 piercing", "3 lb.", "Thrown (range 20/60), versatile (1d8)"},
}

type simpleMeleeWeapons struct{}

func init() {
    table_registry.RegisterTable(&simpleMeleeWeapons{})
}

func (s *simpleMeleeWeapons) Name() string {
    return "simple_melee_weapons"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (s *simpleMeleeWeapons) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p simple_melee_weapons' to view it."
}

// GetFormatted returns an ASCII-formatted table of simple melee weapons.
func (s *simpleMeleeWeapons) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("SIMPLE MELEE WEAPONS (PG 162)\n")
    sb.WriteString("------------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-15s | %-12s | %-15s | %-6s | %s\n", "Name", "Cost", "Damage", "Wt.", "Properties"))
    sb.WriteString("------------------------------------------------------------------------------------------------------\n")

    for _, row := range simpleMeleeWeaponsTable {
        sb.WriteString(fmt.Sprintf("%-15s | %-12s | %-15s | %-6s | %s\n", row.Name, row.Cost, row.Damage, row.Weight, row.Properties))
    }

    return sb.String()
}

func (s *simpleMeleeWeapons) Category() string {
    return "Ruins of Symbaroum 5E"
}
func (s *simpleMeleeWeapons) SubCategory() string {
    return "Weapons"
}
