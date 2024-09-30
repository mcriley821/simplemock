package imported_args

import (
	"os"
)

type ImportedArg interface {
	ImportedArg(*os.File)
}

