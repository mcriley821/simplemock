package main

// NotAnInterface is a struct, not an interface — used to test the
// "type is not an interface" error path.
type NotAnInterface struct{}

// unexported is an unexported interface — used to test the
// "interface is not exported" error path.
type unexported interface {
	Do() error
}

func main() {}
