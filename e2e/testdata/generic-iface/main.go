package main

//go:generate simplemock -iface GenericRepo -out os.Stdout

// GenericRepo is a generic interface used to test generic interface support.
type GenericRepo[T any] interface {
	Get() T
}

func main() {}
