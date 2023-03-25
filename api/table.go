package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/moistsmallvalley/adbin/table"
)

type TableResponse struct {
	Columns []string    `json:"columns"`
	Rows    []table.Row `json:"rows"`
}

type tableHandler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewTableHandler(tables []table.Table, db *sql.DB) http.Handler {
	m := map[string]table.Table{}
	for _, t := range tables {
		m[t.Name] = t
	}
	return &tableHandler{
		tables: m,
		db:     db,
	}
}

func (h *tableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeNotFound(w)
		return
	}

	sp := strings.Split(r.URL.Path, "/")
	if len(sp) < 2 {
		writeBadRequest(w, fmt.Sprintf("invalid path: %s", r.URL.Path))
		return
	}

	name := sp[1]
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

	writeOK(w, TableResponse{
		Columns: columns,
		Rows:    rows,
	})
}
