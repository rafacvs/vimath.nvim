package main

import (
	"fmt"
)

type Evaluator struct {
	symbols map[string]float64
}

func NewEvaluator() *Evaluator {
	return &Evaluator{
		symbols: make(map[string]float64),
	}
}

func (e *Evaluator) Eval(expr Expression) (val float64, err error) {
	switch node := expr.(type) {
	case *Number:
		return node.Value, nil
	case *Identifier:
		val, exists := e.symbols[node.Name]
		if !exists {
			// fmt.Printf("undeclared variable: %s\n", node.Name)
		}

		return val, nil
	case *ParenExpr:
		inner, err := e.Eval(node.Inner)
		if err != nil {
			// fmt.Printf("*ParenExpr: no inner\n")
		}

		return inner, nil
	case *BinaryExpr:
		left, err := e.Eval(node.Left)
		if err != nil {
			return 0, fmt.Errorf("*BinaryExpr: no left val")
		}
		right, err := e.Eval(node.Right)
		if err != nil {
			return 0, fmt.Errorf("*BinaryExpr: no right val")
		}

		switch node.Op {
		case PLUS:
			return left + right, nil
		case MINUS:
			return left - right, nil
		case MULTIPLY:
			return left * right, nil
		case DIVIDE:
			return left / right, nil
		}
	case *UnaryExpr:
		right, err := e.Eval(node.Right)
		if err != nil {
			return 0, fmt.Errorf("*BinaryExpr: no right val")
		}

		switch node.Op {
		case PLUS:
			return right, nil
		case MINUS:
			return right * -1, nil
		}
	case *AssignmentStmt:
		val, err := e.Eval(node.Value)
		if err != nil {
			return 0, fmt.Errorf("*AssignmentStmt: no eval")
		}
		e.symbols[node.Name] = val
		return val, nil
	}
	return 0, fmt.Errorf("unknown expression type: %T", expr)
}
