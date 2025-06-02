package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	//异常处理
	router.NotFound = http.HandlerFunc(app.notfoundErrorResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowErrorResponse)
	//注册路由
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.checkHealth)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	return router
}
