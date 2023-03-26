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

const (
	testDBUser = "root"
	testDBPass = "example"
	testDBName = "apitest"
)

func TestGetTableHandlerServeHTTP(t *testing.T) {
	db, err := testdb.InitTestDB(testDBUser, testDBPass, testDBName,
		`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  age int
)`)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO users VALUES (null, "testuser", 23)`)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO users VALUES (null, "otheruser", null)`)
	require.NoError(t, err)

	tables, err := table.ListDBTables(db)
	require.NoError(t, err)

	h := NewGetTableHandler(tables, db)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r = withURLParams(r, map[string]string{"name": "users"})
	h.ServeHTTP(w, r)

	body, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.JSONEq(t, `
{
	"columns": ["id", "username", "age"],
	"rows": [
		{
			"id": 1,
			"username": "testuser",
			"age": 23
		},
		{
			"id": 2,
			"username": "otheruser",
			"age": null
		}
	]
}`, string(body))
}
