package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// TradeGoodsEntry represents a single row in the Trade Goods table.
type TradeGoodsEntry struct {
    Name  string
    Value string
}

var tradeGoodsTable = map[string][]TradeGoodsEntry{
    "General Trade Goods": {
        {"Cardamom, one box", "5 ortegs"},
        {"Cinnamon, one box", "1 thaler"},
        {"Clove, one box", "7 ortegs"},
        {"Copper, one bar", "5 shillings"},
        {"Cotton fabric, one roll", "1 thaler"},
        {"Cumin, one box", "5 ortegs"},
        {"Ginger, one box", "5 ortegs"},
        {"Gold, one bar", "500 thaler"},
        {"Grain, one sack", "5 ortegs"},
        {"Honey, one jar", "1 thaler"},
        {"Iron, one bar", "1 thaler"},
        {"Mint, one box", "5 ortegs"},
        {"Oil (vegetable), one earthen jug", "5 ortegs"},
        {"Roka berries, one box", "12 ortegs"},
        {"Saffron, one box", "50 thaler"},
        {"Salt, one sack", "1 thaler"},
        {"Silk, one roll", "50 thaler"},
        {"Spices, one box", "3 ortegs"},
        {"Sugar, one sack", "1 thaler"},
        {"Tar, one barrel", "5 shillings"},
        {"Turmeric, one box", "15 ortegs"},
        {"Vinegar, one earthen jug", "1 shilling"},
    },
    "Tobacco Goods": {
        {"Chewing tobacco, one box", "6 ortegs"},
        {"Fruit tobacco, one box", "4 ortegs"},
        {"Herbal tobacco, one box", "2 shillings"},
        {"Longbottom leaf, one box", "8 ortegs"},
        {"Smelling snuff, one box", "2 thaler"},
    },
    "Tobacco Utensils": {
        {"Clay pipe", "3 shillings"},
        {"Long-stemmed seafoam pipe", "5 thaler"},
        {"Smoke tube", "5 ortegs"},
        {"Snuff box", "1 shilling"},
        {"Wooden pipe", "5 shillings"},
    },
}

type tradeGoods struct{}

func init() {
    table_registry.RegisterTable(&tradeGoods{})
}

func (t *tradeGoods) Name() string {
    return "trade_goods"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (t *tradeGoods) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p trade_goods' to view it."
}

// GetFormatted returns an ASCII-formatted table of Trade Goods with category headers.
func (t *tradeGoods) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("TRADE GOODS (PG 180)\n")
    sb.WriteString("============================================================\n")

    for category, items := range tradeGoodsTable {
        sb.WriteString(fmt.Sprintf("\n%s\n", strings.ToUpper(category)))
        sb.WriteString("------------------------------------------------------------\n")
        sb.WriteString(fmt.Sprintf("%-35s | %s\n", "Name", "Value"))
        sb.WriteString("------------------------------------------------------------\n")

        for _, row := range items {
            sb.WriteString(fmt.Sprintf("%-35s | %s\n", row.Name, row.Value))
        }
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (t *tradeGoods) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (t *tradeGoods) SubCategory() string {
    return "Equipment and Services"
}
