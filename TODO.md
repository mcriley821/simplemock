# TODO

## Bugs


- **Output file never closed** (`simplemock.go:234`): `os.Create` opens a file descriptor that is never explicitly closed. Because `os.Exit` is called throughout `main()`, deferred `Close` calls would not run anyway, but even on the success path the file is left open until the process exits, risking incomplete flushes on some platforms.

- **Package load errors silently ignored** (`simplemock.go:173`): After `packages.Load` returns, individual packages may still carry errors in their `pkg.Errors` slice even when the top-level `err` is `nil`. These are never inspected, so a partially-loaded or broken dependency can silently produce an incorrect mock.

- **Redundant parentheses on single-return methods** (`simplemock.go:343`): `typeString(sig.Results())` formats a `*types.Tuple`, which always adds surrounding parentheses. Single-return signatures are emitted as `func Foo() (error)` instead of `func Foo() error`. The output is valid Go but non-idiomatic and would be reformatted by `gofmt`.

## Architecture

- **E2E tests require a pre-installed binary** (`e2e/e2e_test.go:27`): The test suite shells out to `go generate`, which in turn invokes `simplemock` from `$PATH`. The tests therefore depend on an externally-installed binary and cannot verify the code under test in the same build. Running `go test ./e2e/...` on a clean checkout silently tests a stale or unrelated binary.

- **No unit tests for any helper function**: `signature`, `typeString`, `defaultedArgs`, `importsUsedBy`, and `relativeTo` are all untested in isolation. Only eight happy-path e2e scenarios exist, so regressions in these functions are invisible until they affect generated output.

- **`pkgs[0]` assumed to be the source file's package** (`simplemock.go:177,227`): `packages.Config{Tests: true}` instructs the loader to also load test variants of packages. The loader can return multiple packages in undefined order, but the code unconditionally uses `pkgs[0]` for scope lookup, package-name determination, and import resolution. There is no verification that `pkgs[0]` actually corresponds to `GOFILE`.

- **Mock name is fixed as `{Interface}Mock`** (`simplemock.go:245`): There is no `-mock-name` flag. If a type named `FooMock` already exists in the target package the generated file will not compile, and there is no way to work around it without patching the output manually.

- **`os.Exit` prevents deferred cleanup** (`simplemock.go` throughout): `os.Exit` is called in at least ten places inside `main()`. Any resource that would be released via `defer` (e.g., closing the output file, flushing a writer) is silently skipped. Refactoring `main` to return an error and exit once at the top level would fix this.

## Missing Test Coverage

- **No test for the `-out <filename>` code path**: Every e2e test uses `-out os.Stdout`. The `os.Create` branch (writing to a real file) is completely untested.

- **No tests for error conditions**: There are no tests for type-not-found, non-exported interface, non-interface type, or missing required flags. Error messages and exit codes are unverified.

- **No test for interfaces with embedded interfaces**: A common Go pattern such as `type ReadWriter interface { io.Reader; io.Writer }` is not exercised. The AST visitor and import-collection logic may behave incorrectly for embedded types from external packages.

- **No test for interfaces with unnamed parameters**: An interface method like `Read([]byte) (int, error)` (no parameter names) relies on the `arg%d` fallback in `signature()` and `defaultedArgs()`, but this path has no dedicated test.

## Future-Proofing

- **No support for generic interfaces** (Go 1.18+): Interfaces with type parameters (`type Repo[T any] interface { Get() T }`) are not handled. The tool will reject them with a generic "not an interface" or "not a named type" error rather than a clear unsupported-feature message.

- **`justfile` `test` target does not enforce `install`**: The `test` recipe calls `go test ./e2e/...` but the `install` step that builds and places the binary in `$PATH` is a separate recipe with no declared dependency. Running `just test` without `just install` will test whatever binary is currently installed (possibly none, or a previous version).

- **No `-mock-name` flag for naming flexibility**: Beyond collision avoidance, many projects use naming conventions different from `{Interface}Mock` (e.g., `Fake{Interface}`, `{Interface}Stub`, `Mock{Interface}`). Hardcoding the suffix limits adoption without offering an escape hatch.
