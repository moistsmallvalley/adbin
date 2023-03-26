package api

import (
	"database/sql"
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type GetTableResponse struct {
	Columns []string    `json:"columns"`
	Rows    []table.Row `json:"rows"`
}

type getTableHandler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewGetTableHandler(tables []table.Table, db *sql.DB) http.Handler {
	m := map[string]table.Table{}
	for _, t := range tables {
		m[t.Name] = t
	}
	return &getTableHandler{
		tables: m,
		db:     db,
	}
}

func (h *getTableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := urlParam(r, "name")
	tbl, ok := h.tables[name]
	if !ok {
		writeNotFound(w)
		return
	}

	rows, err := table.Select(h.db, tbl)
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

	writeOK(w, GetTableResponse{
		Columns: columns,
		Rows:    rows,
	})
}
