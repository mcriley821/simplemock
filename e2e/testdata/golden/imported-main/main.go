package main

import (
	"github.com/stretchr/testify/suite"
)

//go:generate simplemock -iface suite.TestingSuite -out os.Stdout

func main() {
	var _ suite.TestingSuite
}
