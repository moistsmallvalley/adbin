package api

import (
	"net/http"

	"github.com/moistsmallvalley/adbin/table"
)

type GetTableResponse struct {
	Name    string           `json:"name"`
	Columns []ColumnResponse `json:"columns"`
}

type getTableHandler struct {
	tables map[string]table.Table
}

func NewGetTableHandler(tables []table.Table) http.Handler {
	return &getTableHandler{
		tables: tableMap(tables),
	}
}

func (h *getTableHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := urlParam(r, "name")
	tbl, ok := h.tables[name]
	if !ok {
		writeNotFound(w)
		return
	}

	var columns []ColumnResponse
	for _, c := range tbl.Columns {
		columns = append(columns, toColumnResponse(c))
	}

	writeOK(w, GetTableResponse{
		Name:    name,
		Columns: columns,
	})
}
