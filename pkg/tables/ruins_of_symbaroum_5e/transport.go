package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// TransportEntry represents a single row in the Transport table.
type TransportEntry struct {
    Name        string
    Cost		string
}

var transportTable = []TransportEntry{
    {"Boat, rowing", "15 thaler"},
    {"Canoe", "10 thaler"},
    {"Cart", "5 thaler"},
    {"Galley", "1000 thaler"},
    {"Horse, draft", "25 thaler"},
    {"Horse, riding", "50 thaler"},
    {"Horse, war", "200 thaler"},
    {"Mule", "5 thaler"},
    {"Riverboat", "2000 thaler"},
    {"Sleigh", "10 thaler"},
    {"Wagon", "20 thaler"},
}

type transport struct{}

func init() {
    table_registry.RegisterTable(&transport{})
}

func (t *transport) Name() string {
    return "transport"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (t *transport) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p transport' to view it."
}

// GetFormatted returns an ASCII-formatted table of Transport.
func (t *transport) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("TRANSPORT\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %s\n", "Name", "Cost"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range transportTable {
        sb.WriteString(fmt.Sprintf("%-20s | %s\n", row.Name, row.Cost))
    }

    sb.WriteString("\n* Note: You can rent passage on a willing transport for 1/10 the cost, per person, per day. For rowing boats, canoes and animals you must leave a deposit equal to 1/2 the cost, or your expected bill, whichever is higher.\n")

    return sb.String()
}

// Category assigns the broader game system category.
func (t *transport) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (t *transport) SubCategory() string {
    return "Equipment and Services"
}
