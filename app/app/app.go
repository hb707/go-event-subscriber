package app

import (
	"event/config"
	"event/repository"
)

type App struct{
	config *config.Config

	repository *repository.Repository
}

func NewApp(config *config.Config) *App {
	a := App{
		config: config,
	}

	var err error
	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	return &a
}