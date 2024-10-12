package main

import (
	"empty/empty"
)

//go:generate simplemock empty.Empty os.Stdout

func main() {
	var _ empty.Empty = nil
}
