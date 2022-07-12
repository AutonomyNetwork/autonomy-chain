package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	denom := strings.ToLower(t.Denom)
	tokens := k.GetAllToken(ctx)
	for _, token := range tokens {
		if strings.EqualFold(denom, token.Denom) {
			return nil, sdkerrors.Wrapf(types.ErrDenomAlreadyExist, "invalid token denom: %s;", t.Denom)
		}
	}

	token := types.Token{
		Creator:       t.Creator,
		Id:            count,
		Denom:         strings.ToLower(denom),
		Decimals:      t.Decimals,
		DisplayName:   strings.ToLower(t.DisplayName),
		InitialSupply: t.InitialSupply,
		Holders:       1,
	}

	if err := k.MintToken(ctx, t.Creator, sdk.NewCoin(denom, sdk.NewIntFromUint64(t.InitialSupply))); err != nil {
		return nil, err
	}
	k.SetToken(ctx, token)
	k.SetTokenCount(ctx, count+1)

	metaData := bank.Metadata{
		Description: fmt.Sprintf("The details about %s token", denom),
		DenomUnits: []*bank.DenomUnit{
			{Denom: denom, Exponent: 0},
			{Denom: strings.ToLower(t.DisplayName), Exponent: uint32(t.Decimals)},
		},
		Base:    denom,
		Display: strings.ToLower(t.DisplayName),
	}

	k.SetDenomMetaData(ctx, metaData)

	_ = ctx.EventManager().EmitTypedEvents(
		&types.EventIssueToken{
			Address:       t.Creator,
			Decimals:      t.Decimals,
			InitialSupply: t.InitialSupply,
			Denom:         t.Denom,
		},
		&types.EventModuleName)

	return &types.MsgIssueTokenResponse{Id: count}, nil
}
