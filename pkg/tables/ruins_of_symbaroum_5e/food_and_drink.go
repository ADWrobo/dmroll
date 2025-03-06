// TODO: Look into breaking these out into their own .go files.
// This might make it easier for users to refine their search for specific food or drinks.
// However, this will be kept since this is how the table is organized on PG 178.

package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// FoodAndDrinkEntry represents a single row in the Food and Drink table.
type FoodAndDrinkEntry struct {
    Name string
    Cost string
}

// FoodAndDrinkCategory holds a category name and its corresponding food and drink entries.
type FoodAndDrinkCategory struct {
    Category string
    Entries  []FoodAndDrinkEntry
}

var foodAndDrinkTable = []FoodAndDrinkCategory{
    {"Basics", []FoodAndDrinkEntry{
        {"Bread", "2 ortegs"},
        {"Casserole", "2 ortegs"},
        {"Cheese", "2 ortegs"},
        {"Feast, countryside (per person)", "1 shilling"},
        {"Feast, town (per person)", "1 thaler"},
    }},
    {"Desserts", []FoodAndDrinkEntry{
        {"Hack tray", "3 shillings"},
        {"Fried pastry", "2 shillings"},
        {"Fruit pie", "1 thaler"},
        {"Fruit pudding", "1 thaler"},
        {"Fruit sherbet", "2 thaler"},
        {"Ice cream & fruit", "7 thaler"},
        {"Honey-roasted sorrel", "3 shillings"},
        {"Candied ginger", "3 shillings"},
        {"Marmalade confectionery", "3 shillings"},
        {"Nuts in chocolate", "5 ortegs"},
        {"Truffle-butter biscuit", "5 ortegs"},
        {"Sugar-coated rose petals", "5 ortegs"},
        {"Salty-sweet needles", "5 ortegs"},
        {"Waffles with butter and honey", "3 shillings"},
    }},
    {"Beverages", []FoodAndDrinkEntry{
        {"Table ale (watered stout)", "1 orteg"},
        {"Bottle of ale (unfiltered)", "2 ortegs"},
        {"Bottle of red undrink (mulled wine)", "3 ortegs"},
        {"Bottle of stout", "1 thaler"},
        {"Bottle of wine", "1 thaler"},
        {"Bottle of Southern Slopes (from Alberetor)", "2 thaler"},
        {"Bottle of Veerra’s Red (simple wine)", "4 ortegs"},
        {"Tankard of blackbrew (stout)", "2 ortegs"},
        {"Tankard of stutt (unspecified)", "3 ortegs"},
        {"Tankard of Adersel (triple fermented ale)", "8 ortegs"},
        {"Tankard of Argona (fine stout)", "1 shilling"},
        {"Tankard of The Duke’s Relief (simple stout)", "5 shillings"},
        {"Tankard of Kurun’s Honor (triple fermented ale)", "3 shillings"},
        {"Tankard of Urtal (triple fermented red ale)", "2 shillings"},
        {"Tankard of veloum (barbarian must)", "1 shilling"},
        {"Tankard of Zarekian Blackbrew", "1 shilling"},
    }},
    {"Fish Dishes", []FoodAndDrinkEntry{
        {"Fish sauce & crispbread", "5 ortegs"},
        {"Trout pudding with turnips", "1 shilling"},
        {"Salted herring with turnips", "5 shillings"},
        {"Buttered walleye with mash", "2 thaler"},
    }},
    {"Porridge", []FoodAndDrinkEntry{
        {"Watered porridge", "4 ortegs"},
        {"Ale-porridge with butter", "7 ortegs"},
        {"Spicy cream-porridge", "1 shilling"},
    }},
    {"Meat Dishes", []FoodAndDrinkEntry{
        {"Barbecued young-boar with beets", "6 shillings"},
        {"King’s steak in gravy", "8 shillings"},
        {"Slow-roast with stewed carrot", "8 shillings"},
        {"Hash patties with turnips", "4 shillings"},
        {"Roka sausage with mashed beets", "12 shillings"},
        {"Stuffed lung with black mash", "8 shillings"},
    }},
    {"Pies", []FoodAndDrinkEntry{
        {"Fish pie", "5 shillings"},
        {"Offal pie", "4 shillings"},
        {"Cabbage pie", "4 shillings"},
        {"Meat pie", "8 shillings"},
        {"Trout pie", "8 shillings"},
        {"Kidney pie", "8 shillings"},
        {"Mushroom pie", "4 shillings"},
    }},
    {"Stews", []FoodAndDrinkEntry{
        {"Mixed stew", "5 ortegs"},
        {"Fish & turnips", "1 shilling"},
        {"Cabbage stew", "4 ortegs"},
        {"Meat & beets", "4 ortegs"},
        {"Root vegetable stew", "4 ortegs"},
    }},
    {"Soups", []FoodAndDrinkEntry{
        {"Blood-soup with dark bread", "5 ortegs"},
        {"Onion soup with crispbread", "4 ortegs"},
    }},
    {"Teas", []FoodAndDrinkEntry{
        {"Fruit tea", "1 shilling"},
        {"Iron oak tea", "8 ortegs"},
        {"Spice tea", "1 orteg"},
        {"Smoked tea", "4 ortegs"},
        {"Herbal tea", "2 ortegs"},
    }},
}

type foodAndDrink struct{}

func init() {
    table_registry.RegisterTable(&foodAndDrink{})
}

func (f *foodAndDrink) Name() string {
    return "food_and_drink"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (f *foodAndDrink) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p food_and_drink' to view it."
}

// GetFormatted returns an ASCII-formatted table of Food and Drink with category headers.
func (f *foodAndDrink) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("FOOD AND DRINK (PG 178)\n")
    sb.WriteString("============================================================\n")

    for _, category := range foodAndDrinkTable { // Now iterates in a defined order
        sb.WriteString(fmt.Sprintf("\n%s\n", strings.ToUpper(category.Category)))
        sb.WriteString("------------------------------------------------------------\n")
        sb.WriteString(fmt.Sprintf("%-30s | %s\n", "Name", "Cost"))
        sb.WriteString("------------------------------------------------------------\n")

        for _, row := range category.Entries {
            sb.WriteString(fmt.Sprintf("%-30s | %s\n", row.Name, row.Cost))
        }
    }

    return sb.String()
}

// Category assigns the broader game system category.
func (f *foodAndDrink) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (f *foodAndDrink) SubCategory() string {
    return "Equipment and Services"
}
