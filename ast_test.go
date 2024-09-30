package main

import (
	"bytes"
	"encoding/json"
	"go/parser"
	"go/token"
	"os"
	"path"
	"testing"

	"github.com/mcriley821/simplemock/env"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testData = "testdata"
	descJson = "desc.json"
)

type Desc struct {
	Input    string   `json:"input"`
	Args     []string `json:"args"`
	Expected string   `json:"expected"`
}

func TestGeneratesMocks(t *testing.T) {
	entries, err := os.ReadDir(testData)
	assert.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		assert.True(t, entry.Type().IsDir())

		t.Run(entry.Name(), func(t *testing.T) {
			descPath := path.Join(testData, entry.Name(), descJson)

			f, err := os.Open(descPath)
			assert.NoError(t, err)
			require.NotNil(t, f)

			desc := new(Desc)
			err = json.NewDecoder(f).Decode(desc)
			require.NoError(t, err)
			require.NotEmpty(t, desc.Args)

			cwd, err := os.Getwd()
			require.NoError(t, err)
			assert.NotEmpty(t, cwd)

			input := path.Join(cwd, testData, entry.Name(), desc.Input)

			t.Setenv("GOFILE", input)

			env, err := env.Load()
			assert.NoError(t, err)
			require.NotNil(t, env)

			expected, err := os.ReadFile(path.Join(testData, entry.Name(), desc.Expected))
			assert.NoError(t, err)
			require.NotEmpty(t, expected)

			fset := token.NewFileSet()
			ast, err := parser.ParseFile(fset, input, nil, parser.SkipObjectResolution)
			assert.NoError(t, err)
			require.NotNil(t, ast)

			iface := desc.Args[0]

			actual := &bytes.Buffer{}
			err = GenerateMockFromAst(iface, actual, env)
			assert.NoError(t, err)
			assert.Equal(t, string(expected), actual.String())
		})
	}
}
