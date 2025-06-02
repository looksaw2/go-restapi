package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{
		"err": message,
	}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Println(err)
	msg := "Some error occur in the server internal"
	app.errorResponse(w, r, http.StatusInternalServerError, msg)
}

func (app *application) notfoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := "the page was not found"
	app.errorResponse(w, r, http.StatusNotFound, msg)
}

func (app *application) methodNotAllowErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := "This method not allowed"
	app.errorResponse(w, r, http.StatusMethodNotAllowed, msg)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, err map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, err)
}
