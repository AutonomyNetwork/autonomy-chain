package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) IssueToken(goctx context.Context, t *types.MsgIssueToken) (*types.MsgIssueTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	count := k.GetTokenCount(ctx)
	token := types.Token{
		Creator:       t.Creator,
		Id:            count,
		Denom:         t.Denom,
		InitialSupply: t.InitialSupply,
		Holders:       0,
	}

	if err := k.MintToken(ctx, t.Creator, sdk.NewCoin(t.Denom, sdk.NewIntFromUint64(t.InitialSupply))); err != nil {
		return nil, err
	}
	k.SetToken(ctx, token)
	k.SetTokenCount(ctx, count+1)

	ctx.EventManager().EmitTypedEvents(
		&types.EventIssueToken{
			Address:       t.Creator,
			Decimals:      t.Decimals,
			InitialSupply: t.InitialSupply,
			Denom:         t.Denom,
		},
		&types.EventModuleName)

	return &types.MsgIssueTokenResponse{Id: count}, nil

}
