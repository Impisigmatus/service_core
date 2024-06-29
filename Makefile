PROJECT := $(shell cat go.mod | grep module | awk -F ' ' '{print $$2}' | awk -F '/' '{print $$NF}')

REPORTS := .reports
VENDOR := vendor

GOLANGCI-LINT := github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(REPORTS):
	mkdir -p $(REPORTS)

help:
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: $(REPORTS) ## Настроить рабочее окружение

clean: ## Очистить рабочее окружение
	rm -rf $(VENDOR)
	rm -rf $(REPORTS)
	go clean -r -i -testcache -modcache

mock-gen: setup ## Генерация mock
	go generate ./...

format: setup ## Запуск форматирования исходного кода
	go run $(GOLANGCI-LINT) run ./... --fix

lint: setup ## Запуск линтеров
	go run $(GOLANGCI-LINT) run ./... --out-format checkstyle > $(REPORTS)/golangci-lint.xml

tests: mock-gen ## Запуск авто-тестов
	go test -race -cover -coverprofile $(REPORTS)/coverage.out ./...
	go tool cover -html=$(REPORTS)/coverage.out -o $(REPORTS)/coverage.html

all: format lint tests ## Последовательный запуск основных команд

.DEFAULT_GOAL := help
