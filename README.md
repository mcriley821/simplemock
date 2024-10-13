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
Invoke `simplemock` with `-help` to get up-to-date usage:

```sh
simplemock -help
```

```go
//go:generate simplemock -iface {interface} -out {filename}
```

The provided `interface` must either be in the same package as the directive,
or accessible via imported packages.

`simplemock` generates a struct type that implements the provided interface.
This struct holds `func` data members named after the corresponding interface
method with a `Func` suffix. In each mock method, the corresponding `func` data
member is called if it is non-nil - otherwise `panic` is called.

## Examples
See the [end-to-end](./e2e/testdata) test directories for concrete examples of
the different ways to utilize `simplemock`.

