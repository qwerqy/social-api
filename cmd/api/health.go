package main

import (
	"net/http"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
	Env string `json:"env"`
	Version string `json:"version"`
}

// HealthCheck godoc
//	@Summary		HealthCheck
//	@Description	HealtCheck
//	@Tags			ops
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	HealthCheckResponse
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter,r *http.Request) {
	data := HealthCheckResponse{
		Status: "ok",
		Env: app.config.env,
		Version: version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w,r,err)
	}
}