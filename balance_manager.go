package main

import (
	"fmt"
	"sync"
)

// BalanceManager to manage net balances between users
type BalanceManager struct {
	mu      sync.Mutex
	balance map[string]map[string]float64
}

// Initialize BalanceManager
func NewBalanceManager() *BalanceManager {
	return &BalanceManager{
		balance: make(map[string]map[string]float64),
	}
}

// Update the balance between two users
func (bm *BalanceManager) UpdateBalance(lender, borrower string, amount float64) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	// Update borrower's debt to lender
	if _, exists := bm.balance[borrower]; !exists {
		bm.balance[borrower] = make(map[string]float64)
	}
	bm.balance[borrower][lender] += amount

	// Update lender's credit from borrower
	if _, exists := bm.balance[lender]; !exists {
		bm.balance[lender] = make(map[string]float64)
	}
	bm.balance[lender][borrower] -= amount
}

// Show balances for all users
func (bm *BalanceManager) ShowBalances() {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	for borrower, lendings := range bm.balance {
		for lender, amount := range lendings {
			if amount > 0 {
				fmt.Printf("%s owes %s: %.2f\n", borrower, lender, amount)
			}
		}
	}
}

// Show balances for a single user
func (bm *BalanceManager) ShowUserBalance(userID string) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	found := false
	for borrower, lendings := range bm.balance {
		for lender, amount := range lendings {
			if borrower == userID && amount > 0 {
				fmt.Printf("%s owes %s: %.2f\n", borrower, lender, amount)
				found = true
			} 
		}
	}
	if !found {
		fmt.Println("No balances")
	}
}
