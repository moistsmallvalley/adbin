package table

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateStatement(t *testing.T) {
	t.Run("returns update statement with one primary key", func(t *testing.T) {
		table := Table{
			Name: "users",
			Columns: []Column{
				{Name: "id", Type: TypeInt, PrimaryKey: true},
				{Name: "name", Type: TypeVarChar},
			},
		}

		query, args := UpdateStatement(table, Row{"id": 123, "name": "testname"})

		assert.Equal(t, "UPDATE `users` SET `id` = ?, `name` = ? WHERE `id` = ?", query)
		assert.Equal(t, []any{123, "testname", 123}, args)
	})

	t.Run("returns update statement with two primary keys", func(t *testing.T) {
		table := Table{
			Name: "friendships",
			Columns: []Column{
				{Name: "user_id", Type: TypeInt, PrimaryKey: true},
				{Name: "friend_id", Type: TypeInt, PrimaryKey: true},
				{Name: "closeness", Type: TypeInt},
			},
		}

		query, args := UpdateStatement(table, Row{"user_id": 123, "friend_id": 234, "closeness": 333})

		assert.Equal(t, "UPDATE `friendships` SET `user_id` = ?, `friend_id` = ?, `closeness` = ? WHERE `user_id` = ? AND `friend_id` = ?", query)
		assert.Equal(t, []any{123, 234, 333, 123, 234}, args)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updates existing row", func(t *testing.T) {
		db, err := initTestDB(`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`)
		require.NoError(t, err)
		defer db.Close()

		_, err = db.Exec(`INSERT INTO users VALUES (null, "testuser")`)
		require.NoError(t, err)

		tables, err := ListDBTables(db)
		require.NoError(t, err)
		table := tables[0]

		ctx := context.Background()
		err = Update(ctx, db, table, Row{"id": 1, "username": "updateduser"})
		require.NoError(t, err)

		rows, err := Select(ctx, db, table, nil, nil)
		require.NoError(t, err)
		assert.Equal(t, Row{"id": int32(1), "username": "updateduser"}, rows[0])
	})

	t.Run("doesn't save nonexistent row", func(t *testing.T) {
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
		err = Update(ctx, db, table, Row{"id": 1, "username": "updateduser"})
		require.NoError(t, err)

		rows, err := Select(ctx, db, table, nil, nil)
		require.NoError(t, err)
		assert.Len(t, rows, 0)
	})
}
