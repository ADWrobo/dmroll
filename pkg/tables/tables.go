package tables

import (
    "fmt"
    "math/rand"
    "time"

    "dmroll/pkg/table_registry"
)

type Table = table_registry.Table

// RollOnTable picks a random entry from the named table.
func RollOnTable(tableName string) (string, error) {
    t, found := table_registry.GetTable(tableName) // Get table from registry
    if !found {
        return "", fmt.Errorf("table %q not found", tableName)
    }
    return t.GetRandomEntry(), nil
}

// PrintTable returns the entire table in a formatted ASCII string.
func PrintTable(tableName string) (string, error) {
    t, found := table_registry.GetTable(tableName) // Get table from registry
    if !found {
        return "", fmt.Errorf("table %q not found", tableName)
    }
    return t.GetFormatted(), nil
}

// Utility function for random index in a slice length
func RandomIndex(sliceLength int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(sliceLength)
}

// ListTables now delegates to table_registry.ListTables()
func ListTables(categoryFilter, subCategoryFilter string) []string {
    return table_registry.ListTables(categoryFilter, subCategoryFilter)
}
