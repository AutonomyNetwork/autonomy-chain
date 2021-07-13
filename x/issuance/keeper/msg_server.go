package keeper

import (
	"context"
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	
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
		DisplayName: t.DisplayName,
		InitialSupply: t.InitialSupply,
		Holders:       0,
	}

	if err := k.MintToken(ctx, t.Creator, sdk.NewCoin(t.Denom, sdk.NewIntFromUint64(t.InitialSupply))); err != nil {
		return nil, err
	}
	k.SetToken(ctx, token)
	k.SetTokenCount(ctx, count+1)
	
	metaData:= bank.Metadata{
		Description: fmt.Sprintf("The deatilas about %s token", t.Denom),
		DenomUnits:  []*bank.DenomUnit{
			{Denom: t.Denom, Exponent: 0},
			{Denom: t.DisplayName, Exponent: uint32(t.Decimals)},
		},
		Base:        t.Denom,
		Display:     t.DisplayName,
	}
	
	k.SetDenomMetaData(ctx, metaData)

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
