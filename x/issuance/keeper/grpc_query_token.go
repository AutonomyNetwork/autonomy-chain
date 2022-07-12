package keeper

import (
	"context"
	"strings"

	"reflect"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

func (k Keeper) TokenAll(c context.Context, req *types.QueryAllTokenRequest) (*types.QueryAllTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tokens []*types.Token
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	tokenStore := prefix.NewStore(store, types.TokenKey)

	pageRes, err := query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.Token

		token, err := types.UnMarshalToken(k.cdc, value)
		if err != nil {
			return err
		}
		token.Holders = uint64(len(k.GetDenomHolders(ctx, token.Denom)))
		tokens = append(tokens, &token)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenResponse{Tokens: tokens, Pagination: pageRes}, nil
}

func (k Keeper) Token(c context.Context, req *types.QueryGetTokenRequest) (*types.QueryGetTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if !k.HasToken(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	token := k.GetToken(ctx, req.Id)
	token.Holders = uint64(len(k.GetDenomHolders(ctx, token.Denom)))
	return &types.QueryGetTokenResponse{Token: &token}, nil
}

func (k Keeper) TokenByDenom(c context.Context, req *types.QueryTokenByDenomRequest) (*types.QueryGetTokenResponse, error) {
	empty := &types.QueryTokenByDenomRequest{}
	if req == nil || *req == *empty {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	token := k.GetTokenByDenom(ctx, req.Denom)
	token.Holders = uint64(len(k.GetDenomHolders(ctx, token.Denom)))
	if reflect.DeepEqual(token, types.Token{}) {
		return nil, status.Error(codes.NotFound, "requested denom doesn't exist")
	}

	return &types.QueryGetTokenResponse{
		Token: &token,
	}, nil

}

func (k Keeper) TokensByOwner(c context.Context, req *types.QueryTokensByOwnerRequest) (*types.QueryAllTokenResponse, error) {
	empty := &types.QueryTokensByOwnerRequest{}
	if req == nil || *req == *empty {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var tokens []*types.Token
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	tokenStore := prefix.NewStore(store, types.TokenKey)

	pageRes, err := query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
		var token types.Token

		token, err := types.UnMarshalToken(k.cdc, value)
		if err != nil {
			return err
		}
		if strings.EqualFold(token.Creator, req.Address) {
			token.Holders = uint64(len(k.GetDenomHolders(ctx, token.Denom)))
			tokens = append(tokens, &token)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenResponse{
		Tokens:     tokens,
		Pagination: pageRes,
	}, nil
}
