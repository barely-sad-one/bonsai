BINARY_NAME := bonsai
# Build output directory
BUILD_DIR := build

.PHONY: all build clean cache

all: build

build:
	@echo ">> Building the binary..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/bonsai/main.go

clean:
	@echo ">> Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

cache:
	@echo ">> Cleaning Go build cache..."
	@go clean -cache -modcache -testcache -fuzzcache
