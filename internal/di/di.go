package di

import (
	"go.uber.org/dig"
	"taskbuilder/internal/config"
	"taskbuilder/internal/logger"
	"taskbuilder/internal/storage"
	"taskbuilder/internal/storage/orm"
	"taskbuilder/internal/task"
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
	container.Provide(task.NewTaskService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
