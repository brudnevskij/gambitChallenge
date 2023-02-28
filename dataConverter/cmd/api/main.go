package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

// const webPort = "8080"

type Config struct {
	webPort  int32
	rpcPort  int32
	dataApi  string
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	app := Config{
		webPort:  8008,
		rpcPort:  5001,
		dataApi:  "http://tuftuf.gambitlabs.fi/feed.txt",
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.webPort),
		Handler: app.routes(),
	}

	err := rpc.Register(new(RPCServer))
	go app.rpcListen()

	// fmt.Println(convertLong(65480, 65535))
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *Config) rpcListen() error {
	log.Printf("starting rpc serever on port %d\n", app.rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", app.rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
