package main

import (
	"fmt"
	"os"

	"github.com/elitwilson/beeflang/internal/lexer"
	"github.com/elitwilson/beeflang/internal/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go <file.beef>")
		fmt.Println("  go run main.go --dump-tokens <file.beef>")
		os.Exit(1)
	}

	// Check for --dump-tokens flag
	dumpTokens := false
	filename := os.Args[1]

	if os.Args[1] == "--dump-tokens" {
		if len(os.Args) < 3 {
			fmt.Println("Error: --dump-tokens requires a filename")
			os.Exit(1)
		}
		dumpTokens = true
		filename = os.Args[2]
	}

	// Read source file
	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Dump tokens mode
	if dumpTokens {
		l := lexer.New(string(source))
		fmt.Printf("Tokens for %s:\n", filename)
		fmt.Println("---")
		for {
			tok := l.NextToken()
			fmt.Printf("%-15s %-10s (line %d, col %d)\n", tok.Type, tok.Literal, tok.Line, tok.Column)
			if tok.Type == token.EOF {
				break
			}
		}
		return
	}

	// Normal interpreter mode (not yet implemented)
	fmt.Printf("Read %d bytes from %s\n", len(source), filename)
	fmt.Println("Beeflang interpreter not yet implemented!")
}
