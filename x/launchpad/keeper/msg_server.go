package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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
	tr := k.issuenceKeeper.HasToken(ctx, t.TokenId)

	if !tr {
		fmt.Printf("invalid token id")
		return nil, sdkerrors.Wrapf(nil, "invalid token id: %d", t.TokenId)
	}

	token := k.issuenceKeeper.GetToken(ctx, t.TokenId)

	if token.Creator != t.Creator {
		fmt.Printf("unauthorized creator")
		return nil, sdkerrors.Wrapf(nil, "unauthorized creator")
	}

	if token.InitialSupply < t.Supply {
		fmt.Printf("launchpad supply should be less than token totla supply")
		return nil, sdkerrors.Wrapf(nil, "launchpad supply should be less than token totla supply")
	}

	launchpad := types.Launchpad{
		Id:        count,
		Creator:   t.Creator,
		TokenId:   t.TokenId,
		Supply:    t.Supply,
		Softcap:   t.Softcap,
		Hardcap:   t.Hardcap,
		StartTime: t.StartTime,
		EndTime:   t.EndTime,
		Status:    t.Status,
	}

	k.SetLaunchpad(ctx, launchpad)
	k.SetLaunchpadCount(ctx, count+1)

	createdLaunchpads := k.GetCreatedLaunchpads(ctx)

	launchpads := types.Launchpads{
		CreatedLaunchpads: append(createdLaunchpads.CreatedLaunchpads, launchpad.Id),
	}

	k.SetCreatedLaunchpads(ctx, launchpads)

	_ = ctx.EventManager().EmitTypedEvents(
		&types.EventCreateLaunchpad{
			Creator: t.Creator,
		},
		&types.EventModuleName)

	return &types.MsgCreateLaunchpadResponse{Id: count}, nil
}

func (k msgServer) DepositToLaunchpad(goctx context.Context, t *types.MsgDepositToLaunchpad) (*types.MsgDepositToLaunchpadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	l := k.HasLaunchpad(ctx, t.Id)
	count := k.GetDepositToLaunchpadCount(ctx)
	if !l {
		fmt.Printf("invalid launchpad id")
		return nil, sdkerrors.Wrapf(nil, "invalid launchpad id")
	}

	lp := k.GetLaunchpad(ctx, t.Id)

	if lp.Status != "ACTIVE" {
		fmt.Printf("invalid launchpad status")
		return nil, sdkerrors.Wrapf(nil, "invalid launchpad status")
	}

	deptr, err := sdk.AccAddressFromBech32(t.Depositor)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx,
		deptr, types.ModuleName, sdk.Coins{t.Amount})

	if err != nil {
		fmt.Println("error while subtracting balance from depoistor account", err)
		return nil, sdkerrors.Wrapf(nil, "error while subtracting balance from depoistor account")
	}

	lp.Deposits = lp.Deposits + t.Amount.Amount.Uint64()
	k.SetLaunchpad(ctx, lp)

	dep := types.DepositToLaunchpad{
		Id:        lp.Id,
		Depositor: t.Depositor,
		Amount:    t.Amount,
	}

	k.SetDepositToLaunchpad(ctx, lp.Id, count, dep)
	k.SetDepositLaunchpadCount(ctx, count+1)
	return &types.MsgDepositToLaunchpadResponse{Id: 0}, nil
}
