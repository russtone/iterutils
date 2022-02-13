package iterutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/russtone/iterutils"
)

func Test_Combined(t *testing.T) {
	tests := []struct {
		slice1 []string
		slice2 []string
	}{
		{
			[]string{"one", "two", "three"},
			[]string{"1", "2", "3"},
		},
		{
			[]string{"1"},
			[]string{"2"},
		},
	}

	for _, vec := range tests {
		it := iterutils.Combine(
			iterutils.Strings(vec.slice1...),
			iterutils.Strings(vec.slice2...),
		)

		assert.Equal(t, uint64(len(vec.slice1)+len(vec.slice2)), it.Count())

		var s string
		res := make([]string, 0)

		for it.Next(&s) {
			res = append(res, s)
		}

		assert.Equal(t, vec.slice1, res[0:len(vec.slice1)])
		assert.Equal(t, vec.slice2, res[len(vec.slice1):])

		res = make([]string, 0)
		it.Reset()

		for it.Next(&s) {
			res = append(res, s)
		}

		assert.Equal(t, vec.slice1, res[0:len(vec.slice1)])
		assert.Equal(t, vec.slice2, res[len(vec.slice1):])

		assert.NoError(t, it.Close())
	}
}
