package api

import (
	"database/sql"
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type GetRowsResponse struct {
	Columns []ColumnResponse `json:"columns"`
	Rows    []table.Row      `json:"rows"`
}

type ColumnResponse struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Required      bool   `json:"required"`
	PrimaryKey    bool   `json:"primaryKey"`
	AutoIncrement bool   `json:"autoIncrement"`
}

func toColumnResponse(c table.Column) ColumnResponse {
	return ColumnResponse{
		Name:          c.Name,
		Type:          string(c.Type),
		Required:      c.Required,
		PrimaryKey:    c.PrimaryKey,
		AutoIncrement: c.AutoIncrement,
	}
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

	var columns []ColumnResponse
	for _, c := range tbl.Columns {
		columns = append(columns, toColumnResponse(c))
	}

	writeOK(w, GetRowsResponse{
		Columns: columns,
		Rows:    rows,
	})
}
