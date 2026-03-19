package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	eth     *ethclient.Client
	voting  *Voting
	address common.Address
}

func NewClient(rpcURL, contractAddr string) (*Client, error) {
	eth, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	addr := common.HexToAddress(contractAddr)
	voting, err := NewVoting(addr, eth)
	if err != nil {
		return nil, err
	}

	return &Client{eth: eth, voting: voting, address: addr}, nil
}

func (c *Client) GetResults(ctx context.Context, proposalID uint64) ([]uint64, []uint64, error) {
	opts := &bind.CallOpts{Context: ctx}
	result, err := c.voting.GetResults(opts, new(big.Int).SetUint64(proposalID))
	if err != nil {
		return nil, nil, err
	}
	return toUint64Slice(result.Ids), toUint64Slice(result.Votes), nil
}

func (c *Client) HasVoted(ctx context.Context, proposalID uint64, voter common.Address) (bool, error) {
	opts := &bind.CallOpts{Context: ctx}
	return c.voting.HasVoted(opts, new(big.Int).SetUint64(proposalID), voter)
}

func toUint64Slice(in []*big.Int) []uint64 {
	out := make([]uint64, len(in))
	for i, v := range in {
		out[i] = v.Uint64()
	}
	return out
}

// GetProposal — читает данные голосования напрямую из контракта по ID
func (c *Client) GetProposal(ctx context.Context, id uint64) (struct {
	Id          *big.Int
	Title       string
	Description string
	Creator     common.Address
	StartTime   *big.Int
	EndTime     *big.Int
	Finalized   bool
	TotalVotes  *big.Int
}, error) {
	opts := &bind.CallOpts{Context: ctx}
	return c.voting.Proposals(opts, new(big.Int).SetUint64(id))
}

// Vote - отправляет транзакцию с голосом за кандидата в конкретном голосовании
func (c *Client) Vote(ctx context.Context, privateKeyHex string, proposalID, candidateID uint64) (string, error) {
	// Убираем префикс "0x", если он есть
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// парсим приватный ключ
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("Invalid private key: %w", err)
	}

	// Получаем chainID от ноды
	chainID, err := c.eth.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("get chainID: %w", err)
	}

	// Создаем opts для подписи транзакции
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("create transactor: %w", err)
	}

	// Отправляем транзакцию в контракт
	tx, err := c.voting.Vote(opts,
		new(big.Int).SetUint64(proposalID),
		new(big.Int).SetUint64(candidateID),
	)
	if err != nil {
		return "", fmt.Errorf("send vote transaction: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// GetAllProposals - возвращает все голосования по счетчику proposalCount
func (c *Client) GetAllProposals(ctx context.Context) ([]map[string]any, error) {
	opts := &bind.CallOpts{Context: ctx}

	count, err := c.voting.ProposalCount(opts)
	if err != nil {
		return nil, fmt.Errorf("get proposal count: %w", err)
	}

	proposals := make([]map[string]any, 0, count.Uint64())
	for i := uint64(1); i <= count.Uint64(); i++ {
		p, err := c.voting.Proposals(opts, new(big.Int).SetUint64(i))
		if err != nil {
			return nil, fmt.Errorf("get proposal %d; %w", i, err)
		}
		proposals = append(proposals, map[string]any{
			"id":          p.Id.Uint64(),
			"title":       p.Title,
			"description": p.Description,
			"creator":     p.Creator.Hex(),
			"start_time":  p.StartTime.Uint64(),
			"end_time":    p.EndTime.Uint64(),
			"finalized":   p.Finalized,
			"total_votes": p.TotalVotes.Uint64(),
		})
	}
	return proposals, nil
}

// GetCandidates - возвращает всех кандидатов голосования по ID
func (c *Client) GetCandidates(ctx context.Context, proposalID uint64) ([]map[string]any, error) {
	opts := &bind.CallOpts{Context: ctx}

	count, err := c.voting.CandidateCount(opts, new(big.Int).SetUint64(proposalID))
	if err != nil {
		return nil, fmt.Errorf("get condidate count: %w", err)
	}

	candidates := make([]map[string]any, 0, count.Uint64())
	for i := uint64(1); i <= count.Uint64(); i++ {
		cand, err := c.voting.Candidates(opts,
			new(big.Int).SetUint64(proposalID),
			new(big.Int).SetUint64(i),
		)
		if err != nil {
			return nil, fmt.Errorf("get candidate %d: %w", i, err)
		}
		candidates = append(candidates, map[string]any{
			"id":         cand.Id.Uint64(),
			"name":       cand.Name,
			"vote_count": cand.VoteCount.Uint64(),
		})
	}
	return candidates, nil
}

// FinalizeProposal — финализирует голосование после endTime
func (c *Client) FinalizeProposal(ctx context.Context, privateKeyHex string, proposalID uint64) (string, error) {
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	chainID, err := c.eth.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("get chainID: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("create transactor: %w", err)
	}
	opts.Context = ctx

	tx, err := c.voting.FinalizeProposal(opts, new(big.Int).SetUint64(proposalID))
	if err != nil {
		return "", fmt.Errorf("finalize tx: %w", err)
	}

	return tx.Hash().Hex(), nil
}
