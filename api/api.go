package api

import (
	"strings"

	"github.com/moistsmallvalley/adbin/table"
	"github.com/pkg/errors"
)

func tableMap(tables []table.Table) map[string]table.Table {
	m := map[string]table.Table{}
	for _, t := range tables {
		m[t.Name] = t
	}
	return m
}

func parseKeyPath(keyPath string, tbl table.Table) (keyCols []table.Column, keyVals []any, err error) {
	keyStrs := strings.Split(keyPath, "/")
	keyCols = tbl.PrimaryKeyColumns()
	if len(keyStrs) != len(keyCols) {
		return nil, nil, errors.Errorf("number of keys in path (%d) != number of keys in table (%d)", len(keyStrs), len(keyCols))
	}

	for i, col := range keyCols {
		key, err := col.Type.Parse(keyStrs[i])
		if err != nil {
			return nil, nil, err
		}
		keyVals = append(keyVals, key)
	}
	return keyCols, keyVals, nil
}
