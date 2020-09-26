// Package calculator provides a library for simple calculations in Go.
package calculator

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them together
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two number and returns the result of dividing the first by
// the second
func Divide(a, b float64) (float64, error) {
	return 0, nil
}
