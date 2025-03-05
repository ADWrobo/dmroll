package tables

import (
    "fmt"
    "strings"
)

// WildernessGuide holds a single row of data.
type WildernessGuide struct {
    Name        string
    Wisdom      string
    Perception  string
    Features    string
}

// wildernessGuidesTable is an in-memory list of data entries.
var wildernessGuidesTable = []WildernessGuide{
    {"Fonsiul", "16(+3)", "+5", "Famous goblin with a temper; already contracted by Ordo Magica"},
    {"Arval", "15(+2)", "+4", "Timid youngster from Clan Baiaga; already contracted as a reserve by the Sun Church"},
    {"Crooky", "17(+3)", "+7", "Stoop-shouldered ogre with a sense of humor; planning her own expedition"},
    {"Vindel", "17(+3)", "+6", "Elderly Ambrian, retired from the Queen's Rangers; newly returned, lightly wounded, terrible mood"},
    {"Hygla", "18(+4)", "+7", "Scarred, love-struck goblin; pregnant and reluctant to embark on longer journeys"},
    {"Erdamir", "14(+2)", "+5", "Charismatic Ambrian rogue; was a poacher before taking on guiding work"},
    {"Sula", "12(+1)", "+3", "Thick-skinned barbarian woman; extensive knowledge of edible plants"},
    {"Thraska", "19(+4)", "+8", "One-armed troll, highly experienced but extremely slow-moving"},
    {"Nissel", "11(+0)", "+2", "Young, arrogant Ambrian with a nose for trouble but no patience for fools"},
    {"Torvik", "16(+3)", "+6", "A scarred ex-mercenary who drinks heavily but never gets lost"},
    {"Yzorra", "13(+1)", "+4", "Former witch who abandoned her coven; still feared in some villages"},
    {"Brugg", "10(+0)", "+2", "Brawny ogre who guides for sport, not coin; may leave group behind if bored"},
    {"Keleb", "20(+5)", "+9", "Mysterious Changeling deeply connected to the forest; rumored to be cursed"},
    {"Ordwin", "9(-1)", "+1", "Half-blind veteran tracker, but with sharp intuition"},
    {"Velka", "15(+2)", "+5", "Morose Ambrian scholar studying Davokar’s shifting paths"},
    {"Jorren", "18(+4)", "+8", "Once served as a noble's bodyguard; sole survivor of an expedition"},
    {"Lutha", "14(+2)", "+4", "Soft-spoken barbarian woman with a wolf companion"},
    {"Belsan", "17(+3)", "+6", "Greedy, pragmatic goblin; will sell secrets to highest bidder"},
    {"Hargil", "13(+1)", "+3", "Young Ambrian with a heart full of stories but little experience"},
    {"Vredja", "12(+1)", "+4", "Heavyset ogre known for her intricate Davokar maps"},
    {"Morn", "16(+3)", "+6", "Stern Ambrian ranger who lost his previous group to corruption"},
    {"Tassar", "10(+0)", "+2", "Daring but reckless guide; always chasing treasure rumors"},
    {"Edrik", "19(+4)", "+8", "Elderly goblin with uncanny luck; guided by forest spirits"},
    {"Aska", "8(-1)", "+1", "Wide-eyed young guide desperate to prove herself, but lacks confidence"},
    {"Guldar", "14(+2)", "+5", "Soft-spoken Changeling who claims to hear the trees whisper"},
}

// wildernessGuides implements the Table interface
type wildernessGuides struct {}

// init automatically registers this table when this file is imported.
func init() {
    RegisterTable(&wildernessGuides{})
}

// Name returns a string identifying the table.
func (wg *wildernessGuides) Name() string {
    return "wilderness_guides"
}

// GetRandomEntry returns a single random entry from the table.
func (wg *wildernessGuides) GetRandomEntry() string {
    i := randomIndex(len(wildernessGuidesTable))
    row := wildernessGuidesTable[i]
    return fmt.Sprintf("%s | Wis: %s | Perc: %s | %s",
        row.Name, row.Wisdom, row.Perception, row.Features)
}

// GetFormatted returns an ASCII-formatted table of all entries.
func (wg *wildernessGuides) GetFormatted() string {
    var sb strings.Builder

    // Headers
    sb.WriteString("WILDERNESS GUIDES\n")
    sb.WriteString("------------------------------------------------------------------------------------\n")
    sb.WriteString(fmt.Sprintf("%-10s | %-8s | %-10s | %s\n",
        "Name", "Wisdom", "Perception", "Features"))
    sb.WriteString("------------------------------------------------------------------------------------\n")

    // Rows
    for _, row := range wildernessGuidesTable {
        sb.WriteString(fmt.Sprintf("%-10s | %-8s | %-10s | %s\n",
            row.Name, row.Wisdom, row.Perception, row.Features))
    }

    return sb.String()
}

func (m *wildernessGuides) Category() string {
    return "People"
}