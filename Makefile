BIN            = github-url-previewer
OUTPUT_DIR     = ./build
DOCS           = ./docs
RELEASE_VER   := $(shell git rev-parse --short HEAD)
TEST_PACKAGES  := $(shell go list ./...)
test_go := go

.PHONY: help
.DEFAULT_GOAL := help


build: clean ## Build binary and output to /build
	go build -v -a -installsuffix cgo -ldflags "-X main.version=$(RELEASE_VER)" -o $(OUTPUT_DIR)/$(BIN)
	$(OUTPUT_DIR)/$(BIN) --version

clean: ## Removing binary in output dir and stop and remove the containers
	$(RM) $(OUTPUT_DIR)/$(BIN)

test: test/unit test/fmt ## Perform  tests

test/unit: ## Perform unit tests
	$(test_go) test -cover -v $(TEST_PACKAGES)
