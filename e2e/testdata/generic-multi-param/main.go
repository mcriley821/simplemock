package main

//go:generate simplemock -iface Pair -out os.Stdout

// Pair is a generic interface with two type parameters.
type Pair[K comparable, V any] interface {
	Key() K
	Val() V
}

func main() {}
