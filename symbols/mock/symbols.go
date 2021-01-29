package mock

import (
	"github.com/stretchr/testify/mock"
)

// Symbols - Mock struct (implements symbols.Symbols interface)
type Symbols struct {
	mock.Mock
}

// Seed - Mock
func (s *Symbols) Seed() {
	s.Called()
}

// Put - Mock
func (s *Symbols) Put(symbol string, value int) {
	s.Called(symbol, value)
}

// Contains - Mock
func (s *Symbols) Contains(symbol string) bool {
	args := s.Called(symbol)
	return args.Bool(0)
}

// Get - Mock
func (s *Symbols) Get(symbol string) int {
	args := s.Called(symbol)
	return args.Int(0)
}
