package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	tmpDir, err := os.MkdirTemp("", "simplemock-e2e-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(tmpDir)

	binaryPath := filepath.Join(tmpDir, "simplemock")
	cmd := exec.Command("go", "build", "-o", binaryPath, "..")
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic("failed to build simplemock: " + err.Error())
	}

	origPath := os.Getenv("PATH")
	os.Setenv("PATH", tmpDir+string(os.PathListSeparator)+origPath)
	defer os.Setenv("PATH", origPath)

	os.Exit(m.Run())
}
