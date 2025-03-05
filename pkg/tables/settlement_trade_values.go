package tables

import (
    "fmt"
    "strings"
)

// settlementTradeValuesRow holds data for each settlement row.
type settlementTradeValuesRow struct {
    SettlementType string
    SingleItemMax  string // e.g. "1 shilling"
    TotalValue     string // e.g. "1 thaler's worth"
}

// settlementTradeValuesData is derived from your snippet:
// Settlement Type | Single Item (Max. cost) | Total value available
var settlementTradeValuesData = []settlementTradeValuesRow{
    {"Lonely Farm",     "1 shilling",       "1 thaler’s worth"},
    {"Solitary Village","1 thaler",         "10 thaler’s worth"},
    {"Ordinary Village","3 thaler",         "30 thaler’s worth"},
    {"Trade Station",   "10 thaler",        "100 thaler’s worth"},
    {"Ambrian Town",    "50 thaler",        "500 thaler’s worth"},
    {"Thistle Hold",    "100 thaler",       "1000 thaler’s worth"},
    {"Yndaros",         "1000 thaler",      "10000 thaler’s worth"},
}

// settlementTradeValues is our table struct. We won't roll on it, just print.
type settlementTradeValues struct{}

func init() {
    RegisterTable(&settlementTradeValues{})
}

func (s *settlementTradeValues) Name() string {
    return "settlement_trade_values"
}

// Since we don't intend to roll randomly on it, return a static message.
func (s *settlementTradeValues) GetRandomEntry() string {
    return "This table does not support random rolls. Use '-t -p settlement_trade_values' to view."
}

func (s *settlementTradeValues) GetFormatted() string {
    var sb strings.Builder
    sb.WriteString("SETTLEMENT TRADE VALUES (PG 161)\n")
    sb.WriteString("--------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-17s | %-20s | %s\n",
        "Settlement Type", "Single Item (max)", "Total Value Available"))
    sb.WriteString("--------------------------------------------------------------\n")

    for _, row := range settlementTradeValuesData {
        sb.WriteString(fmt.Sprintf("%-17s | %-20s | %s\n",
            row.SettlementType, row.SingleItemMax, row.TotalValue))
    }

    return sb.String()
}

func (s *settlementTradeValues) Category() string {
    return "DM Tools"
}
