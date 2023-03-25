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

func TestTableListHandler(t *testing.T) {
	tables := []table.Table{
		{Name: "users"},
		{Name: "messages"},
	}

	h := NewTableListHandler(tables)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	h.ServeHTTP(w, r)

	body, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.JSONEq(t, `["users", "messages"]`, string(body))
}
