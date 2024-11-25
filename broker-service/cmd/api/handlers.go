package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Broker message",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
