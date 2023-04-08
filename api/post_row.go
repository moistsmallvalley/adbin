package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type PostTableRequest table.Row

type postRowHandler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewPostRowHandler(tables []table.Table, db *sql.DB) http.Handler {
	return &postRowHandler{
		tables: tableMap(tables),
		db:     db,
	}
}

func (h *postRowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := urlParam(r, "name")
	tbl, ok := h.tables[name]
	if !ok {
		writeNotFound(w)
		return
	}

	var row table.Row
	if err := json.NewDecoder(r.Body).Decode(&row); err != nil {
		writeBadRequest(w, "invalid json format: "+err.Error())
		return
	}

	row, err := row.ParseTimes(tbl)
	if err != nil {
		writeBadRequest(w, err.Error())
		return
	}

	if err := table.ValidateRow(tbl, row); err != nil {
		writeBadRequest(w, err.Error())
		return
	}

	newRow, err := table.Insert(r.Context(), h.db, tbl, row)
	if err != nil {
		writeInternalServerError(w, "db insert error", err)
		return
	}

	writeOK(w, newRow)
}
