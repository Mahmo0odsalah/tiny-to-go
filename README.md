# Tiny Basic Compiler

This is a compiler for the Tiny Basic programming language written in Go. It compiles from Tiny Basic to Go code.

## Features

- Lexical analysis
- Parsing
- Semantic analysis
- Code generation

## Getting Started

To get started with the Tiny Basic Compiler, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/tiny-basic-compiler.git`
2. Install Go: [https://golang.org/dl/](https://golang.org/dl/)
3. Build the compiler: `go build`
4. Run the compiler: `./tiny-basic-compiler input.tb output.asm`

## Examples

Here are some examples of Tiny Basic programs that can be compiled using this compiler:

```basic
10 PRINT "Hello, World!"
20 END
```

```basic
10 LET X = 5
20 PRINT X
30 END
```
