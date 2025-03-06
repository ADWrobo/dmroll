package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// ExpenseEntry represents a single row in the Expenses table.
type ExpenseEntry struct {
    Service string
    Cost    string
}

var expensesTable = []ExpenseEntry{
    {"Bath at an inn", "3 ortegs"},
    {"Bed and meals, countryside (per half-day)", "5 ortegs"},
    {"Bodyguard (per day)", "1 thaler"},
    {"Camp life, foot soldier (per day)", "5 ortegs"},
    {"Camp life, rider (per day)", "1 thaler"},
    {"Camp life, knight (per day)", "5 thalers"},
    {"Cartographer", "10 thalers"},
    {"Inn, countryside (per day)", "5 ortegs"},
    {"Inn, town (per day)", "1 thaler"},
    {"Mystic, ritual", "Variable"},
    {"Road or city toll", "10 ortegs or more"},
    {"Washing of clothes", "7 ortegs"},
}

type expenses struct{}

func init() {
    table_registry.RegisterTable(&expenses{})
}

func (e *expenses) Name() string {
    return "expenses"
}

func (e *expenses) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p expenses' to view the full table."
}

func (e *expenses) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("EXPENSES (PG 173)\n")
    sb.WriteString("----------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-30s | %s\n", "Service", "Cost"))
    sb.WriteString("----------------------------------------------------\n")

    for _, row := range expensesTable {
        sb.WriteString(fmt.Sprintf("%-30s | %s\n", row.Service, row.Cost))
    }

    return sb.String()
}

func (e *expenses) Category() string {
    return "Ruins of Symbaroum 5E"
}

func (e *expenses) SubCategory() string {
    return "DM Tools"
}
