package tables

import (
    "fmt"
    "strings"

    "dmroll/pkg/dice"
)

// monthOfYearRow holds data for each month of the year.
type monthOfYearRow struct {
    Order      int
    Season     string
    MonthGod   string
    CommonName string
}

// monthsOfTheYearData is our in-memory table of the 12 months.
var monthsOfTheYearData = []monthOfYearRow{
    {1, "Winter", "High Summer", "High Summer"},
    {2, "Fall", "Harvest", "Harvest"},
    {3, "Fall", "Andonia", "Andonia"},
    {4, "Fall", "Pathfinder", "Slaughter"},
    {5, "Fall", "Sent", "Sent"},
    {6, "Winter", "Trierla", "Trierla"},
    {7, "Winter", "Moragal", "Moragal"},
    {8, "Spring", "Aindar", "Indigent"},
    {9, "Spring", "Path Builder", "Thaw"},
    {10, "Spring", "Vergon", "Sowing"},
    {11, "Summer", "Koneila", "Blooming"},
    {12, "Summer", "Leandro", "Luscious"},
}

// monthsOfTheYear implements the Table interface.
type monthsOfTheYear struct{}

func init() {
    RegisterTable(&monthsOfTheYear{})
}

func (m *monthsOfTheYear) Name() string {
    return "months_of_the_year"
}

// GetRandomEntry picks a random month (1–12) and returns the details.
func (m *monthsOfTheYear) GetRandomEntry() string {
    roll, _ := dice.RollDice("1d12")
    row := monthsOfTheYearData[roll-1]
    return fmt.Sprintf("Month #%d\nSeason: %s\nMonth God: %s\nCommon Name: %s",
        row.Order, row.Season, row.MonthGod, row.CommonName)
}

// GetFormatted prints all 12 months in a neat ASCII table.
func (m *monthsOfTheYear) GetFormatted() string {
    var sb strings.Builder
    sb.WriteString("MONTHS OF THE YEAR (PG 42)\n")
    sb.WriteString("------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %-10s | %-15s | %s\n", "Order", "Season", "Month God", "Common Name"))
    sb.WriteString("------------------------------------------------------------------\n")

    for _, row := range monthsOfTheYearData {
        sb.WriteString(fmt.Sprintf("%-5d | %-10s | %-15s | %s\n",
            row.Order, row.Season, row.MonthGod, row.CommonName))
    }
    return sb.String()
}

func (m *monthsOfTheYear) Category() string {
    return "DM Tools"
}
