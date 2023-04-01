package table

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertStatement(t *testing.T) {
	table := Table{
		Name: "users",
		Columns: []Column{
			{Name: "id", Type: TypeInt},
			{Name: "name", Type: TypeVarChar},
		},
	}
	row := Row{
		"id":   123,
		"name": "testname",
	}

	query, args := InsertStatement(table, row)

	assert.Equal(t, "INSERT INTO `users` SET `id` = ?, `name` = ?", query)
	assert.Equal(t, []any{123, "testname"}, args)
}

func TestInsert(t *testing.T) {
	db, err := initTestDB(`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`)
	require.NoError(t, err)
	defer db.Close()

	tables, err := ListDBTables(db)
	require.NoError(t, err)
	table := tables[0]

	ctx := context.Background()
	row, err := Insert(ctx, db, table, Row{"username": "testuser"})

	require.NoError(t, err)
	assert.Equal(t, Row{
		"id":       int64(1),
		"username": "testuser",
	}, row)
}
