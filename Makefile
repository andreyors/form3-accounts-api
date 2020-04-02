dc=docker-compose

.DEFAULT_GOAL := help

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: up
up: ## Starts local env
	${dc} up -d --build

.PHONY: down
down: ## Stops local env
	${dc} down

.PHONY: ps
ps: ## Shows list of containers
	${dc} ps

.PHONY: log
log: ## Listen to logs
	${dc} logs -f app

.PHONY: ssh
ssh: ## Enters local env
	${dc} run app ash

.PHONY: test
test: ## Run tests
	${dc} run app go test -v ./...

.PHONY: e2e
e2e: ## Run e2e
	${dc} run app go test -v cmd/form3/form3_e2e_test.go

.PHONY: fix
fix: ## Run black
	${dc} run app gofmt -w cmd

.PHONY: cov
cov: ## Run coverage
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html
