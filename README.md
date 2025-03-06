# dmroll

...is a command-line tool for Dungeon Masters to roll dice, roll on predefined tables, and generate structured content like ruins for TTRPG campaigns. It is specifically populated with material from the Ruins of Symbaroum 5e Core Rulebooks, but can be expanded to include more general or setting specific material.

<i>tl;dr I was tired of GUIs plagued with ads, so I wanted to make my own CLI to streamline my DM'ing.</i>

## Features

- **Roll Dice**: Supports standard dice notation (e.g., `1d20`).
- **Table Operations**:
  - List available tables.
  - Print an entire table in ASCII format.
  - Roll a random entry from a table.
- **Scenario Generation**:
  - Build structured ruins using multiple random tables.

## Installation

Clone the repository and build the executable:

```sh
git clone https://github.com/yourusername/dmroll.git
cd dmroll
go build -o dmroll ./cmd/dmroll
```
## Building Binaries and Retrieving SHAs (For Homebrew Things)

TODO: Implement some kind of build/release automation in concert with the homebrew tap formula. 

```sh
# MacOS or Linux
GOOS=darwin GOARCH=amd64 go build -o dmroll
GOOS=darwin GOARCH=arm64 go build -o dmroll-arm64
GOOS=linux GOARCH=amd64 go build -o dmroll-linux

shasum -a 256 dmroll         # macOS (Intel)
shasum -a 256 dmroll-arm64   # macOS (Apple Silicon)
shasum -a 256 dmroll-linux   # Linux

# Windows (not yet supported, but for future reference)
set GOOS=windows && set GOARCH=amd64 && go build -o dmroll.exe
CertUtil -hashfile dmroll.exe SHA256
```

## Usage

### Rolling Dice

```sh
dmroll -r 1d20
```

### Table Operations

List all available tables:

```sh
dmroll -t -l
```

List available tables in a category or sub-category:

```sh
dmroll -t -l <category>
dmroll -t -l <subcategory>
```

Print an entire table:

```sh
dmroll -t -p <table_name>
```

Roll on a table:

```sh
dmroll -t <table_name>
```

### Building a Ruin
This is an engine from GG78, combining several tables and additional logic for producing random ruins.

```sh
dmroll -b ruin
```

## Example Output

### Rolling Dice

```sh
$ dmroll -r 1d20

14
```

### Rolling on a Table

```sh
$ dmroll -t wilderness_guides

Galdor the Swift: A perceptive and wise elf guide with knowledge of hidden paths.
```

### Printing a Table

```sh
$ dmroll -t -p wilderness_guides

--------------------------------------------------
WILDERNESS GUIDES
--------------------------------------------------
1. Galdor the Swift - Perception +3, Survival +5
2. Ulfgar the Stout - Tough as nails dwarf scout
...
```

### Building a Ruin

```sh
$ dmroll -b ruin

RUIN GENERATION
==================================================

Original Purpose: Ancient Temple

Overall Feature: Overgrown with vines

Overall Trait: Haunted by spectral whispers

...
==================================================
RUIN GENERATION COMPLETE
```

## Directory Structure

```
dmroll/
│── cmd/dmroll/        # Main entry point
│    └── main.go       # Main application logic
│── pkg/               # Core packages
│    ├── dice/         # Dice rolling functions
│    │    ├── dice.go
│    │    └── dx.go files
│    ├── table_registry/ # Table registry
│    │    └── registry.go
│    ├── tables/       # Table management
│    │    ├── ruins_of_symbaroum_5e/
│    │    │    └── [various tables].go
│    │    ├── ruins_of_symbaroum_5e_modifiers.go
│    │    └── tables.go
│── README.md          # Documentation
│── go.mod             # Module dependencies
```

## License

This project is licensed under the MIT License.