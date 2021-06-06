package rest

import (
	health2 "taskbuilder/internal/server/rest/health"
	task2 "taskbuilder/internal/server/rest/task"
	user2 "taskbuilder/internal/server/rest/user"
)

func (s *server) MapRoutes() {
	apiV1 := s.engine.Group("api/v1")
	health2.Bind(apiV1)
	task2.Bind(apiV1, s.container)
	user2.Bind(apiV1, s.container)

}
