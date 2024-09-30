package resolves_multi_args

import (
	"net/url"
	"os"
)

type ResolvesMultiArgs interface {
	ResolvesMultiArgs(*os.File, *url.URL)
}

