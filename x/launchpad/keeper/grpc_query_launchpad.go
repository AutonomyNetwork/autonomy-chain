package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
)

func (k Keeper) LaunchpadAll(c context.Context, req *types.QueryAllLaunchpadRequest) (*types.QueryAllLaunchpadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var launchpads []*types.Launchpad
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	launchpadStore := prefix.NewStore(store, types.LaunchpadKey)

	pageRes, err := query.Paginate(launchpadStore, req.Pagination, func(key []byte, value []byte) error {
		var launchpad types.Launchpad

		launchpad, err := types.UnMarshalLaunchpad(k.cdc, value)
		if err != nil {
			return err
		}

		launchpads = append(launchpads, &launchpad)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLaunchpadResponse{Launchpads: launchpads, Pagination: pageRes}, nil
}

func (k Keeper) CreatedLaunchpads(c context.Context, req *types.QueryCreatedLaunchpadRequest) (*types.QueryCreatedLaunchpadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ids := k.GetCreatedLaunchpads(ctx)
	var launchpads []*types.Launchpad
	for _, id := range ids.CreatedLaunchpads {
		lp := k.GetLaunchpad(ctx, id)
		launchpads = append(launchpads, &lp)
	}

	return &types.QueryCreatedLaunchpadResponse{Launchpads: launchpads}, nil
}

func (k Keeper) ActiveLaunchpads(c context.Context, req *types.QueryActiveLaunchpadRequest) (*types.QueryActiveLaunchpadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ids := k.GetActiveLaunchpads(ctx)
	var launchpads []*types.Launchpad
	for _, id := range ids.ActiveLaunchpads {
		lp := k.GetLaunchpad(ctx, id)
		launchpads = append(launchpads, &lp)
	}

	return &types.QueryActiveLaunchpadResponse{Launchpads: launchpads}, nil
}

func (k Keeper) SuccessLaunchpads(c context.Context, req *types.QuerySuccessLaunchpadRequest) (*types.QuerySuccessLaunchpadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ids := k.GetSuccessLaunchpads(ctx)
	var launchpads []*types.Launchpad
	for _, id := range ids.SuccessLaunchpads {
		lp := k.GetLaunchpad(ctx, id)
		launchpads = append(launchpads, &lp)
	}

	return &types.QuerySuccessLaunchpadResponse{Launchpads: launchpads}, nil
}

func (k Keeper) FailLaunchpads(c context.Context, req *types.QueryFailLaunchpadRequest) (*types.QueryFailLaunchpadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ids := k.GetFailLaunchpads(ctx)
	var launchpads []*types.Launchpad
	for _, id := range ids.SuccessLaunchpads {
		lp := k.GetLaunchpad(ctx, id)
		launchpads = append(launchpads, &lp)
	}

	return &types.QueryFailLaunchpadResponse{Launchpads: launchpads}, nil
}

func (k Keeper) Launchpad(c context.Context, req *types.QueryGetLaunchpadRequest) (*types.QueryGetLaunchpadResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasLaunchpad(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	launchpad := k.GetLaunchpad(ctx, req.Id)
	return &types.QueryGetLaunchpadResponse{Launchpad: &launchpad}, nil
}
