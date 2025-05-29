' Visual Basic .NET example for using g4n library
' This example demonstrates how to use the math library in VB.NET

Imports System
Imports System.Runtime.InteropServices

Public Class MathLibrary
    ' Import declarations for Windows DLL
    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Add(a As Double, b As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Subtract(a As Double, b As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Multiply(a As Double, b As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Divide(a As Double, b As Double, ByRef hasError As Integer) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Power(base As Double, exponent As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function SquareRoot(x As Double, ByRef hasError As Integer) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Logarithm(x As Double, ByRef hasError As Integer) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Sine(x As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Cosine(x As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function Tangent(x As Double) As Double
    End Function

    <DllImport("g4n.dll", CallingConvention:=CallingConvention.Cdecl)>
    Public Shared Function ClearHistory() As Integer
    End Function
End Class

' Example usage class
Public Class Program
    Public Shared Sub Main()
        Console.WriteLine("g4n Math Library Example")
        Console.WriteLine("=============================")
        
        ' Basic arithmetic operations
        Dim result1 As Double = MathLibrary.Add(10.5, 5.3)
        Console.WriteLine($"10.5 + 5.3 = {result1}")
        
        Dim result2 As Double = MathLibrary.Subtract(20.0, 8.0)
        Console.WriteLine($"20.0 - 8.0 = {result2}")
        
        Dim result3 As Double = MathLibrary.Multiply(4.0, 7.0)
        Console.WriteLine($"4.0 * 7.0 = {result3}")
        
        ' Division with error handling
        Dim hasError As Integer = 0
        Dim result4 As Double = MathLibrary.Divide(15.0, 3.0, hasError)
        If hasError = 0 Then
            Console.WriteLine($"15.0 / 3.0 = {result4}")
        Else
            Console.WriteLine("Division error occurred")
        End If
        
        ' Advanced operations
        Dim result5 As Double = MathLibrary.Power(2.0, 8.0)
        Console.WriteLine($"2.0 ^ 8.0 = {result5}")
        
        ' Square root with error handling
        hasError = 0
        Dim result6 As Double = MathLibrary.SquareRoot(16.0, hasError)
        If hasError = 0 Then
            Console.WriteLine($"sqrt(16.0) = {result6}")
        Else
            Console.WriteLine("Square root error occurred")
        End If
        
        ' Trigonometric functions (input in radians)
        Dim angleRadians As Double = Math.PI / 4 ' 45 degrees
        Dim sinResult As Double = MathLibrary.Sine(angleRadians)
        Dim cosResult As Double = MathLibrary.Cosine(angleRadians)
        Dim tanResult As Double = MathLibrary.Tangent(angleRadians)
        
        Console.WriteLine($"sin(π/4) = {sinResult}")
        Console.WriteLine($"cos(π/4) = {cosResult}")
        Console.WriteLine($"tan(π/4) = {tanResult}")
        
        ' Clear calculation history
        Dim clearResult As Integer = MathLibrary.ClearHistory()
        If clearResult = 0 Then
            Console.WriteLine("History cleared successfully")
        Else
            Console.WriteLine("Failed to clear history")
        End If
        
        Console.WriteLine("Press any key to exit...")
        Console.ReadKey()
    End Sub
End Class
