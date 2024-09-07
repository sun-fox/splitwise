package handlers

import (
	"fmt"
)

// Function to handle 'SHOW' command for all users
func HandleShowAll() {
	// Logic to show balances for all users
	fmt.Println("Showing all balances...")
}

// Function to handle 'SHOW <user-id>' command
func HandleShowUser(userID string) {
	// Logic to show balance for a specific user
	fmt.Printf("Showing balance for user %s...\n", userID)
}