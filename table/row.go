package table

type Row map[string]any

func (r Row) PrimaryKeyValues(table Table) []any {
	var vals []any
	for _, col := range table.Columns {
		if col.PrimaryKey {
			vals = append(vals, r[col.Name])
		}
	}
	return vals
}
