package imported_return

import (
	"os"
)

type ImportedReturn interface {
	ImportedReturn() *os.File
}

