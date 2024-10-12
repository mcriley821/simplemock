package main

import (
	"github.com/stretchr/testify/suite"
)

//go:generate simplemock suite.TestingSuite

func main() {
	var _ suite.TestingSuite
}
