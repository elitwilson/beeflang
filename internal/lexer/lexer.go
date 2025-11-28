package lexer

// Lexer performs lexical analysis (tokenization) on source code.
// Lexical analysis is the first phase of an interpreter/compiler - it reads
// raw source code as a string and breaks it into "tokens" (meaningful chunks
// like keywords, identifiers, numbers, operators).
//
// Example: "cut x = 42" becomes tokens [CUT, IDENT("x"), ASSIGN, INT("42")]
//
// The lexer uses a two-pointer technique for reading:
// - position: points to the current character being examined
// - readPosition: points to the next character (lookahead for multi-char tokens like "==")
//
// Position tracking (line/column) is maintained throughout for error reporting.
// When we encounter a syntax error later, we can say "error at line 5, column 12"
// instead of just "syntax error somewhere".
type Lexer struct {
	input        string // the entire source code as a string
	position     int    // current position in input (current char)
	readPosition int    // next reading position (lookahead position)
	ch           byte   // current character under examination
	line         int    // current line number (starts at 1)
	column       int    // current column number (starts at 1)
}

// TODO: Implement lexer methods
