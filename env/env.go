package env

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GoEnv struct {
	GOFILE     string
	GOROOT     string
	GOMOD      string
	GOMODCACHE string
}

// Load environment and 'go env' variables.
func Load() (*GoEnv, error) {
	cmd := exec.Command("go", "env", "GOROOT", "GOMOD", "GOMODCACHE")

	output := &bytes.Buffer{}
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("executing 'go env': %w", err)
	}

	// note: there's always a trailing newline in the output
	envs := strings.Split(output.String(), "\n")
	if len(envs) != 4 {
		return nil, errors.New("failed reading 'go env' output")
	}

	return &GoEnv{
		GOFILE:     os.Getenv("GOFILE"),
		GOROOT:     envs[0],
		GOMOD:      envs[1],
		GOMODCACHE: envs[2],
	}, nil
}
