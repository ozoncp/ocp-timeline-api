PROJECT_ROOT = $(shell pwd )
BIN_DIR ?= $(PROJECT_ROOT)/bin

hello:
	@echo "hello from make"

build: clean
	@echo "Building..."
	@go build -o "$(BIN_DIR)/ocp-timeline-api" "$(PROJECT_ROOT)/cmd/ocp-timeline-api/main.go"

run:
	@echo "Runing..."
	@go run "$(PROJECT_ROOT)/cmd/ocp-timeline-api/main.go"

test:
	@go test ./... -v

clean:
	@echo "Cleaning ... "
	@rm -rf "$(BIN)}"