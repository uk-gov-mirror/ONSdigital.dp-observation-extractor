BUILD=build
BUILD_ARCH=$(BUILD)/$(GOOS)-$(GOARCH)
BIN_DIR?=.

export GOOS?=$(shell go env GOOS)
export GOARCH?=$(shell go env GOARCH)

build:
	@mkdir -p $(BUILD_ARCH)/$(BIN_DIR)
	go build -o $(BUILD_ARCH)/$(BIN_DIR)/dp-observation-extractor cmd/dp-observation-extractor/main.go
debug: build
	HUMAN_LOG=1 go run cmd/dp-observation-extractor/main.go
test:
	go test -cover $(go list ./... | grep -v /vendor/)
.PHONY: build debug test