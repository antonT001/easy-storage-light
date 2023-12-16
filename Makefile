include .env*

BUILD_DIR := ./build
FILES_DIR := ./files
PROJECT_NAME := easy-storage-light
PROJECT_SRC_PATH := ./cmd/service
PROJECT_BINARY := $(FILES_DIR)/$(PROJECT_NAME)

####################################################################################################

build_dir:
	mkdir -p $(BUILD_DIR)
	mkdir -p $(FILES_DIR)

lint: build_dir
	golangci-lint run --timeout 5m -v --config .golangci.yaml | \
	tee $(BUILD_DIR)/lint-errors.log

test: build_dir
	go test ./... -v -count=1 -race -timeout=30s

compile: build_dir
	go build -o $(PROJECT_BINARY) $(PROJECT_SRC_PATH)

build: clean lint test compile

run_local:
	HOST=${HOST} \
	PORT=${PORT} \
	IDLE_TIMEOUT=${IDLE_TIMEOUT} \
	WRITE_TIMEOUT=${WRITE_TIMEOUT} \
	READ_TIMEOUT=${READ_TIMEOUT} \
	STORAGE_BASE_PATH=${STORAGE_BASE_PATH} \
	go run cmd/service/main.go

clean:
	echo "Cleaning all"
	rm -rf $(BUILD_DIR)/
	rm -rf $(FILES_DIR)/
