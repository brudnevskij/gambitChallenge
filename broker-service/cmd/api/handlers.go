package main

import (
	"net/http"
	"net/rpc"
)

func (app *Config) getDataViaRpc(wr http.ResponseWriter, r *http.Request) {

	var jsonPayload jsonResponse

	//dial a connection with rpc server
	client, err := rpc.Dial("tcp", app.rpcApi)
	if err != nil {
		app.errorJSON(wr, err)
		return
	}

	//calling getdata function in the remote rpc server
	var data HumanReadableData
	err = client.Call("RPCServer.GetDataRPC", "", &data)
	if err != nil {
		app.errorJSON(wr, err)
		return
	}
	jsonPayload.Error = false
	jsonPayload.Message = "Data Delivered"
	jsonPayload.Data = data
	app.writeJSON(wr, http.StatusAccepted, jsonPayload)
}
