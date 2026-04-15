package main

// GenericRepo is a generic interface — used to test the
// "generic interfaces are not supported" error path.
type GenericRepo[T any] interface {
	Get() T
}

func main() {}
