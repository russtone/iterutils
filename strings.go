package iterutils

// strings iterates over slice of strings.
type strings struct {
	items []string
	idx   int
}

// Strings create strings iterator.
func Strings(items ...string) Iterator {
	return &strings{items, 0}
}

// Next allows strings to implement Iterator iterface.
func (it *strings) Next(dest *string) bool {
	if it.idx < len(it.items) {
		*dest = it.items[it.idx]
		it.idx++
		return true
	}

	return false
}

// Reset allows strings to implement Iterator iterface.
func (it *strings) Reset() {
	it.idx = 0
}

// Count allows strings to implement Iterator iterface.
func (it *strings) Count() uint64 {
	return uint64(len(it.items))
}

// Close allows strings to implement Iterator iterface.
func (it *strings) Close() error {
	return nil
}
