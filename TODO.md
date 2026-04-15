# TODO

## Bugs

## Architecture

## Future-Proofing

- **No support for generic interfaces** (Go 1.18+): Interfaces with type parameters (`type Repo[T any] interface { Get() T }`) are not handled. The tool will reject them with a generic "not an interface" or "not a named type" error rather than a clear unsupported-feature message.
