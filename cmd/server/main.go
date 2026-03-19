package main

import (
	"log"
	"net/http"
	"os"

	"voting-dapp/internal/api"
	"voting-dapp/internal/blockchain"
)

func main() {
	rpcURL := getEnv("RPC_URL", "http://127.0.0.1:8545")
	contractAddr := getEnv("CONTRACT_ADDRESS", "")
	if contractAddr == "" {
		log.Fatal("CONTRACT_ADDRESS is required")
	}

	// Подключаемся к контракту через go-ethereum
	client, err := blockchain.NewClient(rpcURL, contractAddr)
	if err != nil {
		log.Fatalf("blockchain client: %v", err)
	}

	// Настраиваем роутер и запускаем сервер
	r := api.SetupRouter(client)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}