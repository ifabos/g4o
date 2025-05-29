package math

// Calculator defines the interface for mathematical operations
type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
	Power(base, exponent float64) float64
	SquareRoot(x float64) (float64, error)
	Logarithm(x float64) (float64, error)
	Sine(x float64) float64
	Cosine(x float64) float64
	Tangent(x float64) float64
}

// MathRepository defines the interface for persisting calculation history
type MathRepository interface {
	SaveOperation(operation *MathOperation, result *Result) error
	GetHistory() ([]*MathOperation, error)
	ClearHistory() error
}
