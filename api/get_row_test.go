package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moistsmallvalley/adbin/table"
	"github.com/moistsmallvalley/adbin/testdb"
)

func TestGetRowHandlerServeHTTP(t *testing.T) {
	t.Run("get existing record", func(t *testing.T) {
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

		h := NewGetRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r = withURLParams(r, map[string]string{"name": "users", "keyPath": "1"})
		h.ServeHTTP(w, r)

		body, err := io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.JSONEq(t, `{"id": 1, "username": "testname"}`, string(body))
	})

	t.Run("get existing multi key record", func(t *testing.T) {
		db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
			`
CREATE TABLE friendships (
  user_id INT NOT NULL,
  other_id INT NOT NULL,
  closeness  INT NOT NULL,
  PRIMARY KEY (user_id, other_id)
)`)
		require.NoError(t, err)

		_, err = db.Exec(`INSERT INTO friendships VALUES(10, 20, 100)`)
		require.NoError(t, err)

		tables, err := table.ListDBTables(db)
		require.NoError(t, err)

		h := NewGetRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r = withURLParams(r, map[string]string{"name": "friendships", "keyPath": "10/20"})
		h.ServeHTTP(w, r)

		body, err := io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.JSONEq(t, `{"user_id": 10, "other_id": 20, "closeness": 100}`, string(body))
	})

	t.Run("get not found record", func(t *testing.T) {
		db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
			`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`)
		require.NoError(t, err)

		tables, err := table.ListDBTables(db)
		require.NoError(t, err)

		h := NewGetRowHandler(tables, db)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r = withURLParams(r, map[string]string{"name": "users", "keyPath": "1"})
		h.ServeHTTP(w, r)

		_, err = io.ReadAll(w.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})
}
