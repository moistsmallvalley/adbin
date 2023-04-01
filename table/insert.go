package table

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func InsertStatement(table Table, row Row) (query string, args []any) {
	assign, args := assignmentListClause(table, row)
	query = "INSERT INTO `" + table.Name + "` SET " + assign
	return query, args
}

func assignmentListClause(table Table, row Row) (clause string, args []any) {
	for i, col := range table.Columns {
		if i != 0 {
			clause += ", "
		}
		clause += "`" + col.Name + "` = ?"
		args = append(args, row[col.Name])
	}
	return clause, args
}

func Insert(ctx context.Context, db *sql.DB, table Table, row Row) (Row, error) {
	query, args := InsertStatement(table, row)
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	autoKeyCol := table.AutoIncrementPrimaryKeyColumn()
	if autoKeyCol != nil {
		id, err := result.LastInsertId()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		row[autoKeyCol.Name] = id
	}

	return row, nil
}
