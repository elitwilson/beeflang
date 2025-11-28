package evaluator

import "github.com/elitwilson/beeflang/internal/object"

// Environment stores variable bindings (name -> value mappings).
// It supports nested scopes through the `outer` pointer, enabling block-level scoping.
//
// Example:
//   outer := NewEnvironment()
//   outer.Set("x", &object.Integer{Value: 10})
//
//   inner := NewEnclosedEnvironment(outer)
//   inner.Set("y", &object.Integer{Value: 20})
//   inner.Get("x")  // finds x in outer scope
//   inner.Get("y")  // finds y in inner scope
type Environment struct {
	store map[string]object.Object
	outer *Environment // pointer to enclosing (parent) scope
}

// NewEnvironment creates a new environment with no outer scope (global scope).
func NewEnvironment() *Environment {
	s := make(map[string]object.Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new environment enclosed by an outer environment.
// This is used for creating nested scopes (functions, blocks, etc.).
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get retrieves a variable from the environment.
// It searches the current scope first, then walks up the outer scopes.
// Returns (value, true) if found, (nil, false) if not found.
func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		// Not found in current scope, check outer scope
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set stores a variable in the current environment scope.
// This does NOT modify outer scopes - it creates/updates in the current scope only.
func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}
