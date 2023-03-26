package api

import (
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type getTablesHandler struct {
	tables []table.Table
}

func NewGetTablesHandler(tables []table.Table) http.Handler {
	return &getTablesHandler{
		tables: tables,
	}
}

func (h *getTablesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
