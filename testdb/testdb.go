package testdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func InitTestDB(user, pass, dbname string, ddls ...string) (db *sql.DB, err error) {
	conf := mysql.NewConfig()
	conf.Net = "tcp"
	conf.User = user
	conf.Passwd = pass
	conf.ParseTime = true
	conf.Timeout = 5 * time.Second
	conf.ReadTimeout = conf.Timeout
	conf.WriteTimeout = conf.Timeout

	dsn := conf.FormatDSN()

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func(db *sql.DB) {
		if err != nil {
			db.Close()
		}
	}(db)

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = db.Exec(fmt.Sprintf("USE %s", dbname))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, ddl := range ddls {
		if _, err = db.Exec(ddl); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return db, nil
}
