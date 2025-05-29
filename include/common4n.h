#ifndef COMMON4N_H
#define COMMON4N_H

#ifdef __cplusplus
extern "C" {
#endif

// Basic arithmetic operations
double Add(double a, double b);
double Subtract(double a, double b);
double Multiply(double a, double b);
double Divide(double a, double b, int* hasError);

// Advanced mathematical operations
double Power(double base, double exponent);
double SquareRoot(double x, int* hasError);
double Logarithm(double x, int* hasError);

// Trigonometric functions (input in radians)
double Sine(double x);
double Cosine(double x);
double Tangent(double x);

// Utility functions
int ClearHistory();
char* GetLastErrorMessage();
void FreeErrorMessage(char* msg);

// Version information functions
char* GetVersion();           // Returns full version string (e.g., "1.0-abc1234")
char* GetVersionShort();      // Returns semantic version only (e.g., "1.0")
char* GetBuildInfo();         // Returns formatted build information
char* GetGitCommit();         // Returns git commit hash
char* GetGitBranch();         // Returns git branch name
char* GetBuildTime();         // Returns build timestamp
void FreeVersionString(char* str);  // Free version string memory

#ifdef __cplusplus
}
#endif

#endif // COMMON4N_H
