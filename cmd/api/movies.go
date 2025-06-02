package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/looksaw/greenlight/internal/data"
	"github.com/looksaw/greenlight/internal/validator"
)

type envelope map[string]any

// check health方法
func (app *application) checkHealth(w http.ResponseWriter, r *http.Request) {
	type healthCheckResponse struct {
		Status  string `json:"status"`
		Env     string `json:"env"`
		Version string `json:"version"`
	}
	data := &healthCheckResponse{
		Status:  "available",
		Env:     app.Config.env,
		Version: version,
	}
	err := app.writeJSON(w, http.StatusOK, envelope{"health_info": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// create app方法
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
	// 开始验证
	v := validator.New()
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: data.Runtime(input.Runtime),
		Genres:  input.Genres,
	}
	data.ValidateMovie(v, movie)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v/n", input)
}

// show id的方法
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notfoundErrorResponse(w, r)
	}
	movie := data.Movie{
		ID:       id,
		CreateAt: time.Now(),
		Title:    "Casablanca",
		Runtime:  102,
		Genres:   []string{"drama", "romance", "war"},
		Version:  1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
