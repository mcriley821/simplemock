package main

//go:generate simplemock -iface Empty -mock-name CustomMock -out os.Stdout

type Empty interface{}

func main() {
	//
}
