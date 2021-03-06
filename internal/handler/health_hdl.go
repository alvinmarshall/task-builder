package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthHandler struct {
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Ping(context echo.Context) error {
	body := map[string]interface{}{
		"message": "application is online",
	}
	return context.JSONPretty(http.StatusOK, body, " ")
}
