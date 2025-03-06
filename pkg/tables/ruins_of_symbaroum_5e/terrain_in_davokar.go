package ruins_of_symbaroum_5e

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
    "dmroll/pkg/table_registry"
    "dmroll/pkg/tables"
)

// TerrainEntry represents a single row in the "Terrain in Davokar" table.
type TerrainEntry struct {
    MinRoll   int
    MaxRoll   int
    Examples  string
    Effect    string
}

// terrainTable holds the 1d20 + Mods table from your photo.
var terrainTable = []TerrainEntry{
    {1, 10, "None", 
     "—"},
    {11, 12, "Easily traversable", 
     "The party covers 6 miles (10 km) more than normal."},
    {13, 14, "Swamp/marsh", 
     "The party covers 3 miles (5 km) less, and each member must make a DC 10 Constitution saving throw or gain 1 level of exhaustion."},
    {15, 16, "Sinkhole", 
     "Each member must make a DC 10 Wisdom saving throw or take 2d6 bludgeoning damage from falling."},
    {17, 18, "Poisonous spores", 
     "Each member must make a DC 13 Constitution saving throw. On a failure, take 1d6 poison damage for 1d4+1 rounds."},
    {19, 20, "Vengeful terrain", 
     "Create an enemy group suited to this location."},
    {21, 21, "Slightly corrupted nature", 
     "Make one roll on the Corrupted Nature table in the Ruins of Symbaroum Bestiary, page 23."},
    {22, 22, "Corrupted nature", 
     "Make two rolls on the Corrupted Nature table in the Ruins of Symbaroum Bestiary, page 23."},
    // 23+ => Severely corrupted nature
    {23, 0,  "Severely corrupted nature", 
     "Make three rolls on the Corrupted Nature table in the Ruins of Symbaroum Bestiary, page 23."},
}

// terrainInDavokar implements the Table interface.
type terrainInDavokar struct{}

func init() {
    table_registry.RegisterTable(&terrainInDavokar{})
}

// Name identifies this table in the CLI.
func (t *terrainInDavokar) Name() string {
    return "terrain_in_davokar"
}

// GetRandomEntry prompts for a Davokar location modifier, rolls 1d20 + mod,
// finds the matching row, and returns a formatted string.
func (t *terrainInDavokar) GetRandomEntry() string {
    // Prompt user for location-based mod (from davokar_modifiers.go).
    mod := tables.PromptForDavokarModifier()
    
    fmt.Println() // optional blank line after input

    // Roll 1d20
    baseRoll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling dice for Terrain in Davokar: %v", err)
        baseRoll = 1
    }
    finalRoll := baseRoll + mod

    // Find the corresponding entry in our terrainTable
    entry := findTerrainEntry(finalRoll)

    // Return the final result
    return fmt.Sprintf(
        "Terrain in Davokar result: %d (1d20+%d)\nExamples: %s\nEffect: %s",
        finalRoll, mod, entry.Examples, entry.Effect,
    )
}

// GetFormatted prints a pretty ASCII version of the entire table (for `dmroll -p terrain_in_davokar`).
func (t *terrainInDavokar) GetFormatted() string {
    var sb strings.Builder
    sb.WriteString("TERRAIN IN DAVOKAR (1d20 + Mods)\n")
    sb.WriteString("----------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-10s | %-22s | %s\n", "Roll Range", "Examples", "Effect"))
    sb.WriteString("----------------------------------------------------------------------------\n")

    for _, row := range terrainTable {
        rangeLabel := ""
        if row.MaxRoll == 0 {
            // e.g. 23+
            rangeLabel = fmt.Sprintf("%d+", row.MinRoll)
        } else if row.MinRoll == row.MaxRoll {
            // single value e.g. "21", "22"
            rangeLabel = fmt.Sprintf("%d", row.MinRoll)
        } else {
            // e.g. "1-10", "11-12"
            rangeLabel = fmt.Sprintf("%d-%d", row.MinRoll, row.MaxRoll)
        }

        sb.WriteString(fmt.Sprintf("%-10s | %-22s | %s\n",
            rangeLabel, row.Examples, row.Effect))
    }
    return sb.String()
}

// findTerrainEntry returns the correct row given a final roll value.
func findTerrainEntry(roll int) TerrainEntry {
    for _, row := range terrainTable {
        if row.MaxRoll < row.MinRoll {
            // open-ended row: 23+
            if roll >= row.MinRoll {
                return row
            }
        } else {
            // normal range (MinRoll..MaxRoll)
            if roll >= row.MinRoll && roll <= row.MaxRoll {
                return row
            }
        }
    }

    // If you somehow rolled beyond 23+ (like 30) or below 1 (shouldn’t happen),
    // fallback to a ??? row
    return TerrainEntry{
        MinRoll:  roll,
        MaxRoll:  roll,
        Examples: "???",
        Effect:   "--",
    }
}

func (r *terrainInDavokar) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (r *terrainInDavokar) SubCategory() string {
    return "Davokar"
}
