package math

import (
	domainMath "github.com/ifabos/g4o/internal/domain/math"
)

// MathService handles mathematical operations using DDD principles
type MathService struct {
	calculator domainMath.Calculator
	repository domainMath.MathRepository
}

// NewMathService creates a new MathService instance
func NewMathService(calculator domainMath.Calculator, repository domainMath.MathRepository) *MathService {
	return &MathService{
		calculator: calculator,
		repository: repository,
	}
}

// PerformOperation executes a mathematical operation and stores it in history
func (s *MathService) PerformOperation(operand1, operand2 float64, operator string) (*domainMath.Result, error) {
	operation := domainMath.NewMathOperation(operand1, operand2, operator)

	if err := operation.Validate(); err != nil {
		return &domainMath.Result{Value: 0, Error: err}, err
	}

	var result *domainMath.Result

	switch operator {
	case "+":
		value := s.calculator.Add(operand1, operand2)
		result = &domainMath.Result{Value: value, Error: nil}
	case "-":
		value := s.calculator.Subtract(operand1, operand2)
		result = &domainMath.Result{Value: value, Error: nil}
	case "*":
		value := s.calculator.Multiply(operand1, operand2)
		result = &domainMath.Result{Value: value, Error: nil}
	case "/":
		value, err := s.calculator.Divide(operand1, operand2)
		result = &domainMath.Result{Value: value, Error: err}
	case "^":
		value := s.calculator.Power(operand1, operand2)
		result = &domainMath.Result{Value: value, Error: nil}
	case "sqrt":
		value, err := s.calculator.SquareRoot(operand1)
		result = &domainMath.Result{Value: value, Error: err}
	case "log":
		value, err := s.calculator.Logarithm(operand1)
		result = &domainMath.Result{Value: value, Error: err}
	case "sin":
		value := s.calculator.Sine(operand1)
		result = &domainMath.Result{Value: value, Error: nil}
	case "cos":
		value := s.calculator.Cosine(operand1)
		result = &domainMath.Result{Value: value, Error: nil}
	case "tan":
		value := s.calculator.Tangent(operand1)
		result = &domainMath.Result{Value: value, Error: nil}
	}

	// Store operation in history if successful
	if result.Error == nil {
		s.repository.SaveOperation(operation, result)
	}

	return result, result.Error
}

// GetCalculationHistory returns the history of calculations
func (s *MathService) GetCalculationHistory() ([]*domainMath.MathOperation, error) {
	return s.repository.GetHistory()
}

// ClearCalculationHistory clears the calculation history
func (s *MathService) ClearCalculationHistory() error {
	return s.repository.ClearHistory()
}
