package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func (app *Config) getData(wr http.ResponseWriter, r *http.Request) {
	log.Println("hit")
}

func (app *Config) getDataViaRpc(wr http.ResponseWriter, r *http.Request) {

	client, err := rpc.Dial("tcp", "localhost:5001")
	if err != nil {
		log.Print(err)
		return
	}

	var data HumanReadableData
	err = client.Call("RPCServer.GetDataRPC", "", &data)

	if err != nil {
		log.Print(err)
		return
	}

	log.Print(data)
}
