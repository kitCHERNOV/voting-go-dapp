package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// VerificationResult — результат независимой верификации голосования.
type VerificationResult struct {
	ProposalID       uint64            `json:"proposal_id"`
	Valid            bool              `json:"valid"`
	TotalVotes       uint64            `json:"total_votes"`
	CandidateVotes   map[uint64]uint64 `json:"candidate_votes"`
	WinnerID         uint64            `json:"winner_id"`
	ContractWinnerID uint64            `json:"contract_winner_id"`
	Discrepancies    []string          `json:"discrepancies"`
}

// VerifyProposal читает события из блокчейна и независимо
// пересчитывает результаты голосования.
func (c *Client) VerifyProposal(ctx context.Context, proposalID uint64) (*VerificationResult, error) {
	result := &VerificationResult{
		ProposalID:     proposalID,
		CandidateVotes: make(map[uint64]uint64),
		Discrepancies:  []string{},
	}

	parsedABI, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}

	proposalTopic := common.BigToHash(new(big.Int).SetUint64(proposalID))

	// Шаг 1: читаем CommitMade — кто сделал commit
	commitEventID := parsedABI.Events["CommitMade"].ID
	commitLogs, err := c.eth.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics:    [][]common.Hash{{commitEventID}, {proposalTopic}},
	})
	if err != nil {
		return nil, fmt.Errorf("filter commit logs: %w", err)
	}

	committedVoters := make(map[common.Address]bool)
	for _, l := range commitLogs {
		if len(l.Topics) < 3 {
			continue
		}
		voter := common.HexToAddress(l.Topics[2].Hex())
		committedVoters[voter] = true
	}

	// Шаг 2: читаем VoteRevealed — реальные голоса
	revealEventID := parsedABI.Events["VoteRevealed"].ID
	revealLogs, err := c.eth.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics:    [][]common.Hash{{revealEventID}, {proposalTopic}},
	})
	if err != nil {
		return nil, fmt.Errorf("filter reveal logs: %w", err)
	}

	seenVoters := make(map[common.Address]bool)
	for _, l := range revealLogs {
		if len(l.Topics) < 4 {
			continue
		}
		voter := common.HexToAddress(l.Topics[2].Hex())
		candidateID := new(big.Int).SetBytes(l.Topics[3].Bytes()).Uint64()

		// Проверка 1: reveal без commit
		if !committedVoters[voter] {
			result.Discrepancies = append(result.Discrepancies,
				fmt.Sprintf("reveal without commit: voter %s", voter.Hex()))
		}

		// Проверка 2: двойное голосование
		if seenVoters[voter] {
			result.Discrepancies = append(result.Discrepancies,
				fmt.Sprintf("duplicate reveal: voter %s", voter.Hex()))
			continue
		}
		seenVoters[voter] = true

		result.CandidateVotes[candidateID]++
		result.TotalVotes++
	}

	// Шаг 3: определяем победителя по пересчитанным голосам
	var maxVotes uint64
	for candidateID, votes := range result.CandidateVotes {
		if votes > maxVotes {
			maxVotes = votes
			result.WinnerID = candidateID
		}
	}

	// Шаг 4: сравниваем с состоянием контракта
	contractInfo, err := c.GetProposalInfo(ctx, proposalID)
	if err != nil {
		return nil, fmt.Errorf("get proposal info: %w", err)
	}

	contractTotalVotes := contractInfo["total_votes"].(uint64)
	if contractTotalVotes != result.TotalVotes {
		result.Discrepancies = append(result.Discrepancies,
			fmt.Sprintf("total votes mismatch: contract=%d, recomputed=%d",
				contractTotalVotes, result.TotalVotes))
	}

	// Шаг 5: победитель из события ProposalFinalized
	finalizeEventID := parsedABI.Events["ProposalFinalized"].ID
	finalizeLogs, err := c.eth.FilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{c.address},
		Topics:    [][]common.Hash{{finalizeEventID}, {proposalTopic}},
	})
	if err != nil {
		return nil, fmt.Errorf("filter finalize logs: %w", err)
	}

	if len(finalizeLogs) > 0 && len(finalizeLogs[0].Data) >= 32 {
		winnerBytes := finalizeLogs[0].Data[:32]
		result.ContractWinnerID = new(big.Int).SetBytes(winnerBytes).Uint64()

		if result.TotalVotes > 0 && result.ContractWinnerID != result.WinnerID {
			result.Discrepancies = append(result.Discrepancies,
				fmt.Sprintf("winner mismatch: contract=%d, recomputed=%d",
					result.ContractWinnerID, result.WinnerID))
		}
	}

	result.Valid = len(result.Discrepancies) == 0
	return result, nil
}