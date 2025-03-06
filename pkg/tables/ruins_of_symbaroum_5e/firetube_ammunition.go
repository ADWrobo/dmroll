package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// FiretubeAmmunitionEntry represents a single row in the Firetube Ammunition table.
type FiretubeAmmunitionEntry struct {
    Name       string
    Cost       string
    Weight     string
    Effects    string
}

var firetubeAmmunitionTable = []FiretubeAmmunitionEntry{
    {"Bang Powder (DC 25)", "8/24 thaler", "2 lb./5 lb.", "1d10/2d8 thunder damage, deafened"},
    {"Burning Powder (DC 20)", "12/30 thaler", "1 lb./3 lb.", "1d12/2d10 fire damage"},
    {"Flash Powder (DC 25)", "6/20 thaler", "1 lb./3 lb.", "1d6/2d6 fire damage, blinded"},
    {"Shock Powder (DC 30)", "20/50 thaler", "2 lb./5 lb.", "1d12/4d6 bludgeoning damage, stunned"},
}

type firetubeAmmunition struct{}

func init() {
    table_registry.RegisterTable(&firetubeAmmunition{})
}

func (f *firetubeAmmunition) Name() string {
    return "firetube_ammunition"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (f *firetubeAmmunition) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p firetube_ammunition' to view it."
}

// GetFormatted returns an ASCII-formatted table of Firetube Ammunition.
func (f *firetubeAmmunition) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("FIRETUBE AMMUNITION (PG 166)\n")
    sb.WriteString("-------------------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-25s | %-15s | %-10s | %s\n", "Name", "Cost (Portable/Stationary)", "Weight", "Effects (Portable/Stationary)"))
    sb.WriteString("-------------------------------------------------------------------------------------------------\n")

    for _, row := range firetubeAmmunitionTable {
        sb.WriteString(fmt.Sprintf("%-25s | %-15s | %-10s | %s\n", row.Name, row.Cost, row.Weight, row.Effects))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (f *firetubeAmmunition) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (f *firetubeAmmunition) SubCategory() string {
    return "Weapons"
}
