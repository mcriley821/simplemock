# TODO

## Bugs


## Architecture

- **E2E tests require a pre-installed binary** (`e2e/e2e_test.go:27`): The test suite shells out to `go generate`, which in turn invokes `simplemock` from `$PATH`. The tests therefore depend on an externally-installed binary and cannot verify the code under test in the same build. Running `go test ./e2e/...` on a clean checkout silently tests a stale or unrelated binary.

- **`pkgs[0]` assumed to be the source file's package** (`simplemock.go:177,227`): `packages.Config{Tests: true}` instructs the loader to also load test variants of packages. The loader can return multiple packages in undefined order, but the code unconditionally uses `pkgs[0]` for scope lookup, package-name determination, and import resolution. There is no verification that `pkgs[0]` actually corresponds to `GOFILE`.

- **Mock name is fixed as `{Interface}Mock`** (`simplemock.go:245`): There is no `-mock-name` flag. If a type named `FooMock` already exists in the target package the generated file will not compile, and there is no way to work around it without patching the output manually.


## Future-Proofing

- **No support for generic interfaces** (Go 1.18+): Interfaces with type parameters (`type Repo[T any] interface { Get() T }`) are not handled. The tool will reject them with a generic "not an interface" or "not a named type" error rather than a clear unsupported-feature message.

- **No `-mock-name` flag for naming flexibility**: Beyond collision avoidance, many projects use naming conventions different from `{Interface}Mock` (e.g., `Fake{Interface}`, `{Interface}Stub`, `Mock{Interface}`). Hardcoding the suffix limits adoption without offering an escape hatch.
