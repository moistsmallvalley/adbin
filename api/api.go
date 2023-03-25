package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/moistsmallvalley/adbin/log"
	"github.com/moistsmallvalley/adbin/table"
)

type handler struct {
	tables map[string]table.Table
	db     *sql.DB
}

func NewHandler(tables []table.Table, db *sql.DB) http.Handler {
	m := map[string]table.Table{}
	for _, t := range tables {
		m[t.Name] = t
	}
	return &handler{
		tables: m,
		db:     db,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		rows = []map[string]any{}
	}

	writeOK(w, rows)
}

func writeOK(w http.ResponseWriter, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Warn(err.Error())
	}
}

func writeNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	writeErrorMessage(w, "not found")
}

func writeBadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	writeErrorMessage(w, msg)
}

func writeInternalServerError(w http.ResponseWriter, msg string, err error) {
	log.Error("%s: %+v", msg, err)
	w.WriteHeader(http.StatusInternalServerError)
	writeErrorMessage(w, msg)
}

func writeErrorMessage(w http.ResponseWriter, msg string) {
	if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
		log.Warn(err.Error())
	}
}
