// Code generated by simplemock. DO NOT EDIT.
package input_test

type ReaderMock struct {
	ReadFunc func(p []byte) (n int, err error)
}

func (m *ReaderMock) Read(p []byte) (n int, err error) {
	if m.ReadFunc != nil {
		return m.ReadFunc(p)
	}
	panic("Read called with nil ReadFunc!")
}
