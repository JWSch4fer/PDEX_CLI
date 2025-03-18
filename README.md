# PDEX_CLI
PDEX_CLI is a command-line interface tool for the PokeAPI to allow you to explore the Pokémon world from your terminal. Search for locations, discover Pokémon, try your luck at catching them, and manage a Pokédex—all in one interactive REPL.

## Features
- **Interactive REPL**: A Read-Eval-Print Loop that continuously prompts for user commands.
- **Explore Locations:** Use the map command to view 20 available Pokémon locations at a time.
- **Navigation:** Easily navigate to previous location pages using the mapb command.
- **Discover Encounters:** With explore {area_name}, see all Pokémon that can be encountered in a specific area.
- **Catch Pokémon:** Attempt to capture a Pokémon with the catch {pokemon_name} command. The catch probability is dynamically based on the Pokémon's base experience.
- **Inspect Your Collection:** Use inspect {pokemon_name} to view details of an individual Pokémon or pokedex to see all captured Pokémon.

## Demo
<img alt="PDEX_CLI Demo" src="https://github.com/JWSch4fer/PDEX_CLI/blob/main/demo/demo.gif" width="600" />

## Installation
Make sure you have Go installed on your system (version 1.13+ is recommended).

Clone the repository:
```sh
git clone https://github.com/JWSch4fer/PDEX_CLI.git
cd PDEX_CLI
go build
```
This will generate an executable that you can run directly from your terminal.

## Usage
Run the executable to launch the Pokedex CLI. On startup, the application displays a welcome message and a list of available commands:

plaintext
Copy
```
Pokedex command line interface options:

map                     | Show 20 available locations
mapb                    | Show previous 20 locations
explore {area_name}     | Show all Pokémon encounters in the specified area
catch {pokemon_name}    | Try to catch a Pokémon
inspect {pokemon_name}  | Inspect a caught Pokémon
pokedex                 | Inspect all caught Pokémon
help                    | Print available commands
exit                    | Shut down the Pokedex
```

Simply type any of these commands to start explore

## License
This project is licensed under the MIT License. See the LICENSE file for details.
