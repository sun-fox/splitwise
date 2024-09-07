package main

import (
	"fmt"
	"sync"

	"github.com/sun-fox/splitwise/models"
)

// Declare the global variable
var (
	Splits models.Split
	mu     sync.Mutex
)

// Add a Lending or a Borrowing Transaction
func AddTransaction(lender string, borrower string, amount float64) {
	mu.Lock()
	defer mu.Unlock()

	if Splits[lender] == nil {
		Splits[lender] = make(models.UserLendings)
	}
	Splits[lender][borrower] = amount
}

// Helper Utility: To view the PassBook for debugging or accounting purposes
func PrintSplits() {
	mu.Lock()
	defer mu.Unlock()

	for user, lendings := range Splits {
		fmt.Printf("%s's Balances:\n", user)
		for otherUser, amount := range lendings {
			// Check if the amount is negative
			if amount < 0 {
				// Exchange lender and borrower and print the positive amount
				fmt.Printf("  %s owes %s: %.2f\n", otherUser, user, -amount)
			} else {
				fmt.Printf("  %s owes %s: %.2f\n", otherUser, user, amount)
			}
		}
	}
}

// Main function to run the application
func main() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Expense Sharing Application (type EXIT to quit)")

	// Initialize the global variable
	Splits = make(models.Split)

	AddTransaction("user1", "user2", 50.0)
	AddTransaction("user1", "user3", 30.0)
	AddTransaction("user2", "user1", 20.0)
	PrintSplits()
	// for {
	// 	// Display prompt and read input
	// 	fmt.Print("> ")
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Error reading input:", err)
	// 		continue
	// 	}

	// 	// Trim the newline character and handle the command
	// 	input = strings.TrimSpace(input)
	// 	handleCommand(input)
	// }
}
