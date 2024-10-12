package main

import (
	"empty/pkg"
)

//go:generate simplemock pkg.Suite os.Stdout

func main() {
	var _ pkg.Suite
}
