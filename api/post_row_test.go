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

func TestPostRowHandlerServeHTTP(t *testing.T) {
	t.Run("inserts not null value", func(t *testing.T) {
		db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
			`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`)
		require.NoError(t, err)

		tables, err := table.ListDBTables(db)
		require.NoError(t, err)

		reqBody := bytes.NewBuffer([]byte(`{"username": "testname"}`))

		h := NewPostRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", reqBody)
		r = withURLParams(r, map[string]string{"name": "users"})
		h.ServeHTTP(w, r)

		body, err := io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.JSONEq(t, `{"id": 1, "username": "testname"}`, string(body))

		ctx := context.Background()
		rows, err := table.Select(ctx, db, tables[0], nil, nil)
		require.NoError(t, err)
		assert.Equal(t, []table.Row{
			{"id": int32(1), "username": "testname"},
		}, rows)
	})

	t.Run("inserts null value", func(t *testing.T) {
		db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
			`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  age int
)`)
		require.NoError(t, err)

		tables, err := table.ListDBTables(db)
		require.NoError(t, err)

		reqBody := bytes.NewBuffer([]byte(`{"age": null}`))

		h := NewPostRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", reqBody)
		r = withURLParams(r, map[string]string{"name": "users"})
		h.ServeHTTP(w, r)

		body, err := io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.JSONEq(t, `{"id": 1, "age": null}`, string(body))

		ctx := context.Background()
		rows, err := table.Select(ctx, db, tables[0], nil, nil)
		require.NoError(t, err)
		assert.Equal(t, []table.Row{
			{"id": int32(1), "age": (*int32)(nil)},
		}, rows)
	})
}
