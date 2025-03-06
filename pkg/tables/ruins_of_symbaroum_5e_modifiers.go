package tables

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"
)

// PromptForDavokarRegion prompts the user to pick Bright, Wild, or Dark Davokar.
// Returns one of: "bright", "wild", or "dark".
func PromptForDavokarRegion() string {
    fmt.Println()
    fmt.Println("Where is the party in Davokar?")
    fmt.Println(" (1) Bright Davokar")
    fmt.Println(" (2) Wild Davokar")
    fmt.Println(" (3) Dark Davokar")
    fmt.Print("Enter choice (number or text): ")
    fmt.Println()

    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Println("Error reading input, defaulting to Bright Davokar.")
        return "bright"
    }
    input = strings.TrimSpace(input)

    // Try numeric
    if numericValue, err := strconv.Atoi(input); err == nil {
        switch numericValue {
        case 2: return "wild"
        case 3: return "dark"
        default: return "bright"
        }
    }

    // Otherwise parse textual
    lower := strings.ToLower(input)
    re := regexp.MustCompile(`[^a-z]+`)
    cleaned := re.ReplaceAllString(lower, "")

    switch cleaned {
    case "brightdavokar", "bright":
        return "bright"
    case "wilddavokar", "wild":
        return "wild"
    case "darkdavokar", "dark":
        return "dark"
    }

    log.Println("Unrecognized region; defaulting to Bright Davokar.")
    return "bright"
}

// PromptForDavokarModifier prompts the user for environment (Light, Wild, Dark, or Waterways).
// Returns +0, +2, or +5 based on user input.
func PromptForDavokarModifier() int {
    fmt.Println()
    fmt.Println("The party is in:")
    fmt.Println(" (1) Bright Davokar")
    fmt.Println(" (2) Wild Davokar")
    fmt.Println(" (3) Dark Davokar")
    fmt.Println(" (4) Traveling along waterways")
    fmt.Print("Enter choice (number or text): ")
	fmt.Println()

    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Println("Error reading input, defaulting to Light Davokar (+0).")
        return 0
    }
    input = strings.TrimSpace(input)

    // Try numeric
    if numericValue, err := strconv.Atoi(input); err == nil {
        return environmentModFromNumeric(numericValue)
    }
    // Otherwise, parse text
    return environmentModFromString(input)
}

// environmentModFromNumeric maps 1 => 0, 2 => 2, 3 => 5, 4 => 2
func environmentModFromNumeric(n int) int {
    switch n {
    case 1:
        return 0 // Bright Davokar
    case 2, 4:
        return 2 // Wild Davokar or Waterways
    case 3:
        return 5 // Dark Davokar
    }
    log.Println("Unrecognized numeric environment. Defaulting to Light Davokar (+0).")
    return 0
}

// environmentModFromString handles strings like "dark davokar", "wild", etc.
func environmentModFromString(s string) int {
    lower := strings.ToLower(s)
    reNonAlpha := regexp.MustCompile(`[^a-z]+`)
    cleaned := reNonAlpha.ReplaceAllString(lower, "")

    switch cleaned {
    case "brightdavokar", "bright":
        return 0
    case "wilddavokar", "wild", "waterways", "travelingalongwaterways":
        return 2
    case "darkdavokar", "dark":
        return 5
    }

    log.Println("Unrecognized environment string. Defaulting to Light Davokar (+0).")
    return 0
}
