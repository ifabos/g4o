package math

import (
	"testing"
)

func TestBasicCalculator_Add(t *testing.T) {
	calc := NewBasicCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5.0, 3.0, 8.0},
		{"negative numbers", -2.0, -3.0, -5.0},
		{"mixed signs", 10.0, -4.0, 6.0},
		{"zero", 0.0, 5.0, 5.0},
		{"decimals", 2.5, 3.7, 6.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculator_Subtract(t *testing.T) {
	calc := NewBasicCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10.0, 3.0, 7.0},
		{"negative result", 3.0, 10.0, -7.0},
		{"negative numbers", -5.0, -2.0, -3.0},
		{"zero", 5.0, 0.0, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculator_Multiply(t *testing.T) {
	calc := NewBasicCalculator()

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 4.0, 5.0, 20.0},
		{"negative numbers", -3.0, -2.0, 6.0},
		{"mixed signs", -4.0, 5.0, -20.0},
		{"zero", 0.0, 5.0, 0.0},
		{"decimals", 2.5, 4.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculator_Divide(t *testing.T) {
	calc := NewBasicCalculator()

	t.Run("valid division", func(t *testing.T) {
		result, err := calc.Divide(15.0, 3.0)
		if err != nil {
			t.Errorf("Divide(15.0, 3.0) returned unexpected error: %v", err)
		}
		if result != 5.0 {
			t.Errorf("Divide(15.0, 3.0) = %f; want 5.0", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := calc.Divide(10.0, 0.0)
		if err == nil {
			t.Error("Divide(10.0, 0.0) should return an error")
		}
	})
}

func TestBasicCalculator_SquareRoot(t *testing.T) {
	calc := NewBasicCalculator()

	t.Run("valid square root", func(t *testing.T) {
		result, err := calc.SquareRoot(16.0)
		if err != nil {
			t.Errorf("SquareRoot(16.0) returned unexpected error: %v", err)
		}
		if result != 4.0 {
			t.Errorf("SquareRoot(16.0) = %f; want 4.0", result)
		}
	})

	t.Run("square root of negative number", func(t *testing.T) {
		_, err := calc.SquareRoot(-4.0)
		if err == nil {
			t.Error("SquareRoot(-4.0) should return an error")
		}
	})
}

func TestBasicCalculator_Logarithm(t *testing.T) {
	calc := NewBasicCalculator()

	t.Run("logarithm of zero", func(t *testing.T) {
		_, err := calc.Logarithm(0.0)
		if err == nil {
			t.Error("Logarithm(0.0) should return an error")
		}
	})

	t.Run("logarithm of negative number", func(t *testing.T) {
		_, err := calc.Logarithm(-1.0)
		if err == nil {
			t.Error("Logarithm(-1.0) should return an error")
		}
	})
}
