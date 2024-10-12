package pkg

import "github.com/stretchr/testify/suite"

//go:generate simplemock Suite os.Stdout

type Suite interface {
	TestSuite() suite.TestingSuite
}
