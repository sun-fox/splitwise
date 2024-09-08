package main

import "fmt"

// Helper functions to convert string to int/float
func toInt(s string) int {
	var val int
	fmt.Sscanf(s, "%d", &val)
	return val
}

func toFloat(s string) float64 {
	var val float64
	fmt.Sscanf(s, "%f", &val)
	return val
}
