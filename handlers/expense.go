package handlers

import (
	"fmt"
	"strconv"
	"strings"
)

// Handle the EXPENSE command
func HandleExpense(args []string) {
	if len(args) < 5 {
		fmt.Println("Invalid EXPENSE command")
		return
	}

	payer := args[0]
	amountStr := args[1]
	numUsersStr := args[2]
	users := args[3 : 3+strings.Count(numUsersStr, " ")]
	expenseType := args[3+len(users)]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount:", amountStr)
		return
	}

	_, err = strconv.Atoi(numUsersStr)
	if err != nil {
		fmt.Println("Invalid number of users:", numUsersStr)
		return
	}

	switch expenseType {
	case "EQUAL":
		handleEqualExpense(payer, amount, users)
	case "EXACT":
		if len(args) < 4+len(users) {
			fmt.Println("Invalid EXPENSE command for EXACT")
			return
		}
		exactAmounts := args[4+len(users):]
		handleExactExpense(payer, amount, users, exactAmounts)
	case "PERCENT":
		if len(args) < 4+len(users) {
			fmt.Println("Invalid EXPENSE command for PERCENT")
			return
		}
		percentages := args[4+len(users):]
		handlePercentExpense(payer, amount, users, percentages)
	default:
		fmt.Println("Unknown expense type:", expenseType)
	}
}

// Handle EQUAL expense type
func handleEqualExpense(payer string, amount float64, users []string) {
	fmt.Printf("Handling EQUAL expense: %s paid %.2f and split among %v\n", payer, amount, users)
	// Logic to distribute amount equally among users
}

// Handle EXACT expense type
func handleExactExpense(payer string, totalAmount float64, users []string, amounts []string) {
	fmt.Printf("Handling EXACT expense: %s paid %.2f and amounts are %v\n", payer, totalAmount, amounts)
	// Logic to handle exact amounts specified for each user
}

// Handle PERCENT expense type
func handlePercentExpense(payer string, totalAmount float64, users []string, percentages []string) {
	fmt.Printf("Handling PERCENT expense: %s paid %.2f and percentages are %v\n", payer, totalAmount, percentages)
	// Logic to handle percentage splits among users
}