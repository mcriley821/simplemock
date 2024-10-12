package main

import (
	"github.com/stretchr/testify/suite"
)

//go:generate simplemock suite.TestingSuite os.Stdout

func main() {
	var _ suite.TestingSuite
}
