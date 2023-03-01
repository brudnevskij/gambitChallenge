package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/rpc"
)

func (app *Config) getDataViaRpc(wr http.ResponseWriter, r *http.Request) {

	//getting fields from body of the request
	var jsonPayload jsonResponse
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(wr, r, &requestPayload)
	if err != nil {
		app.errorJSON(wr, err, http.StatusBadRequest)
		return
	}

	jsonData, _ := json.Marshal(requestPayload)

	//authenticating user
	request, err := http.NewRequest(http.MethodPost, app.authApi, bytes.NewBuffer(jsonData))
	request.Header.Set("content-type", "application/json")
	if err != nil {
		app.errorJSON(wr, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		app.errorJSON(wr, err, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(wr, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(wr, err, http.StatusUnauthorized)
		return
	}

	app.infoLog.Print("Authenticated succesfully")

	//dial a connection with rpc server
	RPCclient, err := rpc.Dial("tcp", app.rpcApi)
	if err != nil {
		app.errorJSON(wr, err)
		return
	}

	//calling getdata function in the remote rpc server
	var data HumanReadableData
	err = RPCclient.Call("RPCServer.GetDataRPC", "", &data)
	if err != nil {
		app.errorJSON(wr, err)
		return
	}

	//sending data back
	jsonPayload.Error = false
	jsonPayload.Message = "Data Delivered"
	jsonPayload.Data = data
	app.writeJSON(wr, http.StatusAccepted, jsonPayload)
}
