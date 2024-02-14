package app

import (
	"event/config"
	"event/event"
	"event/repository"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
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
	eventLog := make(chan []ethTypes.Log, 100)


	// client는 app에 추가될 다른 패키지 모듈에 추가
	if a.client, err = ethclient.Dial(config.Node.Uri); err != nil {
		panic(err)
	} else {
		if a.repository, err = repository.NewRepository(config); err != nil {
			panic(err)
		}

		if a.catch, err = event.NewCatch(config, a.client, eventLog); err != nil {
			panic(err)
		}

		if a.scan, _, err = event.NewScan(config, a.client, a.catch.GetEventToCatch(), eventLog); err != nil {
			panic(err)
		}
	}

	for {
	}
}