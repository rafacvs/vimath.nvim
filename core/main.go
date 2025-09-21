package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("examples/example1.txt")
	file, err := os.Open("examples/parser1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lexer := NewLexer()
	stringTokens := lexer.LexicAnalysis(file)
	fmt.Print("########### Starting printing outputs... ###########\n")

	for _, lexerToken := range stringTokens {
		fmt.Printf("Line: %+v\nTokens: %+v\n", lexerToken.String, lexerToken.Tokens)

		if len(lexerToken.Tokens) > 0 && lexerToken.Tokens[0].Type != COMMENT {
			parser := NewParser(lexerToken.Tokens)
			assignments := parser.parseAssignmentStmt()
			fmt.Printf("Assignments: %+v\n\n", assignments)
		}
	}
	fmt.Print("########### Finished printing outputs... ###########\n")

}
