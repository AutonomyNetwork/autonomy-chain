package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx sdk.Context) []bank.Balance
	
	GetDenomMetaData(ctx sdk.Context, denom string) (bank.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData bank.Metadata)
	IterateAllDenomMetaData(ctx sdk.Context, cb func(bank.Metadata) bool)
}
