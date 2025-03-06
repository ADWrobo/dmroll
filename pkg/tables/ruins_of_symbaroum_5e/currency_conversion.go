package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// CurrencyEntry represents a single row in the Currency Conversion table.
type CurrencyEntry struct {
    CoinName string
    GPValue  string
    SPValue  string
    CPValue  string
}

var currencyConversionTable = []CurrencyEntry{
    {"Thaler", "1", "10", "100"},
    {"Shilling", "1/10", "1", "10"},
    {"Orteg", "1/100", "1/10", "1"},
}

type currencyConversion struct{}

func init() {
    table_registry.RegisterTable(&currencyConversion{})
}

func (c *currencyConversion) Name() string {
    return "currency_conversion"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (c *currencyConversion) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p currency_conversion' to view it."
}

// GetFormatted returns an ASCII-formatted table of currency conversion values.
func (c *currencyConversion) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("CURRENCY CONVERSION TABLE (PG 160)\n")
    sb.WriteString("--------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-10s | %-10s | %-10s | %s\n", "Coin Name", "GP Value", "SP Value", "CP Value"))
    sb.WriteString("--------------------------------------------------\n")

    for _, row := range currencyConversionTable {
        sb.WriteString(fmt.Sprintf("%-10s | %-10s | %-10s | %s\n", row.CoinName, row.GPValue, row.SPValue, row.CPValue))
    }

    return sb.String()
}

func (c *currencyConversion) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (c *currencyConversion) SubCategory() string {
    return "DM Tools"
}
