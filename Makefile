BIN_DIR := ./bin

run: $(BIN_DIR)/kvstore $(BIN_DIR)/kvstore-cli
	@nohup $(BIN_DIR)/kvstore > /dev/null 2>&1 &
	@sleep 1
	@$(BIN_DIR)/kvstore-cli
	@pkill -f "$(BIN_DIR)/kvstore" || true

$(BIN_DIR)/kvstore:
	@go build -o $(BIN_DIR)/kvstore ./cmd

$(BIN_DIR)/kvstore-cli:
	@go build -o $(BIN_DIR)/kvstore-cli ./client

stop:
	@pkill -f "$(BIN_DIR)/kvstore" || true
