package tables

import (
    "fmt"
    "strings"
	"log"

	"dmroll/pkg/dice"
)

// RuinSafeguard represents a single row in the Safeguards for Entryways table.
type RuinSafeguard struct {
    Roll      string
    Safeguard string
}

// ruinSafeguardsForEntrywaysTable holds all entries.
var ruinSafeguardsForEntrywaysTable = []RuinSafeguard{
    {"1-2", "No safeguards"},
    {"3-4", "Barricade only"},
    {"5", "Lone guard"},
    {"6", "Guard post"},
    {"7", "Trap"},
    {"8", "Lone guard and barricade"},
    {"9", "Guard post and barricade"},
    {"10", "Double guard post and barricade"},
}

// ruinSafeguardsForEntryways implements the Table interface
type ruinSafeguardsForEntryways struct{}

// Register the table automatically
func init() {
    RegisterTable(&ruinSafeguardsForEntryways{})
}

// Name returns the table name
func (r *ruinSafeguardsForEntryways) Name() string {
    return "ruin_safeguards_for_entryways"
}

// GetRandomEntry returns a random row
func (r *ruinSafeguardsForEntryways) GetRandomEntry() string {
    roll, err := dice.RollDice("1d10")
    if err != nil {
        log.Printf("Error rolling 1d10 for ruin safeguards: %v", err)
        roll = 1
    }
    // Convert the dice result (1-10) to a zero-based index.
    index := roll - 1
    // Safety check in case the table isn't exactly 10 entries.
    if index < 0 || index >= len(ruinSafeguardsForEntrywaysTable) {
        index = 0
    }
    entry := ruinSafeguardsForEntrywaysTable[index]
    return fmt.Sprintf("Safeguard: %s", entry.Safeguard)
}

// GetFormatted returns an ASCII-formatted table of all entries, including detailed notes.
func (r *ruinSafeguardsForEntryways) GetFormatted() string {
    var sb strings.Builder

    sb.WriteString("SAFEGUARDS FOR ENTRYWAYS (1d10)\n")
    sb.WriteString("-------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-7s | %s\n", "Roll", "Safeguard"))
    sb.WriteString("-------------------------------------------\n")

    for _, row := range ruinSafeguardsForEntrywaysTable {
        sb.WriteString(fmt.Sprintf("%-7s | %s\n", row.Roll, row.Safeguard))
    }

    sb.WriteString("\nNOTES:\n")
    sb.WriteString("- **Barricade**: Attackers must spend an entire turn climbing to enter close combat. A DC 13 Strength (Athletics) or Dexterity (Acrobatics) check is required to get past it in a hurry. Moving silently past requires a DC 15 Dexterity (Stealth) check. Defenders can fire projectiles through or over the barricade without penalty, while attackers firing back at them have disadvantage on their attack rolls.\n")
    sb.WriteString("- **Lone guard**: A humanoid or guard beast tasked to raise the alarm in case of intruders.\n")
    sb.WriteString("- **Guard post**: 1d4 humanoids, beasts, abominations, or undead.\n")
    sb.WriteString("- **Trap**: Can be detected with a passive Perception of 15 or a successful DC 15 Intelligence (Investigation) or Wisdom (Perception) [Vigilant -3] check. Deals 1d12 damage plus poison (Roll 1d4: 1-2: 1d6, 3: 1d8, 4: 1d10). Requires a successful DC 17 Strength (Athletics) check to break loose and triggers an alarm heard across the entire level. Disarming it requires a DC 17 Dexterity (Sleight of Hand or Thieves’ Tools) check.\n")

    return sb.String()
}

func (r *ruinSafeguardsForEntryways) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *ruinSafeguardsForEntryways) SubCategory() string {
    return "Ruin"
}
