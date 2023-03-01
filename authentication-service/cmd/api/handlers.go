package main

import (
	"errors"
	"net/http"
)

func (app *Config) Authenticate(rw http.ResponseWriter, r *http.Request) {

	//getting values from body
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(rw, r, &requestPayload)
	app.infoLog.Print(requestPayload)
	if err != nil {
		app.errorJSON(rw, err, http.StatusBadRequest)
		return
	}

	//validating the user
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	app.infoLog.Print(user)
	if err != nil {
		app.errorJSON(rw, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(rw, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Success",
	}

	//writing response
	app.writeJSON(rw, http.StatusAccepted, payload)
}
