package main

//go:generate simplemock -iface Empty -out os.Stdout

type Empty interface{}

func main() {
	//
}
