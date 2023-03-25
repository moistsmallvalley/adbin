package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectStatement(t *testing.T) {
	table := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id"},
			{Name: "name"},
		},
	}

	query, args := SelectStatement(table)

	assert.Equal(t, "SELECT `id`, `name` FROM `users`", query)
	assert.Nil(t, args)
}

func TestSelect(t *testing.T) {
	db, err := initTestDB(`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  age INT
)`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO users values (null, "testuser", null)`)
	require.NoError(t, err)

	tables, err := ListDBTables(db)
	require.NoError(t, err)
	require.Len(t, tables, 1)
	table := tables[0]

	rows, err := Select(db, table)
	require.NoError(t, err)

	assert.Equal(t, []map[string]any{
		{
			"id":       int32(1),
			"username": "testuser",
			"age":      (*int32)(nil),
		},
	}, rows)
}
