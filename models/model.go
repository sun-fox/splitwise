package models

// UserLendings holds the amounts lent by a user to other users.
type UserLendings map[string]float64

// Split represents the splits among users.
type Split map[string]UserLendings
