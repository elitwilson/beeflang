package lexer

import "github.com/elitwilson/beeflang/internal/token"

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

// New creates a new Lexer instance and initializes it by reading the first character
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar() // Initialize by reading first character
	return l
}

// NextToken reads the next token from the input and returns it
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Capture current position for this token
	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Line: tok.Line, Column: tok.Column}
		} else {
			tok = l.newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = l.newToken(token.PLUS, l.ch)
	case '-':
		tok = l.newToken(token.MINUS, l.ch)
	case '*':
		tok = l.newToken(token.ASTERISK, l.ch)
	case '/':
		tok = l.newToken(token.SLASH, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch), Line: tok.Line, Column: tok.Column}
		} else {
			tok = l.newToken(token.NOT, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch), Line: tok.Line, Column: tok.Column}
		} else {
			tok = l.newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch), Line: tok.Line, Column: tok.Column}
		} else {
			tok = l.newToken(token.GT, l.ch)
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch), Line: tok.Line, Column: tok.Column}
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch), Line: tok.Line, Column: tok.Column}
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch)
		}
	case '(':
		tok = l.newToken(token.LPAREN, l.ch)
	case ')':
		tok = l.newToken(token.RPAREN, l.ch)
	case ':':
		tok = l.newToken(token.COLON, l.ch)
	case ',':
		tok = l.newToken(token.COMMA, l.ch)
	case '#':
		l.skipComment()
		return l.NextToken() // Recursively get next token after comment
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // Early return - readIdentifier already advanced
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok // Early return - readNumber already advanced
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readChar advances the lexer position and reads the next character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII NUL - signals EOF
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++

	// Track newlines for line counting
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
}

// peekChar looks ahead at the next character without advancing position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// readIdentifier reads an identifier or keyword (letters and underscores)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads an integer literal
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace skips over whitespace characters (space, tab, newline, carriage return)
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment skips from '#' to the end of the line
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// newToken creates a new token with the current line/column position
func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    l.line,
		Column:  l.column,
	}
}

// isLetter checks if a character is a letter or underscore
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if a character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
