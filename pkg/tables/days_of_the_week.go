package tables

import (
    "fmt"
    "strings"

    "dmroll/pkg/dice"
)

// dayOfWeekRow holds data for each day of the week.
type dayOfWeekRow struct {
    Order     int
    FullName  string
    CommonName string
}

// daysOfTheWeekData is our in-memory table: 7 days, each with an order, full name, and common name.
var daysOfTheWeekData = []dayOfWeekRow{
    {1, "The Day of the Earth",  "Earthsday"},
    {2, "The Day of the Water",  "Watersday"},
    {3, "The Day of the Wind",   "Windsday"},
    {4, "The Day of the Forest", "Forestsday"},
    {5, "The Day of the Mountain","Mountainsday"},
    {6, "The Day of the People", "Peoplseday"},
    {7, "The Day of the Sun",    "Priosday"},
}

// daysOfTheWeek implements the Table interface.
type daysOfTheWeek struct{}

func init() {
    RegisterTable(&daysOfTheWeek{})
}

func (d *daysOfTheWeek) Name() string {
    return "days_of_the_week"
}

// GetRandomEntry picks a random day (1–7) and returns the data.
func (d *daysOfTheWeek) GetRandomEntry() string {
    roll, _ := dice.RollDice("1d7")
    row := daysOfTheWeekData[roll-1]
    return fmt.Sprintf("Day #%d\nFull Name: %s\nCommon Name: %s",
        row.Order, row.FullName, row.CommonName)
}

// GetFormatted prints the entire 7-day table in ASCII.
func (d *daysOfTheWeek) GetFormatted() string {
    var sb strings.Builder
    sb.WriteString("DAYS OF THE WEEK (PG 41)\n")
    sb.WriteString("----------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %-22s | %s\n", "Order", "Full Name", "Common Name"))
    sb.WriteString("----------------------------------------------------\n")

    for _, row := range daysOfTheWeekData {
        sb.WriteString(fmt.Sprintf("%-5d | %-22s | %s\n",
            row.Order, row.FullName, row.CommonName))
    }
    return sb.String()
}

func (d *daysOfTheWeek) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (d *daysOfTheWeek) SubCategory() string {
    return "DM Tools"
}
