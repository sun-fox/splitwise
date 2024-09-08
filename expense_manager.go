package main

import "errors"

// ExpenseManager to manage expenses
type ExpenseManager struct {
	Users    map[string]User
	Balances *BalanceManager
	Expenses []Expense
}

// Initialize ExpenseManager
func NewExpenseManager() *ExpenseManager {
	return &ExpenseManager{
		Users:    make(map[string]User),
		Balances: NewBalanceManager(),
		Expenses: []Expense{},
	}
}

// Add a user
func (em *ExpenseManager) AddUser(id, name, email, phone string) {
	em.Users[id] = User{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}
}

// Add an expense and update balances
func (em *ExpenseManager) AddExpense(paidBy string, amount float64, users []string, splitType SplitType, amounts []float64) error {
	// Validate if all users exist
	for _, userID := range users {
		if _, exists := em.Users[userID]; !exists {
			return errors.New("one or more users do not exist")
		}
	}

	expense := Expense{
		PaidBy:  paidBy,
		Amount:  amount,
		Users:   users,
		Split:   splitType,
		Amounts: amounts,
	}
	em.Expenses = append(em.Expenses, expense)

	// Split logic
	switch splitType {
	case EQUAL:
		equalSplit := amount / float64(len(users))
		for _, user := range users {
			if user != paidBy {
				em.Balances.UpdateBalance(paidBy, user, equalSplit)
			}
		}
	case EXACT:
		for i, user := range users {
			if user != paidBy {
				em.Balances.UpdateBalance(paidBy, user, amounts[i])
			}
		}
	case PERCENT:
		for i, user := range users {
			if user != paidBy {
				splitAmount := amount * (amounts[i] / 100)
				em.Balances.UpdateBalance(paidBy, user, splitAmount)
			}
		}
	default:
		return errors.New("invalid split type")
	}

	return nil
}
