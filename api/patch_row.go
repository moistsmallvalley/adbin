package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type PatchRowRequest table.Row

type patchRowHandler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewPatchRowHandler(tables []table.Table, db *sql.DB) http.Handler {
	return &patchRowHandler{
		tables: tableMap(tables),
		db:     db,
	}
}

func (h *patchRowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var row table.Row
	if err := json.NewDecoder(r.Body).Decode(&row); err != nil {
		writeBadRequest(w, "invalid json format: "+err.Error())
		return
	}
	for i, keyVal := range keyVals {
		row[keyCols[i].Name] = keyVal
	}

	rows, err := table.Select(r.Context(), h.db, tbl, keyCols, keyVals)
	if err != nil {
		writeInternalServerError(w, "db select error", err)
		return
	}
	if len(rows) == 0 {
		writeNotFound(w)
		return
	}

	row, err = row.ParseTimes(tbl)
	if err != nil {
		writeBadRequest(w, err.Error())
		return
	}

	if err := table.ValidateRow(tbl, row); err != nil {
		writeBadRequest(w, err.Error())
		return
	}

	if err := table.Update(r.Context(), h.db, tbl, row); err != nil {
		writeInternalServerError(w, "db insert error", err)
		return
	}

	writeOK(w, map[string]string{})
}
