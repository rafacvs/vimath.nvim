# vimath.nvim Grammar Specification

## Grammar Definition (EBNF-like)

```ebnf
(* Main program structure *)
program = { statement | comment | empty_line } ;

(* Statements *)
statement = assignment ;
assignment = identifier "=" expression ;

(* Expressions with precedence *)
expression = term { ("+" | "-") term } ;
term       = factor { ("*" | "/") factor } ;
factor     = number | identifier | "(" expression ")" | unary_op ;

(* Unary operations *)
unary_op = ("+" | "-") factor ;

(* Basic tokens *)
identifier = letter { letter | digit | "_" } ;
number     = digit { digit } [ "." digit { digit } ] ;
comment    = "#" { any } ;

(* Character classes *)
letter = "a".."z" | "A".."Z" ;
digit  = "0".."9" ;
any = ? any character except newline ? ;
```

## Operator Precedence

From highest to lowest precedence:

1. **Parentheses**: `( expression )`
2. **Unary operators**: `+expression`, `-expression`
3. **Multiplication/Division**: `*`, `/` (left-associative)
4. **Addition/Subtraction**: `+`, `-` (left-associative)
5. **Assignment**: `=` (right-associative)

## Token Types

| Token Type   | Examples                | Description                                  |
| ------------ | ----------------------- | -------------------------------------------- |
| `NUMBER`     | `5`, `3.14`, `0.5`      | Decimal numbers only                         |
| `IDENTIFIER` | `x`, `total`, `price_1` | Variable names (letters, digits, underscore) |
| `EQUAL`      | `=`                     | Assignment operator                          |
| `PLUS`       | `+`                     | Addition/unary plus                          |
| `MINUS`      | `-`                     | Subtraction/unary minus                      |
| `MULTIPLY`   | `*`                     | Multiplication                               |
| `DIVIDE`     | `/`                     | Division                                     |
| `LPAREN`     | `(`                     | Left parenthesis                             |
| `RPAREN`     | `)`                     | Right parenthesis                            |
| `COMMENT`    | `# comment`             | Single-line comments only                    |
