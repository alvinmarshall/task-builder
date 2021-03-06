package server

import (
	"github.com/labstack/echo/v4"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/handler"
)

func (s *server) MapRoutes() {
	apiV1 := s.engine.Group("api/v1")
	s.healthRoutes(apiV1)
	s.taskRoutes(apiV1)
	handler.UserRouter(apiV1, s.container)

}

func (s *server) healthRoutes(api *echo.Group) {
	healthRoutes := api.Group("/health")
	{
		h := handler.NewHealthHandler()
		healthRoutes.GET("", h.Ping)
	}
}

func (s *server) taskRoutes(api *echo.Group) {
	var taskSvc port.TaskService
	s.container.Invoke(func(svc port.TaskService) { taskSvc = svc })
	taskRoutes := api.Group("/tasks")
	{
		h := handler.NewTaskHandler(taskSvc)
		taskRoutes.POST("", h.Create)
		taskRoutes.GET("", h.GetAll)
		taskRoutes.GET("/:id", h.Get)
	}
}
