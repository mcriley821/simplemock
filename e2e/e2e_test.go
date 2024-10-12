package e2e

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimplemock(t *testing.T) {
	testCases, err := os.ReadDir("testdata")
	assert.NoError(t, err)
	require.NotNil(t, testCases)

	root, err := os.Getwd()
	require.NoError(t, err)

	for _, testCase := range testCases {
		t.Run(testCase.Name(), func(t *testing.T) {
			require.True(t, testCase.IsDir(), "expected %s to be a module directory", testCase.Name())

			dirName := path.Join(root, "testdata", testCase.Name())

			cmd := exec.Command("go", "generate", "./...")
			cmd.Dir = dirName

			actual, err := cmd.Output()
			require.NoError(t, err, "output: %s", actual)

			expected, err := os.ReadFile(path.Join(dirName, "expected.txt"))
			assert.NoError(t, err)
			require.NotNil(t, expected)

			assert.Equal(t, string(expected), string(actual))
		})
	}
}
