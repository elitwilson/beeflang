package ast

import (
	"testing"

	"github.com/elitwilson/beeflang/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestIntegerLiteralNode(t *testing.T) {
	tok := token.Token{Type: token.INT, Literal: "42", Line: 1, Column: 1}
	intLiteral := &IntegerLiteral{
		Token: tok,
		Value: 42,
	}

	assert.Equal(t, "42", intLiteral.TokenLiteral())
	assert.Equal(t, int64(42), intLiteral.Value)

	// Verify it implements Expression interface
	var _ Expression = intLiteral
}

func TestIdentifierNode(t *testing.T) {
	tok := token.Token{Type: token.IDENT, Literal: "foo", Line: 1, Column: 1}
	ident := &Identifier{
		Token: tok,
		Value: "foo",
	}

	assert.Equal(t, "foo", ident.TokenLiteral())
	assert.Equal(t, "foo", ident.Value)

	// Verify it implements Expression interface
	var _ Expression = ident
}

func TestPrefixExpressionNode(t *testing.T) {
	// -5
	tok := token.Token{Type: token.MINUS, Literal: "-", Line: 1, Column: 1}
	prefix := &PrefixExpression{
		Token:    tok,
		Operator: "-",
		Right: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "5"},
			Value: 5,
		},
	}

	assert.Equal(t, "-", prefix.TokenLiteral())
	assert.Equal(t, "-", prefix.Operator)
	assert.NotNil(t, prefix.Right)

	// Verify it implements Expression interface
	var _ Expression = prefix
}

func TestInfixExpressionNode(t *testing.T) {
	// 5 + 3
	tok := token.Token{Type: token.PLUS, Literal: "+", Line: 1, Column: 1}
	infix := &InfixExpression{
		Token:    tok,
		Left:     &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
		Operator: "+",
		Right:    &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "3"}, Value: 3},
	}

	assert.Equal(t, "+", infix.TokenLiteral())
	assert.Equal(t, "+", infix.Operator)
	assert.NotNil(t, infix.Left)
	assert.NotNil(t, infix.Right)

	// Verify it implements Expression interface
	var _ Expression = infix
}

func TestVariableDeclarationNode(t *testing.T) {
	// cut x = 42
	tok := token.Token{Type: token.CUT, Literal: "cut", Line: 1, Column: 1}
	varDecl := &VariableDeclaration{
		Token: tok,
		Name: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "x"},
			Value: "x",
		},
		Value: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "42"},
			Value: 42,
		},
	}

	assert.Equal(t, "cut", varDecl.TokenLiteral())
	assert.Equal(t, "x", varDecl.Name.Value)
	assert.NotNil(t, varDecl.Value)

	// Verify it implements Statement interface
	var _ Statement = varDecl
}

func TestReturnStatementNode(t *testing.T) {
	// serve 42
	tok := token.Token{Type: token.SERVE, Literal: "serve", Line: 1, Column: 1}
	returnStmt := &ReturnStatement{
		Token: tok,
		ReturnValue: &IntegerLiteral{
			Token: token.Token{Type: token.INT, Literal: "42"},
			Value: 42,
		},
	}

	assert.Equal(t, "serve", returnStmt.TokenLiteral())
	assert.NotNil(t, returnStmt.ReturnValue)

	// Verify it implements Statement interface
	var _ Statement = returnStmt
}

func TestIfStatementNode(t *testing.T) {
	// if x > 0: ... beef
	tok := token.Token{Type: token.IF, Literal: "if", Line: 1, Column: 1}
	ifStmt := &IfStatement{
		Token: tok,
		Condition: &InfixExpression{
			Token:    token.Token{Type: token.GT, Literal: ">"},
			Left:     &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
			Operator: ">",
			Right:    &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "0"}, Value: 0},
		},
		Consequence: &BlockStatement{
			Statements: []Statement{},
		},
		Alternative: nil, // optional else block
	}

	assert.Equal(t, "if", ifStmt.TokenLiteral())
	assert.NotNil(t, ifStmt.Condition)
	assert.NotNil(t, ifStmt.Consequence)

	// Verify it implements Statement interface
	var _ Statement = ifStmt
}

func TestFunctionDeclarationNode(t *testing.T) {
	// praise add(x, y): ... beef
	tok := token.Token{Type: token.PRAISE, Literal: "praise", Line: 1, Column: 1}
	fnDecl := &FunctionDeclaration{
		Token: tok,
		Name: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "add"},
			Value: "add",
		},
		Parameters: []*Identifier{
			{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
			{Token: token.Token{Type: token.IDENT, Literal: "y"}, Value: "y"},
		},
		Body: &BlockStatement{
			Statements: []Statement{},
		},
	}

	assert.Equal(t, "praise", fnDecl.TokenLiteral())
	assert.Equal(t, "add", fnDecl.Name.Value)
	assert.Len(t, fnDecl.Parameters, 2)
	assert.NotNil(t, fnDecl.Body)

	// Verify it implements Statement interface
	var _ Statement = fnDecl
}

func TestFunctionCallNode(t *testing.T) {
	// preach(42)
	tok := token.Token{Type: token.IDENT, Literal: "preach", Line: 1, Column: 1}
	fnCall := &FunctionCall{
		Token: tok,
		Function: &Identifier{
			Token: token.Token{Type: token.IDENT, Literal: "preach"},
			Value: "preach",
		},
		Arguments: []Expression{
			&IntegerLiteral{
				Token: token.Token{Type: token.INT, Literal: "42"},
				Value: 42,
			},
		},
	}

	assert.Equal(t, "preach", fnCall.TokenLiteral())
	// Check that Function is an Identifier
	ident, ok := fnCall.Function.(*Identifier)
	assert.True(t, ok, "Function should be an Identifier")
	assert.Equal(t, "preach", ident.Value)
	assert.Len(t, fnCall.Arguments, 1)

	// Verify it implements Expression interface
	var _ Expression = fnCall
}

func TestBlockStatementNode(t *testing.T) {
	// Block with multiple statements
	block := &BlockStatement{
		Token: token.Token{Type: token.COLON, Literal: ":", Line: 1, Column: 1},
		Statements: []Statement{
			&VariableDeclaration{
				Token: token.Token{Type: token.CUT, Literal: "cut"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
				Value: &IntegerLiteral{Token: token.Token{Type: token.INT, Literal: "5"}, Value: 5},
			},
			&ReturnStatement{
				Token:       token.Token{Type: token.SERVE, Literal: "serve"},
				ReturnValue: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "x"}, Value: "x"},
			},
		},
	}

	assert.Equal(t, ":", block.TokenLiteral())
	assert.Len(t, block.Statements, 2)

	// Verify it implements Statement interface
	var _ Statement = block
}
