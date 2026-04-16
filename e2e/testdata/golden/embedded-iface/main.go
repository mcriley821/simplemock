package main

import "io"

//go:generate simplemock -iface ReadWriter -out os.Stdout

type ReadWriter interface {
	io.Reader
	io.Writer
}

func main() {
	var _ io.ReadWriter
}
