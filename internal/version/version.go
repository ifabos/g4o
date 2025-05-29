package version

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// Build-time variables (injected during compilation)
var (
	// Version holds the semantic version (e.g., "1.0")
	Version = "unknown"
	// GitCommit holds the short git commit hash
	GitCommit = "unknown"
	// BuildTime holds the build timestamp
	BuildTime = "unknown"
	// GitBranch holds the git branch name
	GitBranch = "unknown"
)

// VersionInfo contains complete version information
type VersionInfo struct {
	Version     string `json:"version"`
	GitCommit   string `json:"git_commit"`
	BuildTime   string `json:"build_time"`
	GitBranch   string `json:"git_branch"`
	FullVersion string `json:"full_version"`
}

// GetVersionInfo returns complete version information
func GetVersionInfo() *VersionInfo {
	fullVersion := fmt.Sprintf("%s-%s", Version, GitCommit)
	if Version == "unknown" || GitCommit == "unknown" {
		// Try to extract from runtime git info if build-time injection failed
		if runtimeVersion, err := ExtractVersionFromGit(); err == nil {
			if runtimeVersion.Version != "unknown" {
				Version = runtimeVersion.Version
				GitCommit = runtimeVersion.GitCommit
				GitBranch = runtimeVersion.GitBranch
				fullVersion = fmt.Sprintf("%s-%s", Version, GitCommit)
			}
		}
	}

	return &VersionInfo{
		Version:     Version,
		GitCommit:   GitCommit,
		BuildTime:   BuildTime,
		GitBranch:   GitBranch,
		FullVersion: fullVersion,
	}
}

// GetVersion returns the version string (e.g., "1.0-abc1234")
func GetVersion() string {
	info := GetVersionInfo()
	return info.FullVersion
}

// GetVersionShort returns just the semantic version (e.g., "1.0")
func GetVersionShort() string {
	info := GetVersionInfo()
	return info.Version
}

// ExtractVersionFromGit extracts version information from git repository
func ExtractVersionFromGit() (*VersionInfo, error) {
	// Get current branch name
	branchBytes, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get branch name: %w", err)
	}
	branch := strings.TrimSpace(string(branchBytes))

	// Extract version from branch name (format: feature/x.y)
	version, err := ExtractVersionFromBranch(branch)
	if err != nil {
		return nil, fmt.Errorf("failed to extract version from branch '%s': %w", branch, err)
	}

	// Get short commit hash
	commitBytes, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit hash: %w", err)
	}
	commit := strings.TrimSpace(string(commitBytes))

	fullVersion := fmt.Sprintf("%s-%s", version, commit)

	return &VersionInfo{
		Version:     version,
		GitCommit:   commit,
		BuildTime:   "runtime",
		GitBranch:   branch,
		FullVersion: fullVersion,
	}, nil
}

// ExtractVersionFromBranch extracts version from branch name
// Expected format: "feature/x.y" where x.y is the version
func ExtractVersionFromBranch(branchName string) (string, error) {
	if branchName == "" {
		return "", errors.New("branch name is empty")
	}

	// Regex pattern to match "feature/x.y" format
	pattern := `^feature/(\d+\.\d+)$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(branchName)
	if len(matches) != 2 {
		return "", fmt.Errorf("branch name '%s' does not match expected format 'feature/x.y'", branchName)
	}

	return matches[1], nil
}

// IsValidVersionFormat checks if a version string follows semver format (x.y)
func IsValidVersionFormat(version string) bool {
	pattern := `^\d+\.\d+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(version)
}

// GetBuildInfo returns formatted build information
func GetBuildInfo() string {
	info := GetVersionInfo()
	return fmt.Sprintf("g4n v%s (commit: %s, branch: %s, built: %s)",
		info.FullVersion, info.GitCommit, info.GitBranch, info.BuildTime)
}
