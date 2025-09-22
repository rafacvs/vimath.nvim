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
		val, exists := e.symbols[node.String()]
		if !exists {
			fmt.Printf("undeclared variable: %s\n", node.String())
		}

		return val, nil
	case *ParenExpr:
		inner, err := e.Eval(node.Inner)
		if err != nil {
			fmt.Printf("*ParenExpr: no inner\n")
		}

		return inner, nil
	}

	return 0, fmt.Errorf("unknown expression type: %T", expr)
}
