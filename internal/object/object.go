package object

import "fmt"

// Object represents a runtime value in the Beeflang interpreter.
//
// NOTE: "Object" here refers to runtime values (integers, booleans, strings, etc.),
// NOT Object-Oriented Programming. This is standard interpreter terminology - even
// non-OOP languages represent runtime values as "objects."
//
// Every value that exists during program execution implements this interface.
type Object interface {
	Type() string   // Returns the type of the object (e.g., "INTEGER", "BOOLEAN")
	Inspect() string // Returns a string representation for debugging/printing
}

// Integer represents an integer value at runtime.
type Integer struct {
	Value int64
}

func (i *Integer) Type() string {
	return "INTEGER"
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Boolean represents a boolean value at runtime.
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string {
	return "BOOLEAN"
}

func (b *Boolean) Inspect() string {
	if b.Value {
		return "true"
	}
	return "false"
}

// String represents a string value at runtime.
type String struct {
	Value string
}

func (s *String) Type() string {
	return "STRING"
}

func (s *String) Inspect() string {
	return s.Value
}

// Null represents the absence of a value.
// Used for functions that don't return anything, uninitialized variables, etc.
type Null struct{}

func (n *Null) Type() string {
	return "NULL"
}

func (n *Null) Inspect() string {
	return "null"
}

// Singleton instances used throughout the interpreter for efficiency.
// Instead of creating new objects, we reuse these single instances.
var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)
