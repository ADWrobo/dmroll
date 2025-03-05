package tables

import (
    "fmt"
    "strings"
)

// TravelTimeEntry represents a single row in the Modes of Travel and Time table.
type TravelTimeEntry struct {
    Mode            string
    PlainsOfAmbria  string
    LightDavokar    string
    DarkDavokar     string
}

var travelTimeTable = []TravelTimeEntry{
    {"Day's March", "10 miles", "10 miles", "5 miles"},
    {"Day's Forced March", "20 miles", "15 miles", "8 miles"},
    {"Day's Death March", "30 miles", "20 miles", "10 miles"},
    {"Day's Ride", "20 miles", "15 miles", "5 miles"},
    {"Day's Forced Ride", "30 miles", "25 miles", "8 miles"},
    {"Day's Death Ride", "35 miles", "30 miles", "10 miles"},
}

type travelTime struct{}

func init() {
    RegisterTable(&travelTime{})
}

func (t *travelTime) Name() string {
    return "modes_of_travel_and_time"
}

// GetFormatted returns an ASCII-formatted table of travel times.
func (t *travelTime) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("MODES OF TRAVEL AND TIME\n")
    sb.WriteString("-----------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %-15s | %-15s | %s\n", "Way of Travel", "Plains of Ambria", "Light Davokar", "Dark Davokar"))
    sb.WriteString("-----------------------------------------------------------------\n")

    for _, row := range travelTimeTable {
        sb.WriteString(fmt.Sprintf("%-20s | %-15s | %-15s | %s\n", row.Mode, row.PlainsOfAmbria, row.LightDavokar, row.DarkDavokar))
    }

    return sb.String()
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (t *travelTime) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p modes_of_travel_and_time' to view it."
}

func (t *travelTime) Category() string {
    return "DM Tools"
}
