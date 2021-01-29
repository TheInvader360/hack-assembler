package symbols

// Symbols - interface
type Symbols interface {
	Seed()
	Put(symbol string, value int)
	Contains(symbol string) bool
	Get(symbol string) int
}
