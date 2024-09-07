package models

// Transactions holds the amounts lent by a user to other users.
type Transactions map[string]float64

// Entry represents a transaction in the Passbook .
type Passbook map[string]Transactions
