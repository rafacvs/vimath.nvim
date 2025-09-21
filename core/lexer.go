package main

import (
	"bufio"
	"log"
	"os"
)

type TokenType = int

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

type Token struct {
	Type   TokenType
	Lexeme string
}

func lexer(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokenizeLine(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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

		if isAlpha(char) {

		}

		if isDigit(char) {

		}

		// TODO: handle other token types

	}

	return tokens
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
