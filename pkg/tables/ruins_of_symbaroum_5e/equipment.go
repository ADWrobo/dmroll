package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// EquipmentEntry represents a single row in the Equipment table.
type EquipmentEntry struct {
    Item   string
    Cost   string
    Weight string
}

var equipmentTable = []EquipmentEntry{
    {"Bandages", "5 ortegs", "-"},
    {"Bear Trap", "1 thaler", "25 lb."},
    {"Bedroll", "5 ortegs", "7 lb."},
    {"Blanket", "5 ortegs", "3 lb."},
    {"Brass bell", "6 shillings", "-"},
    {"Climbing equipment", "1 thaler", "12 lb."},
    {"Cooking pan", "1 shilling", "5 lb."},
    {"Crayons", "1 orteg", "-"},
    {"Crowbar", "2 thaler", "5 lb."},
    {"Drinking horn", "2 ortegs", "3 lb. (full)"},
    {"Firewood", "2 ortegs", "3 lb."},
    {"Fishing line, lures and hooks", "5 ortegs", "-"},
    {"Fishing net", "1 shilling", "4 lb."},
    {"Grappling hook*", "1 thaler", "4 lb."},
    {"Holy Symbol", "5 thaler", "-"},
    {"Hourglass", "10 thaler", "1 lb."},
    {"Incense (1 hour’s worth)", "1 thaler", "-"},
    {"Ink and quill", "1 shilling", "-"},
    {"Ladder", "4 shillings", "25 lb."},
    {"Lamp Oil", "7 ortegs", "-"},
    {"Lantern", "4 shillings", "2 lb."},
    {"Lockpicks", "1 thaler", "-"},
    {"Needle and thread", "3 ortegs", "-"},
    {"Paper", "7 ortegs", "-"},
    {"Parchment", "4 ortegs", "-"},
    {"Perfume", "4 thaler", "1 lb."},
    {"Pocket mirror", "10 thaler", "-"},
    {"Rations (per day, per person)", "5 ortegs", "2 lb."},
    {"Rope (per 50-foot section)", "4 shillings", "4 lb."},
    {"Rope ladder (per 10-foot section)", "5 shillings", "2 lb."},
    {"Snare", "5 shillings", "10 lb."},
    {"Snow shoes", "5 shillings", "10 lb."},
    {"Soap", "5 ortegs", "-"},
    {"Spy glass", "50 thaler", "1 lb."},
    {"Tankard", "1 orteg", "1 lb."},
    {"Tent", "5 shillings", "20 lb."},
    {"Torch", "3 ortegs", "1 lb."},
	{"Waterskin", "1 shilling", "5 lb. (full)"},
    {"Wax candle", "2 ortegs", "-"},
    {"Wax, seal", "1 thaler", "1 lb."},
    {"Weapon maintenance kit", "5 shillings", "2 lb."},
    {"Whetstone", "4 ortegs", "1 lb."},
    {"Whistle", "2 shillings", "-"},
}

type equipment struct{}

func init() {
    table_registry.RegisterTable(&equipment{})
}

func (e *equipment) Name() string {
    return "equipment"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (e *equipment) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p equipment' to view it."
}

// GetFormatted returns an ASCII-formatted table of Equipment.
func (e *equipment) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("EQUIPMENT\n")
    sb.WriteString("------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-25s | %-10s | %s\n", "Item", "Cost", "Weight"))
    sb.WriteString("------------------------------------------------------\n")

    for _, row := range equipmentTable {
        sb.WriteString(fmt.Sprintf("%-25s | %-10s | %s\n", row.Item, row.Cost, row.Weight))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (e *equipment) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (e *equipment) SubCategory() string {
    return "Equipment and Services"
}
