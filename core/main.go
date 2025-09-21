package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("examples/example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lexer(file)
}
