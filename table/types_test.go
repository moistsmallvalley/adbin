package table

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeParse(t *testing.T) {
	tests := []struct {
		name     string
		t        Type
		s        string
		expected any
	}{
		{name: "int", t: TypeInt, s: "123", expected: int64(123)},
		{name: "datetime", t: TypeDateTime, s: "2020-01-02T10:20:30Z", expected: time.Date(2020, 1, 2, 10, 20, 30, 0, time.UTC)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.t.Parse(test.s)
			require.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
