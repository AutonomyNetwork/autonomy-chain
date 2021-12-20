package launchpad

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/keeper"
	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.LaunchpadList {
		k.SetLaunchpad(ctx, *elem)
	}

	k.SetLaunchpadCount(ctx, genState.LaunchpadCount)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	return genesis
}
