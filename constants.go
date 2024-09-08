package main

// Split types: EQUAL, EXACT, PERCENT
type SplitType string

const (
	EQUAL   SplitType = "EQUAL"
	EXACT   SplitType = "EXACT"
	PERCENT SplitType = "PERCENT"
)