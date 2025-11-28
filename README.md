# Beeflang

A meme language interpreter honoring the Church of Beef 🥩

## What is this?

Beeflang is a learning project to build a minimal Turing-complete programming language from scratch in Go. It features beef-themed keywords and a Python-like syntax.

## Example

```beeflang
# Hello, Beef!
praise genesis():
   cut x = 42
   preach(x)

   if x > 0:
      cut y = 5
   beef
beef
```

## Development

**Prerequisites:** Go 1.21+

**Quick Start:**
```bash
# Run tests
./dev.sh test

# Run a program (when implemented)
./dev.sh run examples/hello.beef

# Dump tokens for debugging
./dev.sh lex examples/hello.beef
```

## Status

🚧 **Early Development** - Currently building the lexer. See [CLAUDE.md](CLAUDE.md) for development workflow and [BEEFLANG_SPEC.md](BEEFLANG_SPEC.md) for language specification.

## Keywords

- `praise` - function declaration
- `beef` - block terminator
- `feast while` - while loop
- `if` / `else` - conditionals
- `cut` - variable declaration
- `serve` - return statement
- `genesis` - entry point (main function)
- `preach()` - print/output function

## License

This is a learning project. Do whatever you want with it.
