package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const batchSize = uint64(500)

type Indexer struct {
	eth         *ethclient.Client
	addr        common.Address
	lastBlock   uint64
	blockHashes map[uint64]common.Hash
	voting      *Voting
}

func NewIndexer(eth *ethclient.Client, addr common.Address, voting *Voting) *Indexer {
	return &Indexer{
		eth:         eth,
		addr:        addr,
		lastBlock:   0,
		blockHashes: make(map[uint64]common.Hash),
		voting:      voting,
	}
}

func (idx *Indexer) Sync(ctx context.Context) error {
	latest, err := idx.eth.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("get latest block: %w", err)
	}

	for from := idx.lastBlock; from <= latest; from += batchSize {
		to := min(from+batchSize-1, latest)
		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(from),
			ToBlock:   new(big.Int).SetUint64(to),
			Addresses: []common.Address{idx.addr},
		}
		logs, err := idx.eth.FilterLogs(ctx, query)
		if err != nil {
			return fmt.Errorf("batch %d-%d: %w", from, to, err)
		}

		for _, l := range logs {
			if stored, ok := idx.blockHashes[l.BlockNumber]; ok {
				if stored != l.BlockHash {
					idx.reindex(ctx, l.BlockNumber)
				}
			}
			idx.blockHashes[l.BlockNumber] = l.BlockHash
			idx.dispatch(l)
		}
		idx.lastBlock = to + 1
	}

	return idx.subscribeNewBlocks(ctx)
}

func (idx *Indexer) dispatch(l types.Log) {
	// Определяем тип события по первому топику
	if len(l.Topics) == 0 {
		return
	}

	abi, err := VotingMetaData.GetAbi()
	if err != nil {
		return
	}

	event, err := abi.EventByID(l.Topics[0])
	if err != nil {
		return
	}

	switch event.Name {
	case "VoteCast":
		fmt.Printf("[event] VoteCast block=%d tx=%s\n", l.BlockNumber, l.TxHash.Hex())
	case "ProposalCreated":
		fmt.Printf("[event] ProposalCreated block=%d tx=%s\n", l.BlockNumber, l.TxHash.Hex())
	case "CandidateAdded":
		fmt.Printf("[event] CandidateAdded block=%d tx=%s\n", l.BlockNumber, l.TxHash.Hex())
	case "ProposalFinalized":
		fmt.Printf("[event] ProposalFinalized block=%d tx=%s\n", l.BlockNumber, l.TxHash.Hex())
	}
}

func (idx *Indexer) reindex(ctx context.Context, fromBlock uint64) {
	fmt.Printf("[reorg] detected at block %d, reindexing...\n", fromBlock)
	idx.lastBlock = fromBlock
}

func (idx *Indexer) subscribeNewBlocks(ctx context.Context) error {
	headers := make(chan *types.Header)
	sub, err := idx.eth.SubscribeNewHead(ctx, headers)
	if err != nil {
		// Hardhat не всегда поддерживает подписки — не фатально
		fmt.Println("[indexer] subscription not available, polling mode only")
		return nil
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				fmt.Printf("[indexer] subscription error: %v\n", err)
				return
			case header := <-headers:
				fmt.Printf("[indexer] new block: %d\n", header.Number.Uint64())
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}