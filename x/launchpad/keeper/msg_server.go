package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
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

func (k msgServer) CreateLaunchpad(goctx context.Context, t *types.MsgCreateLaunchpad) (*types.MsgCreateLaunchpadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	count := k.GetLaunchpadCount(ctx)

	launchpad := types.Launchpad{
		Id:              count,
		Creator:         t.Creator,
		TokenId:         t.TokenId,
		AccepetedDenoms: t.AccepetedDenoms,
		Softcap:         t.Softcap,
		Hardcap:         t.Hardcap,
		StartTime:       t.StartTime,
		EndTime:         t.EndTime,
		Status:          t.Status,
	}

	k.SetLaunchpad(ctx, launchpad)
	k.SetLaunchpadCount(ctx, count+1)

	_ = ctx.EventManager().EmitTypedEvents(
		&types.EventCreateLaunchpad{
			Creator: t.Creator,
		},
		&types.EventModuleName)

	return &types.MsgCreateLaunchpadResponse{Id: count}, nil

}

func (m msgServer) DepositToLaunchpad(goctx context.Context, t *types.MsgDepositToLaunchpad) (*types.MsgDepositToLaunchpadResponse, error) {
	return &types.MsgDepositToLaunchpadResponse{Id: 0}, nil
}
