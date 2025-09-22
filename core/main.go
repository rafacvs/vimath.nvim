package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "example1.txt", "Filename to parse (searched in examples/ directory)")
	flag.Parse()

	filePath := filepath.Join("examples", filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File not found: %s", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("########### Parsing file: %s ###########\n", filePath)

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
