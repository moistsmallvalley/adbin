package api

import (
	"database/sql"
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type GetRowResponse struct {
	Columns []ColumnResponse `json:"columns"`
	Row     table.Row        `json:"row"`
}

type getRowHandler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewGetRowHandler(tables []table.Table, db *sql.DB) http.Handler {
	return &getRowHandler{
		tables: tableMap(tables),
		db:     db,
	}
}

func (h *getRowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := urlParam(r, "name")
	tbl, ok := h.tables[name]
	if !ok {
		writeNotFound(w)
		return
	}

	keyPath := urlParam(r, "keyPath")
	keyCols, keyVals, err := parseKeyPath(keyPath, tbl)
	if err != nil {
		writeNotFound(w)
		return
	}

	rows, err := table.Select(r.Context(), h.db, tbl, keyCols, keyVals)
	if err != nil {
		writeInternalServerError(w, "select error", err)
		return
	}
	if len(rows) == 0 {
		writeNotFound(w)
		return
	}

	var columns []ColumnResponse
	for _, c := range tbl.Columns {
		columns = append(columns, toColumnResponse(c))
	}

	writeOK(w, GetRowResponse{
		Columns: columns,
		Row:     rows[0],
	})
}
