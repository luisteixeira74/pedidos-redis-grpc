PROTO_SRC=proto/email.proto
PROTO_OUT=proto

proto:
	protoc \
		--go_out=$(PROTO_OUT) \
		--go-grpc_out=$(PROTO_OUT) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_SRC)
