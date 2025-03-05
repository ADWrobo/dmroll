package tables

import (
    "fmt"
    "strings"
)

// LifestyleExpenseEntry represents a single row in the Lifestyle Expenses table.
type LifestyleExpenseEntry struct {
    Lifestyle string
    Cost      string
}

var lifestyleExpensesTable = []LifestyleExpenseEntry{
    {"Wretched", "0 ortegs/day"},
    {"Squalid", "5 ortegs/day"},
    {"Modest", "1 shilling/day"},
    {"Comfortable", "1 thaler/day"},
    {"Wealthy", "5 thalers/day"},
    {"Aristocratic", "3 thalers/day (minimum)"},
}

type lifestyleExpenses struct{}

func init() {
    RegisterTable(&lifestyleExpenses{})
}

func (l *lifestyleExpenses) Name() string {
    return "lifestyle_expenses"
}

func (l *lifestyleExpenses) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p lifestyle_expenses' to view the full table."
}

func (l *lifestyleExpenses) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("LIFESTYLE EXPENSES (PG 177)\n")
    sb.WriteString("------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-15s | %s\n", "Lifestyle", "Cost/Day"))
    sb.WriteString("------------------------------------\n")

    for _, row := range lifestyleExpensesTable {
        sb.WriteString(fmt.Sprintf("%-15s | %s\n", row.Lifestyle, row.Cost))
    }

    return sb.String()
}

func (l *lifestyleExpenses) Category() string {
    return "DM Tools"
}
