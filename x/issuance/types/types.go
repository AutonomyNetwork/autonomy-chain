package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "issuance"
	
	// StoreKey defines the primary module store key
	StoreKey = ModuleName
	
	// RouterKey is the message route for slashing
	RouterKey = ModuleName
	
	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var EventModuleName = EventModule{Name: ModuleName}

// UInt64Bytes uses the SDK byte marshaling to encode a uint64
func UInt64Bytes(n uint64) []byte {
	return sdk.Uint64ToBigEndian(n)
}

func MustMarshalToken(cdc codec.BinaryCodec, token *Token) []byte {
	return cdc.MustMarshal(token)
}

func MustUnmashalToken(cdc codec.BinaryCodec, value []byte) (token Token) {
	token, err := UnMarshalToken(cdc, value)
	if err != nil {
		panic(err)
	}
	
	return token
}

func UnMarshalToken(cdc codec.BinaryCodec, value []byte) (token Token, err error) {
	err = cdc.Unmarshal(value, &token)
	return token, err
}


