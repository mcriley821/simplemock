// Code generated by simplemock. DO NOT EDIT.
package imported_return_test

import "os"

type ImportedReturnMock struct {
	ImportedReturnFunc func() *os.File
}

func (m *ImportedReturnMock) ImportedReturn() *os.File {
	if m.ImportedReturnFunc != nil {
		return m.ImportedReturnFunc()
	}
	panic("ImportedReturn called with nil ImportedReturnFunc!")
}
