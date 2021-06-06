package server

import (
	"taskbuilder/internal/handler/health"
	"taskbuilder/internal/handler/task"
	"taskbuilder/internal/handler/user"
)

func (s *server) MapRoutes() {
	apiV1 := s.engine.Group("api/v1")
	health.Bind(apiV1)
	task.Bind(apiV1, s.container)
	user.Bind(apiV1, s.container)

}
