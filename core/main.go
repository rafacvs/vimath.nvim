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
		fmt.Printf("Line: %+v\nTokens: %+v\n", lexerToken.String, lexerToken.Tokens)
	}
	fmt.Print("########### Finished printing outputs... ###########\n")
}
