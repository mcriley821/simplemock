package main

import (
	"os"

	"github.com/mcriley821/simplemock/env"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: simplemock interface")
		os.Exit(1)
	}

	env, err := env.Load()
	if err != nil {
		println("Failed to load environment: " + err.Error())
		os.Exit(1)
	}

	if env.GOFILE == "" {
		println("Expected GOFILE environment variable to be set.")
		println("You should be using a //go:generate directive.")
		os.Exit(1)
	}

	if err := GenerateMockFromAst(os.Args[1], os.Stdout, env); err != nil {
		println("Failed to generate: %v" + err.Error())
		os.Exit(1)
	}
}
