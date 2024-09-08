package main

// Expense struct to store the details of an expense
type Expense struct {
	PaidBy  string    // User who paid
	Amount  float64   // Total amount
	Users   []string  // Users involved
	Split   SplitType // Split type (EQUAL, EXACT, PERCENT)
	Amounts []float64 // For EXACT or PERCENT, we store amounts or percentages
}

//Constructor
func NewExpense(paidby string, amount float64, users []string, splitType SplitType, amounts []float64) Expense{
	return Expense{
		PaidBy: paidby,
		Amount: amount,
		Users: users,
		Split: splitType,
		Amounts: amounts,
	}
}