PACKAGES := $(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TENDERMINT_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

BUILD_TAGS := $(strip netgo,ledger)
LD_FLAGS := -s -w \
    -X github.com/cosmos/cosmos-sdk/version.Name=Autonomy \
    -X github.com/cosmos/cosmos-sdk/version.AppName=autonomy \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TENDERMINT_VERSION)

clean:
	rm -rf ./bin ./vendor

install: mod-vendor
	go install -mod=readonly -tags="${BUILD_TAGS}" -ldflags="${LD_FLAGS}" ./cmd/autonomy

go-lint:
	@golangci-lint run --fix

mod-vendor:
	@go mod vendor

test:
	@go test -mod=readonly -v -cover ${PACKAGES}

tools:
	@go install github.com/bufbuild/buf/cmd/buf@v0.37.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0
	@go install github.com/goware/modvendor@v0.3.0
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0

.PHONY: tools test proto-lint proto-gen mod-vendor go-lint install