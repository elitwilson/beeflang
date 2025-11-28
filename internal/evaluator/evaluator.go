package evaluator

import (
	"github.com/elitwilson/beeflang/internal/ast"
	"github.com/elitwilson/beeflang/internal/object"
)

// Eval evaluates an AST node and returns the resulting runtime object.
// This is the core of the interpreter - it walks the AST and executes the code.
func Eval(node ast.Node) object.Object {
	switch n := node.(type) {

	// Program: evaluate all statements and return the last result
	case *ast.Program:
		return evalProgram(n)

	// Literals: convert AST literals to runtime objects
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}

	case *ast.BooleanLiteral:
		return nativeBoolToBooleanObject(n.Value)

	case *ast.StringLiteral:
		return &object.String{Value: n.Value}

	// Expressions: evaluate recursively
	case *ast.PrefixExpression:
		right := Eval(n.Right)
		return evalPrefixExpression(n.Operator, right)

	case *ast.InfixExpression:
		left := Eval(n.Left)
		right := Eval(n.Right)
		return evalInfixExpression(n.Operator, left, right)

	// Expression statement: evaluate the expression
	case *ast.ExpressionStatement:
		return Eval(n.Expression)
	}

	return nil
}

// evalProgram evaluates all statements in a program and returns the last result
func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)
	}

	return result
}

// evalPrefixExpression evaluates prefix expressions like -5 or !true
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)
	default:
		return object.NULL
	}
}

// evalBangOperator implements the ! (not) operator
func evalBangOperator(right object.Object) object.Object {
	switch right {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	case object.NULL:
		return object.TRUE
	default:
		return object.FALSE
	}
}

// evalMinusPrefixOperator implements the - (negation) operator
func evalMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != "INTEGER" {
		return object.NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

// evalInfixExpression evaluates infix expressions like 5 + 3 or 10 > 5
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	// Integer operations
	case left.Type() == "INTEGER" && right.Type() == "INTEGER":
		return evalIntegerInfixExpression(operator, left, right)

	// String concatenation
	case left.Type() == "STRING" && right.Type() == "STRING":
		return evalStringInfixExpression(operator, left, right)

	// Boolean comparison (using pointer equality optimization)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)

	default:
		return object.NULL
	}
}

// evalIntegerInfixExpression handles arithmetic and comparison on integers
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	// Arithmetic
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}

	// Comparison
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)

	default:
		return object.NULL
	}
}

// evalStringInfixExpression handles string operations
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return object.NULL
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

// nativeBoolToBooleanObject converts a Go bool to a Boolean object
// Uses singleton TRUE/FALSE for efficiency
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return object.TRUE
	}
	return object.FALSE
}
