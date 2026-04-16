package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimplemockFileOut(t *testing.T) {
	root, err := os.Getwd()
	require.NoError(t, err)

	dir := filepath.Join(root, "testdata", "golden", "empty-main")
	outFile := filepath.Join(dir, "empty_mock.go")

	t.Cleanup(func() { os.Remove(outFile) })

	cmd := exec.Command("simplemock", "-iface", "Empty", "-out", "empty_mock.go")
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFILE="+filepath.Join(dir, "main.go"))

	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "simplemock output: %s", output)

	actual, err := os.ReadFile(outFile)
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join(dir, "expected.txt"))
	require.NoError(t, err)

	assert.Equal(t, string(expected), string(actual))
}
