PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GRPC_GO = protoc-gen-go-grpc

.PHONY: prepare
prepare: install-deps protogen tidy

.PHONY: install-deps
install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest;\
	go install github.com/automation-co/husky@latest;\
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest;\
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest;\
	(husky init || true) && husky install;

.PHONY: protogen
protogen:
	protoc proto/*.proto --proto_path=./proto \
         --go_out=./proto --go_opt=module=github.com/prettyboiiii/bouncer/proto \
         --go-grpc_out=./proto --go-grpc_opt=module=github.com/prettyboiiii/bouncer/proto

.PHONY: tidy
tidy: 
	go mod tidy