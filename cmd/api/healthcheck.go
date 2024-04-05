package main

import (
	"net/http"
)

type payload struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	pl := payload{
		Status:      "available",
		Environment: "development",
		Version:     "1.0.0",
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"healthdata": pl}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
