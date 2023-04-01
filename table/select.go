package table

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/moistsmallvalley/adbin/log"
	"github.com/pkg/errors"
)

func SelectStatement(table Table, keyCols []Column, keyVals []any) (query string, args []any) {
	query = "SELECT "
	for i, col := range table.Columns {
		if i != 0 {
			query += ", "
		}
		query += "`" + col.Name + "`"
	}
	query += " FROM " + "`" + table.Name + "`"
	if len(keyCols) > 0 {
		where, whereArgs := whereClause(keyCols, keyVals)
		query += " " + where
		args = append(args, whereArgs...)
	}
	return query, args
}

func whereClause(keyCols []Column, keyVals []any) (clause string, args []any) {
	clause = "WHERE "
	for i, key := range keyCols {
		if i != 0 {
			clause += " AND "
		}
		clause += "`" + key.Name + "` = ?"
		args = append(args, keyVals[i])
	}
	return clause, args
}

func Select(ctx context.Context, db *sql.DB, table Table, keyCols []Column, keyVals []any) ([]Row, error) {
	query, args := SelectStatement(table, keyCols, keyVals)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	return ScanRows(table, rows)
}

func ScanRows(table Table, rows *sql.Rows) ([]Row, error) {
	var objs []Row
	for rows.Next() {
		fields, err := allocFields(table)
		if err != nil {
			return nil, err
		}
		if err := rows.Scan(fields...); err != nil {
			return nil, err
		}
		obj := map[string]any{}
		for i, c := range table.Columns {
			obj[c.Name] = deref(fields[i])
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func allocFields(table Table) ([]any, error) {
	var fields []any
	for _, c := range table.Columns {
		if c.Required {
			switch c.Type {
			case TypeInteger, TypeInt:
				fields = append(fields, new(int32))
			case TypeSmallInt:
				fields = append(fields, new(int16))
			case TypeTinyInt:
				fields = append(fields, new(int8))
			case TypeMediumInt:
				fields = append(fields, new(int32))
			case TypeBigInt:
				fields = append(fields, new(int64))
			case TypeDecimal, TypeNumeric:
				fields = append(fields, new(float64)) // TODO: use fixed floating point value
			case TypeFloat:
				fields = append(fields, new(float32))
			case TypeDouble:
				fields = append(fields, new(float64))
			case TypeBit:
				fields = append(fields, new(int8))
			case TypeDate, TypeDateTime, TypeTimestamp, TypeTime:
				fields = append(fields, new(time.Time))
			case TypeYear:
				fields = append(fields, new(int32))
			case TypeChar, TypeVarChar:
				fields = append(fields, new(string))
			case TypeBinary, TypeVarBinary, TypeBlob:
				fields = append(fields, new([]byte))
			case TypeText:
				fields = append(fields, new(string))
			case TypeEnum:
				fields = append(fields, new(string))
			case TypeSet:
				return nil, errors.New("set columns not supported yet")
			case TypeJson:
				fields = append(fields, new(string))
			default:
				return nil, errors.Errorf("unsupported column type: %s", c.Type)
			}
		} else {
			switch c.Type {
			case TypeInteger, TypeInt:
				fields = append(fields, new(*int32))
			case TypeSmallInt:
				fields = append(fields, new(*int16))
			case TypeTinyInt:
				fields = append(fields, new(*int8))
			case TypeMediumInt:
				fields = append(fields, new(*int32))
			case TypeBigInt:
				fields = append(fields, new(*int64))
			case TypeDecimal, TypeNumeric:
				fields = append(fields, new(*float64)) // TODO: use fixed floating point value
			case TypeFloat:
				fields = append(fields, new(*float32))
			case TypeDouble:
				fields = append(fields, new(*float64))
			case TypeBit:
				fields = append(fields, new(*int8))
			case TypeDate, TypeDateTime, TypeTimestamp, TypeTime:
				fields = append(fields, new(*time.Time))
			case TypeYear:
				fields = append(fields, new(*int32))
			case TypeChar, TypeVarChar:
				fields = append(fields, new(*string))
			case TypeBinary, TypeVarBinary, TypeBlob:
				fields = append(fields, new(*[]byte))
			case TypeText:
				fields = append(fields, new(*string))
			case TypeEnum:
				fields = append(fields, new(*string))
			case TypeSet:
				return nil, errors.New("set columns not supported yet")
			case TypeJson:
				fields = append(fields, new(*string))
			default:
				return nil, errors.Errorf("unsupported column type: %s", c.Type)
			}
		}
	}
	return fields, nil
}

func deref(ptr any) any {
	switch v := ptr.(type) {
	case *int8:
		return *v
	case *int16:
		return *v
	case *int32:
		return *v
	case *int64:
		return *v
	case *float32:
		return *v
	case *float64:
		return *v
	case *string:
		return *v
	case *[]byte:
		return *v
	case *time.Time:
		return *v

	case **int8:
		return *v
	case **int16:
		return *v
	case **int32:
		return *v
	case **int64:
		return *v
	case **float32:
		return *v
	case **float64:
		return *v
	case **string:
		return *v
	case **[]byte:
		return *v
	case **time.Time:
		return *v

	default:
		msg := fmt.Sprintf("deref unsupported type: %+v", ptr)
		log.Error(msg)
		panic(msg)
	}
}
