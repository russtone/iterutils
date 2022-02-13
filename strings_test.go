package iterutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/russtone/iterutils"
)

func Test_Slice(t *testing.T) {
	tests := []struct {
		strings []string
	}{
		{
			[]string{"one", "two", "three"},
		},
	}

	for _, vec := range tests {
		it := iterutils.Strings(vec.strings...)

		assert.Equal(t, uint64(len(vec.strings)), it.Count())

		var s string
		res := make([]string, 0)

		for it.Next(&s) {
			res = append(res, s)
		}

		assert.Equal(t, vec.strings, res)

		it.Reset()
		res = make([]string, 0)

		for it.Next(&s) {
			res = append(res, s)
		}

		assert.Equal(t, vec.strings, res)

		assert.NoError(t, it.Close())
	}
}
