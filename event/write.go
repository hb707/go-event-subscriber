package event

import (
	"context"
	"event/config"
	"event/types"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Catch struct {
	config *config.Config
	client *ethclient.Client

	needToCatchEvent map[common.Hash]types.NeedToCatchEvent
}

func NewCatch(config *config.Config, client *ethclient.Client, eventChan chan []ethTypes.Log) (*Catch, error) {
	c := &Catch{
		config: config,
		client: client,
	}

	// 캐치해야하는 이벤트 정의!
	// Transfer이벤트 : Transfer(address, address, uint256)에 대한 해시값에 대해 transfer함수를 할당함
	c.needToCatchEvent = map[common.Hash]types.NeedToCatchEvent{
		common.BytesToHash(crypto.Keccak256([]byte("Transfer(address,address,uint256)"))) : {
			NeedToCatchEventFunc: c.Transfer,
		},
	}

	go c.startToCatch(eventChan)

	return c, nil
}

func (c *Catch) Transfer(e *ethTypes.Log, tx *ethTypes.Transaction) {
	// fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", tx)
}

func (c *Catch) startToCatch(events <- chan []ethTypes.Log) {
	for event := range events {
		ctx := context.Background()
		txList := make(map[common.Hash]*ethTypes.Transaction)

		for _, e := range event {
			hash := e.TxHash

			if _, ok := txList[hash]; !ok {
				if tx, pending, err := c.client.TransactionByHash(ctx, hash); err == nil {
					if !pending {
						txList[hash] = tx
					}
				} else {
					continue
				}
			}

			// reverted tx 제외
			if e.Removed {
				continue
			} else if et, ok := c.needToCatchEvent[e.Topics[0]]; !ok {
				// @TODO : 캐치하지 않는 이벤트 발생시 로그로만 남겨두기
			} else {
				// 캐치할 이벤트 발생시 NeedToCatchEventFunc 함수 실행
				et.NeedToCatchEventFunc(&e, txList[hash])
			}
		}
	}
}

func (c *Catch) GetEventToCatch() []common.Hash {
	eventsToCatch := make([]common.Hash, 0)

	for e := range c.needToCatchEvent {
		eventsToCatch = append(eventsToCatch, e)
	}
	return eventsToCatch
}	