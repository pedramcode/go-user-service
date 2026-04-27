GO			:= go
SQLC		:= sqlc

CMD_DIR		:= cmd
API_DIR		:= $(CMD_DIR)/api
ENTRY_FILE	:= $(API_DIR)/main.go

.PHONY: run
run:
	$(GO) run $(ENTRY_FILE)

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: generate
generate:
	$(SQLC) generate