package tables

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
)

// MarkOfCorruptionEntry represents a single row in the Marks of Corruption table.
type MarkOfCorruptionEntry struct {
    Roll   int
    Detail string
}

var marksOfCorruptionTable = []MarkOfCorruptionEntry{
    {1, "No mark, but you gain a point of permanent Corruption. Erase all current temporary Corruption."},
    {2, "No mark, but you gain a point of permanent Corruption. Erase all current temporary Corruption."},
    {3, "Thirst for warm blood, must drink warm blood every day or gain a level of exhaustion."},
    {4, "Taste for cadavers, must feed on something rancid every day or gain a level of exhaustion."},
    {5, "Fangs, or nails in the shape of claws or talons."},
    {6, "Taste for raw meat, must feed on fresh meat every day or gain a level of exhaustion."},
    {7, "Discolored skin, blemishes and severe rashes."},
    {8, "A festering wound that does not heal."},
    {9, "Boils in mouth, throat, and ears that burst at inopportune times. You are deafened."},
    {10, "Eyes that actually blacken with anger, hunger, or lust. Reduce temporary Corruption by 1."},
    {11, "Speaks in an unknown, evil-sounding language while sleeping. Reduce temporary Corruption by 1."},
    {12, "Degeneration of sight (you are blinded). This condition forces you to rely on other senses instead, like smell and touch. Reduce temporary Corruption by 1."},
    {13, "Cold as a corpse or feverishly hot without any signs of sickness or disease. Reduce temporary Corruption by 1."},
    {14, "Drawn to tainted artifacts and evil places while sleepwalking. Reduce temporary Corruption by 1."},
    {15, "Veins that bulge black when experiencing anger or other strong emotions. Reduce temporary Corruption to half, rounding down."},
    {16, "A birthmark that with imagination may look like a dark rune or an evil symbol. Reduce temporary Corruption to half, rounding down."},
    {17, "A faint odor of decay follows you, despite you feeling healthy. Reduce temporary Corruption to half, rounding down."},
    {18, "Dark streaks in the blood, visible when you bleed. Reduce temporary Corruption to half, rounding down."},
    {19, "Breath that stinks of sulfur. Disadvantage on social checks. Reduce temporary Corruption to half, rounding down."},
    {20, "Eyes that glitter in the dark. Reduce temporary Corruption to half, rounding down."},
}

type marksOfCorruption struct{}

func init() {
    RegisterTable(&marksOfCorruption{})
}

func (m *marksOfCorruption) Name() string {
    return "marks_of_corruption"
}

// GetRandomEntry rolls 1d20, then returns the matching item's description.
func (m *marksOfCorruption) GetRandomEntry() string {
    roll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling 1d20: %v", err)
        roll = 1
    }
    entry := findMarkOfCorruption(roll)
    return fmt.Sprintf("Mark of Corruption (1d20) roll: %d\n%s", roll, entry.Detail)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (m *marksOfCorruption) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("MARKS OF CORRUPTION (1d20)\n")
    sb.WriteString("---------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Effect"))
    sb.WriteString("---------------------------------------------------------------\n")

    for _, row := range marksOfCorruptionTable {
        sb.WriteString(fmt.Sprintf("%-5d | %s\n", row.Roll, row.Detail))
    }

    return sb.String()
}

// findMarkOfCorruption retrieves the corresponding entry for a given roll.
func findMarkOfCorruption(roll int) MarkOfCorruptionEntry {
    for _, entry := range marksOfCorruptionTable {
        if entry.Roll == roll {
            return entry
        }
    }
    // Fallback in case of an invalid roll (should not happen)
    return MarkOfCorruptionEntry{Roll: roll, Detail: "Unknown corruption effect."}
}

func (m *marksOfCorruption) Category() string {
    return "DM Tools"
}
