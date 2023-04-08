package table

import (
	"time"

	"github.com/pkg/errors"
)

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

func (r Row) ParseTimes(table Table) (Row, error) {
	parsed := Row{}
	for k, v := range r {
		c, ok := table.ColumnByName(k)
		if !ok {
			return nil, errors.Errorf("unknow column %s", k)
		}
		if !c.IsTime() {
			parsed[k] = v
			continue
		}
		if v == nil {
			parsed[k] = v
			continue
		}
		if _, ok := v.(time.Time); ok {
			parsed[k] = v
			continue
		}
		if _, ok := v.(*time.Time); ok {
			parsed[k] = v
			continue
		}
		s, ok := v.(string)
		if !ok {
			return nil, errors.Errorf("%s is not time string", k)
		}
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		parsed[k] = t
	}
	return parsed, nil
}
