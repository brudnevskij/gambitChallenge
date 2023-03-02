package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

type Config struct {
	webPort  int32
	rpcPort  int32
	dataApi  string
	infoLog  *log.Logger
	errorLog *log.Logger
}

var app Config = Config{
	webPort:  80,
	rpcPort:  5001,
	dataApi:  os.Getenv("DATA_API"),
	infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
}

func main() {

	//Code comented here starts a HTTP server
	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", app.webPort),
	// 	Handler: app.routes(),
	// }

	err := rpc.Register(new(RPCServer))
	app.rpcListen()

	if err != nil {
		app.errorLog.Print("Failed to start a RPC server")
		app.errorLog.Print(err)
	}

	// err = srv.ListenAndServe()
	// if err != nil {
	// app.errorLog.Print("Failed to start a web server")
	// app.errorLog.Print(err)
	// }
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
