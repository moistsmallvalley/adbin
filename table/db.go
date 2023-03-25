package table

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func ListDBTables(db *sql.DB) ([]Table, error) {
	names, err := listDBTableNames(db)
	if err != nil {
		return nil, err
	}

	var tables []Table
	for _, name := range names {
		table, err := getTable(db, name)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func listDBTableNames(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, errors.WithStack(err)
		}
		names = append(names, name)
	}
	return names, nil
}

func getTable(db *sql.DB, name string) (Table, error) {
	row := db.QueryRow(fmt.Sprintf("SHOW CREATE TABLE `%s`", name))
	var (
		tableName string
		ddl       string
	)
	if err := row.Scan(&tableName, &ddl); err != nil {
		return Table{}, errors.WithStack(err)
	}
	tables, err := Parse(ddl)
	if err != nil {
		return Table{}, errors.WithStack(err)
	}
	if len(tables) != 1 {
		return Table{}, errors.Errorf("invalid ddl: %s", ddl)
	}
	return tables[0], nil
}
