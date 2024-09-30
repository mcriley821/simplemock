# simplemock
`simplemock` is a simplistic and opinionated Go mock generator designed to be
used with [`//go:generate`](https://go.dev/blog/generate) directives. It
generates minimalistic mock implementations for interfaces, making unit tests
more clear and descriptive without complicated frameworks and configurations.

## Installation
```bash
go install github.com/mcriley821/simplemock@latest
```

## Usage
To generate a mock for an interface, simply use a `//go:generate` directive:

```go
//go:generate simplemock interface
```

The provided `interface` must either be in the same package as the directive,
or accessible via imported packages.

## Example
`simplemock` generates a struct type that implements the provided interface.
This struct holds `func` data members named after the corresponding interface
method with a `Func` suffix. In each interface method, the corresponding `func`
data member is called if it is non-nil otherwise `panic` is called.

Here is a clarifying example:

**app.go**
```go
package app

import "errors"

type Store interface {
  Get(key string) (value string, err error)
  Set(key, value string) error
}

func Foo(store Store) error {
  value, err := store.Get("mykey")
  //
  err = store.Set("mykey", newValue)
}
```

**app_test.go**
```go
package app_test

import (
  "errors"

  "github.com/example/app"
)

// generates store_mock_test.go
//go:generate simplemock app.Store

func TestFoo(t *testing.T) {
  store := &StoreMock{
    GetFunc: func(key string) (value string, err error) { return "testvalue" },
    SetFunc: func(key, value string) error {
      // assert key and value are as expected
    },
  }
  // ...
  Foo(store)
}
```

**store_mock_test.go**
```go
// Code generated by simplemock. DO NOT EDIT.
package app_test

import (
  "errors"
  "github.com/example/app"
)

type StoreMock struct {
  GetFunc func(key string) (value string, err error)
  SetFunc func(key string, value string) error
}

func (m *StoreMock) Get(key string) (value string, err error) {
  if m.GetFunc != nil {
    return m.GetFunc(key)
  }
  panic("Get called with nil GetFunc")
}

func (m *StoreMock) Set(key string, value string) error {
  if m.SetFunc != nil {
    return m.SetFunc(key, value)
  }
  panic("Set called with nil SetFunc")
}
```
