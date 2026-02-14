package main

import (
	"fmt"
	"tpl/lexer"
	"tpl/token"
)

func main() {

	input := `
	thalapathy count = 0

	IamWaiting (count < 5) {
		alliswell()
		count = count + 1
	}
	`

	fmt.Println("🔥 Thalapathy Programming Language Lexer 🔥")
	fmt.Println("===========================================")

	l := lexer.New(input)

	for {
		tok := l.NextToken()

		fmt.Printf("Type: %-15s Literal: %s\n", tok.Type, tok.Literal)

		if tok.Type == token.EOF {
			break
		}
	}
}
