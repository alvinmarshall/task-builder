package migration

import "taskbuilder/internal/core/domain"

var models = []interface{}{
	&domain.User{},
	&domain.Task{},
	&domain.Role{},
}
