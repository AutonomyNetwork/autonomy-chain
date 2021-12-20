package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "launchpad"

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

func MustMarshalLaunchpad(cdc codec.BinaryCodec, launchpad *Launchpad) []byte {
	return cdc.MustMarshal(launchpad)
}

func MustUnmashalLaunchpad(cdc codec.BinaryCodec, value []byte) (launchpad Launchpad) {
	token, err := UnMarshalLaunchpad(cdc, value)
	if err != nil {
		panic(err)
	}

	return token
}

func UnMarshalLaunchpad(cdc codec.BinaryCodec, value []byte) (launchpad Launchpad, err error) {
	err = cdc.Unmarshal(value, &launchpad)
	return launchpad, err
}
