package keeper

import (
	"testing"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	
	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

func createNToken(keeper *Keeper, ctx sdk.Context, n int) []types.Token {
	items := make([]types.Token, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.IssueToken(ctx, items[i])
	}
	return items
}

func TestTokenGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetToken(ctx, item.Id))
	}
}

func TestTokenExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasToken(ctx, item.Id))
	}
}

func TestTokenRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNToken(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveToken(ctx, item.Id)
		assert.False(t, keeper.HasToken(ctx, item.Id))
	}
}

func TestTokenGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNToken(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllToken(ctx))
}

func TestTokenCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNToken(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetTokenCount(ctx))
}
