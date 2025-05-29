// filepath: /workspaces/g4n/cmd/library/main.go
package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	mathApp "github.com/ifabos/g4o/internal/application/math"
	mathInfra "github.com/ifabos/g4o/internal/infrastructure/math"
	"github.com/ifabos/g4o/internal/version"
)

// Global service instance
var mathService *mathApp.MathService

func init() {
	calculator := mathInfra.NewBasicCalculator()
	repository := mathInfra.NewInMemoryMathRepository()
	mathService = mathApp.NewMathService(calculator, repository)
}

// CResult represents a C-compatible result structure
type CResult struct {
	value    C.double
	hasError C.int
	errorMsg *C.char
}

//export Add
func Add(a, b C.double) C.double {
	result, _ := mathService.PerformOperation(float64(a), float64(b), "+")
	return C.double(result.Value)
}

//export Subtract
func Subtract(a, b C.double) C.double {
	result, _ := mathService.PerformOperation(float64(a), float64(b), "-")
	return C.double(result.Value)
}

//export Multiply
func Multiply(a, b C.double) C.double {
	result, _ := mathService.PerformOperation(float64(a), float64(b), "*")
	return C.double(result.Value)
}

//export Divide
func Divide(a, b C.double, hasError *C.int) C.double {
	result, err := mathService.PerformOperation(float64(a), float64(b), "/")
	if err != nil {
		*hasError = 1
		return 0.0
	}
	*hasError = 0
	return C.double(result.Value)
}

//export Power
func Power(base, exponent C.double) C.double {
	result, _ := mathService.PerformOperation(float64(base), float64(exponent), "^")
	return C.double(result.Value)
}

//export SquareRoot
func SquareRoot(x C.double, hasError *C.int) C.double {
	result, err := mathService.PerformOperation(float64(x), 0, "sqrt")
	if err != nil {
		*hasError = 1
		return 0.0
	}
	*hasError = 0
	return C.double(result.Value)
}

//export Logarithm
func Logarithm(x C.double, hasError *C.int) C.double {
	result, err := mathService.PerformOperation(float64(x), 0, "log")
	if err != nil {
		*hasError = 1
		return 0.0
	}
	*hasError = 0
	return C.double(result.Value)
}

//export Sine
func Sine(x C.double) C.double {
	result, _ := mathService.PerformOperation(float64(x), 0, "sin")
	return C.double(result.Value)
}

//export Cosine
func Cosine(x C.double) C.double {
	result, _ := mathService.PerformOperation(float64(x), 0, "cos")
	return C.double(result.Value)
}

//export Tangent
func Tangent(x C.double) C.double {
	result, _ := mathService.PerformOperation(float64(x), 0, "tan")
	return C.double(result.Value)
}

//export ClearHistory
func ClearHistory() C.int {
	err := mathService.ClearCalculationHistory()
	if err != nil {
		return 1
	}
	return 0
}

//export GetLastErrorMessage
func GetLastErrorMessage() *C.char {
	// This is a simplified implementation
	// In a real-world scenario, you'd want to store the last error message
	return C.CString("No error")
}

//export FreeErrorMessage
func FreeErrorMessage(msg *C.char) {
	C.free(unsafe.Pointer(msg))
}

// Version Information Functions

//export GetVersion
func GetVersion() *C.char {
	versionStr := version.GetVersion()
	return C.CString(versionStr)
}

//export GetVersionShort
func GetVersionShort() *C.char {
	versionShort := version.GetVersionShort()
	return C.CString(versionShort)
}

//export GetBuildInfo
func GetBuildInfo() *C.char {
	buildInfo := version.GetBuildInfo()
	return C.CString(buildInfo)
}

//export GetGitCommit
func GetGitCommit() *C.char {
	info := version.GetVersionInfo()
	return C.CString(info.GitCommit)
}

//export GetGitBranch
func GetGitBranch() *C.char {
	info := version.GetVersionInfo()
	return C.CString(info.GitBranch)
}

//export GetBuildTime
func GetBuildTime() *C.char {
	info := version.GetVersionInfo()
	return C.CString(info.BuildTime)
}

//export FreeVersionString
func FreeVersionString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func main() {
	// This is required for building as a C shared library
}
