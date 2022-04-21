package types

import (
	issuence "github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, from, to sdk.AccAddress, amount sdk.Coins) error

	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx sdk.Context) []bank.Balance

	GetDenomMetaData(ctx sdk.Context, denom string) (bank.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData bank.Metadata)
	IterateAllDenomMetaData(ctx sdk.Context, cb func(bank.Metadata) bool)

	BlockedAddr(to sdk.AccAddress) bool
}

type AccountKeeper interface {
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI
	GetAccount(sdk.Context, sdk.AccAddress) types.AccountI
	SetAccount(sdk.Context, types.AccountI)
	GetModuleAccount(ctx sdk.Context, modeuleName string) types.ModuleAccountI
}

type IssuenceKeeper interface {
	GetToken(ctx sdk.Context, id uint64) (token issuence.Token)
	HasToken(ctx sdk.Context, id uint64) bool
}
