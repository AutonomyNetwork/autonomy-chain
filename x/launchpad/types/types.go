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
	launch, err := UnMarshalLaunchpad(cdc, value)
	if err != nil {
		panic(err)
	}

	return launch
}

func UnMarshalLaunchpad(cdc codec.BinaryCodec, value []byte) (launchpad Launchpad, err error) {
	err = cdc.Unmarshal(value, &launchpad)
	return launchpad, err
}

func MustMarshalDepositToLaunchpad(cdc codec.BinaryCodec, deposit *DepositToLaunchpad) []byte {
	return cdc.MustMarshal(deposit)
}

func MustUnmashalDepositToLaunchpad(cdc codec.BinaryCodec, value []byte) (deposit DepositToLaunchpad) {
	depo, err := UnMarshalDepositToLaunchpad(cdc, value)
	if err != nil {
		panic(err)
	}

	return depo
}

func UnMarshalDepositToLaunchpad(cdc codec.BinaryCodec, value []byte) (deposit DepositToLaunchpad, err error) {
	err = cdc.Unmarshal(value, &deposit)
	return deposit, err
}

func MustMarshalCreatedLaunchpads(cdc codec.BinaryCodec, launchpads *Launchpads) []byte {
	return cdc.MustMarshal(launchpads)
}

func MustUnmashalCreatedLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads) {
	laun, err := UnMarshalCreatedLaunchpads(cdc, value)
	if err != nil {
		panic(err)
	}

	return laun
}

func UnMarshalCreatedLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads, err error) {
	err = cdc.Unmarshal(value, &launchpads)
	return launchpads, err
}

func MustMarshalActiveLaunchpads(cdc codec.BinaryCodec, launchpads *Launchpads) []byte {
	return cdc.MustMarshal(launchpads)
}

func MustUnmashalActiveLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads) {
	laun, err := UnMarshalActiveLaunchpads(cdc, value)
	if err != nil {
		panic(err)
	}

	return laun
}

func UnMarshalActiveLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads, err error) {
	err = cdc.Unmarshal(value, &launchpads)
	return launchpads, err
}

func MustMarshalSuccessLaunchpads(cdc codec.BinaryCodec, launchpads *Launchpads) []byte {
	return cdc.MustMarshal(launchpads)
}

func MustUnmashalSuccessLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads) {
	laun, err := UnMarshalSuccessLaunchpads(cdc, value)
	if err != nil {
		panic(err)
	}

	return laun
}

func UnMarshalSuccessLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads, err error) {
	err = cdc.Unmarshal(value, &launchpads)
	return launchpads, err
}
func MustMarshalFailLaunchpads(cdc codec.BinaryCodec, launchpads *Launchpads) []byte {
	return cdc.MustMarshal(launchpads)
}

func MustUnmashalFailLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads) {
	laun, err := UnMarshalFailLaunchpads(cdc, value)
	if err != nil {
		panic(err)
	}

	return laun
}

func UnMarshalFailLaunchpads(cdc codec.BinaryCodec, value []byte) (launchpads Launchpads, err error) {
	err = cdc.Unmarshal(value, &launchpads)
	return launchpads, err
}
