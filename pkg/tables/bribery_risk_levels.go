package tables

import (
    "fmt"
    "strings"
)

// bribeRiskRow holds one row of data about risk level, intent, and size of bribe.
type bribeRiskRow struct {
    RiskLevel string
    Intent    string
    SizeOfBribe string
}

// bribeRiskData is our in-memory table derived from your snippet.
var bribeRiskData = []bribeRiskRow{
    {
        RiskLevel:  "Low",
        Intent:     "To do one’s job faster",
        SizeOfBribe: "One tenth the person’s daily income. If less than an orteg, bribes might be made with food instead",
    },
    {
        RiskLevel:  "Moderate",
        Intent:     "To break a rule",
        SizeOfBribe: "10–100% of a day's income, often used as an alibi/excuse for the bribed person",
    },
    {
        RiskLevel:  "High",
        Intent:     "To break a law",
        SizeOfBribe: "100%+ a day’s income, as well as an alibi or excuse for the bribed person. Very few people will risk death or exile just for money",
    },
}

// briberyRiskLevels is our table struct.
type briberyRiskLevels struct{}

func init() {
    RegisterTable(&briberyRiskLevels{})
}

func (b *briberyRiskLevels) Name() string {
    return "bribery_risk_levels"
}

// This table is not for random rolls, so we return a static note.
func (b *briberyRiskLevels) GetRandomEntry() string {
    return "This table does not support random rolls. Use '-t -p bribery_risk_levels' to view."
}

// GetFormatted prints the entire table in ASCII format.
func (b *briberyRiskLevels) GetFormatted() string {
    var sb strings.Builder
    sb.WriteString("BRIBERY RISK LEVELS (PG 161)\n")
    sb.WriteString("--------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-10s | %-25s | %s\n",
        "Risk", "Intent", "Size of Bribe"))
    sb.WriteString("--------------------------------------------------------------------------\n")

    for _, row := range bribeRiskData {
        sb.WriteString(fmt.Sprintf("%-10s | %-25s | %s\n",
            row.RiskLevel, row.Intent, row.SizeOfBribe))
    }
    return sb.String()
}

func (b *briberyRiskLevels) Category() string {
    return "DM Tools"
}
