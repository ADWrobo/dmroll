package tables

import (
    "fmt"
    "math/rand"
    "time"
    "strings"
)

// Table is the interface that all table files must implement.
type Table interface {
    Name() string
    GetRandomEntry() string
    GetFormatted() string
    Category() string
}

// knownTables is a registry of all tables your CLI can interact with.
var knownTables = map[string]Table{}

// RegisterTable allows each table file to register itself at init().
func RegisterTable(t Table) {
    tableName := strings.ToLower(t.Name())
    knownTables[tableName] = t
}

// RollOnTable picks a random entry from the named table.
func RollOnTable(tableName string) (string, error) {
    t, found := knownTables[strings.ToLower(tableName)]
    if !found {
        return "", fmt.Errorf("table %q not found", tableName)
    }
    return t.GetRandomEntry(), nil
}

// PrintTable returns the entire table in a formatted ASCII string.
func PrintTable(tableName string) (string, error) {
    t, found := knownTables[strings.ToLower(tableName)]
    if !found {
        return "", fmt.Errorf("table %q not found", tableName)
    }
    return t.GetFormatted(), nil
}

// Utility function for random index in a slice length
func randomIndex(sliceLength int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(sliceLength)
}

func ListTables() []string {
    // We'll create a map: category -> slice of table names
    categoryMap := make(map[string][]string)
    for _, t := range knownTables {
        cat := t.Category()
        name := t.Name()
        categoryMap[cat] = append(categoryMap[cat], name)
    }

    // Build a final slice of strings, grouped by category
    var result []string
    for cat, tables := range categoryMap {
        result = append(result, fmt.Sprintf("[%s]", cat))
        for _, tableName := range tables {
            result = append(result, "  - " + tableName)
        }
    }

    return result
}