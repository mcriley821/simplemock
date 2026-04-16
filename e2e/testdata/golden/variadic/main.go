package main

import "io"

//go:generate simplemock -iface Variadic -out os.Stdout

type Variadic interface {
	Foo(x1, x2 int, xs ...int) error
	Bar(xs ...io.Reader) (string, error)
}

func main() {
	//
}
