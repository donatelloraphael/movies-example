package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/debug/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/users", app.registerUserHandler)

	router.HandlerFunc(http.MethodPost, "/user/login", app.createAuthenticationHandler)

	return router
}
