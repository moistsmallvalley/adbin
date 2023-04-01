package api

import (
	"database/sql"
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type GetRowsResponse struct {
	Columns []string    `json:"columns"`
	Rows    []table.Row `json:"rows"`
}

type getRowsHandler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewGetRowsHandler(tables []table.Table, db *sql.DB) http.Handler {
	return &getRowsHandler{
		tables: tableMap(tables),
		db:     db,
	}
}

func (h *getRowsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := urlParam(r, "name")
	tbl, ok := h.tables[name]
	if !ok {
		writeNotFound(w)
		return
	}

	rows, err := table.Select(r.Context(), h.db, tbl, nil, nil)
	if err != nil {
		writeInternalServerError(w, "db scan rows error", err)
		return
	}
	if rows == nil {
		rows = []table.Row{}
	}

	var columns []string
	for _, c := range tbl.Columns {
		columns = append(columns, c.Name)
	}

	writeOK(w, GetRowsResponse{
		Columns: columns,
		Rows:    rows,
	})
}
