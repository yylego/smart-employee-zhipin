GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always --abbrev=40 --dirty=+code)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
	THIRD_PARTY_GOOGLE_API_PROTO_FILES=$(shell find third_party/google/api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/yylego/kratos-orz/cmd/orzkratos-srv-proto@latest

.PHONY: config
# generate internal proto
config:
	buf generate --template buf.gen.config.yaml

.PHONY: api
# generate api proto
api:
	buf generate --template buf.gen.yaml

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy
	@# Remove obsolete +build tags from wire_gen.go
	@sed -i '' '/^\/\/ +build/d' ./cmd/*/wire_gen.go

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

.PHONY: orz
# Proto-Service auto sync
orz:
	make all && orzkratos-srv-proto -auto

# generate TypeScript gRPC clients from proto files
# npm install -g @protobuf-ts/plugin
web_api_grpc_ts:
	mkdir -p ./bin/web_api_grpc_ts.out
	buf generate --template buf.gen.ts.yaml --include-imports

# convert gRPC clients to HTTP clients
# go install github.com/yylego/kratos-vue3/cmd/vue3kratos@latest
web_api_grpc_to_http:
	vue3kratos gen-grpc-via-http-in-root --grpc-ts-root=./bin/web_api_grpc_ts.out

# cleanup generated TypeScript files
web_api_cleanup:
	rm -rf ./bin/web_api_grpc_ts.out

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
