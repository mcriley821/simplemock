# TODO

## Bugs

- **Void-method return statement**: The generated mock always emits `return m.FuncName(args)` for every method. For methods with no return values this is a compile error (`(no value) used as value`). The template must distinguish between void and non-void methods and emit a bare `return` (or no return at all) for the former.

## Architecture

## Future-Proofing
