package pkg

import "github.com/stretchr/testify/suite"

//go:generate simplemock -iface suite.TestingSuite -out os.Stdout

var _ suite.TestingSuite
