# ─── Переменные ───────────────────────────────────────────────────────
RPC_URL         ?= http://127.0.0.1:8545
CONTRACT_ADDRESS ?= $(shell cat deploy.json | jq -r .address)

# ─── Hardhat ──────────────────────────────────────────────────────────

## Запустить локальную ноду Hardhat
node:
	npx hardhat node

## Скомпилировать контракт
compile:
	npx hardhat compile

## Задеплоить контракт на локальную сеть
deploy:
	npx hardhat run scripts/deploy.js --network localhost

## Создать тестовые данные (proposal + кандидаты)
seed:
	npx hardhat run scripts/seed.js --network localhost

## Запустить тесты контракта
test-contract:
	npx hardhat test

# ─── ABI + биндинг ────────────────────────────────────────────────────

## Экспортировать ABI из артефактов
abi:
	node -e "const a=require('./artifacts/contracts/Voting.sol/Voting.json');require('fs').writeFileSync('Voting.abi',JSON.stringify(a.abi,null,2))"

## Сгенерировать Go-биндинг из ABI
gen:
	abigen --abi=Voting.abi --pkg=blockchain --type=Voting --out=internal/blockchain/voting.go

## Полный цикл: компиляция + ABI + биндинг
build-contract: compile abi gen

# ─── Go ───────────────────────────────────────────────────────────────

## Собрать Go-проект
build:
	go build ./...

## Запустить Go-сервер
server:
	RPC_URL=$(RPC_URL) CONTRACT_ADDRESS=$(CONTRACT_ADDRESS) go run cmd/server/main.go

## Запустить тесты Go
test-go:
	go test ./... -v

## Подтянуть зависимости
tidy:
	go mod tidy

# ─── Полный цикл разработки ───────────────────────────────────────────

## Деплой + seed + запуск сервера
dev: deploy seed server

## Показать все доступные команды
help:
	@grep -E '^##' Makefile | sed 's/## //'

.PHONY: node compile deploy seed test-contract abi gen build-contract build server test-go tidy dev help