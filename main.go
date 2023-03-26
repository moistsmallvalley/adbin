package main

import (
	"database/sql"
	"flag"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moistsmallvalley/adbin/api"
	"github.com/moistsmallvalley/adbin/log"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", "http://localhost:5173"))
	r.Use(middleware.SetHeader("Access-Control-Allow-Method", "*"))

	r.Get("/api/tables", api.NewGetTablesHandler(tables).ServeHTTP)
	r.Get("/api/tables/{name}", api.NewGetTableHandler(tables, db).ServeHTTP)

	log.Info("starting server")
	if err := http.ListenAndServe(":8080", r); err != nil {
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
