package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/sun-fox/splitwise/handlers"
)

// Function to handle terminal commands
func handleCommand(input string) {
	// Split the input to get the command and arguments
	args := strings.Fields(input)
	if len(args) == 0 {
		fmt.Println("No input provided")
		return
	}

	// Example: Handle different commands
	switch args[0] {
	case "EXPENSE":
		handlers.HandleExpense(args[1:])
	case "SHOW":
		if len(args) > 1 {
			handlers.HandleShowUser(args[1])
		} else {
			handlers.HandleShowAll()
		}
	case "EXIT":
		fmt.Println("Exiting application...")
		os.Exit(0)
	default:
		fmt.Println("Unknown command:", args[0])
	}
}
