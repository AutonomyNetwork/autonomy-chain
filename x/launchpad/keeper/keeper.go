package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
)

type (
	Keeper struct {
		cdc            codec.Codec
		storeKey       sdk.StoreKey
		bankKeeper     types.BankKeeper
		issuenceKeeper types.IssuenceKeeper
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	issuenceKeeper types.IssuenceKeeper,
	bankKeeper types.BankKeeper,

) Keeper {
	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		bankKeeper:     bankKeeper,
		issuenceKeeper: issuenceKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
