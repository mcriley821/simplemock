package main

import (
	"empty/empty"
)

//go:generate simplemock -iface empty.Empty -out os.Stdout

func main() {
	var _ empty.Empty = nil
}
