package main

import (
	"errors"
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

func (p *Parser) Empty() bool {
	return p.Pos >= len(p.Tokens)
}

func (p *Parser) Peek() (token Token, err error) {
	if p.Empty() {
		return Token{}, errors.New("[peek] No more tokens")
	}

	return p.Tokens[p.Pos], nil
}

func (p *Parser) Advance() (token Token, err error) {
	if p.Empty() {
		return Token{}, errors.New("[advance] No more tokens")
	}

	current := p.Tokens[p.Pos]
	p.Pos++

	return current, nil
}

func (p *Parser) parseAssignmentStmt() *AssignmentStmt {
	currentToken, err := p.Advance()
	if err != nil {
		fmt.Printf("[parseAssignmentStmt] %s", err)
		return nil
	}
	if p.Empty() || currentToken.Type != IDENTIFIER {
		return nil
	}

	name := currentToken.Lexeme

	currentToken, err = p.Advance()
	if err != nil {
		fmt.Printf("[parseAssignmentStmt] %s", err)
		return nil
	}
	if p.Empty() || currentToken.Type != EQUAL {
		return nil
	}

	val := p.parseExpression()
	if val != nil {
		return &AssignmentStmt{Name: name, Value: val}
	}

	fmt.Println("Error parsing expression:")
	return nil
}

/*
* TODO: Refactor ensuring it's dyanamic, supporting
* chaining expressions instead of a sequential algorithm
 */
func (p *Parser) parseExpression() Expression {
	currentToken, err := p.Peek()
	if err != nil {
		fmt.Printf("[parseExpression] %s\n", err)
		return nil
	}

	if currentToken.Type == NUMBER {
		currentToken, err := p.Advance()
		if err != nil {
			fmt.Printf("[parseExpression] %s\n", err)
			return nil
		}

		lval, _ := strconv.ParseFloat(currentToken.Lexeme, 64)

		if p.Empty() {
			return &Number{Value: lval}
		}

		// p.Empty() check above ensures Peek won't error
		currentToken, _ = p.Peek()

		IS_OP := currentToken.Type == PLUS ||
			currentToken.Type == MINUS ||
			currentToken.Type == MULTIPLY ||
			currentToken.Type == DIVIDE

		if IS_OP {
			op := currentToken.Type

			currentToken, err = p.Advance()
			if err != nil {
				fmt.Printf("[parseExpression] %s\n", err)
				return nil
			}

			if currentToken.Type == NUMBER {
				rval, _ := strconv.ParseFloat(currentToken.Lexeme, 64)
				p.Advance()
				return &BinaryExpr{
					Left:  &Number{Value: lval},
					Op:    op,
					Right: &Number{Value: rval},
				}
			}
		}

		return &Number{Value: lval}
	}

	// handles single identifier not preceeded by an expression
	if currentToken.Type == IDENTIFIER {
		name := currentToken.Lexeme
		p.Advance()
		return &Identifier{Name: name}
	}

	if currentToken.Type == LPAREN {
		currentToken, err = p.Advance()
		if err != nil {
			fmt.Printf("[parseExpression] %s\n", err)
			return nil
		}

		expr := p.parseExpression()
		currentToken, err = p.Peek() // p.parseExpression() modifies currentToken
		if err != nil {
			fmt.Printf("[parseExpression] %s\n", err)
		}
		if currentToken.Type != RPAREN {
			fmt.Printf("[parseExpression] Expected ')'\n")
			return nil
		}

		p.Advance() // consume RPAREN
		return &ParenExpr{Inner: expr}
	}

	return nil
}

func (p *Parser) parseTerm() {

}

func (p *Parser) parseFactor() Expression {
	currentToken, err := p.Peek()
	if err != nil {
		fmt.Printf("[parseFactor] %s\n", err)
	}

	switch currentToken.Type {
	case NUMBER:
		value, _ := strconv.ParseFloat(currentToken.Lexeme, 64)
		p.Advance()

		return &Number{Value: value}
	case IDENTIFIER:
		p.Advance()
		name := currentToken.Lexeme

		return &Identifier{Name: name}
	case LPAREN:
		p.Advance()
		expr := p.parseExpression()

		currentToken, err = p.Advance()
		if err != nil {
			fmt.Printf("[parseFactor] %s\n", err)
		}
		if currentToken.Type != RPAREN {
			fmt.Printf("[parseFactor] Expected ')'\n")
			return nil
		}

		return &ParenExpr{Inner: expr}
	case PLUS, MINUS:
		p.Advance()
		right := p.parseFactor()
		return &UnaryExpr{Op: currentToken.Type, Right: right}
	default:
		return nil
	}
}
