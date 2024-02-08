package event

import (
	"context"
	"event/config"
	"event/types"

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
		common.BytesToHash(crypto.Keccak256([]byte("Transfer(address, address, uint256)"))) : {
			NeedToCatchEventFunc: c.Transfer,
		},
	}

	go c.startToCatch(eventChan)

	return c, nil
}

func (c *Catch) Transfer(e *ethTypes.Log, tx *ethTypes.Transaction) {}

func (c *Catch) startToCatch(events <- chan []ethTypes.Log) {
	for event := range events {
		ctx := context.Background()
		txList := make(map[common.Hash]*ethTypes.Transaction)

		for _, e := range event {
			if _, ok := txList[e.TxHash]; !ok {
				if tx, pending, err := c.client.TransactionByHash(ctx, e.TxHash); err != nil {
					if !pending {
						txList[e.TxHash] = tx
					}
				}
			}
		}
	}
}