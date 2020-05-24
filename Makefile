PROJECT_NAME := "demo_for_ci"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
 
.PHONY: all dep lint vet test test-coverage build clean
 
all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@echo "$(cat cover.out)"

all-tests-coverage: ## Run all tests with coverage
	@go test -tags=integration -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@echo "$(cat cover.out)"

build: dep ## Build the binary file
	@go build -i -o build/main $(PKG)
 
clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)/build