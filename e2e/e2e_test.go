package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimplemock(t *testing.T) {
	testCases, err := os.ReadDir("testdata/golden")
	require.NoError(t, err)

	root, err := os.Getwd()
	require.NoError(t, err)

	for _, testCase := range testCases {
		t.Run(testCase.Name(), func(t *testing.T) {
			require.True(t, testCase.IsDir(), "expected %s to be a module directory", testCase.Name())

			dirName := filepath.Join(root, "testdata", "golden", testCase.Name())

			cmd := exec.Command("go", "generate", "./...")
			cmd.Dir = dirName

			actual, err := cmd.CombinedOutput()
			require.NoError(t, err, "output: %s", actual)

			expected, err := os.ReadFile(filepath.Join(dirName, "expected.txt"))
			require.NoError(t, err)

			assert.Equal(t, string(expected), string(actual))
		})
	}
}
