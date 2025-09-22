package main

import (
	"fmt"
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
	Index  int
}

func (l *Lexer) LexicAnalysis(lines []string) []LexerTokens {
	var lexerLines []LexerTokens
	for index, line := range lines {
		tokens := l.tokenizeLine(line)

		lexerTokens := LexerTokens{String: line, Tokens: tokens, Index: index}
		lexerLines = append(lexerLines, lexerTokens)
	}

	return lexerLines
}

func (l *Lexer) tokenizeLine(line string) []Token {
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
			tokens, i = l.handleDigit(tokens, line, i)
			continue
		}

		if isAlpha(char) {
			tokens, i = l.handleAlpha(tokens, line, i)
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
			// fmt.Printf("Unexpected character: %c\n", char)
		}
		i++
	}

	return tokens
}

func (l *Lexer) handleDigit(tokens []Token, line string, i int) ([]Token, int) {
	start := i
	for i < len(line) && (isDigit(line[i]) || line[i] == '.') {
		if start == i && line[i] == '.' {
			i++
			continue
		}
		i++
	}

	tokens = append(tokens, Token{Type: NUMBER, Lexeme: line[start:i]})
	return tokens, i
}

func (l *Lexer) handleAlpha(tokens []Token, line string, i int) ([]Token, int) {
	start := i
	for i < len(line) && (isAlpha(line[i]) || isDigit(line[i]) || line[i] == '_') {
		i++
	}

	tokens = append(tokens, Token{Type: IDENTIFIER, Lexeme: line[start:i]})
	return tokens, i
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
