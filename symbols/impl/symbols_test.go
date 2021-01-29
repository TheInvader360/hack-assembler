package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeed(t *testing.T) {
	s := NewSymbols()
	s.Seed()
	assert.Equal(t, 23, len(s.Map))
	assert.Equal(t, 0, s.Map["SP"])
	assert.Equal(t, 15, s.Map["R15"])
	assert.Equal(t, 24576, s.Map["KBD"])
}

func TestPut(t *testing.T) {
	s := Symbols{Map: map[string]int{"a": 0}}
	s.Put("z", 100)
	assert.Equal(t, 2, len(s.Map))
	assert.Equal(t, 100, s.Map["z"])
}

func TestContains(t *testing.T) {
	s := Symbols{Map: map[string]int{"y": 0}}
	assert.True(t, s.Contains("y"))
	assert.False(t, s.Contains("z"))
}

func TestGet(t *testing.T) {
	s := Symbols{Map: map[string]int{"a": 0, "b": 1, "c": 2}}
	assert.Equal(t, 1, s.Get("b"))
}
