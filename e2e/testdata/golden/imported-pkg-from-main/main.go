package main

import (
	"empty/pkg"
)

//go:generate simplemock -iface pkg.Suite -out os.Stdout

func main() {
	var _ pkg.Suite
}
