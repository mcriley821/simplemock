package main

import (
	"empty/empty"
)

//go:generate simplemock empty.Empty

func main() {
	var _ empty.Empty = nil
}
