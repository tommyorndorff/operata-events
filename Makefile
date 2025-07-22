.PHONY: all build test clean lint fmt vet example help install-tools commit-help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOVET=$(GOCMD) vet

# Build info
BINARY_NAME=operata-events
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date +%FT%T%z)

all: fmt vet lint test ## Run fmt, vet, lint and test

build: ## Build the application
	$(GOBUILD) -v ./...

test: ## Run tests
	$(GOTEST) -v -race -coverprofile=coverage.out ./...

clean: ## Clean build cache
	$(GOCLEAN)
	rm -f coverage.out
	rm -rf bin/

lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	$(GOFMT) -s -w .

vet: ## Run go vet
	$(GOVET) ./...

example: ## Run example
	$(GOCMD) run examples/main.go

coverage: test ## Generate coverage report
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

install-tools: ## Install development tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	npm install -g @commitlint/cli @commitlint/config-conventional
	npm install -g semantic-release @semantic-release/changelog @semantic-release/git @semantic-release/github

mod-tidy: ## Tidy module dependencies
	$(GOMOD) tidy

mod-verify: ## Verify module dependencies
	$(GOMOD) verify

commit-help: ## Show conventional commit format
	@echo "Conventional Commit Format:"
	@echo "  <type>[optional scope]: <description>"
	@echo ""
	@echo "Types:"
	@echo "  feat:     A new feature"
	@echo "  fix:      A bug fix"
	@echo "  docs:     Documentation only changes"
	@echo "  style:    Changes that do not affect the meaning of the code"
	@echo "  refactor: A code change that neither fixes a bug nor adds a feature"
	@echo "  perf:     A code change that improves performance"
	@echo "  test:     Adding missing tests or correcting existing tests"
	@echo "  build:    Changes that affect the build system or external dependencies"
	@echo "  ci:       Changes to our CI configuration files and scripts"
	@echo "  chore:    Other changes that don't modify src or test files"
	@echo ""
	@echo "Scopes: events, examples, utils, ci, deps"
	@echo ""
	@echo "Examples:"
	@echo "  feat(events): add support for new event type"
	@echo "  fix(utils): correct validation logic"
	@echo "  docs: update README with installation instructions"

validate-commit: ## Validate the last commit message
	@if command -v npx >/dev/null 2>&1; then \
		npx commitlint --from HEAD~1 --to HEAD --verbose; \
	else \
		echo "commitlint not installed, skipping commit validation"; \
	fi

setup-git-hooks: ## Setup git commit message template
	git config commit.template .gitmessage

check-deps: ## Check for outdated dependencies
	$(GOCMD) list -u -m all

ci: fmt vet lint test ## Run all CI checks

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo "  lint        - Run golangci-lint (if available)"
	@echo "  clean       - Clean build artifacts"
	@echo "  check       - Run fmt, vet, and test"
	@echo "  help        - Show this help"
