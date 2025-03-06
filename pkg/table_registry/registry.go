package table_registry

import (
    "fmt"
    "sort"
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

// ListTables lists all registered tables, sorted alphabetically by category, subcategory, and table name.
func ListTables(categoryFilter, subCategoryFilter string) []string {
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

    // Sort categories
    var sortedCategories []string
    for cat := range categoryMap {
        sortedCategories = append(sortedCategories, cat)
    }
    sort.Strings(sortedCategories)

    var result []string
    for _, cat := range sortedCategories {
        if categoryFilter != "" && !strings.EqualFold(cat, categoryFilter) {
            continue
        }

        result = append(result, fmt.Sprintf("[%s]", cat))

        // Sort subcategories
        var sortedSubCategories []string
        for subCat := range categoryMap[cat] {
            sortedSubCategories = append(sortedSubCategories, subCat)
        }
        sort.Strings(sortedSubCategories)

        for _, subCat := range sortedSubCategories {
            if subCategoryFilter != "" && !strings.Contains(strings.ToLower(subCat), strings.ToLower(subCategoryFilter)) {
                continue
            }
            result = append(result, fmt.Sprintf("  - %s:", subCat))

            // Sort table names
            sort.Strings(categoryMap[cat][subCat])
            for _, tableName := range categoryMap[cat][subCat] {
                result = append(result, "    - " + tableName)
            }
        }
    }

    return result
}