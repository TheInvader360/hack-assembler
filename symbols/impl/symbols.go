package impl

// Symbols - wraps a symbol(string):value(int) map (implements symbols.Symbols interface)
type Symbols struct {
	Map map[string]int
}

// NewSymbols - returns a pointer to new symbols
func NewSymbols() *Symbols {
	return &Symbols{map[string]int{}}
}

// Seed - seeds the symbols map with the 23 pre-defined symbols
func (s *Symbols) Seed() {
	s.Put("SP", 0)
	s.Put("LCL", 1)
	s.Put("ARG", 2)
	s.Put("THIS", 3)
	s.Put("THAT", 4)
	s.Put("R0", 0)
	s.Put("R1", 1)
	s.Put("R2", 2)
	s.Put("R3", 3)
	s.Put("R4", 4)
	s.Put("R5", 5)
	s.Put("R6", 6)
	s.Put("R7", 7)
	s.Put("R8", 8)
	s.Put("R9", 9)
	s.Put("R10", 10)
	s.Put("R11", 11)
	s.Put("R12", 12)
	s.Put("R13", 13)
	s.Put("R14", 14)
	s.Put("R15", 15)
	s.Put("SCREEN", 16384)
	s.Put("KBD", 24576)
}

// Put - adds an entry to the map
func (s *Symbols) Put(symbol string, value int) {
	s.Map[symbol] = value
}

// Contains - is this symbol in the map?
func (s *Symbols) Contains(symbol string) bool {
	_, ok := s.Map[symbol]
	return ok
}

// Get - returns the value for the specified symbol
func (s *Symbols) Get(symbol string) int {
	return s.Map[symbol]
}
