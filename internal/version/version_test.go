package version

import (
	"strings"
	"testing"
)

func TestExtractVersionFromBranch(t *testing.T) {
	tests := []struct {
		name        string
		branchName  string
		expected    string
		expectError bool
	}{
		{
			name:        "valid feature branch 1.0",
			branchName:  "feature/1.0",
			expected:    "1.0",
			expectError: false,
		},
		{
			name:        "valid feature branch 2.5",
			branchName:  "feature/2.5",
			expected:    "2.5",
			expectError: false,
		},
		{
			name:        "valid feature branch 10.15",
			branchName:  "feature/10.15",
			expected:    "10.15",
			expectError: false,
		},
		{
			name:        "invalid format - no feature prefix",
			branchName:  "1.0",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format - main branch",
			branchName:  "main",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format - develop branch",
			branchName:  "develop",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format - three part version",
			branchName:  "feature/1.0.1",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format - no version",
			branchName:  "feature/",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid format - non-numeric version",
			branchName:  "feature/v1.0",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty branch name",
			branchName:  "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "feature branch with additional path",
			branchName:  "feature/1.0/test",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExtractVersionFromBranch(tt.branchName)

			if tt.expectError {
				if err == nil {
					t.Errorf("ExtractVersionFromBranch(%s) expected error but got none", tt.branchName)
				}
			} else {
				if err != nil {
					t.Errorf("ExtractVersionFromBranch(%s) unexpected error: %v", tt.branchName, err)
				}
				if result != tt.expected {
					t.Errorf("ExtractVersionFromBranch(%s) = %s; want %s", tt.branchName, result, tt.expected)
				}
			}
		})
	}
}

func TestIsValidVersionFormat(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected bool
	}{
		{"valid 1.0", "1.0", true},
		{"valid 2.5", "2.5", true},
		{"valid 10.15", "10.15", true},
		{"invalid v1.0", "v1.0", false},
		{"invalid 1.0.1", "1.0.1", false},
		{"invalid 1", "1", false},
		{"invalid empty", "", false},
		{"invalid alpha", "1.0-alpha", false},
		{"invalid letters", "a.b", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidVersionFormat(tt.version)
			if result != tt.expected {
				t.Errorf("IsValidVersionFormat(%s) = %v; want %v", tt.version, result, tt.expected)
			}
		})
	}
}

func TestGetVersionInfo(t *testing.T) {
	// Test with default unknown values
	info := GetVersionInfo()

	if info == nil {
		t.Error("GetVersionInfo() returned nil")
		return
	}

	// Check that struct is properly populated
	if info.Version == "" {
		t.Error("Version should not be empty")
	}
	if info.GitCommit == "" {
		t.Error("GitCommit should not be empty")
	}
	if info.FullVersion == "" {
		t.Error("FullVersion should not be empty")
	}

	// FullVersion should be in format "version-commit"
	parts := strings.Split(info.FullVersion, "-")
	if len(parts) != 2 {
		t.Errorf("FullVersion should be in format 'version-commit', got: %s", info.FullVersion)
	}
}

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion() should not return empty string")
	}

	// Should contain a dash for version-commit format
	if !strings.Contains(version, "-") {
		t.Errorf("GetVersion() should return version-commit format, got: %s", version)
	}
}

func TestGetVersionShort(t *testing.T) {
	versionShort := GetVersionShort()
	if versionShort == "" {
		t.Error("GetVersionShort() should not return empty string")
	}
}

func TestGetBuildInfo(t *testing.T) {
	buildInfo := GetBuildInfo()
	if buildInfo == "" {
		t.Error("GetBuildInfo() should not return empty string")
	}

	// Should contain "g4n" and "v"
	if !strings.Contains(buildInfo, "g4n") {
		t.Error("GetBuildInfo() should contain 'g4n'")
	}
	if !strings.Contains(buildInfo, "v") {
		t.Error("GetBuildInfo() should contain version prefix 'v'")
	}
}

// Test with mocked build-time variables
func TestGetVersionInfoWithBuildTime(t *testing.T) {
	// Save original values
	originalVersion := Version
	originalCommit := GitCommit
	originalBuildTime := BuildTime
	originalBranch := GitBranch

	// Set test values
	Version = "1.5"
	GitCommit = "abc1234"
	BuildTime = "2025-01-01T12:00:00Z"
	GitBranch = "feature/1.5"

	// Test
	info := GetVersionInfo()

	if info.Version != "1.5" {
		t.Errorf("Expected version 1.5, got %s", info.Version)
	}
	if info.GitCommit != "abc1234" {
		t.Errorf("Expected commit abc1234, got %s", info.GitCommit)
	}
	if info.FullVersion != "1.5-abc1234" {
		t.Errorf("Expected full version 1.5-abc1234, got %s", info.FullVersion)
	}

	// Restore original values
	Version = originalVersion
	GitCommit = originalCommit
	BuildTime = originalBuildTime
	GitBranch = originalBranch
}

func BenchmarkGetVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetVersion()
	}
}

func BenchmarkExtractVersionFromBranch(b *testing.B) {
	branchName := "feature/1.0"
	for i := 0; i < b.N; i++ {
		ExtractVersionFromBranch(branchName)
	}
}
