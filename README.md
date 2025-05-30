# Monter Programming Language

**Monter** is a simple, lightweight, and easy-to-use programming language designed for learning and experimentation. It features a custom lexer, parser, and interpreter, making it a great tool for understanding the fundamentals of programming language design.

## Features

- **Custom Lexer**: Tokenizes input strings into meaningful tokens.
- **Parser**: Converts tokens into an Abstract Syntax Tree (AST).
- **Interpreter**: Executes the parsed code.
- **REPL**: Interactive Read-Eval-Print Loop for live coding and testing.

## Getting Started

### Prerequisites

- Go 1.18 or later installed on your system.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/vickychhetri/mon-ter.git
   cd mon-ter
   ```

2. Build the project:
   ```bash
   go build -o monter
   ```

3. Run the REPL:
   ```bash
   ./monter
   ```

### Usage

Once you start the REPL, you will see the following welcome message:

```text
Hello <username>! This is mon-ter programming language!
Feel free to type in commands
```

You can then type in commands and see the results immediately. For example:

```text
>> let x = 5;
>> x + 10;
15
```

### Example Code

Here’s an example of a simple program in Monter:

```monter
let add = fn(x, y) {
    x + y;
};

let result = add(5, 10);
result;
```

## Development

### Running Tests

To run the tests, use:

```bash
go test ./...
```

### File Structure

- `lexer/`: Contains the lexer implementation.
- `parser/`: Contains the parser implementation.
- `repl/`: Contains the REPL implementation.
- `token/`: Defines token types and utilities.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Inspired by the book *Writing An Interpreter In Go* by Thorsten Ball.