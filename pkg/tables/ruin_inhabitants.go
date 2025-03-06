package tables

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
)

// RuinInhabitant represents a single row in the Ruin Inhabitants table.
type RuinInhabitant struct {
    Roll        string
    Inhabitants string
}

// ruinInhabitantsTable holds all entries.
var ruinInhabitantsTable = []RuinInhabitant{
    {"1-3", "Treasure hunters (×2) / Plunderers (+2, minimum 1)"},
    {"4-6", "Goblins (×3)"},
    {"7-8", "Beasts (×2)"},
    {"9-10", "Arachnids (×2)"},
    {"11-12", "Bestials (×1)"},
    {"13-14", "Trolls (+2, minimum 1)"},
    {"15", "Elves (×1)"},
    {"16", "Undead (×2)"},
    {"17", "Abominations (+2, minimum 1)"},
    {"18-20", "Make two rolls and keep both creatures; ignore further results of 18-20"},
}

// ruinInhabitants implements the Table interface
type ruinInhabitants struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinInhabitants{})
}

// Name returns the table name
func (r *ruinInhabitants) Name() string {
    return "ruin_inhabitants"
}

// GetRandomEntry returns a random row
func (r *ruinInhabitants) GetRandomEntry() string {
    roll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling 1d20 for ruin inhabitants: %v", err)
        roll = 1
    }
    // Convert the dice result (1-20) to a zero-based index.
    index := roll - 1
    if index < 0 || index >= len(ruinInhabitantsTable) {
        index = 0
    }
    entry := ruinInhabitantsTable[index]
    return fmt.Sprintf("Inhabitants: %s", entry.Inhabitants)
}

// GetFormatted returns an ASCII-formatted table of all entries, including notes.
func (r *ruinInhabitants) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("RUIN INHABITANTS (1d20)\n")
    sb.WriteString("----------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-7s | %s\n", "Roll", "Inhabitants"))
    sb.WriteString("----------------------------------------------------------\n")

    for _, row := range ruinInhabitantsTable {
        sb.WriteString(fmt.Sprintf("%-7s | %s\n", row.Roll, row.Inhabitants))
    }

    sb.WriteString("\nNOTES:\n")
    sb.WriteString("- **Treasure hunters**: A group of explorers or looters camped inside the ruin, searching for valuables.\n")
    sb.WriteString("- **Goblins**: These creatures have either made the ruin their home or were forced into servitude by another faction.\n")
    sb.WriteString("- **Beasts**: A lone predator or a pack of beasts has sought refuge or is nesting in the ruin.\n")
    sb.WriteString("- **Arachnids**: A colony of arachnids uses the ruin as a lair while hunting for artifacts and prey.\n")
    sb.WriteString("- **Bestials**: One or more bestials have temporarily taken shelter in the ruin.\n")
    sb.WriteString("- **Trolls**: A group of trolls may be resting inside, or sneaking around looking for lost artifacts.\n")
    sb.WriteString("- **Elves**: Elves have claimed the ruin as their base or guard it against intruders.\n")
    sb.WriteString("- **Undead**: The previous inhabitants of the ruin have not left; they are restless and unwelcome visitors.\n")
    sb.WriteString("- **Abominations**: The ruin is corrupted and attracts—or spawns—abominations.\n")
    sb.WriteString("- **Multiple groups (18-20)**: Roll twice and keep both results; the groups coexist, interact, or are in conflict.\n")

    return sb.String()
}

func (r *ruinInhabitants) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinInhabitants) SubCategory() string {
    return "Ruin"
}
