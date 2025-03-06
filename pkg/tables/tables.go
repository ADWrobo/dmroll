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
    SubCategory() string
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
    // Create a nested map: category -> subcategory -> slice of table names
    categoryMap := make(map[string]map[string][]string)
    for _, t := range knownTables {
        cat := t.Category()
        subCat := t.SubCategory()
        name := t.Name()

        if categoryMap[cat] == nil {
            categoryMap[cat] = make(map[string][]string)
        }
        categoryMap[cat][subCat] = append(categoryMap[cat][subCat], name)
    }

    // Build a final slice of strings, grouped by category and subcategory
    var result []string
    for cat, subCategories := range categoryMap {
        result = append(result, fmt.Sprintf("[%s]", cat))
        for subCat, tables := range subCategories {
            result = append(result, fmt.Sprintf("  - %s:", subCat))
            for _, tableName := range tables {
                result = append(result, "    - " + tableName)
            }
        }
    }

    return result
}