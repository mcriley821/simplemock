package empty_test

import (
	"empty/empty"
	"testing"
)

//go:generate simplemock -iface empty.Empty -out os.Stdout

func TestWithMock(t *testing.T) {
	var _ empty.Empty = &EmptyMock{}
}
