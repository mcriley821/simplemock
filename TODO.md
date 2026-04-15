# TODO

## Bugs


## Architecture

- **E2E tests require a pre-installed binary** (`e2e/e2e_test.go:27`): The test suite shells out to `go generate`, which in turn invokes `simplemock` from `$PATH`. The tests therefore depend on an externally-installed binary and cannot verify the code under test in the same build. Running `go test ./e2e/...` on a clean checkout silently tests a stale or unrelated binary.

## Future-Proofing

- **No support for generic interfaces** (Go 1.18+): Interfaces with type parameters (`type Repo[T any] interface { Get() T }`) are not handled. The tool will reject them with a generic "not an interface" or "not a named type" error rather than a clear unsupported-feature message.
