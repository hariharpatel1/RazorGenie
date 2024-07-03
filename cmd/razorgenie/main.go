package main

import (
	"fmt"
	"os"
	"razor-genie/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: razorgenie <command> [args...]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	var err error
	switch command {
	case "/generate":
		if len(args) < 1 {
			err = fmt.Errorf("usage: /generate <prompt>")
		} else {
			err = commands.HandleGenerate(args[0])
		}
	case "/refactor":
		// Implement handler for refactoring
	case "/docgen":
		// Implement handler for documentation generation
	case "/testgen":
		// Implement handler for test generation
	case "/review":
		// Implement handler for code review
	case "/addfile":
		// Implement handler for adding files
	case "/dropfile":
		// Implement handler for dropping files
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error executing command '%s': %v\n", command, err)
		os.Exit(1)
	}
}
