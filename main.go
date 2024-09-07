package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// Main function to run the application
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Expense Sharing Application (type EXIT to quit)")

	for {
		// Display prompt and read input
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the newline character and handle the command
		input = strings.TrimSpace(input)
		handleCommand(input)
	}
}
