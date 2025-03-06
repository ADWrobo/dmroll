// TODO: Same as food and drink--should this be broken out into smaller table files?
// I could see the value in making it easier for users to find a specific item...

package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// AlchemicalElixirEntry represents a single row in the Alchemical Elixirs table.
type AlchemicalElixirEntry struct {
    Name   string
    Value  string
    Amount string
}

// AlchemicalElixirCategory holds a category name and its corresponding elixirs.
type AlchemicalElixirCategory struct {
    Category string
    Entries  []AlchemicalElixirEntry
}

var alchemicalElixirsTable = []AlchemicalElixirCategory{
    {"Antidotes", []AlchemicalElixirEntry{
        {"Weak (DC 5, +2)", "5 thaler", "1d6"},
        {"Moderate (DC 10, +5)", "10 thaler", "2d6"},
        {"Strong (DC 15, +10)", "20 thaler", "3d6"},
    }},
    {"Antidote Candle", []AlchemicalElixirEntry{
        {"Weak (DC 10, +2)", "10 thaler", "1d6"},
        {"Moderate (DC 15, +5)", "20 thaler", "1d6"},
    }},
    {"Poisons", []AlchemicalElixirEntry{
        {"Weak (DC 10, 1 round [+1], 1d6 poison damage [+3])", "10 thaler", "1"},
        {"Moderate (DC 15, 1 round [+1], 1d6 poison damage [+3])", "20 thaler", "1"},
        {"Strong (DC 20, 1 minute [+3], 1d6 poison damage [+9])", "30 thaler*", "1"},
    }},
    {"Poison Candle", []AlchemicalElixirEntry{
        {"Weak (DC 13, 1 round [+1], 1d6 poison damage [+3])", "40 thaler*", "1d4"},
        {"Moderate (DC 17, 1 round [+1], 1d6 poison damage [+3])", "60 thaler*", "1d4"},
        {"Protective oil (DC 17)", "20 thaler", "1"},
    }},
    {"Purple Sap", []AlchemicalElixirEntry{
        {"Weak (DC 15)", "20 thaler", "1d6"},
        {"Moderate (DC 17)", "40 thaler", "1d6"},
        {"Strong (DC 20)", "60 thaler", "1d6"},
    }},
    {"Transforming Drought", []AlchemicalElixirEntry{
        {"Weak (DC 13)", "10 thaler*", "1d4"},
        {"Moderate (DC 15)", "30 thaler*", "1d4"},
        {"Strong (DC 17)", "60 thaler*", "1d4"},
    }},
    {"Miscellaneous Alchemical Items", []AlchemicalElixirEntry{
        {"Revealing light (DC 13)", "10 thaler", "1d4"},
        {"Shadow tint (DC 15)", "15 thaler", "1d6"},
        {"Smoke bomb (DC 10)", "10 thaler", "1d6"},
        {"Spirit friend (DC 13)", "60 thaler", "1"},
        {"Spore bomb (DC 25)", "5 thaler", "1d6"},
        {"Stun bolt (DC 13)", "20 thaler*", "1"},
        {"Thorn beasties (DC 15)", "5 thaler", "1d6"},
        {"Thunder ball (DC 10)", "15 thaler", "1d6"},
        {"Choking spores (DC 20)", "20 thaler", "1d6"},
        {"Concentrated magic (DC 15)", "10 thaler", "1d6"},
        {"Drone dew (DC 10)", "10 thaler", "1d4"},
        {"Elemental essence (DC 15)", "10 thaler", "1d4"},
        {"Elixir of Life (DC 15)", "10 thaler", "1d6"},
        {"Eye drops (DC 13)", "5 thaler", "1d6"},
        {"Fire dye (DC 10)", "5 thaler", "1d6"},
        {"Flash powder (DC 15)", "10 thaler", "1d4"},
        {"Ghost candle (DC 15)", "10 thaler", "1d4"},
        {"Herbal cure (DC 10)", "10 thaler", "1d4"},
        {"Holy water (DC 17)", "10 thaler", "1d6"},
        {"Homing arrow (DC 13)", "10 thaler", "1d6"},
        {"Homunculus (DC 17)", "20 thaler*", "1"},
        {"Twilight tincture (DC 17)", "60 thaler", "1d4"},
        {"War paint (DC 15)", "10 thaler", "1d6"},
        {"Way bread (DC 10)", "5 thaler", "1d6"},
        {"Wild elix (DC 13)", "10 thaler", "1d6"},
        {"Wraith dust", "20 thaler", "1d6"},
    }},
}

type alchemicalElixirs struct{}

func init() {
    table_registry.RegisterTable(&alchemicalElixirs{})
}

func (a *alchemicalElixirs) Name() string {
    return "alchemical_elixirs"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (a *alchemicalElixirs) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p alchemical_elixirs' to view it."
}

// GetFormatted returns an ASCII-formatted table of Alchemical Elixirs with category headers and notes.
func (a *alchemicalElixirs) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("ALCHEMICAL ELIXIRS (PG 181)\n")
    sb.WriteString("================================================================\n")

    for _, category := range alchemicalElixirsTable { // Now iterates in a defined order
        sb.WriteString(fmt.Sprintf("\n%s\n", strings.ToUpper(category.Category)))
        sb.WriteString("----------------------------------------------------------------\n")
        sb.WriteString(fmt.Sprintf("%-35s | %-10s | %s\n", "Name", "Value", "Amount Available"))
        sb.WriteString("----------------------------------------------------------------\n")

        for _, row := range category.Entries {
            sb.WriteString(fmt.Sprintf("%-35s | %-10s | %s\n", row.Name, row.Value, row.Amount))
        }
    }

    // Append important notes at the end
    sb.WriteString("\nNOTES:\n")
    sb.WriteString("*  This can only be purchased on the black market.\n")
    sb.WriteString("†  This is the minimum DC; if you set a higher DC to make the elixir, that higher DC is also used to save against the elixir’s effect.\n")

    return sb.String()
}

// Category assigns the broader game system category.
func (a *alchemicalElixirs) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (a *alchemicalElixirs) SubCategory() string {
    return "Equipment and Services"
}
