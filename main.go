package main

import (
	"database/sql"
	"flag"
	"net/http"
	"strconv"
	"time"

	"github.com/moistsmallvalley/adbin/api"
	"github.com/moistsmallvalley/adbin/log"
	"github.com/moistsmallvalley/adbin/middleware"
	"github.com/moistsmallvalley/adbin/table"
	"github.com/pkg/errors"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	args, err := parseCommandArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := openDB(args)
	if err != nil {
		log.Fatal(err.Error())
	}

	tables, err := table.ListDBTables(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/api/tables", http.StripPrefix("/api/tables", api.NewTableListHandler(tables)))
	mux.Handle("/api/tables/", http.StripPrefix("/api/tables", api.NewTableHandler(tables, db)))

	handler := middleware.NewCORSMiddleware(mux, "http://localhost:5173", "*")

	log.Info("starting server")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Info(err.Error())
	}
}

type commandArgs struct {
	host     string
	port     int
	dbname   string
	user     string
	password string
}

func parseCommandArgs() (*commandArgs, error) {
	host := flag.String("host", "localhost", "db host")
	port := flag.Int("port", 3306, "db port")
	dbname := flag.String("dbname", "", "db name")
	user := flag.String("user", "root", "db user")
	password := flag.String("password", "", "db password")

	flag.Parse()

	if *dbname == "" {
		return nil, errors.New("dbname not set")
	}

	return &commandArgs{
		host:     *host,
		port:     *port,
		dbname:   *dbname,
		user:     *user,
		password: *password,
	}, nil
}

func openDB(args *commandArgs) (*sql.DB, error) {
	conf := mysql.NewConfig()
	conf.Net = "tcp"
	conf.Addr = args.host + ":" + strconv.Itoa(args.port)
	conf.DBName = args.dbname
	conf.User = args.user
	conf.Passwd = args.password
	conf.ParseTime = true
	conf.Timeout = 5 * time.Second
	conf.ReadTimeout = conf.Timeout
	conf.WriteTimeout = conf.Timeout

	dsn := conf.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil

}
