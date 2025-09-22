package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "example1.txt", "Filename to parse (based on project's root dir)")
	flag.Parse()

	filePath := filepath.Join(filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File not found: %s", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("########### Parsing file: %s ###########\n", filePath)

	lexer := NewLexer()
	stringTokens := lexer.LexicAnalysis(lines)
	// fmt.Print("########### Starting printing outputs... ###########\n")

	evaluator := NewEvaluator()
	for _, lexerToken := range stringTokens {
		if len(lexerToken.Tokens) > 0 {
			// fmt.Printf("Line %02d > %+v\n", lexerToken.Index, lexerToken.String)
			// fmt.Printf("Tokens: %+v\n",  lexerToken.Tokens)

			if lexerToken.Tokens[0].Type != COMMENT {
				parser := NewParser(lexerToken.Tokens)
				assignments := parser.parseAssignmentStmt()
				if assignments != nil {
					// fmt.Printf("        > %s\n", assignments)
					val, err := evaluator.Eval(assignments.Value)
					if err == nil {
						evaluator.symbols[assignments.Name] = val
						// fmt.Printf("        > %f\n\n", val)
						fmt.Printf("%d %f\n", lexerToken.Index, val)
					} else {
						// fmt.Printf("[Error on EVAL]")
					}

				}
			}
		}
	}
	// fmt.Print("########### Finished printing outputs... ###########\n")
}
