package table_registry

import (
    "fmt"
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

// knownTables stores all registered tables.
var knownTables = map[string]Table{}

// RegisterTable registers a table globally.
func RegisterTable(t Table) {
    tableName := strings.ToLower(t.Name())
    knownTables[tableName] = t
}

// GetTable returns a table by name.
func GetTable(name string) (Table, bool) {
    t, found := knownTables[strings.ToLower(name)]
    return t, found
}

// ListTables lists all registered tables, filtered by category and/or subcategory.
func ListTables(categoryFilter, subCategoryFilter string) []string {
    // Create a nested map: category -> subcategory -> slice of table names
    categoryMap := make(map[string]map[string][]string)
    for _, t := range knownTables {
        cat := strings.ToLower(t.Category())
        subCat := strings.ToLower(t.SubCategory())
        name := t.Name()

        if categoryMap[cat] == nil {
            categoryMap[cat] = make(map[string][]string)
        }
        categoryMap[cat][subCat] = append(categoryMap[cat][subCat], name)
    }

    var result []string
    for cat, subCategories := range categoryMap {
        if categoryFilter != "" && strings.ToLower(cat) != categoryFilter {
            continue
        }        

        result = append(result, fmt.Sprintf("[%s]", cat))
        for subCat, tables := range subCategories {
            if subCategoryFilter != "" && !strings.Contains(strings.ToLower(subCat), subCategoryFilter) {
                continue
            }
            result = append(result, fmt.Sprintf("  - %s:", subCat))
            for _, tableName := range tables {
                result = append(result, "    - " + tableName)
            }
        }
    }
    
    return result
}