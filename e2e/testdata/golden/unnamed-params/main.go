package main

//go:generate simplemock -iface Reader -out os.Stdout

type Reader interface {
	Read([]byte) (int, error)
}

func main() {}
