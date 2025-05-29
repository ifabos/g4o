package math

import "errors"

// MathOperation represents a mathematical operation
type MathOperation struct {
	Operand1 float64
	Operand2 float64
	Operator string
}

// Result represents the result of a mathematical operation
type Result struct {
	Value float64
	Error error
}

// NewMathOperation creates a new MathOperation
func NewMathOperation(operand1, operand2 float64, operator string) *MathOperation {
	return &MathOperation{
		Operand1: operand1,
		Operand2: operand2,
		Operator: operator,
	}
}

// Validate checks if the operation is valid
func (m *MathOperation) Validate() error {
	switch m.Operator {
	case "+", "-", "*", "/", "^", "sqrt", "log", "sin", "cos", "tan":
		if m.Operator == "/" && m.Operand2 == 0 {
			return errors.New("division by zero")
		}
		if m.Operator == "log" && m.Operand1 <= 0 {
			return errors.New("logarithm of non-positive number")
		}
		if m.Operator == "sqrt" && m.Operand1 < 0 {
			return errors.New("square root of negative number")
		}
		return nil
	default:
		return errors.New("unsupported operation")
	}
}
