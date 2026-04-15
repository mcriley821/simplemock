package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimplemockErrors(t *testing.T) {
	root, err := os.Getwd()
	require.NoError(t, err)

	emptyMainDir := filepath.Join(root, "testdata", "empty-main")
	errorTypesDir := filepath.Join(root, "testdata", "error-types")

	testCases := []struct {
		name               string
		dir                string
		goFile             string // value for GOFILE env var; empty means omit GOFILE
		args               []string
		wantExitCode       int
		wantStderrContains string
	}{
		{
			name:               "missing iface flag",
			dir:                emptyMainDir,
			goFile:             "main.go",
			args:               []string{"-out", "os.Stdout"},
			wantExitCode:       1,
			wantStderrContains: "option '-iface' is required",
		},
		{
			name:               "missing out flag",
			dir:                emptyMainDir,
			goFile:             "main.go",
			args:               []string{"-iface", "Empty"},
			wantExitCode:       1,
			wantStderrContains: "option '-out' is required",
		},
		{
			name:               "missing GOFILE env",
			dir:                emptyMainDir,
			goFile:             "", // GOFILE intentionally omitted
			args:               []string{"-iface", "Empty", "-out", "os.Stdout"},
			wantExitCode:       1,
			wantStderrContains: "Expected GOFILE environment variable to be set.",
		},
		{
			name:               "type not found",
			dir:                emptyMainDir,
			goFile:             "main.go",
			args:               []string{"-iface", "NonExistent", "-out", "os.Stdout"},
			wantExitCode:       1,
			wantStderrContains: "Could not find type 'NonExistent'",
		},
		{
			name:               "non-interface type",
			dir:                errorTypesDir,
			goFile:             "main.go",
			args:               []string{"-iface", "NotAnInterface", "-out", "os.Stdout"},
			wantExitCode:       1,
			wantStderrContains: "is not an interface",
		},
		{
			name:               "non-exported interface",
			dir:                errorTypesDir,
			goFile:             "main.go",
			args:               []string{"-iface", "unexported", "-out", "os.Stdout"},
			wantExitCode:       1,
			wantStderrContains: "is not exported",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("simplemock", tc.args...)
			cmd.Dir = tc.dir

			// Build env: start from the current process env, strip any existing
			// GOFILE, then add the test-specific value (if any).
			env := make([]string, 0, len(os.Environ())+1)
			for _, e := range os.Environ() {
				if !strings.HasPrefix(e, "GOFILE=") {
					env = append(env, e)
				}
			}
			if tc.goFile != "" {
				env = append(env, "GOFILE="+tc.goFile)
			}
			cmd.Env = env

			var stderr bytes.Buffer
			cmd.Stderr = &stderr

			runErr := cmd.Run()
			require.Error(t, runErr, "expected simplemock to exit non-zero")

			exitErr, ok := runErr.(*exec.ExitError)
			require.True(t, ok, "expected *exec.ExitError, got %T: %v", runErr, runErr)
			assert.Equal(t, tc.wantExitCode, exitErr.ExitCode())
			assert.Contains(t, stderr.String(), tc.wantStderrContains)
		})
	}
}
