package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// TobaccoGoodsEntry represents a single row in the Tobacco Goods table.
type TobaccoGoodsEntry struct {
    Name  string
    Value string
}

var tobaccoGoodsTable = map[string][]TobaccoGoodsEntry{
    "Tobacco": {
        {"Chewing tobacco, one box", "6 ortegs"},
        {"Fruit tobacco, one box", "4 ortegs"},
        {"Herbal tobacco, one box", "2 shillings"},
        {"Longbottom leaf, one box", "8 ortegs"},
        {"Smelling snuff, one box", "2 thaler"},
    },
    "Utensils": {
        {"Clay pipe", "3 shillings"},
        {"Long-stemmed seafoam pipe", "5 thaler"},
        {"Smoke tube", "5 ortegs"},
        {"Snuff box", "1 shilling"},
        {"Wooden pipe", "5 shillings"},
    },
}

type tobaccoGoods struct{}

func init() {
    table_registry.RegisterTable(&tobaccoGoods{})
}

func (t *tobaccoGoods) Name() string {
    return "tobacco_goods"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (t *tobaccoGoods) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p tobacco_goods' to view it."
}

// GetFormatted returns an ASCII-formatted table of Tobacco Goods with category headers.
func (t *tobaccoGoods) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("TOBACCO GOODS (PG 180)\n")
    sb.WriteString("============================================================\n")

    for category, items := range tobaccoGoodsTable {
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
func (t *tobaccoGoods) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (t *tobaccoGoods) SubCategory() string {
    return "Equipment and Services"
}
