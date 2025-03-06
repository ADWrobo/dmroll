package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// MartialMeleeWeaponEntry represents a single row in the Martial Melee Weapons table.
type MartialMeleeWeaponEntry struct {
    Name       string
    Cost       string
    Damage     string
    Weight     string
    Properties string
}

var martialMeleeWeaponsTable = []MartialMeleeWeaponEntry{
    {"Assassin’s Blade", "20 thaler", "1d6 piercing", "1 lb.", "Concealed, finesse"},
    {"Axe", "20 thaler", "1d8 slashing", "3 lb.", "—"},
    {"Chain Staff", "20 thaler", "1d8 bludgeoning", "3 lb.", "Ensnaring, reach"},
    {"Crow’s Beak", "5 thaler", "1d8 piercing", "3 lb.", "Deep impact"},
    {"Double-axe", "30 thaler", "1d12 slashing", "7 lb.", "Massive, two-handed"},
    {"Estoc", "30 thaler", "1d8 piercing", "3 lb.", "Deep impact, finesse"},
    {"Executioner’s Sword", "60 thaler", "1d12 slashing", "4 lb.", "Massive, two-handed"},
    {"Fencing Sword", "25 thaler", "1d8 piercing", "2 lb.", "Finesse"},
    {"Flail", "10 thaler", "1d8 bludgeoning", "6 lb.", "Ensnaring"},
    {"Grappling Axe", "30 thaler", "1d8 slashing", "4 lb.", "Versatile (1d10)"},
    {"Great Flail", "20 thaler", "1d12 bludgeoning", "8 lb.", "Ensnaring, heavy, two-handed"},
    {"Greatsword", "50 thaler", "2d6 slashing", "6 lb.", "Heavy, two-handed"},
    {"Halberd", "20 thaler", "1d10 slashing", "10 lb.", "Heavy, reach, two-handed"},
    {"Lance", "10 thaler", "1d12 piercing", "6 lb.", "Heavy, reach, special"},
    {"Long Hammer", "25 thaler", "1d8 bludgeoning", "6 lb.", "—"},
    {"Longsword", "15 thaler", "1d8 slashing", "3 lb.", "Versatile (1d10)"},
    {"Maul", "10 thaler", "2d6 bludgeoning", "10 lb.", "Heavy, two-handed"},
    {"Parrying Dagger", "5 thaler", "1d4 piercing", "1 lb.", "Balanced, light"},
    {"Pike", "15 thaler", "1d10 piercing", "18 lb.", "Heavy, reach, two-handed"},
    {"Shortsword", "10 thaler", "1d6 piercing", "2 lb.", "Finesse, light"},
    {"Stiletto", "5 thaler", "1d4 piercing", "2 lb.", "Deep impact, finesse"},
    {"Whip", "2 thaler", "1d4 slashing", "3 lb.", "Ensnaring, finesse, reach"},
}

type martialMeleeWeapons struct{}

func init() {
    table_registry.RegisterTable(&martialMeleeWeapons{})
}

func (m *martialMeleeWeapons) Name() string {
    return "martial_melee_weapons"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (m *martialMeleeWeapons) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p martial_melee_weapons' to view it."
}

// GetFormatted returns an ASCII-formatted table of martial melee weapons.
func (m *martialMeleeWeapons) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("MARTIAL MELEE WEAPONS (PG 163)\n")
    sb.WriteString("------------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %-12s | %-15s | %-6s | %s\n", "Name", "Cost", "Damage", "Wt.", "Properties"))
    sb.WriteString("------------------------------------------------------------------------------------------------------\n")

    for _, row := range martialMeleeWeaponsTable {
        sb.WriteString(fmt.Sprintf("%-20s | %-12s | %-15s | %-6s | %s\n", row.Name, row.Cost, row.Damage, row.Weight, row.Properties))
    }

    return sb.String()
}

func (m *martialMeleeWeapons) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (m *martialMeleeWeapons) SubCategory() string {
    return "Weapons"
}
