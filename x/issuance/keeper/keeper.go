package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

type (
	Keeper struct {
		cdc        codec.Marshaler
		storeKey   sdk.StoreKey
		BankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		BankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
