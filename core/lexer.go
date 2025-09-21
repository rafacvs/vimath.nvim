package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type TokenType int

const (
	NUMBER TokenType = iota
	IDENTIFIER
	EQUAL
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	COMMENT
)

func (t TokenType) String() string {
	switch t {
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case EQUAL:
		return "EQUAL"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULTIPLY:
		return "MULTIPLY"
	case DIVIDE:
		return "DIVIDE"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case COMMENT:
		return "COMMENT"
	default:
		return "unknown"
	}
}

type Token struct {
	Type   TokenType
	Lexeme string
}

func (tok Token) String() string {
	return fmt.Sprintf("%v(%q)", tok.Type, tok.Lexeme)
}

type Lexer struct{}

func NewLexer() *Lexer {
	return &Lexer{}
}

type LexerTokens struct {
	String string
	Tokens []Token
}

/*
* TODO: move file reading to main.go. Kept here for
* now to make testing easier, as I'm not really sure
* what I'm doing yet. Wont move to NewLexer either
* as eventually it'll go to main.
 */
func (l *Lexer) LexicAnalysis(file *os.File) []LexerTokens {
	fmt.Print("########### Starting lexical analysis... ###########\n")
	var lexerLines []LexerTokens
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := tokenizeLine(line)
		fmt.Printf("[line] %s\n", line)
		fmt.Printf("[tokens] %v\n", tokens)

		lexerLines = append(lexerLines, LexerTokens{String: line, Tokens: tokens})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("########### Finished lexical analysis... ###########\n")
	fmt.Print("\n\n\n")
	return lexerLines
}

func tokenizeLine(line string) []Token {
	var tokens []Token

	if len(line) == 0 {
		return tokens
	}

	if line[0] == '#' {
		tokens = append(tokens, Token{Type: COMMENT, Lexeme: line})
		return tokens
	}

	i := 0
	for i < len(line) {
		char := line[i]

		if char == ' ' || char == '\t' {
			i++
			continue
		}

		if isDigit(char) {
			start := i
			for i < len(line) && (isDigit(line[i]) || line[i] == '.') {
				if start == i && line[i] == '.' {
					i++
					continue
				}
				i++
			}

			tokens = append(tokens, Token{Type: NUMBER, Lexeme: line[start:i]})
			continue
		}

		if isAlpha(char) {
			start := i
			for i < len(line) && (isAlpha(line[i]) || isDigit(line[i]) || line[i] == '_') && char != ' ' {
				i++
			}

			tokens = append(tokens, Token{Type: IDENTIFIER, Lexeme: line[start:i]})
			continue
		}

		switch char {
		case '=':
			tokens = append(tokens, Token{Type: EQUAL, Lexeme: string(char)})
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Lexeme: string(char)})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Lexeme: string(char)})
		case '*':
			tokens = append(tokens, Token{Type: MULTIPLY, Lexeme: string(char)})
		case '/':
			tokens = append(tokens, Token{Type: DIVIDE, Lexeme: string(char)})
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Lexeme: string(char)})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Lexeme: string(char)})
		default:
			fmt.Printf("Unexpected character: %c\n", char)
		}

		i++
	}

	return tokens
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
