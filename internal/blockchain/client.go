package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

	// addr := common.HexToAddress(contractAddr)
	addr := common.HexToAddress(strings.ToLower(contractAddr))
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
	Id              *big.Int
	Title           string
	Description     string
	Creator         common.Address
	StartTime       *big.Int
	EndTime         *big.Int
	Finalized       bool
	TotalVotes      *big.Int
	CommitDeadline  *big.Int
	RevealDeadline  *big.Int
	DepositRequired *big.Int
	Phase           uint8
}, error) {
	opts := &bind.CallOpts{Context: ctx}
	return c.voting.Proposals(opts, new(big.Int).SetUint64(id))
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
			"id":               p.Id.Uint64(),
			"title":            p.Title,
			"description":      p.Description,
			"creator":          p.Creator.Hex(),
			"start_time":       p.StartTime.Uint64(),
			"end_time":         p.EndTime.Uint64(),
			"finalized":        p.Finalized,
			"total_votes":      p.TotalVotes.Uint64(),
			"commit_deadline":  p.CommitDeadline.Uint64(),
			"reveal_deadline":  p.RevealDeadline.Uint64(),
			"deposit_required": p.DepositRequired.String(),
			"phase":            p.Phase,
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

// Commit — отправляет транзакцию commit с хешем голоса и депозитом
func (c *Client) Commit(ctx context.Context, privateKeyHex string, proposalID uint64, commitHash [32]byte, depositWei *big.Int) (string, error) {
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
	opts.Value = depositWei

	tx, err := c.voting.Commit(opts, new(big.Int).SetUint64(proposalID), commitHash)
	if err != nil {
		return "", fmt.Errorf("commit tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// Reveal — раскрывает голос, верификация происходит в контракте
func (c *Client) Reveal(ctx context.Context, privateKeyHex string, proposalID, candidateID uint64, salt [32]byte) (string, error) {
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

	tx, err := c.voting.Reveal(opts,
		new(big.Int).SetUint64(proposalID),
		new(big.Int).SetUint64(candidateID),
		salt,
	)
	if err != nil {
		return "", fmt.Errorf("reveal tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// SlashNoReveal — срезает депозит участника не раскрывшего голос
func (c *Client) SlashNoReveal(ctx context.Context, privateKeyHex string, proposalID uint64, voter common.Address) (string, error) {
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

	tx, err := c.voting.SlashNoReveal(opts, new(big.Int).SetUint64(proposalID), voter)
	if err != nil {
		return "", fmt.Errorf("slash tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// AdvancePhase — переводит голосование в следующую фазу
func (c *Client) AdvancePhase(ctx context.Context, privateKeyHex string, proposalID uint64) (string, error) {
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

	tx, err := c.voting.AdvancePhase(opts, new(big.Int).SetUint64(proposalID))
	if err != nil {
		return "", fmt.Errorf("advance phase tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// GetPhase — возвращает текущую фазу голосования (0=Commit, 1=Reveal, 2=Finalized)
func (c *Client) GetPhase(ctx context.Context, proposalID uint64) (uint8, error) {
	opts := &bind.CallOpts{Context: ctx}
	p, err := c.voting.Proposals(opts, new(big.Int).SetUint64(proposalID))
	if err != nil {
		return 0, fmt.Errorf("get proposal: %w", err)
	}
	return p.Phase, nil
}

// GetProposalInfo — возвращает полную информацию о голосовании включая поля Stage 2
func (c *Client) GetProposalInfo(ctx context.Context, id uint64) (map[string]any, error) {
	opts := &bind.CallOpts{Context: ctx}
	info, err := c.voting.GetProposalInfo(opts, new(big.Int).SetUint64(id))
	if err != nil {
		return nil, fmt.Errorf("get proposal info: %w", err)
	}

	phaseNames := map[uint8]string{0: "commit", 1: "reveal", 2: "finalized"}

	return map[string]any{
		"id":               info.Id.Uint64(),
		"title":            info.Title,
		"description":      info.Description,
		"creator":          info.Creator.Hex(),
		"start_time":       info.StartTime.Uint64(),
		"commit_deadline":  info.CommitDeadline.Uint64(),
		"reveal_deadline":  info.RevealDeadline.Uint64(),
		"deposit_required": info.DepositRequired.String(),
		"phase":            phaseNames[info.Phase],
		"finalized":        info.Finalized,
		"total_votes":      info.TotalVotes.Uint64(),
	}, nil
}

// GetProposalCommitters — возвращает адреса, сделавшие commit для голосования.
// Читает события CommitMade из блокчейна.
func (c *Client) GetProposalCommitters(ctx context.Context, proposalID uint64) ([]string, error) {
	parsedABI, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}

	eventID := parsedABI.Events["CommitMade"].ID
	proposalTopic := common.BigToHash(new(big.Int).SetUint64(proposalID))
	
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics:    [][]common.Hash{{eventID}, {proposalTopic}},
	}

	logs, err := c.eth.FilterLogs(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("filter logs: %w", err)
	}

	seen := make(map[common.Address]bool)
	for _, l := range logs {
		if len(l.Topics) < 3 {
			continue
		}
		// Topics[2] = voter address (indexed)
		voter := common.HexToAddress(l.Topics[2].Hex())
		seen[voter] = true
	}

	result := make([]string, 0, len(seen))
	for addr := range seen {
		result = append(result, addr.Hex())
	}
	return result, nil
}