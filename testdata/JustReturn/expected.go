// Code generated by simplemock. DO NOT EDIT.
package just_return_test

type JustReturnMock struct {
	JustReturnFunc func() int
}

func (m *JustReturnMock) JustReturn() int {
	if m.JustReturnFunc != nil {
		return m.JustReturnFunc()
	}
	panic("JustReturn called with nil JustReturnFunc!")
}
