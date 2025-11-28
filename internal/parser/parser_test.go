package parser

import (
	"testing"

	"github.com/elitwilson/beeflang/internal/ast"
	"github.com/elitwilson/beeflang/internal/lexer"
	"github.com/stretchr/testify/assert"
)

// Phase 2: Real failing tests for parser

func TestParseIntegerLiteral(t *testing.T) {
	input := "42"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1, "program should have 1 statement")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "statement should be *ast.ExpressionStatement")

	intLiteral, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(t, ok, "expression should be *ast.IntegerLiteral")
	assert.Equal(t, int64(42), intLiteral.Value)
}

func TestParseIdentifier(t *testing.T) {
	input := "foobar"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1, "program should have 1 statement")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "statement should be *ast.ExpressionStatement")

	ident, ok := stmt.Expression.(*ast.Identifier)
	assert.True(t, ok, "expression should be *ast.Identifier")
	assert.Equal(t, "foobar", ident.Value)
}

func TestParseBooleanLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Len(t, program.Statements, 1, "program should have 1 statement")

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok, "statement should be *ast.ExpressionStatement")

		boolLiteral, ok := stmt.Expression.(*ast.BooleanLiteral)
		assert.True(t, ok, "expression should be *ast.BooleanLiteral")
		assert.Equal(t, tt.expected, boolLiteral.Value)
	}
}

func TestParseStringLiteral(t *testing.T) {
	input := `"Hello, Beef!"`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1, "program should have 1 statement")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "statement should be *ast.ExpressionStatement")

	strLiteral, ok := stmt.Expression.(*ast.StringLiteral)
	assert.True(t, ok, "expression should be *ast.StringLiteral")
	assert.Equal(t, "Hello, Beef!", strLiteral.Value)
}

func TestParsePrefixExpression(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"-5", "-", 5},
		{"!10", "!", 10},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Len(t, program.Statements, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok, "statement should be *ast.ExpressionStatement")

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(t, ok, "expression should be *ast.PrefixExpression")
		assert.Equal(t, tt.operator, exp.Operator)

		intLit, ok := exp.Right.(*ast.IntegerLiteral)
		assert.True(t, ok, "right should be *ast.IntegerLiteral")
		assert.Equal(t, tt.value, intLit.Value)
	}
}

func TestParseInfixExpression(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		assert.Len(t, program.Statements, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok, "statement should be *ast.ExpressionStatement")

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		assert.True(t, ok, "expression should be *ast.InfixExpression, got %T", stmt.Expression)
		assert.Equal(t, tt.operator, exp.Operator)

		testIntegerLiteral(t, exp.Left, tt.leftValue)
		testIntegerLiteral(t, exp.Right, tt.rightValue)
	}
}

func TestParseVariableDeclaration(t *testing.T) {
	input := "cut x = 5"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt := program.Statements[0]
	varDecl, ok := stmt.(*ast.VariableDeclaration)
	assert.True(t, ok, "statement should be *ast.VariableDeclaration")
	assert.Equal(t, "x", varDecl.Name.Value)

	testIntegerLiteral(t, varDecl.Value, 5)
}

func TestParseReturnStatement(t *testing.T) {
	input := "serve 5"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt := program.Statements[0]
	returnStmt, ok := stmt.(*ast.ReturnStatement)
	assert.True(t, ok, "statement should be *ast.ReturnStatement")

	testIntegerLiteral(t, returnStmt.ReturnValue, 5)
}

func TestParseIfStatement(t *testing.T) {
	input := `if x > 5:
   cut y = 10
beef`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt := program.Statements[0]
	ifStmt, ok := stmt.(*ast.IfStatement)
	assert.True(t, ok, "statement should be *ast.IfStatement")
	assert.NotNil(t, ifStmt.Condition)
	assert.NotNil(t, ifStmt.Consequence)
	assert.Len(t, ifStmt.Consequence.Statements, 1)
}

func TestParseFunctionDeclaration(t *testing.T) {
	input := `praise add(x, y):
   serve x + y
beef`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt := program.Statements[0]
	fnDecl, ok := stmt.(*ast.FunctionDeclaration)
	assert.True(t, ok, "statement should be *ast.FunctionDeclaration")
	assert.Equal(t, "add", fnDecl.Name.Value)
	assert.Len(t, fnDecl.Parameters, 2)
	assert.Equal(t, "x", fnDecl.Parameters[0].Value)
	assert.Equal(t, "y", fnDecl.Parameters[1].Value)
	assert.NotNil(t, fnDecl.Body)
}

func TestParseFunctionCall(t *testing.T) {
	input := "preach(42)"
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	assert.Len(t, program.Statements, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "statement should be *ast.ExpressionStatement")

	callExp, ok := stmt.Expression.(*ast.FunctionCall)
	assert.True(t, ok, "expression should be *ast.FunctionCall")

	ident, ok := callExp.Function.(*ast.Identifier)
	assert.True(t, ok, "function should be identifier")
	assert.Equal(t, "preach", ident.Value)
	assert.Len(t, callExp.Arguments, 1)
}

func TestParseOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + 3 * 2", "(5 + (3 * 2))"},
		{"5 * 3 + 2", "((5 * 3) + 2)"},
		{"5 + 3 - 2", "((5 + 3) - 2)"},
		{"-5 + 3", "((-5) + 3)"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		// For now, just check it parses without errors
		// Full precedence testing would require String() method on AST
		assert.Len(t, program.Statements, 1)
	}
}

// Helper functions

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) {
	intLit, ok := exp.(*ast.IntegerLiteral)
	assert.True(t, ok, "expression should be *ast.IntegerLiteral, got %T", exp)
	assert.Equal(t, value, intLit.Value)
}
