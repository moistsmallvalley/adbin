package api

import (
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type tableListHandler struct {
	tables []table.Table
}

func NewTableListHandler(tables []table.Table) http.Handler {
	return &tableListHandler{
		tables: tables,
	}
}

func (h *tableListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeNotFound(w)
		return
	}

	var names []string
	for _, t := range h.tables {
		names = append(names, t.Name)
	}
	writeOK(w, names)
}
