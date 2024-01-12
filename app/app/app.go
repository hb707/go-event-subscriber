package app

import (
	"event/config"
	"event/event"
	"event/repository"
)

type App struct{
	config *config.Config

	repository *repository.Repository
	scan *event.Scan
}

func NewApp(config *config.Config) *App {
	a := App{
		config: config,
	}

	var err error
	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}
	if a.scan, err = event.NewScan(config); err != nil {
		panic(err)
	}

	return &a
}