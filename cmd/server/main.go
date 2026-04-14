package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"voting-dapp/internal/api"
	"voting-dapp/internal/blockchain"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	rpcURL := getEnv("RPC_URL", "http://127.0.0.1:8545")

	// Voting contract
	contractAddr := strings.TrimSpace(getEnv("CONTRACT_ADDRESS", ""))
	if contractAddr == "" {
		log.Fatal("CONTRACT_ADDRESS is required")
	} else if !common.IsHexAddress(contractAddr) {
		log.Fatalf("invalid contract address: %q", contractAddr)
	}

	// VoterRegistry contract (Stage 3)
	registryAddr := strings.TrimSpace(getEnv("REGISTRY_ADDRESS", ""))
	if registryAddr == "" {
		log.Fatal("REGISTRY_ADDRESS is required")
	} else if !common.IsHexAddress(registryAddr) {
		log.Fatalf("invalid registry address: %q", registryAddr)
	}

	client, err := blockchain.NewClient(rpcURL, contractAddr)
	if err != nil {
		log.Fatalf("blockchain client: %v", err)
	}

	registryClient, err := blockchain.NewRegistryClient(rpcURL, registryAddr)
	if err != nil {
		log.Fatalf("registry client: %v", err)
	}

	r := api.SetupRouter(client, registryClient)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}