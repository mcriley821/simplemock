package main

import (
	"empty/pkg"
)

//go:generate simplemock pkg.Suite

func main() {
	var _ pkg.Suite
}
