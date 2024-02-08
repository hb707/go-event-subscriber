package event

import (
	"context"
	"event/config"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 이더리움 클라이언트와 연결하여 이벤트를 가져오는 역할
type Scan struct {
	config *config.Config

	FilterQuery ethereum.FilterQuery

	client *ethclient.Client
}

func NewScan(config *config.Config, client *ethclient.Client) (*Scan, chan []ethTypes.Log, error) {
	s := &Scan{
		config: config,
		client: client,
	}

	eventLog := make(chan []ethTypes.Log, 100)
	go s.lookingScan(config.Node.StartBlock, eventLog) // 백그라운드에서 계속 실행되어야하므로 고루틴 생성
	
	return s, eventLog, nil
}

func (s *Scan) lookingScan(
	startBlock int64, 
	// @TODO : 스캔할 Collection,
	// @TODO : 캐치할 Event,
	eventLog chan <- []ethTypes.Log, // 스캔한 이벤트를 담아서 보낼 채널
) {
	startReadBlock, to := startBlock, uint64(0)

	// fromBlock, toBlock, Addresses, Topics
	s.FilterQuery = ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics: [][]common.Hash{},
		FromBlock: big.NewInt(startReadBlock),
	}

	for {
		time.Sleep(1e8) // 일정주기를 두고 go루틴 실행

		ctx := context.Background()
		if maxBlock, err := s.client.BlockNumber(ctx); err != nil {
			fmt.Println("Get BlockNumber", "err", err.Error())
			continue
		} else {
			to = maxBlock
			if to > uint64(startReadBlock) {
				s.FilterQuery.FromBlock = big.NewInt(startReadBlock)
				s.FilterQuery.ToBlock = big.NewInt(int64(to))

				tryCount := 1
				
				Retry: 
				if logs, err := s.client.FilterLogs(ctx, s.FilterQuery); err != nil {
					if tryCount == 3 {
						fmt.Println("fail to get filtered Logs", "err", err.Error())
						break
					} else {
					// From, To 블록만 변경해서 다시 호출
					newTo := big.NewInt(int64(to) - 1)
					newFrom := big.NewInt((startBlock - 1))

					s.FilterQuery.ToBlock = newTo
					s.FilterQuery.FromBlock = newFrom

					tryCount++
					goto Retry

					}
				} else if len(logs) > 0 {
					eventLog <- logs
					startReadBlock = int64(to)
				}
			}
		}
	}

}