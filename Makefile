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

.PHONY: generate_sqlc
generate_sqlc:
	$(SQLC) generate

.PHONY: generate_swagger
generate_swagger:
	swag init -g cmd/api/main.go -o ./docs