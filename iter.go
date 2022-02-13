package iterutils

// Iterator is common iterator interface.
type Iterator interface {
	// Next stores next value to provided pointer.
	// Returns false if there is no more values.
	Next(*string) bool

	// Resets resets Iterator state.
	Reset()

	// Count returns number of values in Iterator.
	Count() uint64

	// Close frees iternal iterator resources.
	Close() error
}
