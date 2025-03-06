package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// EquipmentPackEntry represents a single row in the Equipment Packs table.
type EquipmentPackEntry struct {
    Name        string
    Description string
}

var equipmentPacksTable = []EquipmentPackEntry{
    {"Burglar’s Pack", "Includes a backpack, 10 feet of string, a bell, 5 candles, a crowbar, a hammer, 10 pitons, a lantern, 2 flasks of oil, 5 days rations, a tinderbox, and a waterskin. The pack also has 50 feet of rope strapped to the side of it."},
    {"Diplomat’s Pack", "Includes a chest, a set of fine clothes, a bottle of ink, an ink pen, a lamp, 2 flasks of oil, 5 sheets of paper, a vial of perfume, sealing wax, and soap."},
    {"Entertainer’s Pack", "Includes a backpack, a bedroll, 2 sets of clothes (your choice), 5 candles, 5 days of rations, a waterskin, and a disguise kit."},
    {"Explorer’s Pack", "Includes a backpack, a bedroll, a mess kit, a tinderbox, 10 torches, 10 days of rations, and a waterskin. The pack also has 50 feet of rope strapped to the side of it."},
    {"Priest’s Pack", "Includes a backpack, a blanket, 10 candles, a tinderbox, an alms box, 2 blocks of incense, a censer, vestments, 2 days of rations, and a waterskin."},
    {"Scholar’s Pack", "Includes a backpack, a book of lore, a bottle of ink, an ink pen, 10 sheets of parchment, a little bag of sand, and a small knife."},
}

type equipmentPacks struct{}

func init() {
    table_registry.RegisterTable(&equipmentPacks{})
}

func (e *equipmentPacks) Name() string {
    return "equipment_packs"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (e *equipmentPacks) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p equipment_packs' to view it."
}

// GetFormatted returns an ASCII-formatted table of Equipment Packs.
func (e *equipmentPacks) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("EQUIPMENT PACKS\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %s\n", "Name", "Description"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range equipmentPacksTable {
        sb.WriteString(fmt.Sprintf("%-20s | %s\n", row.Name, row.Description))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (e *equipmentPacks) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (e *equipmentPacks) SubCategory() string {
    return "Equipment and Services"
}
