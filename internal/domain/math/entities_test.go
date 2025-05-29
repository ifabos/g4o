package math

import (
	"testing"
)

func TestMathOperation_Validate(t *testing.T) {
	tests := []struct {
		name      string
		operation *MathOperation
		wantError bool
	}{
		{
			name:      "valid addition",
			operation: NewMathOperation(5.0, 3.0, "+"),
			wantError: false,
		},
		{
			name:      "valid subtraction",
			operation: NewMathOperation(5.0, 3.0, "-"),
			wantError: false,
		},
		{
			name:      "valid multiplication",
			operation: NewMathOperation(5.0, 3.0, "*"),
			wantError: false,
		},
		{
			name:      "valid division",
			operation: NewMathOperation(15.0, 3.0, "/"),
			wantError: false,
		},
		{
			name:      "division by zero",
			operation: NewMathOperation(15.0, 0.0, "/"),
			wantError: true,
		},
		{
			name:      "valid square root",
			operation: NewMathOperation(16.0, 0.0, "sqrt"),
			wantError: false,
		},
		{
			name:      "square root of negative",
			operation: NewMathOperation(-4.0, 0.0, "sqrt"),
			wantError: true,
		},
		{
			name:      "valid logarithm",
			operation: NewMathOperation(10.0, 0.0, "log"),
			wantError: false,
		},
		{
			name:      "logarithm of zero",
			operation: NewMathOperation(0.0, 0.0, "log"),
			wantError: true,
		},
		{
			name:      "logarithm of negative",
			operation: NewMathOperation(-5.0, 0.0, "log"),
			wantError: true,
		},
		{
			name:      "unsupported operation",
			operation: NewMathOperation(5.0, 3.0, "unsupported"),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.operation.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("MathOperation.Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
