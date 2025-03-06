package ruins_of_symbaroum_5e

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
    "dmroll/pkg/table_registry"
)

// ArtifactEntry holds a row of data for our 1d12 Artifacts table.
type ArtifactEntry struct {
    MinRoll   int
    MaxRoll   int
    Name      string
    Page      int
}

// artifactsTable: each entry covers a single number range for 1..12.
var artifactsTable = []ArtifactEntry{
    {1, 1,  "Ashiki's Twin Blades",           125},
    {2, 2,  "Buneifer's Death Mask",          126},
    {3, 3,  "Feud Crystal",                   129},
    {4, 4,  "Girakosh's Steel Circle",        130},
    {5, 5,  "Iloona's Mirror Shield",         134},
    {6, 6,  "The Living Chain of Hogarz",     132},
    {7, 7,  "Nema's Orbit Statue",            136},
    {8, 8,  "The War Horn of the Stormhowler",137},
    {9, 9,  "The Whispering Necklace of Lapi-Esha", 135},
    {10, 10,"The Wraith Mark of the Eternity Legion",128},
    {11, 11,"Worldcleaver",                   140},
    {12, 12,"Xtala's Hourglass",              143},
}

// artifacts implements the Table interface, so it’s recognized by our CLI.
type artifacts struct {}

func init() {
    table_registry.RegisterTable(&artifacts{})
}

func (a *artifacts) Name() string {
    return "artifacts"
}

// GetRandomEntry: rolls 1d12, finds the matching artifact, and returns it.
func (a *artifacts) GetRandomEntry() string {
    roll, err := dice.RollDice("1d12")
    if err != nil {
        log.Printf("Error rolling 1d12 for artifacts: %v", err)
        roll = 1
    }
    entry := findArtifactEntry(roll)

    // Return a nicely formatted string
    return fmt.Sprintf(
        "Artifacts (1d12) roll: %d\nArtifact: %s\nPage: %d",
        roll, entry.Name, entry.Page,
    )
}

// GetFormatted: prints the entire table as ASCII.
func (a *artifacts) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("ARTIFACTS (1d12)\n")
    sb.WriteString("----------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %-35s | %s\n", "Roll", "Artifact", "Page"))
    sb.WriteString("----------------------------------------------\n")
    for _, row := range artifactsTable {
        rangeLabel := fmt.Sprintf("%d", row.MinRoll)
        sb.WriteString(fmt.Sprintf("%-5s | %-35s | %d\n", rangeLabel, row.Name, row.Page))
    }
    return sb.String()
}

// findArtifactEntry looks up the correct artifact in artifactsTable based on the roll.
func findArtifactEntry(roll int) ArtifactEntry {
    for _, row := range artifactsTable {
        if roll >= row.MinRoll && roll <= row.MaxRoll {
            return row
        }
    }
    // Fallback if somehow the roll is out of range (shouldn’t happen with 1..12).
    return ArtifactEntry{
        MinRoll: roll, MaxRoll: roll,
        Name:    "???",
        Page:    0,
    }
}

func (m *artifacts) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (m *artifacts) SubCategory() string {
    return "Treasure"
}
