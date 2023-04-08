package table

import (
	"testing"
	"time"

	"github.com/moistsmallvalley/adbin/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRowPrimaryKeyValues(t *testing.T) {
	table := Table{
		Columns: []Column{
			{Name: "id", PrimaryKey: true},
			{Name: "id2", PrimaryKey: true},
			{Name: "value"}},
	}
	row := Row{
		"id":    1,
		"id2":   2,
		"value": 3,
	}
	assert.Equal(t, []any{1, 2}, row.PrimaryKeyValues(table))
}

func TestRowParseTimes(t *testing.T) {
	table := Table{
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "datetime", Type: TypeDateTime},
			{Name: "timestamp", Type: TypeTimestamp},
			{Name: "date", Type: TypeDate},
			{Name: "time", Type: TypeTime},
			{Name: "niltime", Type: TypeTime},
		},
	}
	row := Row{
		"id":        1,
		"datetime":  "2000-04-01T10:20:30Z",
		"timestamp": "2000-05-01T10:20:30Z",
		"date":      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		"time":      util.Ptr(time.Date(2000, 2, 2, 0, 0, 0, 0, time.UTC)),
		"niltime":   nil,
	}

	actual, err := row.ParseTimes(table)
	require.NoError(t, err)

	assert.Equal(t, Row{
		"id":        1,
		"datetime":  time.Date(2000, 4, 1, 10, 20, 30, 0, time.UTC),
		"timestamp": time.Date(2000, 5, 1, 10, 20, 30, 0, time.UTC),
		"date":      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		"time":      util.Ptr(time.Date(2000, 2, 2, 0, 0, 0, 0, time.UTC)),
		"niltime":   nil,
	}, actual)
}
