package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// FarmAnimalEntry represents a single row in the Farm Animals table.
type FarmAnimalEntry struct {
    Name        string
    Cost        string
}

var farmAnimalsTable = []FarmAnimalEntry{
    {"Bull", "10 thaler"},
    {"Chicken", "8 ortegs"},
    {"Cow", "1 thaler"},
    {"Dog", "1 shilling"},
    {"Ox", "3 thaler"},
    {"Pig", "1 thaler"},
    {"Rooster", "5 shillings"},
    {"Sheep", "15 ortegs"},
}

type farmAnimals struct{}

func init() {
    table_registry.RegisterTable(&farmAnimals{})
}

func (f *farmAnimals) Name() string {
    return "farm_animals"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (f *farmAnimals) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p farm_animals' to view it."
}

// GetFormatted returns an ASCII-formatted table of Farm Animals.
func (f *farmAnimals) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("FARM ANIMALS (PG 173)\n")
    sb.WriteString("------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-20s | %s\n", "Name", "Cost"))
    sb.WriteString("------------------------------------------------------------\n")

    for _, row := range farmAnimalsTable {
        sb.WriteString(fmt.Sprintf("%-20s | %s\n", row.Name, row.Cost))
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (f *farmAnimals) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (f *farmAnimals) SubCategory() string {
    return "Equipment and Services"
}
