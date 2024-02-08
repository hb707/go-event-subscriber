package app

import (
	"event/config"
	"event/event"
	"event/repository"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type App struct{
	config *config.Config

	client *ethclient.Client

	repository *repository.Repository
	scan *event.Scan
	catch *event.Catch
}

func NewApp(config *config.Config) *App {
	a := App{
		config: config,
	}

	var err error
	var eventChan chan []types.Log

	// client는 app에 추가될 다른 패키지 모듈에 추가
	if a.client, err = ethclient.Dial(config.Node.Uri); err != nil {
		panic(err)
	} else {
		if a.repository, err = repository.NewRepository(config); err != nil {
			panic(err)
		}
		if a.scan, eventChan, err = event.NewScan(config, a.client); err != nil {
			panic(err)
		}
		if a.catch, err = event.NewCatch(config, a.client, eventChan); err != nil {
			panic(err)
		}

	}

	return &a
}