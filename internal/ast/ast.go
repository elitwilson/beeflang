package ast

import "github.com/elitwilson/beeflang/internal/token"

// Node is the base interface for all AST nodes
type Node interface {
	TokenLiteral() string
}

// Statement represents a statement node in the AST
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node in the AST
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// IntegerLiteral represents an integer literal like 42
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// BooleanLiteral represents a boolean literal like true or false
type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }

// StringLiteral represents a string literal like "Hello, Beef!"
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// Identifier represents a variable or function name
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// PrefixExpression represents prefix operators like -5 or !true
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// InfixExpression represents binary operators like 5 + 3
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// VariableDeclaration represents: prep x = 42
type VariableDeclaration struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (vd *VariableDeclaration) statementNode()       {}
func (vd *VariableDeclaration) TokenLiteral() string { return vd.Token.Literal }

// AssignmentStatement represents: x = 42 (reassignment, no prep keyword)
type AssignmentStatement struct {
	Token token.Token // The identifier token
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }

// ReturnStatement represents: serve x
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// IfStatement represents: if condition: consequence beef else alternative beef
type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }

// WhileLoop represents: feast while condition: body beef
type WhileLoop struct {
	Token     token.Token // The 'feast' or 'while' token
	Condition Expression
	Body      *BlockStatement
}

func (wl *WhileLoop) statementNode()       {}
func (wl *WhileLoop) TokenLiteral() string { return wl.Token.Literal }

// FunctionDeclaration represents: praise name(params): body beef
type FunctionDeclaration struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fd *FunctionDeclaration) statementNode()       {}
func (fd *FunctionDeclaration) TokenLiteral() string { return fd.Token.Literal }

// FunctionCall represents: preach(42)
type FunctionCall struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (fc *FunctionCall) expressionNode()      {}
func (fc *FunctionCall) TokenLiteral() string { return fc.Token.Literal }

// BlockStatement represents a block of statements
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// ExpressionStatement wraps an expression so it can be used as a statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
