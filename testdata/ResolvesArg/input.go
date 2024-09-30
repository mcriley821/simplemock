package resolves_args

import (
	"net/url"
	"os"
)

var _ = url.URL{}

type ResolvesArg interface {
	ResolvesArg(file *os.File)
}

