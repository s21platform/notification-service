.PHONY: protogen

protogen:
	protoc --go_out=. --go-grpc_out=. ./api/notification.proto
	protoc --doc_out=. --doc_opt=markdown,GRPC_API.md ./api/notification.proto