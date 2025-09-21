package main

import (
	"fmt"
	"strconv"
)

type Node struct{}

type Statement struct{ Node }
type Expression interface{}
type AssignmentStmt struct {
	Name  string
	Value Expression
}

type UnaryExpr struct {
	Op    TokenType
	Right Expression
}
type BinaryExpr struct {
	Op    TokenType
	Left  Expression
	Right Expression
}
type ParenExpr struct {
	Inner Expression
}

type Number struct{ Value float64 }
type Identifier struct{ Name string }

type Parser struct {
	Tokens []Token
	Pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{Tokens: tokens, Pos: 0}
}

func (p *Parser) parseAssignmentStmt() *AssignmentStmt {
	currentToken := p.Tokens[p.Pos]
	if p.Pos >= len(p.Tokens) || currentToken.Type != IDENTIFIER {
		return nil
	}
	name := currentToken.Lexeme
	p.Pos++

	if p.Pos >= len(p.Tokens) || p.Tokens[p.Pos].Type != EQUAL {
		return nil
	}
	p.Pos++

	val := p.parseExpression()
	if val != nil {
		fmt.Println("Error parsing expression:")
	}

	return &AssignmentStmt{Name: name, Value: val}
}

func (p *Parser) parseExpression() Expression {
	currentToken := p.Tokens[p.Pos]
	nextToken := p.Tokens[p.Pos+1]
	nextNextToken := p.Tokens[p.Pos+1]

	if currentToken.Type == NUMBER {
		if (nextToken.Type == PLUS || nextToken.Type == MINUS) && (nextNextToken.Type == NUMBER) {
			return &BinaryExpr{
				Left:  currentToken,
				Op:    nextToken.Type,
				Right: nextNextToken,
			}
		}
	}
	return nil
}
