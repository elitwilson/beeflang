package evaluator

import (
	"testing"

	"github.com/elitwilson/beeflang/internal/lexer"
	"github.com/elitwilson/beeflang/internal/object"
	"github.com/elitwilson/beeflang/internal/parser"
	"github.com/stretchr/testify/assert"
)

// Phase 2: Real failing tests - basic expressions

// Helper function to parse and evaluate Beeflang source code
func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func TestEvalIntegerLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"42", 42},
		{"0", 0},
		{"999", 999},
		{"-5", -5},
		{"-100", -100},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Eval should return a value")

		integer, ok := result.(*object.Integer)
		assert.True(t, ok, "Result should be an Integer object")
		assert.Equal(t, tt.expected, integer.Value)
	}
}

func TestEvalBooleanLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Eval should return a value")

		boolean, ok := result.(*object.Boolean)
		assert.True(t, ok, "Result should be a Boolean object")
		assert.Equal(t, tt.expected, boolean.Value)
	}
}

func TestEvalStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"Hello, Beef!"`, "Hello, Beef!"},
		{`"Praise the Beef!"`, "Praise the Beef!"},
		{`""`, ""},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Eval should return a value")

		str, ok := result.(*object.String)
		assert.True(t, ok, "Result should be a String object")
		assert.Equal(t, tt.expected, str.Value)
	}
}

func TestEvalPrefixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Integer negation
		{"-5", int64(-5)},
		{"-42", int64(-42)},
		{"--10", int64(10)},

		// Boolean negation
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!false", false},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Eval should return a value for input: %s", tt.input)

		switch expected := tt.expected.(type) {
		case int64:
			integer, ok := result.(*object.Integer)
			assert.True(t, ok, "Result should be an Integer for input: %s", tt.input)
			assert.Equal(t, expected, integer.Value, "Input: %s", tt.input)
		case bool:
			boolean, ok := result.(*object.Boolean)
			assert.True(t, ok, "Result should be a Boolean for input: %s", tt.input)
			assert.Equal(t, expected, boolean.Value, "Input: %s", tt.input)
		}
	}
}

func TestEvalInfixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Integer arithmetic
		{"5 + 5", int64(10)},
		{"10 - 5", int64(5)},
		{"2 * 3", int64(6)},
		{"10 / 2", int64(5)},
		{"10 % 3", int64(1)},

		// Integer comparisons
		{"5 < 10", true},
		{"10 > 5", true},
		{"5 == 5", true},
		{"5 != 10", true},
		{"5 <= 5", true},
		{"10 >= 5", true},

		// Boolean operations
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},

		// String concatenation
		{`"Hello" + " " + "Beef"`, "Hello Beef"},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Eval should return a value for input: %s", tt.input)

		switch expected := tt.expected.(type) {
		case int64:
			integer, ok := result.(*object.Integer)
			assert.True(t, ok, "Result should be an Integer for input: %s", tt.input)
			assert.Equal(t, expected, integer.Value, "Input: %s", tt.input)
		case bool:
			boolean, ok := result.(*object.Boolean)
			assert.True(t, ok, "Result should be a Boolean for input: %s", tt.input)
			assert.Equal(t, expected, boolean.Value, "Input: %s", tt.input)
		case string:
			str, ok := result.(*object.String)
			assert.True(t, ok, "Result should be a String for input: %s", tt.input)
			assert.Equal(t, expected, str.Value, "Input: %s", tt.input)
		}
	}
}
