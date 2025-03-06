package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"
    
    "dmroll/pkg/table_registry"
)

// AlchemicalWeaponEntry represents a single row in the Alchemical Weapons table.
type AlchemicalWeaponEntry struct {
    Name       string
    Cost       string
    Damage     string
    Weight     string
    Properties string
}

var alchemicalWeaponsTable = []AlchemicalWeaponEntry{
    {"Breaching Pot (buried)", "100 thaler*", "3d10 bludgeoning", "20 lb.", "Siege"},
    {"Breaching Pot (ground)", "100 thaler*", "3d8 fire", "20 lb.", "Area effect (20-foot radius)"},
    {"Firetube, Portable", "100 thaler*", "See Firetube Ammunition", "20 lb.", "Ammunition (range 0), area effect (20-foot cone), heavy, loading, two-handed"},
    {"Firetube, Stationary", "250 thaler*", "See Firetube Ammunition", "50 lb.", "Ammunition (range 0), area effect (60-foot cone), immobile, loading, siege"},
    {"Grenade", "25 thaler", "1d10 fire", "1 lb.", "Area effect (5-foot radius), thrown (range 30/90)"},
}

type alchemicalWeapons struct{}

func init() {
    table_registry.RegisterTable(&alchemicalWeapons{})
}

func (a *alchemicalWeapons) Name() string {
    return "alchemical_weapons"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (a *alchemicalWeapons) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p alchemical_weapons' to view it."
}

// GetFormatted returns an ASCII-formatted table of alchemical weapons.
func (a *alchemicalWeapons) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("ALCHEMICAL WEAPONS (PG 166)\n")
    sb.WriteString("------------------------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-25s | %-12s | %-15s | %-6s | %s\n", "Name", "Cost", "Damage", "Wt.", "Properties"))
    sb.WriteString("------------------------------------------------------------------------------------------------------------------\n")

    for _, row := range alchemicalWeaponsTable {
        sb.WriteString(fmt.Sprintf("%-25s | %-12s | %-15s | %-6s | %s\n", row.Name, row.Cost, row.Damage, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (a *alchemicalWeapons) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (a *alchemicalWeapons) SubCategory() string {
    return "Weapons"
}
