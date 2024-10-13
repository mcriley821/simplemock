package pkg

import "github.com/stretchr/testify/suite"

//go:generate simplemock -iface Suite -out os.Stdout

type Suite interface {
	TestSuite() suite.TestingSuite
}
