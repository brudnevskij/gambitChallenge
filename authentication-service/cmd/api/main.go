package main

import (
	"authentication-service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var counts int64

type Config struct {
	webPort  int32
	DB       *sql.DB
	Models   data.Models
	infoLog  *log.Logger
	errorLog *log.Logger
}

// Configuring application
var app Config = Config{
	webPort:  80,
	infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
}

func main() {
	app.infoLog.Print("Starting auth service")

	//connect to DB
	conn := connectToDB()
	if conn == nil {
		app.errorLog.Panic("Can't connect to Postgres")
	}
	app.DB = conn
	app.Models = data.New(conn)

	//Configuring http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		app.errorLog.Print("Could not start a server")
		app.errorLog.Panic(err)
	}
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDb(dsn)
		if err != nil {
			app.infoLog.Print("DB is not yet ready")
			counts++
		} else {
			app.infoLog.Print("Connected to DB")
			return connection
		}
		if counts > 10 {
			app.errorLog.Print(err)
			return nil
		}
		app.infoLog.Print("Backing off for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
