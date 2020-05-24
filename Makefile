PROJECT_NAME := "demo_for_ci"
PKG := "github.com/2cloudlab/$(PROJECT_NAME)"
PKG_LIST := "demo_for_ci/mylib demo_for_ci/myapp"
 
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
    @ls -al
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}

build: dep ## Build the binary file
	@go build -i -o build/main $(PKG)
 
clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)/build