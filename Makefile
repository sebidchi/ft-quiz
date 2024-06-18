# Go parameters
MAIN_PATH=cmd/quiz/

.PHONY: setup
setup: SHELL:=/bin/bash
setup: ## Setup Project
	[ -f ./.env ] || cp .env.dist .env

.PHONY: test
test: ## Run tests
	go test ./... -race

.PHONY: ci-coverage
ci-coverage: SHELL:=/bin/bash
ci-coverage:
	$(MAKE) arrange
	@echo "==> Checking test coverage..."
	go install github.com/kyoh86/richgo@latest
	@richgo test -failfast -coverprofile=coverage.out ./... -coverpkg ./...
	@go tool cover -html=coverage.out -o coverage.html

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
