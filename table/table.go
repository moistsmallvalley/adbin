package table

import (
	"strings"

	"vitess.io/vitess/go/vt/sqlparser"
)

type Table struct {
	Name    string
	Columns []Column
}

func (t Table) ColumnByName(name string) (column Column, ok bool) {
	for _, column = range t.Columns {
		if column.Name == name {
			return column, true
		}
	}
	return Column{}, false
}

func (t Table) PrimaryKeyColumns() []Column {
	var keys []Column
	for _, c := range t.Columns {
		if c.PrimaryKey {
			keys = append(keys, c)
		}
	}
	return keys
}

func (t Table) AutoIncrementPrimaryKeyColumn() *Column {
	for _, c := range t.Columns {
		if c.PrimaryKey && c.AutoIncrement {
			return &c
		}
	}
	return nil
}

type Column struct {
	Name          string
	Type          Type
	Required      bool
	PrimaryKey    bool
	AutoIncrement bool
}

func (c Column) IsTime() bool {
	return c.Type == TypeTime || c.Type == TypeDate || c.Type == TypeDateTime || c.Type == TypeTimestamp
}

func Parse(ddl string) ([]Table, error) {
	var tables []Table
	for {
		ddl = strings.TrimSpace(ddl)
		if ddl == "" {
			break
		}
		first, rest, err := sqlparser.SplitStatement(ddl)
		if err != nil {
			return nil, err
		}
		stmt, _, err := sqlparser.Parse2(first)
		if err != nil {
			return nil, err
		}
		if ct, ok := stmt.(*sqlparser.CreateTable); ok {
			table, err := makeTable(ct)
			if err != nil {
				return nil, err
			}
			tables = append(tables, table)
		}
		ddl = rest
	}
	return tables, nil
}

func makeTable(parserTable *sqlparser.CreateTable) (Table, error) {
	var columns []Column
	for _, pc := range parserTable.TableSpec.Columns {
		primaryKey := isPrimaryKey(pc, parserTable.TableSpec.Indexes)
		typ, err := makeType(pc.Type.Type)
		if err != nil {
			return Table{}, err
		}
		column := Column{
			Name:          pc.Name.String(),
			Type:          typ,
			Required:      (pc.Type.Options.Null != nil && !*pc.Type.Options.Null) || primaryKey,
			PrimaryKey:    primaryKey,
			AutoIncrement: pc.Type.Options.Autoincrement,
		}
		columns = append(columns, column)
	}
	return Table{
		Name:    parserTable.Table.Name.String(),
		Columns: columns,
	}, nil
}

func isPrimaryKey(column *sqlparser.ColumnDefinition, indexes []*sqlparser.IndexDefinition) bool {
	if column.Type.Options.KeyOpt == sqlparser.ColKeyPrimary {
		return true
	}
	for _, index := range indexes {
		if !index.Info.Primary {
			continue
		}
		for _, indexColumn := range index.Columns {
			if indexColumn.Column.Equal(column.Name) {
				return true
			}
		}
	}
	return false
}
