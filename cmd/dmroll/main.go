package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "math/rand"
    "time"

    "dmroll/pkg/dice"
    "dmroll/pkg/tables"
    _ "dmroll/pkg/tables/ruins_of_symbaroum_5e"
)

var (
    // For rolling dice
    rollDiceFlag = flag.String("r", "", "Roll dice notation (e.g., 1d20)")

    // Table operations
    tableOpFlag    = flag.Bool("t", false, "Operate on tables instead of dice")
    listTablesFlag = flag.Bool("l", false, "List all available tables (use with -t)")
    printTableFlag = flag.String("p", "", "Print out a specified table (use with -t)")
    buildFlag = flag.String("b", "", "Build a large structure/scenario (e.g. 'ruin')")
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

// Helper function that prints one blank line, then prints `msg`.
func printWithNewline(msg string) {
    fmt.Println()
    fmt.Println(msg)
}

func main() {
    flag.Parse()

    // If no flags are provided, show usage and exit
    if len(os.Args) < 2 {
        usage()
        os.Exit(1)
    }

    // 1) Dice Rolling (e.g. dmroll -r 1d20)
    if *rollDiceFlag != "" {
        result, err := dice.RollDice(*rollDiceFlag)
        if err != nil {
            log.Fatalf("Error rolling dice: %v\n", err)
        }
        // Print with a blank line before
        printWithNewline(fmt.Sprintf("%d", result))
        return
    }

    // 2) Table Operations: must have -t
    if *tableOpFlag {
        leftoverArgs := flag.Args()

        // 2A) List Tables: dmroll -t -l
        if *listTablesFlag {
            var categoryFilter, subCategoryFilter string

            // Check if additional arguments are provided for filtering
            if len(leftoverArgs) > 0 {
                if len(leftoverArgs) == 1 {
                    subCategoryFilter = strings.ToLower(leftoverArgs[0]) // Only subcategory
                } else if len(leftoverArgs) == 2 {
                    categoryFilter = strings.ToLower(leftoverArgs[0])    // Category
                    subCategoryFilter = strings.ToLower(leftoverArgs[1]) // Subcategory
                }
            }

            tablesList := tables.ListTables(categoryFilter, subCategoryFilter)
            if len(tablesList) == 0 {
                printWithNewline("No tables match the given filters.")
                return
            }

            msg := "Available tables:"
            for _, line := range tablesList {
                msg += "\n" + line
            }
            printWithNewline(msg)
            return
        }

        // 2B) Print a Table: dmroll -t -p table_name
        if *printTableFlag != "" {
            if len(leftoverArgs) > 0 {
                log.Fatalf("Too many arguments after -p. Usage: dmroll -t -p <table_name>")
            }
            tableName := *printTableFlag
            output, err := tables.PrintTable(tableName)
            if err != nil {
                log.Fatalf("Error printing table %q: %v\n", tableName, err)
            }
            printWithNewline(output)
            return
        }

        // 2C) Roll on a specific table: dmroll -t <table_name>
        if len(leftoverArgs) > 0 {
            tableName := leftoverArgs[0]
            entry, err := tables.RollOnTable(tableName)
            if err != nil {
                log.Fatalf("Error rolling on table %q: %v\n", tableName, err)
            }
            printWithNewline(entry)
            return
        }

        // 2D) If user typed only `dmroll -t` with no further args
        log.Fatalf("No table name specified. Try 'dmroll -t -l' to list tables or 'dmroll -t -p <table>' to print one.")
    }

    // 3) Build a large structure/scenario
    if *buildFlag != "" {
        switch *buildFlag {
        case "ruin":
            result := buildRuin()
            printWithNewline(result)
            return
        default:
            log.Fatalf("Unknown build type: %s (try 'ruin')", *buildFlag)
        }
    }

    // If we reach here, the user didn't provide valid flags
    usage()
}

// usage prints command usage instructions
func usage() {
    fmt.Println()
    fmt.Println("Usage examples:")
    fmt.Println("  dmroll -r 1d20               Roll dice notation (e.g. '1d20')")
    fmt.Println("  dmroll -t -l                 List all registered tables")
    fmt.Println("  dmroll -t <table_name>       Roll a random entry from <table_name>")
    fmt.Println("  dmroll -t -p <table_name>    Print the entire <table_name> in ASCII")
    fmt.Println("  dmroll -b ruin                Build a ruin scenario from multiple tables")
}

func buildRuin() string {
    var sb strings.Builder

    // Step 1: Roll on Ruin Original Purpose Table
    originalPurpose, _ := tables.RollOnTable("ruin_original_purpose")
    sb.WriteString("RUIN GENERATION\n")
    sb.WriteString("==================================================\n\n")
    sb.WriteString("Original Purpose: " + originalPurpose + "\n\n")

    // Step 2: Roll on Ruin Overall Features Table
    overallFeature, _ := tables.RollOnTable("ruin_overall_features")
    sb.WriteString("Overall " + overallFeature + "\n")

    // Step 3: Roll on Ruin Overall Traits Table
    overallTrait, _ := tables.RollOnTable("ruin_overall_traits")
    sb.WriteString("Overall " + overallTrait + "\n\n")

    // Step 4: Roll on Ruin Inhabitants Table (with 18-20 handling)
    inhabitants, _ := tables.RollOnTable("ruin_inhabitants")
    var secondInhabitants string
    var relationship string
    if strings.Contains(inhabitants, "18-20") {
        // Keep rolling until two results (that aren't "18-20") are obtained
        for {
            first, _ := tables.RollOnTable("ruin_inhabitants")
            if strings.Contains(first, "18-20") {
                continue
            }
            second, _ := tables.RollOnTable("ruin_inhabitants")
            if strings.Contains(second, "18-20") {
                continue
            }
            inhabitants, secondInhabitants = first, second
            break
        }
        // Roll on the Inhabitants Relationship Table only if the first roll was 18-20
        relationship, _ = tables.RollOnTable("ruin_inhabitants_relationship")
    }
    sb.WriteString(inhabitants)
    if secondInhabitants != "" {
        sb.WriteString(" & " + strings.TrimPrefix(secondInhabitants, "Inhabitants: ") + "\n")
        sb.WriteString(relationship + "\n")
    } else {
        sb.WriteString("\n")
    }

    // Step 5: Determine levels
    aboveGroundRoll, _ := dice.RollDice("1d20")
    belowGroundRoll, _ := dice.RollDice("1d10")
    aboveGroundLevels := aboveGroundRoll - 10
    if aboveGroundLevels < 0 {
        aboveGroundLevels = 0
    }
    // Include the main level (0)
    totalLevels := aboveGroundLevels + belowGroundRoll + 1

    sb.WriteString(fmt.Sprintf("\nAbove Ground Levels: %d\n", aboveGroundLevels))
    sb.WriteString(fmt.Sprintf("Below Ground Levels: %d\n\n", belowGroundRoll))

    // Build a slice of level numbers.
    // Example: if aboveGroundLevels is 2 and belowGroundRoll is 3,
    // levels will be: [2, 1, 0, -1, -2, -3]
   	levels := make([]int, totalLevels)
    for i := 0; i < aboveGroundLevels; i++ {
        levels[i] = aboveGroundLevels - i
    }
    levels[aboveGroundLevels] = 0
    for i := aboveGroundLevels + 1; i < totalLevels; i++ {
        levels[i] = -1 * (i - aboveGroundLevels)
    }

    // Process each level and, if applicable, the transition to the next level
    for idx, levelNum := range levels {
        // Level details
        sb.WriteString(fmt.Sprintf("LEVEL %d\n", levelNum))
        sb.WriteString("--------------------------------------------------\n")

        // Step 6.2: Roll for the number of rooms on this level (1d8)
        numRooms, _ := dice.RollDice("1d8")
        sb.WriteString(fmt.Sprintf("\n  Rooms on this level: %d\n", numRooms))

        // Step 6.3: Process each room on this level
        for room := 1; room <= numRooms; room++ {
            roomEntryway, _ := tables.RollOnTable("ruin_entryways_to_other_rooms")
            roomDetail, _ := tables.RollOnTable("ruin_details_regarding_the_room")
            sb.WriteString(fmt.Sprintf("    Room %d:\n", room))
            sb.WriteString(fmt.Sprintf("      %s\n", roomEntryway))
            sb.WriteString(fmt.Sprintf("      %s\n", roomDetail))
        }
        sb.WriteString("\n")

        // If there's a next level, process the transition (level connection)
        if idx < len(levels)-1 {
            nextLevel := levels[idx+1]
            sb.WriteString(fmt.Sprintf("TRANSITION to LEVEL %d\n", nextLevel))
            sb.WriteString("--------------------------------------------------\n")
            // Roll 1d2 entryways for this level transition
            numEntryways, _ := dice.RollDice("1d2")
            for i := 0; i < numEntryways; i++ {
                entryway, _ := tables.RollOnTable("ruin_entryways_to_other_levels")
                safeguard, _ := tables.RollOnTable("ruin_safeguards_for_entryways")
                sb.WriteString(fmt.Sprintf("  %s\n", entryway))
                sb.WriteString(fmt.Sprintf("  %s\n", safeguard))
            }
            sb.WriteString("\n")
        }
    }

    sb.WriteString("==================================================\n")
    sb.WriteString("RUIN GENERATION COMPLETE\n")

    return sb.String()
}
