module github.com/AutonomyNetwork/autonomy-chain

go 1.16

require (
	github.com/AutonomyNetwork/nft v0.0.0-20220105130144-8a5d43e2a4ec
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/cosmos/ibc-go v1.2.2
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gravity-devs/liquidity v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/spm v0.1.8
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.40.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
