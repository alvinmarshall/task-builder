package health

import "github.com/labstack/echo/v4"

func Bind(api *echo.Group) {
	healthRoutes := api.Group("/health")
	{
		h := NewHealthHandler()
		healthRoutes.GET("", h.Ping)
	}
}
