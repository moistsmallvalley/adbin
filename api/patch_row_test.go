package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moistsmallvalley/adbin/table"
	"github.com/moistsmallvalley/adbin/testdb"
)

func TestPatchRowHandlerServeHTTP(t *testing.T) {
	t.Run("updates existing record", func(t *testing.T) {
		db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
			`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`)
		require.NoError(t, err)

		_, err = db.Exec(`INSERT INTO users VALUES(null, "testname")`)
		require.NoError(t, err)

		tables, err := table.ListDBTables(db)
		require.NoError(t, err)

		reqBody := bytes.NewBuffer([]byte(`{"username": "updatedname"}`))

		h := NewPatchRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", reqBody)
		r = withURLParams(r, map[string]string{"name": "users", "keyPath": "1"})
		h.ServeHTTP(w, r)

		body, err := io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.JSONEq(t, `{}`, string(body))

		ctx := context.Background()
		rows, err := table.Select(ctx, db, tables[0], nil, nil)
		require.NoError(t, err)
		assert.Equal(t, []table.Row{
			{"id": int32(1), "username": "updatedname"},
		}, rows)
	})

	t.Run("denys updating nonexisting record", func(t *testing.T) {
		db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
			`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`)
		require.NoError(t, err)

		tables, err := table.ListDBTables(db)
		require.NoError(t, err)

		reqBody := bytes.NewBuffer([]byte(`{"username": "updatedname"}`))

		h := NewPatchRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", reqBody)
		r = withURLParams(r, map[string]string{"name": "users", "keyPath": "1"})
		h.ServeHTTP(w, r)

		_, err = io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)

		ctx := context.Background()
		rows, err := table.Select(ctx, db, tables[0], nil, nil)
		require.NoError(t, err)
		assert.Len(t, rows, 0)
	})
}
