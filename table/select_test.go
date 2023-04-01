package table

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectStatement(t *testing.T) {
	t.Run("creates query without where clause if keys are not given", func(t *testing.T) {
		table := Table{
			Name: "users",
			Columns: []Column{
				{Name: "id"},
				{Name: "name"},
			},
		}

		query, args := SelectStatement(table, nil, nil)

		assert.Equal(t, "SELECT `id`, `name` FROM `users`", query)
		assert.Nil(t, args)
	})

	t.Run("creates query with where clause if one key is given", func(t *testing.T) {
		table := Table{
			Name: "users",
			Columns: []Column{
				{Name: "id"},
				{Name: "name"},
			},
		}
		idCol := table.Columns[0]

		query, args := SelectStatement(table, []Column{idCol}, []any{123})

		assert.Equal(t, "SELECT `id`, `name` FROM `users` WHERE `id` = ?", query)
		assert.Equal(t, []any{123}, args)
	})

	t.Run("creates query with where clause if two keys are given", func(t *testing.T) {
		table := Table{
			Name: "users",
			Columns: []Column{
				{Name: "id"},
				{Name: "name"},
			},
		}
		idCol := table.Columns[0]
		nameCol := table.Columns[1]

		query, args := SelectStatement(table, []Column{idCol, nameCol}, []any{123, "testuser"})

		assert.Equal(t, "SELECT `id`, `name` FROM `users` WHERE `id` = ? AND `name` = ?", query)
		assert.Equal(t, []any{123, "testuser"}, args)
	})
}

func TestSelect(t *testing.T) {
	db, err := initTestDB(`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  age INT
)`)
	require.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users VALUES (null, "testuser", null)`)
	require.NoError(t, err)

	tables, err := ListDBTables(db)
	require.NoError(t, err)
	table := tables[0]

	ctx := context.Background()
	rows, err := Select(ctx, db, table, nil, nil)
	require.NoError(t, err)

	assert.Equal(t, []Row{
		{
			"id":       int32(1),
			"username": "testuser",
			"age":      (*int32)(nil),
		},
	}, rows)
}
