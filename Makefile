BINARY_NAME := bonsai
BUILD_DIR := build

DEBUG_FLAGS := -gcflags "all=-N -l"
RELEASE_FLAGS := -ldflags "-s -w"

.PHONY: all build debug release clean cache

all: release

build: release

debug:
	@echo ">> Building debug binary with GIN_MODE=debug..."
	@mkdir -p $(BUILD_DIR)
	@GIN_MODE=debug go build $(DEBUG_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_debug ./cmd/bonsai/main.go

release:
	@echo ">> Building release binary with GIN_MODE=release..."
	@mkdir -p $(BUILD_DIR)
	@GIN_MODE=release go build $(RELEASE_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/bonsai/main.go

clean:
	@echo ">> Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

cache:
	@echo ">> Cleaning Go build cache..."
	@go clean -cache -modcache -testcache -fuzzcache

