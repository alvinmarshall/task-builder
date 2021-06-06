package rest

import (
	"taskbuilder/internal/server/rest/auth"
	"taskbuilder/internal/server/rest/health"
	"taskbuilder/internal/server/rest/task"
	"taskbuilder/internal/server/rest/user"
)

func (s *server) MapRoutes() {
	apiV1 := s.engine.Group("api/v1")
	health.Bind(apiV1)
	task.Bind(apiV1, s.container)
	user.Bind(apiV1, s.container)
	auth.Bind(apiV1, s.container)
}
