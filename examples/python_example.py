#!/usr/bin/env python3
"""
Python example for using g4n library
This example demonstrates how to use the math library in Python using ctypes
"""

import ctypes
import os
import platform

def load_library():
    """Load the appropriate library based on the operating system"""
    system = platform.system()
    
    if system == "Windows":
        # Try to load Windows DLL
        lib_paths = ["build/g4n.dll", "g4n.dll"]
        for lib_path in lib_paths:
            if os.path.exists(lib_path):
                return ctypes.CDLL(lib_path)
        raise OSError("Could not find g4n.dll")
    else:
        # Try to load Linux shared library
        lib_paths = ["build/libcommon4n.so", "./libcommon4n.so", "libcommon4n.so"]
        for lib_path in lib_paths:
            if os.path.exists(lib_path):
                return ctypes.CDLL(lib_path)
        raise OSError("Could not find libcommon4n.so")

def setup_functions(lib):
    """Setup function signatures for the library"""
    
    # Basic arithmetic operations
    lib.Add.argtypes = [ctypes.c_double, ctypes.c_double]
    lib.Add.restype = ctypes.c_double
    
    lib.Subtract.argtypes = [ctypes.c_double, ctypes.c_double]
    lib.Subtract.restype = ctypes.c_double
    
    lib.Multiply.argtypes = [ctypes.c_double, ctypes.c_double]
    lib.Multiply.restype = ctypes.c_double
    
    lib.Divide.argtypes = [ctypes.c_double, ctypes.c_double, ctypes.POINTER(ctypes.c_int)]
    lib.Divide.restype = ctypes.c_double
    
    # Advanced mathematical functions
    lib.Power.argtypes = [ctypes.c_double, ctypes.c_double]
    lib.Power.restype = ctypes.c_double
    
    lib.SquareRoot.argtypes = [ctypes.c_double, ctypes.POINTER(ctypes.c_int)]
    lib.SquareRoot.restype = ctypes.c_double
    
    lib.Logarithm.argtypes = [ctypes.c_double, ctypes.POINTER(ctypes.c_int)]
    lib.Logarithm.restype = ctypes.c_double
    
    # Trigonometric functions
    lib.Sine.argtypes = [ctypes.c_double]
    lib.Sine.restype = ctypes.c_double
    
    lib.Cosine.argtypes = [ctypes.c_double]
    lib.Cosine.restype = ctypes.c_double
    
    lib.Tangent.argtypes = [ctypes.c_double]
    lib.Tangent.restype = ctypes.c_double
    
    # Utility functions
    lib.ClearHistory.argtypes = []
    lib.ClearHistory.restype = ctypes.c_int

class MathLibrary:
    """Python wrapper for the g4n math library"""
    
    def __init__(self):
        self.lib = load_library()
        setup_functions(self.lib)
    
    def add(self, a, b):
        """Add two numbers"""
        return self.lib.Add(ctypes.c_double(a), ctypes.c_double(b))
    
    def subtract(self, a, b):
        """Subtract two numbers"""
        return self.lib.Subtract(ctypes.c_double(a), ctypes.c_double(b))
    
    def multiply(self, a, b):
        """Multiply two numbers"""
        return self.lib.Multiply(ctypes.c_double(a), ctypes.c_double(b))
    
    def divide(self, a, b):
        """Divide two numbers with error handling"""
        has_error = ctypes.c_int(0)
        result = self.lib.Divide(ctypes.c_double(a), ctypes.c_double(b), ctypes.byref(has_error))
        if has_error.value != 0:
            raise ZeroDivisionError("Division by zero or invalid operation")
        return result
    
    def power(self, base, exponent):
        """Calculate base raised to the power of exponent"""
        return self.lib.Power(ctypes.c_double(base), ctypes.c_double(exponent))
    
    def square_root(self, x):
        """Calculate square root with error handling"""
        has_error = ctypes.c_int(0)
        result = self.lib.SquareRoot(ctypes.c_double(x), ctypes.byref(has_error))
        if has_error.value != 0:
            raise ValueError("Square root of negative number")
        return result
    
    def logarithm(self, x):
        """Calculate natural logarithm with error handling"""
        has_error = ctypes.c_int(0)
        result = self.lib.Logarithm(ctypes.c_double(x), ctypes.byref(has_error))
        if has_error.value != 0:
            raise ValueError("Logarithm of non-positive number")
        return result
    
    def sine(self, x):
        """Calculate sine (input in radians)"""
        return self.lib.Sine(ctypes.c_double(x))
    
    def cosine(self, x):
        """Calculate cosine (input in radians)"""
        return self.lib.Cosine(ctypes.c_double(x))
    
    def tangent(self, x):
        """Calculate tangent (input in radians)"""
        return self.lib.Tangent(ctypes.c_double(x))
    
    def clear_history(self):
        """Clear calculation history"""
        result = self.lib.ClearHistory()
        return result == 0

def main():
    """Main example function"""
    print("g4n Math Library - Python Example")
    print("======================================\n")
    
    try:
        # Initialize the math library
        math_lib = MathLibrary()
        
        # Basic arithmetic operations
        print("Basic Arithmetic Operations:")
        print("----------------------------")
        
        print(f"10.5 + 5.3 = {math_lib.add(10.5, 5.3):.2f}")
        print(f"20.0 - 8.0 = {math_lib.subtract(20.0, 8.0):.2f}")
        print(f"4.0 * 7.0 = {math_lib.multiply(4.0, 7.0):.2f}")
        print(f"15.0 / 3.0 = {math_lib.divide(15.0, 3.0):.2f}")
        
        # Test division by zero
        try:
            result = math_lib.divide(10.0, 0.0)
            print(f"10.0 / 0.0 = {result:.2f}")
        except ZeroDivisionError as e:
            print(f"Division by zero error: {e}")
        
        print("\nAdvanced Mathematical Operations:")
        print("---------------------------------")
        
        print(f"2.0 ^ 8.0 = {math_lib.power(2.0, 8.0):.2f}")
        print(f"sqrt(16.0) = {math_lib.square_root(16.0):.2f}")
        print(f"ln(e) = {math_lib.logarithm(2.718281828):.6f}")
        
        # Test square root of negative number
        try:
            result = math_lib.square_root(-4.0)
            print(f"sqrt(-4.0) = {result:.2f}")
        except ValueError as e:
            print(f"Square root error: {e}")
        
        print("\nTrigonometric Functions (radians):")
        print("----------------------------------")
        
        import math
        pi = math.pi
        angle_radians = pi / 4.0  # 45 degrees
        
        print(f"sin(π/4) = {math_lib.sine(angle_radians):.6f}")
        print(f"cos(π/4) = {math_lib.cosine(angle_radians):.6f}")
        print(f"tan(π/4) = {math_lib.tangent(angle_radians):.6f}")
        
        # Additional trigonometric values
        print(f"sin(0) = {math_lib.sine(0.0):.6f}")
        print(f"cos(0) = {math_lib.cosine(0.0):.6f}")
        print(f"sin(π/2) = {math_lib.sine(pi / 2.0):.6f}")
        print(f"cos(π/2) = {math_lib.cosine(pi / 2.0):.6f}")
        
        print("\nDegree to Radian Conversion Examples:")
        print("------------------------------------")
        
        degrees_list = [0, 30, 45, 60, 90, 180, 270, 360]
        
        for degrees in degrees_list:
            radians = math.radians(degrees)
            sin_val = math_lib.sine(radians)
            cos_val = math_lib.cosine(radians)
            
            print(f"{degrees:3.0f}° ({radians:.4f} rad): sin={sin_val:.6f}, cos={cos_val:.6f}")
        
        print("\nUtility Functions:")
        print("------------------")
        
        # Clear calculation history
        if math_lib.clear_history():
            print("History cleared successfully")
        else:
            print("Failed to clear history")
        
        print("\nExample completed successfully!")
        
    except OSError as e:
        print(f"Error loading library: {e}")
        print("Make sure the library file is in the correct location:")
        print("  Windows: build/g4n.dll or g4n.dll")
        print("  Linux: build/libcommon4n.so or ./libcommon4n.so")
    except Exception as e:
        print(f"Unexpected error: {e}")

if __name__ == "__main__":
    main()
