' Visual Basic 6 example for using g4n library
' This example demonstrates how to use the math library in VB6

Option Explicit

' Declare external functions from the DLL
Private Declare Function Add Lib "g4n.dll" (ByVal a As Double, ByVal b As Double) As Double
Private Declare Function Subtract Lib "g4n.dll" (ByVal a As Double, ByVal b As Double) As Double
Private Declare Function Multiply Lib "g4n.dll" (ByVal a As Double, ByVal b As Double) As Double
Private Declare Function Divide Lib "g4n.dll" (ByVal a As Double, ByVal b As Double, ByRef hasError As Long) As Double
Private Declare Function Power Lib "g4n.dll" (ByVal base As Double, ByVal exponent As Double) As Double
Private Declare Function SquareRoot Lib "g4n.dll" (ByVal x As Double, ByRef hasError As Long) As Double
Private Declare Function Logarithm Lib "g4n.dll" (ByVal x As Double, ByRef hasError As Long) As Double
Private Declare Function Sine Lib "g4n.dll" (ByVal x As Double) As Double
Private Declare Function Cosine Lib "g4n.dll" (ByVal x As Double) As Double
Private Declare Function Tangent Lib "g4n.dll" (ByVal x As Double) As Double
Private Declare Function ClearHistory Lib "g4n.dll" () As Long

' Example form with buttons for each operation
Private Sub Form_Load()
    Me.Caption = "g4n Math Library - VB6 Example"
End Sub

Private Sub cmdAdd_Click()
    Dim result As Double
    result = Add(10.5, 5.3)
    MsgBox "10.5 + 5.3 = " & CStr(result)
End Sub

Private Sub cmdSubtract_Click()
    Dim result As Double
    result = Subtract(20#, 8#)
    MsgBox "20.0 - 8.0 = " & CStr(result)
End Sub

Private Sub cmdMultiply_Click()
    Dim result As Double
    result = Multiply(4#, 7#)
    MsgBox "4.0 * 7.0 = " & CStr(result)
End Sub

Private Sub cmdDivide_Click()
    Dim result As Double
    Dim hasError As Long
    
    hasError = 0
    result = Divide(15#, 3#, hasError)
    
    If hasError = 0 Then
        MsgBox "15.0 / 3.0 = " & CStr(result)
    Else
        MsgBox "Division error occurred"
    End If
End Sub

Private Sub cmdPower_Click()
    Dim result As Double
    result = Power(2#, 8#)
    MsgBox "2.0 ^ 8.0 = " & CStr(result)
End Sub

Private Sub cmdSquareRoot_Click()
    Dim result As Double
    Dim hasError As Long
    
    hasError = 0
    result = SquareRoot(16#, hasError)
    
    If hasError = 0 Then
        MsgBox "sqrt(16.0) = " & CStr(result)
    Else
        MsgBox "Square root error occurred"
    End If
End Sub

Private Sub cmdTrigonometry_Click()
    Dim angleRadians As Double
    Dim sinResult As Double
    Dim cosResult As Double
    Dim tanResult As Double
    
    ' 45 degrees in radians (π/4)
    angleRadians = 3.14159265358979 / 4
    
    sinResult = Sine(angleRadians)
    cosResult = Cosine(angleRadians)
    tanResult = Tangent(angleRadians)
    
    MsgBox "sin(π/4) = " & CStr(sinResult) & vbCrLf & _
           "cos(π/4) = " & CStr(cosResult) & vbCrLf & _
           "tan(π/4) = " & CStr(tanResult)
End Sub

Private Sub cmdClearHistory_Click()
    Dim result As Long
    result = ClearHistory()
    
    If result = 0 Then
        MsgBox "History cleared successfully"
    Else
        MsgBox "Failed to clear history"
    End If
End Sub

' Utility function to convert degrees to radians
Private Function DegreesToRadians(degrees As Double) As Double
    DegreesToRadians = degrees * 3.14159265358979 / 180#
End Function

' Utility function to convert radians to degrees
Private Function RadiansToDegrees(radians As Double) As Double
    RadiansToDegrees = radians * 180# / 3.14159265358979
End Function
