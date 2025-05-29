/*
 * C example for using g4n library
 * This example demonstrates how to use the math library in C/C++
 */

#include <stdio.h>
#include <math.h>
#include "g4n.h"

int main() {
    printf("g4n Math Library - C Example\n");
    printf("==================================\n\n");
    
    // Basic arithmetic operations
    printf("Basic Arithmetic Operations:\n");
    printf("----------------------------\n");
    
    double result1 = Add(10.5, 5.3);
    printf("10.5 + 5.3 = %.2f\n", result1);
    
    double result2 = Subtract(20.0, 8.0);
    printf("20.0 - 8.0 = %.2f\n", result2);
    
    double result3 = Multiply(4.0, 7.0);
    printf("4.0 * 7.0 = %.2f\n", result3);
    
    // Division with error handling
    int hasError = 0;
    double result4 = Divide(15.0, 3.0, &hasError);
    if (hasError == 0) {
        printf("15.0 / 3.0 = %.2f\n", result4);
    } else {
        printf("Division error occurred\n");
    }
    
    // Test division by zero
    hasError = 0;
    double result5 = Divide(10.0, 0.0, &hasError);
    if (hasError == 0) {
        printf("10.0 / 0.0 = %.2f\n", result5);
    } else {
        printf("Division by zero error (expected)\n");
    }
    
    printf("\nAdvanced Mathematical Operations:\n");
    printf("---------------------------------\n");
    
    // Power calculation
    double result6 = Power(2.0, 8.0);
    printf("2.0 ^ 8.0 = %.2f\n", result6);
    
    // Square root with error handling
    hasError = 0;
    double result7 = SquareRoot(16.0, &hasError);
    if (hasError == 0) {
        printf("sqrt(16.0) = %.2f\n", result7);
    } else {
        printf("Square root error occurred\n");
    }
    
    // Test square root of negative number
    hasError = 0;
    double result8 = SquareRoot(-4.0, &hasError);
    if (hasError == 0) {
        printf("sqrt(-4.0) = %.2f\n", result8);
    } else {
        printf("Square root of negative number error (expected)\n");
    }
    
    // Natural logarithm with error handling
    hasError = 0;
    double result9 = Logarithm(2.718281828, &hasError);
    if (hasError == 0) {
        printf("ln(e) = %.6f\n", result9);
    } else {
        printf("Logarithm error occurred\n");
    }
    
    printf("\nTrigonometric Functions (radians):\n");
    printf("----------------------------------\n");
    
    // Trigonometric functions
    double pi = 3.14159265358979323846;
    double angleRadians = pi / 4.0; // 45 degrees
    
    double sinResult = Sine(angleRadians);
    double cosResult = Cosine(angleRadians);
    double tanResult = Tangent(angleRadians);
    
    printf("sin(π/4) = %.6f\n", sinResult);
    printf("cos(π/4) = %.6f\n", cosResult);
    printf("tan(π/4) = %.6f\n", tanResult);
    
    // Additional trigonometric values
    printf("sin(0) = %.6f\n", Sine(0.0));
    printf("cos(0) = %.6f\n", Cosine(0.0));
    printf("sin(π/2) = %.6f\n", Sine(pi / 2.0));
    printf("cos(π/2) = %.6f\n", Cosine(pi / 2.0));
    
    printf("\nUtility Functions:\n");
    printf("------------------\n");
    
    // Clear calculation history
    int clearResult = ClearHistory();
    if (clearResult == 0) {
        printf("History cleared successfully\n");
    } else {
        printf("Failed to clear history\n");
    }
    
    // Demonstration of degree to radian conversion
    printf("\nDegree to Radian Conversion Examples:\n");
    printf("------------------------------------\n");
    
    double degrees[] = {0, 30, 45, 60, 90, 180, 270, 360};
    int numAngles = sizeof(degrees) / sizeof(degrees[0]);
    
    for (int i = 0; i < numAngles; i++) {
        double radians = degrees[i] * pi / 180.0;
        double sinVal = Sine(radians);
        double cosVal = Cosine(radians);
        
        printf("%.0f° (%.4f rad): sin=%.6f, cos=%.6f\n", 
               degrees[i], radians, sinVal, cosVal);
    }
    
    printf("\nExample completed successfully!\n");
    return 0;
}

/*
 * Compilation instructions:
 * 
 * Linux:
 *   gcc -o c_example c_example.c -L./build -lcommon4n -lm
 *   ./c_example
 * 
 * Windows (MinGW):
 *   gcc -o c_example.exe c_example.c -L./build -lcommon4n
 *   ./c_example.exe
 * 
 * Note: Make sure the library file (libcommon4n.so or g4n.dll) 
 * is in your library path or the same directory as the executable.
 */
