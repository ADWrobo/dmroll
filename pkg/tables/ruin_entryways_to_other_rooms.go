package tables

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
)

// RuinEntrywayRoom represents a single row in the Entryways to Other Rooms table.
type RuinEntrywayRoom struct {
    Roll     int
    Entryway string
}

// ruinEntrywaysToOtherRoomsTable holds all entries.
var ruinEntrywaysToOtherRoomsTable = []RuinEntrywayRoom{
    {1, "Opening, no door"},
    {2, "Broken wooden door. Opening it requires a DC 13 Dexterity (Stealth) check."},
    {3, "Wooden door, intact, unlocked (broken lock)."},
    {4, "Wooden door, intact, locked. Picking the lock requires a DC 13 Dexterity (Thieves’ Tools) check. Forcing the door requires a DC 15 Strength (Athletics) check."},
    {5, "Reinforced wooden door, intact, unlocked (broken lock). Opening it requires a DC 13 Dexterity (Stealth) check."},
    {6, "Reinforced wooden door, intact, locked. Picking the lock requires a DC 17 Dexterity (Thieves’ Tools) check. Cannot be forced, but can be destroyed: AC 15, 20 hp, damage threshold 5."},
    {7, "Copper or iron door, unlocked. Opening it without creaking loudly requires a DC 17 Dexterity (Stealth) check."},
    {8, "Copper or iron door, locked and rusted/corroded shut. Picking the lock requires a DC 19 Dexterity (Thieves’ Tools) check. The door can be opened with a DC 17 Strength (Athletics) check. Anyone within 300 feet of the entryway hears the ensuing noise."},
}

// ruinEntrywaysToOtherRooms implements the Table interface
type ruinEntrywaysToOtherRooms struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinEntrywaysToOtherRooms{})
}

// Name returns the table name
func (r *ruinEntrywaysToOtherRooms) Name() string {
    return "ruin_entryways_to_other_rooms"
}

// GetRandomEntry returns a random row
func (r *ruinEntrywaysToOtherRooms) GetRandomEntry() string {
    roll, err := dice.RollDice("1d8")
    if err != nil {
        log.Printf("Error rolling 1d8 for ruin entryways to other rooms: %v", err)
        roll = 1
    }
    // Convert the dice result (1-8) to a zero-based index.
    index := roll - 1
    if index < 0 || index >= len(ruinEntrywaysToOtherRoomsTable) {
        index = 0
    }
    entry := ruinEntrywaysToOtherRoomsTable[index]
    return fmt.Sprintf("Entryway: %s", entry.Entryway)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinEntrywaysToOtherRooms) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("ENTRYWAYS TO OTHER ROOMS (1d8)\n")
    sb.WriteString("-------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Entryway"))
    sb.WriteString("-------------------------------------------------------------\n")

    for _, row := range ruinEntrywaysToOtherRoomsTable {
        sb.WriteString(fmt.Sprintf("%-5d | %s\n", row.Roll, row.Entryway))
    }

    return sb.String()
}

func (r *ruinEntrywaysToOtherRooms) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinEntrywaysToOtherRooms) SubCategory() string {
    return "Ruin"
}
