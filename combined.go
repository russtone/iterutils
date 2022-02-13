package iterutils

import "github.com/hashicorp/go-multierror"

// combined represents multiple iterators.
type combined struct {
	items []Iterator
	idx   int
}

// Combine combines multiple iterators.
func Combine(iterators ...Iterator) Iterator {
	return &combined{iterators, 0}
}

// Reset allows combined to implement Iterator iterface.
func (it *combined) Reset() {
	for _, item := range it.items {
		item.Reset()
	}
}

// Next allows combined to implement Iterator iterface.
func (it *combined) Next(dest *string) bool {
	for i := it.idx; i < len(it.items); i++ {
		if it.items[i].Next(dest) {
			return true
		}
	}

	return false
}

// Count allows combined to implement Iterator iterface.
func (it *combined) Count() uint64 {
	count := uint64(0)
	for _, item := range it.items {
		count += item.Count()
	}
	return count
}

// Close allows combined to implement Iterator iterface.
func (it *combined) Close() error {
	var errs error

	for _, item := range it.items {
		if err := item.Close(); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	return errs
}
