package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moistsmallvalley/adbin/table"
)

func TestGetTableHandler(t *testing.T) {
	tables := []table.Table{
		{
			Name: "users",
			Columns: []table.Column{
				{Name: "id", Type: table.TypeInt, Required: true, PrimaryKey: true},
				{Name: "username", Type: table.TypeVarChar, Required: true},
			},
		},
	}

	h := NewGetTableHandler(tables)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r = withURLParams(r, map[string]string{"name": "users"})
	h.ServeHTTP(w, r)

	body, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.JSONEq(t, `
{
	"name": "users",
	"columns": [
		{"name": "id", "type": "int", "required": true, "primaryKey": true, "autoIncrement": false},
		{"name": "username", "type": "varchar", "required": true, "primaryKey": false, "autoIncrement": false}
	]
}`, string(body))
}
