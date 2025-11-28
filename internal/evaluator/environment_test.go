package evaluator

import (
	"testing"

	"github.com/elitwilson/beeflang/internal/object"
	"github.com/stretchr/testify/assert"
)

// Phase 1.2: Real failing Environment tests

func TestEnvironmentSet(t *testing.T) {
	env := NewEnvironment()

	// Set a variable
	val := &object.Integer{Value: 42}
	env.Set("x", val)

	// Verify it was stored
	result, ok := env.Get("x")
	assert.True(t, ok, "Variable 'x' should exist")
	assert.Equal(t, val, result, "Should get back the same object")
}

func TestEnvironmentGet(t *testing.T) {
	env := NewEnvironment()

	// Getting non-existent variable should return false
	result, ok := env.Get("nonexistent")
	assert.False(t, ok, "Non-existent variable should return false")
	assert.Nil(t, result, "Non-existent variable should return nil")

	// Set and get
	env.Set("y", &object.Integer{Value: 100})
	result, ok = env.Get("y")
	assert.True(t, ok, "Variable 'y' should exist")

	integer, isInt := result.(*object.Integer)
	assert.True(t, isInt, "Should be an Integer object")
	assert.Equal(t, int64(100), integer.Value)
}

func TestEnvironmentNestedScope(t *testing.T) {
	// Outer scope
	outer := NewEnvironment()
	outer.Set("x", &object.Integer{Value: 10})

	// Inner scope (nested)
	inner := NewEnclosedEnvironment(outer)
	inner.Set("y", &object.Integer{Value: 20})

	// Inner scope can see outer variables
	x, ok := inner.Get("x")
	assert.True(t, ok, "Inner scope should see outer variable 'x'")
	assert.Equal(t, int64(10), x.(*object.Integer).Value)

	// Inner scope can see its own variables
	y, ok := inner.Get("y")
	assert.True(t, ok, "Inner scope should see its own variable 'y'")
	assert.Equal(t, int64(20), y.(*object.Integer).Value)

	// Outer scope CANNOT see inner variables
	_, ok = outer.Get("y")
	assert.False(t, ok, "Outer scope should NOT see inner variable 'y'")
}

func TestEnvironmentShadowing(t *testing.T) {
	// Outer scope
	outer := NewEnvironment()
	outer.Set("x", &object.Integer{Value: 10})

	// Inner scope shadows 'x'
	inner := NewEnclosedEnvironment(outer)
	inner.Set("x", &object.Integer{Value: 999})

	// Inner scope sees shadowed value
	x, ok := inner.Get("x")
	assert.True(t, ok)
	assert.Equal(t, int64(999), x.(*object.Integer).Value, "Inner scope should see shadowed value")

	// Outer scope still has original value
	x, ok = outer.Get("x")
	assert.True(t, ok)
	assert.Equal(t, int64(10), x.(*object.Integer).Value, "Outer scope should still have original value")
}
