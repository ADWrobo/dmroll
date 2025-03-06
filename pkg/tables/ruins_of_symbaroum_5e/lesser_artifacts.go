// TODO: Should this table be rollable? It is technically not in the PG...

package ruins_of_symbaroum_5e

import (
    "fmt"
    "strings"

    "dmroll/pkg/table_registry"
)

// LesserArtifactEntry represents a single row in the Lesser Artifacts table.
type LesserArtifactEntry struct {
    Name  string
    Cost  string
}

// LesserArtifactCategory holds a category name and its corresponding artifact entries.
type LesserArtifactCategory struct {
    Category string
    Entries  []LesserArtifactEntry
}

var lesserArtifactsTable = []LesserArtifactCategory{
    {"Artifacts", []LesserArtifactEntry{
        {"Animal Mask", "50 thaler"},
        {"Bark Mask", "50 thaler"},
        {"Burial Shroud", "50 thaler"},
        {"Death Mask", "50 thaler"},
        {"Healing Spider", "20 thaler"},
        {"Iron Crown", "50 thaler"},
        {"Lucky Coin", "40 thaler"},
        {"Marlit Cape", "20 thaler"},
        {"Meeting Stone", "10 thaler"},
        {"Mind Prism", "50 thaler"},
        {"Mystical Focus", "60 thaler"},
        {"Order Medallion", "10 thaler"},
        {"Pest Mask", "40 thaler"},
        {"Ritual Codex", "20 thaler"},
        {"Ritual Focus", "40 thaler"},
        {"Ritual Seal", "60 thaler"},
        {"Ruler’s Ring", "50 thaler"},
        {"Rune Staff", "60 thaler"},
        {"Soul Stone", "100 thaler"},
        {"Spark Stone", "50 thaler"},
		{"Staff Foot*", "50 thaler"},
        {"Staff Head*", "50 thaler"},
        {"Sun Mask", "50 thaler"},
        {"Toad Guard", "5 thaler"},
        {"Transcendental Weapon", "60 thaler"},
        {"Witch Braid", "40 thaler"},
    }},
    {"Spell Seal", []LesserArtifactEntry{
        {"Novice", "15 thaler per spell level"},
        {"Adept", "25 thaler per spell level"},
    }},
    {"Spell Scroll", []LesserArtifactEntry{
        {"Novice", "10 thaler per spell level"},
        {"Adept", "20 thaler per spell level"},
        {"Master", "30 thaler per spell level"},
    }},
}

type lesserArtifacts struct{}

func init() {
    table_registry.RegisterTable(&lesserArtifacts{})
}

func (l *lesserArtifacts) Name() string {
    return "lesser_artifacts"
}

// GetRandomEntry is not applicable since this is a non-rollable table.
func (l *lesserArtifacts) GetRandomEntry() string {
    return "This table is not rollable. Use 'dmroll -t -p lesser_artifacts' to view it."
}

// GetFormatted returns an ASCII-formatted table of Lesser Artifacts with category headers.
func (l *lesserArtifacts) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("LESSER ARTIFACTS (PG 184)\n")
    sb.WriteString("================================================================\n")

    for _, category := range lesserArtifactsTable { // Ensures ordered iteration
        sb.WriteString(fmt.Sprintf("\n%s\n", strings.ToUpper(category.Category)))
        sb.WriteString("----------------------------------------------------------------\n")
        sb.WriteString(fmt.Sprintf("%-30s | %s\n", "Name", "Cost"))
        sb.WriteString("----------------------------------------------------------------\n")

        for _, row := range category.Entries {
            sb.WriteString(fmt.Sprintf("%-30s | %s\n", row.Name, row.Cost))
        }
    }

	// Append important notes at the end
	sb.WriteString("\nNOTES:\n")
	sb.WriteString("* These artifacts do not need attunement.\n")

    return sb.String()
}

// Category assigns the broader game system category.
func (l *lesserArtifacts) Category() string {
    return "Ruins of Symbaroum 5E"
}

// SubCategory assigns a more specific classification under the main category.
func (l *lesserArtifacts) SubCategory() string {
    return "Treasure"
}
