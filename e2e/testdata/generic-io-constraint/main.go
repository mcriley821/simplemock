package main

import "io"

//go:generate simplemock -iface Wrapper -out os.Stdout

// Wrapper is a generic interface with io.Reader as a type constraint,
// verifying that the io package is imported in the generated mock.
type Wrapper[T io.Reader] interface {
	Wrap() T
}

func main() {}
