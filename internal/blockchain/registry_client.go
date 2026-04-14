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

// RegistryClient — высокоуровневый клиент для работы с VoterRegistry.
type RegistryClient struct {
	eth      *ethclient.Client
	registry *VoterRegistry
	address  common.Address
}

// NewRegistryClient создаёт клиент реестра по URL ноды и адресу контракта.
func NewRegistryClient(rpcURL, contractAddr string) (*RegistryClient, error) {
	eth, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("dial rpc: %w", err)
	}

	addr := common.HexToAddress(strings.ToLower(contractAddr))
	registry, err := NewVoterRegistry(addr, eth)
	if err != nil {
		return nil, fmt.Errorf("bind registry contract: %w", err)
	}

	return &RegistryClient{eth: eth, registry: registry, address: addr}, nil
}

// IsRegistered — проверяет, зарегистрирован ли адрес.
func (r *RegistryClient) IsRegistered(ctx context.Context, addr common.Address) (bool, error) {
	opts := &bind.CallOpts{Context: ctx}
	return r.registry.IsRegistered(opts, addr)
}

// GetRegisteredCount — считает зарегистрированных участников через события.
func (r *RegistryClient) GetRegisteredCount(ctx context.Context) (uint64, error) {
	parsedABI, err := VoterRegistryMetaData.GetAbi()
	if err != nil {
		return 0, fmt.Errorf("parse abi: %w", err)
	}

	registeredID := parsedABI.Events["VoterRegistered"].ID
	revokedID := parsedABI.Events["VoterRevoked"].ID

	query := ethereum.FilterQuery{
		Addresses: []common.Address{r.address},
		Topics:    [][]common.Hash{{registeredID, revokedID}},
	}

	logs, err := r.eth.FilterLogs(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("filter logs: %w", err)
	}

	registered := make(map[common.Address]bool)
	for _, l := range logs {
		if len(l.Topics) < 2 {
			continue
		}
		voter := common.HexToAddress(l.Topics[1].Hex())
		switch l.Topics[0] {
		case registeredID:
			registered[voter] = true
		case revokedID:
			delete(registered, voter)
		}
	}

	return uint64(len(registered)), nil
}

// Register — регистрирует участника. Требует приватный ключ с REGISTRAR_ROLE.
func (r *RegistryClient) Register(ctx context.Context, privateKey string, voter common.Address) (string, error) {
	opts, err := r.buildTransactOpts(ctx, privateKey)
	if err != nil {
		return "", err
	}
	tx, err := r.registry.Register(opts, voter)
	if err != nil {
		return "", fmt.Errorf("register tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// Revoke — отзывает регистрацию участника. Требует приватный ключ с REGISTRAR_ROLE.
func (r *RegistryClient) Revoke(ctx context.Context, privateKey string, voter common.Address) (string, error) {
	opts, err := r.buildTransactOpts(ctx, privateKey)
	if err != nil {
		return "", err
	}
	tx, err := r.registry.Revoke(opts, voter)
	if err != nil {
		return "", fmt.Errorf("revoke tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// RegisterBatch — регистрирует список участников последовательно.
// Возвращает хеш последней транзакции.
func (r *RegistryClient) RegisterBatch(ctx context.Context, privateKey string, voters []common.Address) (string, error) {
	if len(voters) == 0 {
		return "", fmt.Errorf("voters list is empty")
	}
	var lastHash string
	for _, voter := range voters {
		hash, err := r.Register(ctx, privateKey, voter)
		if err != nil {
			return "", fmt.Errorf("register %s: %w", voter.Hex(), err)
		}
		lastHash = hash
	}
	return lastHash, nil
}

// SelfRegister — самостоятельная регистрация (только если selfRegistrationOpen == true).
func (r *RegistryClient) SelfRegister(ctx context.Context, privateKey string) (string, error) {
	opts, err := r.buildTransactOpts(ctx, privateKey)
	if err != nil {
		return "", err
	}
	tx, err := r.registry.SelfRegister(opts)
	if err != nil {
		return "", fmt.Errorf("self register tx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

// buildTransactOpts — строит TransactOpts из hex-ключа.
func (r *RegistryClient) buildTransactOpts(ctx context.Context, privateKeyHex string) (*bind.TransactOpts, error) {
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	pk, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}
	chainID, err := r.eth.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get chainID: %w", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}
	opts.Context = ctx
	opts.Value = big.NewInt(0)
	return opts, nil
}