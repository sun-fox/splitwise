package main

// User struct to store user details
type User struct {
	ID    string
	Name  string
	Email string
	Phone string
}

// Split types: EQUAL, EXACT, PERCENT
type SplitType string

const (
	EQUAL   SplitType = "EQUAL"
	EXACT   SplitType = "EXACT"
	PERCENT SplitType = "PERCENT"
)

// Expense struct to store the details of an expense
type Expense struct {
	PaidBy  string    // User who paid
	Amount  float64   // Total amount
	Users   []string  // Users involved
	Split   SplitType // Split type (EQUAL, EXACT, PERCENT)
	Amounts []float64 // For EXACT or PERCENT, we store amounts or percentages
}
