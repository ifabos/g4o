package math

import (
	"errors"
	"math"
)

// BasicCalculator implements the Calculator interface
type BasicCalculator struct{}

// NewBasicCalculator creates a new BasicCalculator instance
func NewBasicCalculator() *BasicCalculator {
	return &BasicCalculator{}
}

// Add performs addition
func (c *BasicCalculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction
func (c *BasicCalculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication
func (c *BasicCalculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide performs division with error handling
func (c *BasicCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Power calculates base raised to the power of exponent
func (c *BasicCalculator) Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// SquareRoot calculates the square root
func (c *BasicCalculator) SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(x), nil
}

// Logarithm calculates natural logarithm
func (c *BasicCalculator) Logarithm(x float64) (float64, error) {
	if x <= 0 {
		return 0, errors.New("logarithm of non-positive number")
	}
	return math.Log(x), nil
}

// Sine calculates sine of x (in radians)
func (c *BasicCalculator) Sine(x float64) float64 {
	return math.Sin(x)
}

// Cosine calculates cosine of x (in radians)
func (c *BasicCalculator) Cosine(x float64) float64 {
	return math.Cos(x)
}

// Tangent calculates tangent of x (in radians)
func (c *BasicCalculator) Tangent(x float64) float64 {
	return math.Tan(x)
}
