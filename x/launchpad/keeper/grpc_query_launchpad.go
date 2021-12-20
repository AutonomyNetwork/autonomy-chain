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
		// token.Holders = uint64(len(k.GetDenomHolders(ctx, token.Denom)))
		launchpads = append(launchpads, &launchpad)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLaunchpadResponse{Launchpads: launchpads, Pagination: pageRes}, nil
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
