package main

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeString(t *testing.T) {
	ioReader := func() types.Type {
		pkg := types.NewPackage("io", "io")
		iface := types.NewInterfaceType(nil, nil)
		iface.Complete()
		return types.NewNamed(types.NewTypeName(token.NoPos, pkg, "Reader", nil), iface, nil)
	}()

	tests := []struct {
		name string
		typ  types.Type
		want string
	}{
		{"int", types.Typ[types.Int], "int"},
		{"string", types.Typ[types.String], "string"},
		{"bool", types.Typ[types.Bool], "bool"},
		{"byte/uint8", types.Typ[types.Byte], "uint8"},
		{"[]byte/[]uint8", types.NewSlice(types.Typ[types.Byte]), "[]uint8"},
		{"map[string]int", types.NewMap(types.Typ[types.String], types.Typ[types.Int]), "map[string]int"},
		{"*int", types.NewPointer(types.Typ[types.Int]), "*int"},
		{"error", types.Universe.Lookup("error").Type(), "error"},
		{"io.Reader", ioReader, "io.Reader"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, typeString(tc.typ))
		})
	}
}

func TestSignature(t *testing.T) {
	noPos := token.NoPos
	intType := types.Typ[types.Int]
	errType := types.Universe.Lookup("error").Type()
	byteSlice := types.NewSlice(types.Typ[types.Byte])
	strType := types.Typ[types.String]

	tests := []struct {
		name string
		sig  *types.Signature
		want string
	}{
		{
			"no params no results",
			types.NewSignatureType(nil, nil, nil, nil, nil, false),
			"()",
		},
		{
			"no params one result",
			types.NewSignatureType(nil, nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "", intType)),
				false),
			"() int",
		},
		{
			"no params two results",
			types.NewSignatureType(nil, nil, nil, nil,
				types.NewTuple(
					types.NewVar(noPos, nil, "", intType),
					types.NewVar(noPos, nil, "", errType),
				),
				false),
			"() (int, error)",
		},
		{
			"named param",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "n", intType)),
				nil, false),
			"(n int)",
		},
		{
			"unnamed param",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "", byteSlice)),
				types.NewTuple(
					types.NewVar(noPos, nil, "", intType),
					types.NewVar(noPos, nil, "", errType),
				),
				false),
			"(arg0 []uint8) (int, error)",
		},
		{
			"multiple named params",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(
					types.NewVar(noPos, nil, "a", strType),
					types.NewVar(noPos, nil, "b", strType),
				),
				nil, false),
			"(a string, b string)",
		},
		{
			"variadic named",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "args", types.NewSlice(intType))),
				nil, true),
			"(args ...int)",
		},
		{
			"variadic unnamed",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "", types.NewSlice(strType))),
				nil, true),
			"(arg0 ...string)",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, signature(tc.sig))
		})
	}
}

func TestDefaultedArgs(t *testing.T) {
	noPos := token.NoPos
	intType := types.Typ[types.Int]
	strType := types.Typ[types.String]

	tests := []struct {
		name string
		sig  *types.Signature
		want string
	}{
		{
			"no params",
			types.NewSignatureType(nil, nil, nil, nil, nil, false),
			"",
		},
		{
			"named params",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(
					types.NewVar(noPos, nil, "a", intType),
					types.NewVar(noPos, nil, "b", intType),
				),
				nil, false),
			"a, b",
		},
		{
			"unnamed params",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(
					types.NewVar(noPos, nil, "", intType),
					types.NewVar(noPos, nil, "", strType),
				),
				nil, false),
			"arg0, arg1",
		},
		{
			"mixed named and unnamed",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(
					types.NewVar(noPos, nil, "a", intType),
					types.NewVar(noPos, nil, "", strType),
				),
				nil, false),
			"a, arg1",
		},
		{
			"variadic named",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "args", types.NewSlice(intType))),
				nil, true),
			"args...",
		},
		{
			"variadic unnamed",
			types.NewSignatureType(nil, nil, nil,
				types.NewTuple(types.NewVar(noPos, nil, "", types.NewSlice(strType))),
				nil, true),
			"arg0...",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, defaultedArgs(tc.sig))
		})
	}
}

func TestRelativeTo(t *testing.T) {
	mainPkg := types.NewPackage("main", "main")
	ioPkg := types.NewPackage("io", "io")

	tests := []struct {
		name string
		pkg  *types.Package
		want string
	}{
		{"main package returns empty string", mainPkg, ""},
		{"non-main package returns package name", ioPkg, "io"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, relativeTo(tc.pkg))
		})
	}
}
