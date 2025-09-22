package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("examples/example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lexer := NewLexer()
	stringTokens := lexer.LexicAnalysis(file)
	fmt.Print("########### Starting printing outputs... ###########\n")

	for _, lexerToken := range stringTokens {
		if len(lexerToken.Tokens) > 0 {
			fmt.Printf("Line: %+v\n", lexerToken.String)
			// fmt.Printf("Tokens: %+v\n",  lexerToken.Tokens)

			if lexerToken.Tokens[0].Type != COMMENT {
				parser := NewParser(lexerToken.Tokens)
				assignments := parser.parseAssignmentStmt()
				fmt.Printf("Assignment: %s\n\n", assignments)
			}
		}
	}
	fmt.Print("########### Finished printing outputs... ###########\n")

}
