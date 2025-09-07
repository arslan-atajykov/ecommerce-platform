.PHONY: proto

PROTO_DIR=./proto
OUT_DIR=./proto/gen

proto:
	mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) \
		$(PROTO_DIR)/*.proto
