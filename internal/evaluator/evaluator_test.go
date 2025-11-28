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
	env := NewEnvironment()
	return Eval(program, env)
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

// Phase 1.5: Real failing tests for variables

func TestEvalVariableDeclaration(t *testing.T) {
	input := `
cut x = 42
cut y = x + 8
y
`
	result := testEval(input)
	assert.NotNil(t, result)

	integer, ok := result.(*object.Integer)
	assert.True(t, ok, "Result should be an Integer")
	assert.Equal(t, int64(50), integer.Value)
}

func TestEvalIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"cut a = 5; a", 5},
		{"cut a = 5 * 5; a", 25},
		{"cut a = 5; cut b = a; b", 5},
		{"cut a = 5; cut b = a; cut c = a + b + 5; c", 15},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Input: %s", tt.input)

		integer, ok := result.(*object.Integer)
		assert.True(t, ok, "Result should be an Integer for input: %s", tt.input)
		assert.Equal(t, tt.expected, integer.Value, "Input: %s", tt.input)
	}
}

// Phase 2: Control Flow - Real failing tests

func TestEvalBlockStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		// Block should evaluate all statements and return the last value
		{`
cut x = 5
cut y = 10
x + y
`, 15},
		// Single statement in a block
		{`
42
`, 42},
		// Multiple statements, last one is the result
		{`
cut a = 1
cut b = 2
cut c = 3
c
`, 3},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		assert.NotNil(t, result, "Input: %s", tt.input)

		integer, ok := result.(*object.Integer)
		assert.True(t, ok, "Result should be an Integer for input: %s", tt.input)
		assert.Equal(t, tt.expected, integer.Value, "Input: %s", tt.input)
	}
}

func TestEvalIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{} // can be int64, bool, or nil (for NULL)
	}{
		// Basic if with true condition
		{"if true: 10 beef", int64(10)},
		// Basic if with false condition (should return NULL)
		{"if false: 10 beef", nil},
		// If with expression condition
		{"if 1 < 2: 10 beef", int64(10)},
		{"if 1 > 2: 10 beef", nil},
		// If-else with true condition
		{"if 1 < 2: 10 else: 20 beef", int64(10)},
		// If-else with false condition
		{"if 1 > 2: 10 else: 20 beef", int64(20)},
		// Boolean result
		{"if true: true else: false beef", true},
		// Nested expressions in consequence
		{"if true: 5 + 5 beef", int64(10)},
		// Nested expressions in alternative
		{"if false: 5 else: 10 + 10 beef", int64(20)},
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
		case nil:
			null, ok := result.(*object.Null)
			assert.True(t, ok, "Result should be NULL for input: %s", tt.input)
			assert.NotNil(t, null, "NULL object should not be nil pointer")
		}
	}
}
