package tables

import (
    "fmt"
    "strings"
)

// difficultyClassRangesRow holds one row of the DC ranges table.
type difficultyClassRangesRow struct {
    LevelRange  string
    DC90        string
    DC75        string
    DC50        string
    DC25        string
    DC10        string
}

// We populate it from the snippet in the image (reading approximate data).
// Columns: Level Range, 90% success, 75% success, 50% success, 25% success, 10% success.
var difficultyClassRangesData = []difficultyClassRangesRow{
    {
        LevelRange: "1–4",
        DC90:       "4–8",
        DC75:       "9–11",
        DC50:       "12–13",
        DC25:       "14–15",
        DC10:       "16–17",
    },
    {
        LevelRange: "5–8",
        DC90:       "5–9",
        DC75:       "10–12",
        DC50:       "13–14",
        DC25:       "15–16",
        DC10:       "17–18",
    },
    {
        LevelRange: "9–12",
        DC90:       "6–10",
        DC75:       "11–13",
        DC50:       "14–15",
        DC25:       "16–18",
        DC10:       "19–20",
    },
    {
        LevelRange: "13–16",
        DC90:       "7–11",
        DC75:       "12–14",
        DC50:       "15–17",
        DC25:       "18–19",
        DC10:       "20–21",
    },
    {
        LevelRange: "17–20",
        DC90:       "8–12",
        DC75:       "13–15",
        DC50:       "16–18",
        DC25:       "20–21",
        DC10:       "22–23",
    },
}

// difficultyClassRanges implements the Table interface, but is “print-only.”
type difficultyClassRanges struct{}

func init() {
    RegisterTable(&difficultyClassRanges{})
}

func (d *difficultyClassRanges) Name() string {
    return "difficulty_class_ranges"
}

// Since the table isn’t meant to be rolled on, we return a message
func (d *difficultyClassRanges) GetRandomEntry() string {
    return "This table cannot be rolled. Use '-t -p difficulty_class_ranges' to view it."
}

// GetFormatted prints the DC ranges in an ASCII format.
func (d *difficultyClassRanges) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("DIFFICULTY CLASS RANGES (GG 16)\n")
    sb.WriteString("------------------------------------------------------------------------\n")
    sb.WriteString("Level  | 90% Success | 75% Success | 50% Success | 25% Success | 10% Success\n")
    sb.WriteString("------------------------------------------------------------------------\n")

    for _, row := range difficultyClassRangesData {
        sb.WriteString(fmt.Sprintf(
            "%-6s | %-11s | %-11s | %-12s | %-12s | %s\n",
            row.LevelRange, row.DC90, row.DC75, row.DC50, row.DC25, row.DC10,
        ))
    }
    return sb.String()
}

func (d *difficultyClassRanges) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (d *difficultyClassRanges) SubCategory() string {
    return "DM Tools"
}
