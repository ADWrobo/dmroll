package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// ToolEntry represents a single row in the Tools table.
type ToolEntry struct {
    Name string
    Cost string
}

var toolsTable = []ToolEntry{
    {"Alchemist’s supplies", "50 thaler"},
    {"Artifact Catalog", "20 thaler"},
    {"Bestiary", "10 thaler"},
    {"Brewer’s supplies", "10 thaler"},
    {"Calligrapher’s supplies", "20 thaler"},
    {"Carpenter’s tools", "10 thaler"},
    {"Cartographer’s instruments", "15 thaler"},
    {"Cobbler’s tools", "5 thaler"},
    {"Cook’s utensils", "5 thaler"},
    {"Disguise kit", "25 thaler"},
    {"Excavation tools", "10 thaler"},
    {"Field Laboratory", "25 thaler"},
    {"Field Library", "20 thaler"},
    {"Forgery kit", "15 thaler"},
    {"Cheating kit", "1 thaler"},
    {"Dice set", "5 shillings"},
    {"Playing card set", "5 shillings"},
    {"Glassblower’s tools", "30 thaler"},
    {"Herbalism kit", "5 thaler"},
    {"Jeweler’s tools", "25 thaler"},
    {"Leatherworker’s tools", "5 thaler"},
    {"Mason’s tools", "10 thaler"},
    {"Bagpipe", "3 thaler"},
    {"Birch-bark horn", "2 thaler"},
    {"Brass horn", "3 thaler"},
    {"Drum", "3 thaler"},
    {"Fiddle", "3 thaler"},
    {"Flute", "2 shillings"},
    {"Hurdy-gurdy", "3 thaler"},
    {"Lute", "15 shillings"},
    {"Mouth-harp", "5 shillings"},
    {"Shawm", "5 shillings"},
    {"Spinet", "15 thaler"},
    {"Navigator’s tools", "25 thaler"},
    {"Painter’s supplies", "10 thaler"},
    {"Poisoner’s kit", "50 thaler"},
    {"Potter’s tools", "3 thaler"},
    {"Smith’s tools", "20 thaler"},
    {"Tinker’s tools", "10 thaler"},
    {"Thieves’ tools", "25 thaler"},
    {"Trapper’s Manual", "50 thaler"},
    {"Weaver’s tools", "1 thaler"},
    {"Woodcarver’s tools", "1 thaler"},
}

type tools struct{}

func init() {
    table_registry.RegisterTable(&tools{})
}

func (t *tools) Name() string {
    return "tools"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (t *tools) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p tools' to view it."
}

// GetFormatted returns an ASCII-formatted table of Tools.
func (t *tools) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("TOOLS (PG 174)\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-25s | %s\n", "Name", "Cost"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range toolsTable {
        sb.WriteString(fmt.Sprintf("%-25s | %s\n", row.Name, row.Cost))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (t *tools) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (t *tools) SubCategory() string {
    return "Equipment and Services"
}
