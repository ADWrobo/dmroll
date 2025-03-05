package tables

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
)

// CuriosityEntry represents a single row in the 1d100 Curiosities table.
type CuriosityEntry struct {
    MinRoll int
    MaxRoll int
    Detail  string
}

// Full Curiosities table: 50 rows, each covering two consecutive dice results.
// E.g. (1–2), (3–4), etc.
var curiositiesTable = []CuriosityEntry{
    {1, 2, "Leather pouch containing stone marbles, fourteen shiny black ones and eight off-white, to play with or use as sling stones"},
    {3, 4, "Richly ornamented lyre of silver that can be fitted with new strings"},
    {5, 6, "Nail file with one rough side and one smooth, and a handle shaped like the silhouette of a wolf’s head"},
    {7, 8, "Box containing twenty-four paper clips of patinated copper wire"},
    {9, 10, "Glass bottle containing a wooden model castle in a moorland landscape"},
    {11, 12, "Simple frying pan of rusted iron on which nothing ever sticks"},
    {13, 14, "Set of five bronze measuring cups, all in different sizes"},
    {15, 16, "Box containing thirty-six small, bone discs engraved with Symbarian cuneiform characters"},
    {17, 18, "Small basket woven out of miraculously living vines with healthy green leaves"},
    {19, 20, "Necklace made entirely from linked gold teeth"},

    {21, 22, "Silver brooch depicting a wraith owl, the rings around its eyes made from black, faceted stones"},
    {23, 24, "White-glazed and somewhat cracked porcelain plate, painted with a weathered motif of a woman in a long dress with a poleaxe resting against her shoulder"},
    {25, 26, "Comb made of lindworm bone with cuneiform on each tooth, which not only untangles one’s hair but clears away dirt and grease"},
    {27, 28, "Cracked ceramic serving platter, decorated with various fruit on a white background – two of the fruits are unfamiliar to present-day gardeners"},
    {29, 30, "Box in the shape of a rose with twenty-four golden petals, which must be unfolded in sequence to expose its empty interior"},
    {31, 32, "Silver fountain pen with golden details, which never seems to run out of ink"},
    {33, 34, "Bronze caliper, marked in cuneiform with an unknown unit of measurement"},
    {35, 36, "Flowerpot of dark blue ceramic, decorated with a white-and-red motif of daemonic couples in obscene positions"},
    {37, 38, "Copper box containing twelve ceramic figurines, each depicting the same goblin’s head but with different facial expressions"},
    {39, 40, "Patinated sheet of copper listing the recipe for brewing the alcoholic honey beverage called Noid"},

    {41, 42, "Ornamented dagger with a dull steel blade that slides into the hilt when pressed against something"},
    {43, 44, "White-glazed porcelain mask covering the upper part of one’s face, shaped with an owl’s beak, small branched horns, and red-lined eye openings"},
    {45, 46, "Wax stamp with a handle of pure gold whose mark represents the royal rune Labrys"},
    {47, 48, "Thick, cracked tallow candle which, when lit, emits a dark yellow smoke"},
    {49, 50, "Green ceramic oil lamp with a patterned-glass shade that rotates while the lamp is lit"},
    {51, 52, "Brown leather eyepatch set with different-colored gemstones forming an open eye"},
    {53, 54, "Model of a twin-masted stone ship; the stern can be opened to reveal sixteen tin figures spread across two decks"},
    {55, 56, "Black leather leash with a matching collar large enough for a bull’s neck, decorated with studs of pure gold"},
    {57, 58, "Badly damaged wooden box with moldy fabric inside; contains five well-preserved seeds that can grow into cherry trees with multicolored blossoms (DC 13 Wisdom (Alchemist’s Supplies) to handle properly)"},
    {59, 60, "Long, thin wooden box of flat, patinated copper puzzle pieces; washing off the patina reveals a daemonic abomination motif"},

    {61, 62, "Spoon of black bone, carved with runes filled with red paint that has partially worn off; nothing eaten with it is absorbed by the body"},
    {63, 64, "A 12×9 game board of differently sized squares in various materials (bone, wood, copper), framed in patterned hardwood"},
    {65, 66, "The clappers and tubes of a wind chime made from melodious ettermite"},
    {67, 68, "Fist-sized, irregular, glistening black stone with magnetic properties"},
    {69, 70, "Silver case containing sixteen dried-up, multicolored pieces of chalk; can be restored for drawing via DC 13 Intelligence (Alchemist’s Supplies)"},
    {71, 72, "Carved bone statue of an andrik with a cracked beak, armed with a crossbow and a hatchet"},
    {73, 74, "Silver tuning fork in the key of D"},
    {75, 76, "Stiletto with a wavy blade of strange green metal; the hilt is shaped as a human child with clasped hands"},
    {77, 78, "Statuette of an ox chiseled in black stone; raising its tail tilts the horns"},
    {79, 80, "Blackened walking stick patterned like reptilian scales, topped with a fist-sized pearl for a knob"},

    {81, 82, "Circular belt or cloak buckle made of silver, with gold details depicting three crossed arrows"},
    {83, 84, "Copper spyglass with rotating segments; aiming it at the sun reveals colorful shifting patterns"},
    {85, 86, "Hourglass of misty white crystal in a blue-gray metal stand, whose sand runs exactly twice as fast in one direction as the other"},
    {87, 88, "Silver watering can with a broken spout, decorated with motifs of grotesque birds"},
    {89, 90, "Stuffed, mystically preserved ferret from a Symbarian rite of exaltation, equipped with predatory jaws, a crown of horns, and sickly yellow eyes"},
    {91, 92, "Bone flute, as long as one’s palm and as thin as a finger; if played by someone proficient in Performance, it attracts 1d20 small birds"},
    {93, 94, "A shimmering pink pearl with red streaks, eyeball-sized, sweet-sour taste like berry compote; never loses size or flavor"},
    {95, 96, "Parts of a broken dream catcher that can be repaired (DC 13 Dexterity (Smith’s Tools)); grants advantage on saves vs. Nightmares"},
    {97, 98, "A clay cruse with a wax-sealed cork, containing centuries-old red wine of excellent quality from a superb vintage"},
    {99, 100, "A glossy dark gray stone statue the size of a forearm, depicting a daemon. The base bears a cuneiform inscription: “Jeberaja”"},
}

// curiosities is our table struct implementing the Table interface.
type curiosities struct{}

func init() {
    // Register the table so the CLI knows about it.
    RegisterTable(&curiosities{})
}

func (c *curiosities) Name() string {
    return "curiosities"
}

// GetRandomEntry rolls 1d100, finds the matching row, and returns its description.
func (c *curiosities) GetRandomEntry() string {
    roll, err := dice.RollDice("1d100")
    if err != nil {
        log.Printf("Error rolling 1d100: %v", err)
        roll = 1
    }

    entry := findCuriosity(roll)
    return fmt.Sprintf("Curiosities (1d100) roll: %d\n%s", roll, entry.Detail)
}

// GetFormatted prints the entire table in ASCII form.
func (c *curiosities) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("CURIOSTIES (1d100)\n")
    sb.WriteString("-----------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-7s | %s\n", "Roll", "Detail"))
    sb.WriteString("-----------------------------------------------------------\n")

    for _, row := range curiositiesTable {
        rangeLabel := ""
        if row.MinRoll == row.MaxRoll {
            // e.g. "99" if min=99, max=99
            rangeLabel = fmt.Sprintf("%d", row.MinRoll)
        } else {
            // e.g. "1-2"
            rangeLabel = fmt.Sprintf("%d-%d", row.MinRoll, row.MaxRoll)
        }

        sb.WriteString(fmt.Sprintf("%-7s | %s\n", rangeLabel, row.Detail))
    }
    return sb.String()
}

// findCuriosity locates the correct CuriosityEntry for a given d100 roll.
func findCuriosity(roll int) CuriosityEntry {
    for _, row := range curiositiesTable {
        if roll >= row.MinRoll && roll <= row.MaxRoll {
            return row
        }
    }
    // Fallback if something is out of range
    return CuriosityEntry{
        MinRoll: roll,
        MaxRoll: roll,
        Detail:  "??? (Out of range)",
    }
}

func (m *curiosities) Category() string {
    return "Treasure"
}