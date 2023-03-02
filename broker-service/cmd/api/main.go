package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	webPort  int32
	rpcApi   string
	authApi  string
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//Configuring application
	app := Config{
		webPort:  8080,
		rpcApi:   os.Getenv("RPC_DATA_API"),
		authApi:  os.Getenv("AUTH_API"),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate),
	}
	//Configuring http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.webPort),
		Handler: app.routes(),
	}
	app.infoLog.Printf("Starting server at port %d\n", app.webPort)
	err := srv.ListenAndServe()
	if err != nil {
		app.errorLog.Print("could not start a server")
		app.errorLog.Print(err)
	}
}
