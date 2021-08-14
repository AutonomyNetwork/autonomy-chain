package issuance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/keeper"
	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, *elem)
	}
	
	k.SetTokenCount(ctx, genState.TokenCount)
	
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	
	tokenList := k.GetAllToken(ctx)
	for _, elem := range tokenList {
		elem := elem
		genesis.TokenList = append(genesis.TokenList, &elem)
	}
	
	genesis.TokenCount = k.GetTokenCount(ctx)
	
	return genesis
}
