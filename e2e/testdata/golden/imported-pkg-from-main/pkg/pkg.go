package pkg

import "github.com/stretchr/testify/suite"

type Suite interface {
	TestSuite() suite.TestingSuite
}
