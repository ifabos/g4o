# g4n - Cross-Platform Math Library

A professional Go module implemented using Domain-Driven Design (DDD) that provides basic mathematical operations as a dynamic library for Visual Basic and other applications on both Windows and Linux platforms.

## Features

- **Basic Arithmetic Operations**: Addition, Subtraction, Multiplication, Division
- **Advanced Mathematical Functions**: Power, Square Root, Natural Logarithm
- **Trigonometric Functions**: Sine, Cosine, Tangent (input in radians)
- **Error Handling**: Proper error handling for invalid operations
- **Cross-Platform**: Builds for both Windows (.dll) and Linux (.so)
- **Visual Basic Compatible**: C-compatible exports for easy integration
- **DDD Architecture**: Clean, maintainable code structure
- **History Management**: Optional operation history tracking

## Architecture

This project follows Domain-Driven Design principles:

```plaintext
├── internal/
│   ├── domain/math/          # Domain layer - business logic and entities
│   ├── application/math/     # Application layer - use cases and services
│   └── infrastructure/math/  # Infrastructure layer - implementations
├── cmd/library/              # Main entry point for library build
├── include/                  # C header files for external integration
├── examples/                 # Usage examples for different platforms
└── build/                    # Generated build artifacts
```

## Building the Library

### Prerequisites

- Go 1.24+ installed
- For Windows cross-compilation on Linux: MinGW-w64 toolchain

#### Installing MinGW-w64 (for Windows builds on Linux)

**Ubuntu/Debian:**

```bash
sudo apt-get install gcc-mingw-w64
```

**CentOS/RHEL:**

```bash
sudo yum install mingw64-gcc
```

**Arch Linux:**

```bash
sudo pacman -S mingw-w64-gcc
```

### Build Commands

#### Build for all platforms

```bash
make all
```

#### Build for specific platforms

```bash
make linux    # Build for Linux (.so)
make windows  # Build for Windows (.dll)
make local    # Build for current platform only
```

#### Other useful commands

```bash
make test     # Run unit tests
make clean    # Clean build artifacts
make package  # Create package with headers
make help     # Show all available commands
```

### Build Output

- **Linux**: `build/libcommon4n.so`
- **Windows**: `build/g4n.dll`
- **Headers**: `include/g4n.h`

## Usage Examples

### Visual Basic .NET

```vb
Imports System.Runtime.InteropServices

Public Class MathLibrary
    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Add(a As Double, b As Double) As Double
    End Function
    
    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Divide(a As Double, b As Double, ByRef hasError As Integer) As Double
    End Function
End Class

' Usage
Dim result As Double = MathLibrary.Add(10.5, 5.3)
Console.WriteLine($"Result: {result}")
```

### Visual Basic 6

```vb
Private Declare Function Add Lib "g4n.dll" (ByVal a As Double, ByVal b As Double) As Double
Private Declare Function Divide Lib "g4n.dll" (ByVal a As Double, ByVal b As Double, ByRef hasError As Long) As Double

' Usage
Dim result As Double
result = Add(10.5, 5.3)
MsgBox "Result: " & CStr(result)
```

### C/C++

```c
#include "g4n.h"

int main() {
    double result = Add(10.5, 5.3);
    printf("Result: %f\n", result);
    
    int hasError = 0;
    double divResult = Divide(15.0, 3.0, &hasError);
    if (hasError == 0) {
        printf("Division result: %f\n", divResult);
    }
    
    return 0;
}
```

## API Reference

### Basic Arithmetic Operations

- `double Add(double a, double b)` - Addition
- `double Subtract(double a, double b)` - Subtraction  
- `double Multiply(double a, double b)` - Multiplication
- `double Divide(double a, double b, int* hasError)` - Division with error handling

### Advanced Mathematical Functions

- `double Power(double base, double exponent)` - Power calculation
- `double SquareRoot(double x, int* hasError)` - Square root with error handling
- `double Logarithm(double x, int* hasError)` - Natural logarithm with error handling

### Trigonometric Functions

- `double Sine(double x)` - Sine (input in radians)
- `double Cosine(double x)` - Cosine (input in radians)
- `double Tangent(double x)` - Tangent (input in radians)

### Utility Functions

- `int ClearHistory()` - Clear calculation history
- `char* GetLastErrorMessage()` - Get last error message
- `void FreeErrorMessage(char* msg)` - Free error message memory

## Error Handling

Functions that can fail (like division by zero) use an `int* hasError` parameter:

- `0` = Success
- `1` = Error occurred

For trigonometric functions, all angles should be provided in radians.

## Development

### Running Tests

```bash
make test
```

### Project Structure

```plaintext
g4n/
├── cmd/library/main.go               # C exports and main entry point
├── internal/
│   ├── domain/math/
│   │   ├── entities.go               # Domain entities
│   │   └── interfaces.go             # Domain interfaces
│   ├── application/math/
│   │   └── service.go                # Application services
│   └── infrastructure/math/
│       ├── calculator.go             # Calculator implementation
│       └── repository.go             # Repository implementation
├── include/g4n.h                # C header file
├── examples/                         # Usage examples
├── Makefile                          # Build configuration
└── README.md                         # This file
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

This project is available under the MIT License. See the LICENSE file for more details.

## Platform Support

- **Windows**: x64 (64-bit)
- **Linux**: x64 (64-bit)
- **Visual Basic**: VB6, VB.NET
- **Other**: Any language that can call C functions from dynamic libraries

## Troubleshooting

### Common Issues

1. **Missing MinGW-w64**: Install the MinGW-w64 toolchain for Windows cross-compilation
2. **CGO Disabled**: Ensure `CGO_ENABLED=1` environment variable is set
3. **Library Not Found**: Make sure the library is in your system's library path or the same directory as your executable

### Getting Help

- Check the examples in the `examples/` directory
- Review the test files for usage patterns
- Open an issue on GitHub for bug reports or feature requests
