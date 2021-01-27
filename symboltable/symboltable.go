package symboltable

// SymbolTable - wraps a symbol(string):value(int) map
type SymbolTable struct {
	Map map[string]int
}

// NewSymbolTable - returns a pointer to new symbol table, populated with the 23 pre-defined symbols
func NewSymbolTable() *SymbolTable {
	m := map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    24576,
	}
	return &SymbolTable{Map: m}
}

// Put - adds an entry to the table
func (st SymbolTable) Put(symbol string, value int) {
	st.Map[symbol] = value
}

// Contains - is this symbol in the table?
func (st SymbolTable) Contains(symbol string) bool {
	_, ok := st.Map[symbol]
	return ok
}

// Get - returns the value for the specified symbol
func (st SymbolTable) Get(symbol string) int {
	return st.Map[symbol]
}
