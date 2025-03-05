package tables

import (
    "fmt"
    "log"
    "strings"

    "dmroll/pkg/dice"
)

// TreasureEntry represents one row in the Mystical Treasure table.
type TreasureEntry struct {
    RollRangeMin int
    RollRangeMax int
    Detail       string
}

// mysticalTreasureTable holds each entry for rolls 1..20.
var mysticalTreasureTable = []TreasureEntry{
    {
        1, 1,
        "Tall, thin vase of glazed porcelain, purple with a motif of bright red vines. " +
            "Plants placed in the vase will not age, wither, bloom, or bear fruit.",
    },
    {
        2, 2,
        "Facial mask made of seashells in pink, white, and light blue. " +
            "Its wearer can hold their breath underwater at depths up to 30 ft (10m) for up to 10 minutes.",
    },
    {
        3, 3,
        "A handheld fan painted with stormy mountain scenery, actively cooling the air when waved. " +
            "Roll 1d6 each time it’s used; on a 6 it inflicts 1d4 cold damage and freezes facial hair.",
    },
    {
        4, 4,
        "A bronze bell whose inside and clapper are engraved with cuneiform characters. " +
            "If carried so it hangs freely, it faintly tinkles whenever there is an abomination within 100 m, " +
            "but at the risk of being detected (DC 15 Wis(Perception)).",
    },
    {
        5, 5,
        "Notebook with thin pages of parchment and a pearl-covered binding. " +
            "Text within is only visible to the person who wrote it.",
    },
    {
        6, 6,
        "Metal canister whose white paint has mostly peeled off. " +
            "It’s half full of a powder that conceals signs of aging on human skin for 24 hours, " +
            "but at the cost of 1d4 temporary Corruption. Contains enough for 10 + 1d10 applications.",
    },
    {
        7, 7,
        "An amulet set with a gray stone that exudes a slimy secretion, drying into a useless resin. " +
            "The secretion is enough for two individuals per day. If applied to skin, " +
            "it repels insects (disadvantage on an insect’s melee attacks).",
    },
    {
        8, 8,
        "Eight sticks of incense in a cylindrical leather case. " +
            "They can be burned during negotiations to grant the user advantage on Persuasion checks.",
    },
    {
        9, 9,
        "A wedding ring of pure gold adorned with six sapphires. " +
            "A loving spirit is bound to it and can defend the wearer: once per short rest, " +
            "the wearer has advantage on a saving throw (including death saves). " +
            "However, the jealous spirit will haunt the wearer’s romantic partner with nightmares.",
    },
    {
        10, 10,
        "A smooth wooden box containing two frameless spectacle lenses. " +
            "They provide advantage (or a bonus) on Survival checks but impose disadvantage " +
            "(or a penalty) on Perception checks for the wearer.",
    },
    {
        11, 11,
        "Three iron-cast animal heads (wolf, bear, boar), each fist-sized and rust-free. " +
            "Loops underneath allow them to be tied into a bola, granting advantage when thrown at Beast creatures.",
    },
    {
        12, 12,
        "A long-stemmed briar pipe shaped like a woman’s face with closed eyes. " +
            "Used during a commune ritual, its smoke conveys visions, allowing either a 4th question " +
            "or more clarity on previous answers. Requires a rest (short or long) to recharge.",
    },
    {
        13, 13,
        "A lantern with red glass in a metal frame, the size of a tankard. When lit in dim light, " +
            "it can expose people trying to hide intense feelings (love, outrage, grief). " +
            "It takes a DC 15 Wis(Perception) check per target to notice their glowing outlines.",
    },
    {
        14, 14,
        "A simple right-handed chainmail gauntlet needing relining. " +
            "The wearer can never involuntarily lose their grip on a right-hand weapon.",
    },
    {
        15, 15,
        "A modest ring of low-karat gold inscribed with cuneiform runes (“The cleansing one”). " +
            "Once per day it can conceal any skin damage (scars, tattoos) for one scene, though " +
            "true sight or similar can still detect them.",
    },
    {
        16, 16,
        "A key of blackened silver that can slide into most locks. Roll 1d6: " +
            "on 1–4, the lock opens; on 5–6, the lock melts around the key, making it impossible to unlock.",
    },
    {
        17, 17,
        "A staff of smoke-colored crystal, about as long as one’s palm, " +
            "with one end shaped like a rounded point. Pointing it at the North Star " +
            "causes a faint glow, granting advantage on Wis(Survival) checks to navigate, even underground.",
    },
    {
        18, 18,
        "A small copper pot, large enough to brew a single dose of elixir. " +
            "Adding silver equal to 1 thaler makes the elixir one level stronger.",
    },
    {
        19, 19,
        "A knife with a scarlet wooden handle and a 10 cm curved blade, engraved with a running predator. " +
            "Grants advantage on Dexterity checks when harvesting monster trophies.",
    },
    {
        20, 20,
        "A cracked, uncut ruby that emits enough heat to steam in cold/damp weather. " +
            "If set into a one-handed weapon’s hilt, it can be activated with a bonus action. " +
            "While active, each successful attack deals +1d10 necrotic damage and " +
            "the bearer gains +1 temporary Corruption each round.",
    },
}

// mysticalTreasure implements the Table interface so the CLI can use it.
type mysticalTreasure struct{}

func init() {
    RegisterTable(&mysticalTreasure{})
}

func (m *mysticalTreasure) Name() string {
    return "mystical_treasure"
}

// GetRandomEntry rolls 1d20, then returns the matching item’s description.
func (m *mysticalTreasure) GetRandomEntry() string {
    roll, err := dice.RollDice("1d20")
    if err != nil {
        log.Printf("Error rolling 1d20: %v", err)
        roll = 1
    }
    entry := findMysticalTreasure(roll)
    return fmt.Sprintf("Mystical Treasure (1d20) roll: %d\n%s", roll, entry.Detail)
}

// GetFormatted prints an ASCII version of the entire table for reference.
func (m *mysticalTreasure) GetFormatted() string {
    sb := &strings.Builder{}
    sb.WriteString("MYSTICAL TREASURE (1d20)\n")
    sb.WriteString("---------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-5s | %s\n", "Roll", "Detail"))
    sb.WriteString("---------------------------------------------------------\n")
    for _, row := range mysticalTreasureTable {
        rangeLabel := fmt.Sprintf("%d", row.RollRangeMin)
        sb.WriteString(fmt.Sprintf("%-5s | %s\n", rangeLabel, row.Detail))
    }
    return sb.String()
}

// findMysticalTreasure locates the correct row in mysticalTreasureTable.
func findMysticalTreasure(roll int) TreasureEntry {
    for _, e := range mysticalTreasureTable {
        if roll >= e.RollRangeMin && roll <= e.RollRangeMax {
            return e
        }
    }
    // Fallback if the roll is out of range (shouldn't happen with 1..20).
    return TreasureEntry{
        RollRangeMin: roll,
        RollRangeMax: roll,
        Detail:       "??? (No entry found)",
    }
}

func (m *mysticalTreasure) Category() string {
    return "Treasure"
}