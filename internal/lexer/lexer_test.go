package lexer

import (
	"testing"

	"github.com/elitwilson/beeflang/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestTokenizeIntegers(t *testing.T) {
	input := "42"
	l := New(input)

	tok := l.NextToken()
	assert.Equal(t, token.INT, tok.Type)
	assert.Equal(t, "42", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.EOF, tok.Type)
}

func TestTokenizeIdentifiers(t *testing.T) {
	input := "foo bar"
	l := New(input)

	tok := l.NextToken()
	assert.Equal(t, token.IDENT, tok.Type)
	assert.Equal(t, "foo", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.IDENT, tok.Type)
	assert.Equal(t, "bar", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.EOF, tok.Type)
}

func TestTokenizeKeywords(t *testing.T) {
	input := "cut praise beef genesis serve if else"
	l := New(input)

	expectedTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CUT, "cut"},
		{token.PRAISE, "praise"},
		{token.BEEF, "beef"},
		{token.GENESIS, "genesis"},
		{token.SERVE, "serve"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.EOF, ""},
	}

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		assert.Equal(t, expected.expectedType, tok.Type, "token %d type mismatch", i)
		assert.Equal(t, expected.expectedLiteral, tok.Literal, "token %d literal mismatch", i)
	}
}

func TestTokenizeSingleCharOperators(t *testing.T) {
	input := "+ - * / = < >"
	l := New(input)

	expectedTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.ASSIGN, "="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EOF, ""},
	}

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		assert.Equal(t, expected.expectedType, tok.Type, "token %d type mismatch", i)
		assert.Equal(t, expected.expectedLiteral, tok.Literal, "token %d literal mismatch", i)
	}
}

func TestTokenizeTwoCharOperators(t *testing.T) {
	input := "== != <= >= && ||"
	l := New(input)

	expectedTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.EOF, ""},
	}

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		assert.Equal(t, expected.expectedType, tok.Type, "token %d type mismatch", i)
		assert.Equal(t, expected.expectedLiteral, tok.Literal, "token %d literal mismatch", i)
	}
}

func TestTokenizeDelimiters(t *testing.T) {
	input := "( ) : ,"
	l := New(input)

	expectedTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		assert.Equal(t, expected.expectedType, tok.Type, "token %d type mismatch", i)
		assert.Equal(t, expected.expectedLiteral, tok.Literal, "token %d literal mismatch", i)
	}
}

func TestSkipWhitespace(t *testing.T) {
	input := "   42   "
	l := New(input)

	tok := l.NextToken()
	assert.Equal(t, token.INT, tok.Type)
	assert.Equal(t, "42", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.EOF, tok.Type)
}

func TestSkipComments(t *testing.T) {
	input := `# This is a comment
42  # inline comment`
	l := New(input)

	tok := l.NextToken()
	assert.Equal(t, token.INT, tok.Type)
	assert.Equal(t, "42", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.EOF, tok.Type)
}

func TestHandleFeastWhileKeyword(t *testing.T) {
	input := "feast while"
	l := New(input)

	// For now, we'll treat "feast while" as two separate tokens
	// This test documents the behavior - we can refine later if needed
	tok := l.NextToken()
	assert.Equal(t, token.FEAST_WHILE, tok.Type)
	assert.Equal(t, "feast", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.FEAST_WHILE, tok.Type)
	assert.Equal(t, "while", tok.Literal)

	tok = l.NextToken()
	assert.Equal(t, token.EOF, tok.Type)
}

func TestEOFToken(t *testing.T) {
	input := ""
	l := New(input)

	tok := l.NextToken()
	assert.Equal(t, token.EOF, tok.Type)
}

func TestSimpleVariableDeclaration(t *testing.T) {
	// Integration test: complete statement
	input := "cut x = 42"
	l := New(input)

	expectedTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CUT, "cut"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "42"},
		{token.EOF, ""},
	}

	for i, expected := range expectedTokens {
		tok := l.NextToken()
		assert.Equal(t, expected.expectedType, tok.Type, "token %d type mismatch", i)
		assert.Equal(t, expected.expectedLiteral, tok.Literal, "token %d literal mismatch", i)
	}
}

func TestTrackLineAndColumn(t *testing.T) {
	// Critical for error messages
	input := `cut x = 42
cut y = 5`
	l := New(input)

	// First line tokens
	l.NextToken()        // cut
	l.NextToken()        // x
	l.NextToken()        // =
	tok := l.NextToken() // 42
	assert.Equal(t, 1, tok.Line, "first line should be 1")

	// Second line tokens
	tok = l.NextToken() // cut on line 2
	assert.Equal(t, 2, tok.Line, "second line should be 2")
}
