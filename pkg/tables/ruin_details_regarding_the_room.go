package tables

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
)

// RuinRoomDetail represents a single row in the Details Regarding the Room table.
type RuinRoomDetail struct {
    Roll    string
    Detail  string
}

// ruinDetailsRegardingTheRoomTable holds all entries.
var ruinDetailsRegardingTheRoomTable = []RuinRoomDetail{
    {"1-6", "Nothing of interest."},
    {"7", "Lone creature: A single creature of the group dominating the level or a nearby area."},
    {"8-10", "Group: A number of creatures equal to the number of characters."},
    {"11-12", "Crowd: A number of creatures equal to twice the number of characters."},
    {"13", "Trap: See page 146 or use poisoned spike trap; each creature crossing the room must make three DC 15 Dexterity saving throws, receiving 1d6 piercing damage and 2d6 poison damage on a failure. Creatures that treat the area as difficult terrain get advantage on the saving throws."},
    {"14", "Collapsing ceiling (DC 15 Dexterity saving throw, receiving 1d12 bludgeoning damage on a failure or half as much on a success, creates a hole in the ceiling)."},
    {"15", "Collapsing floor (DC 15 Dexterity saving throw or fall 15 feet and take 2d6 bludgeoning damage, creates a hole in the floor)."},
    {"16", "Remains: Fresh or old corpse/skeleton of a cultural being, wearing jewelry and/or coins worth 1d10 thaler."},
    {"17", "Wealthy remains: Fresh or old corpse/skeleton of a cultural being, wearing jewelry and/or coins worth 5d20 thaler."},
    {"18", "Assorted items: The Gamemaster decides or allows each character to roll once on the Curiosities table (page 73)."},
    {"19", "Treasure trove: Requires DC 13 Intelligence (Investigation) check to find. Roll 1d4 + 1 times on the Treasure Type table on page 73."},
    {"20", "Roll twice on the table, ignoring any further results of 20."},
}

// ruinDetailsRegardingTheRoom implements the Table interface
type ruinDetailsRegardingTheRoom struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinDetailsRegardingTheRoom{})
}

// Name returns the table name
func (r *ruinDetailsRegardingTheRoom) Name() string {
    return "ruin_details_regarding_the_room"
}

// GetRandomEntry returns a random row
func (r *ruinDetailsRegardingTheRoom) GetRandomEntry() string {
    roll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling 1d20 for ruin details: %v", err)
        roll = 1 // default if there's an error
    }
    // Convert the dice result (1-20) to a zero-based index
    index := roll - 1
    if index < 0 || index >= len(ruinDetailsRegardingTheRoomTable) {
        index = 0 // fallback if roll is out of range
    }
    entry := ruinDetailsRegardingTheRoomTable[index]
    return fmt.Sprintf("Room Detail: %s", entry.Detail)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (r *ruinDetailsRegardingTheRoom) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("DETAILS REGARDING THE ROOM (1d20)\n")
    sb.WriteString("-------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Detail"))
    sb.WriteString("-------------------------------------------------------------\n")

    for _, row := range ruinDetailsRegardingTheRoomTable {
        sb.WriteString(fmt.Sprintf("%-5s | %s\n", row.Roll, row.Detail))
    }

    return sb.String()
}

func (r *ruinDetailsRegardingTheRoom) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinDetailsRegardingTheRoom) SubCategory() string {
    return "Ruin"
}
