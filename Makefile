compile:
	protoc ./api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--proto_path=.

test:
	go test -race ./...