proto:
	protoc api/*.proto --go-grpc_out=pkg --go_out=pkg
