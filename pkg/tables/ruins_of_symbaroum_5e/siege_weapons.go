package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// SiegeWeaponsEntry represents a single row in the Siege Weapons table.
type SiegeWeaponsEntry struct {
    Name       string
    Cost       string
    Damage     string
    Weight     string
    Properties string
}

var siegeWeaponsTable = []SiegeWeaponsEntry{
    {"Ballista", "200 thaler", "2d12 piercing", "450 lb.", "Ammunition (range 200/800), crewed (2), immobile, reload, siege"},
    {"Catapult", "400 thaler", "3d8 bludgeoning", "2000 lb.", "Ammunition (range 300/1200), area effect (5-foot radius), crewed (4), immobile, reload, siege"},
    {"Missile Battery", "150 thaler", "3d8 fire", "300 lb.", "Ammunition (range 150/600), area effect (10-foot radius), crewed (2), immobile, reload"},
    {"Trebuchet", "350 thaler", "3d12 bludgeoning", "2500 lb.", "Ammunition (range 300/1200), area effect (5-foot radius), crewed (4), immobile, reload, siege"},
}

type siegeWeapons struct{}

func init() {
    table_registry.RegisterTable(&siegeWeapons{})
}

func (s *siegeWeapons) Name() string {
    return "siege_weapons"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (s *siegeWeapons) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p siege_weapons' to view it."
}

// GetFormatted returns an ASCII-formatted table of Siege Weapons.
func (s *siegeWeapons) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("SIEGE WEAPONS\n")
    sb.WriteString("------------------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-15s | %-10s | %-15s | %-10s | %s\n", "Name", "Cost", "Damage", "Weight", "Properties"))
    sb.WriteString("------------------------------------------------------------------------------------------------------------\n")

    for _, row := range siegeWeaponsTable {
        sb.WriteString(fmt.Sprintf("%-15s | %-10s | %-15s | %-10s | %s\n", row.Name, row.Cost, row.Damage, row.Weight, row.Properties))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (s *siegeWeapons) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (s *siegeWeapons) SubCategory() string {
    return "Weapons"
}
