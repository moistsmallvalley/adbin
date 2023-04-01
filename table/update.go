package table

import (
	"context"
	"database/sql"
)

func UpdateStatement(table Table, row Row) (query string, args []any) {
	assign, assignArgs := assignmentListClause(table, row)
	where, whereArgs := whereClause(table.PrimaryKeyColumns(), row.PrimaryKeyValues(table))
	query = "UPDATE `" + table.Name + "` SET " + assign + " " + where
	args = append(args, assignArgs...)
	args = append(args, whereArgs...)
	return query, args
}

func Update(ctx context.Context, db *sql.DB, table Table, row Row) error {
	query, args := UpdateStatement(table, row)
	if _, err := db.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
