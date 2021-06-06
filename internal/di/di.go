package di

import (
	"go.uber.org/dig"
	"taskbuilder/internal/config"
	"taskbuilder/internal/core/service"
	"taskbuilder/internal/logger"
	"taskbuilder/internal/storage"
	"taskbuilder/internal/storage/orm"
	"taskbuilder/internal/types"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// data source
	container.Provide(storage.NewDataSource)

	// logger
	container.Provide(logger.NewLogger)

	// task
	container.Provide(orm.NewTaskRepo)
	container.Provide(service.NewTaskService)

	// user
	container.Provide(orm.NewUserRepo)
	container.Provide(service.NewUserService)

	// JwtWrapper
	container.Provide(types.NewJwtWrapper)
	// jwt
	container.Provide(service.NewAuthService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
