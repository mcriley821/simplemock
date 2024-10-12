package pkg

import "github.com/stretchr/testify/suite"

//go:generate simplemock Suite

type Suite interface {
	TestSuite() suite.TestingSuite
}
