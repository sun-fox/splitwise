package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Handle commands: EXPENSE, SHOW, SHOW <user-id>
func handleCommand(em *ExpenseManager, command string) {
	args := strings.Split(command, " ")

	switch args[0] {
	case "EXPENSE":
		paidBy := args[1]
		amount := toFloat(args[2])
		numUsers := toInt(args[3])
		users := args[4 : 4+numUsers]
		splitType := SplitType(args[4+numUsers])

		var amounts []float64
		if splitType != EQUAL {
			for _, v := range args[5+numUsers:] {
				amounts = append(amounts, toFloat(v))
			}
		}
		err := em.AddExpense(paidBy, amount, users, splitType, amounts)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case "SHOW":
		if len(args) == 1 {
			em.Balances.ShowBalances()
		} else {
			userID := args[1]
			em.Balances.ShowUserBalance(userID)
		}
	case "ADDUSER":
		if len(args) != 5{
			fmt.Println("Invalid Command")
		}else{
			userId := args[1]
			userName := args[2]
			userEmailID := args[3]
			userPhoneNo := args[4]
			em.AddUser(userId, userName, userEmailID, userPhoneNo)
		}
	default:
		fmt.Println("Invalid command")
	}
}


// Main method to simulate usage
func main() {
	em := NewExpenseManager()

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
		handleCommand(em, input)
	}
}
