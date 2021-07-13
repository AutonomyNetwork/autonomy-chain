package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// GetTokenCount get the total number of token
func (k Keeper) GetTokenCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TokenCountKey)

	if bz == nil {
		return 0
	}

	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		panic("cannot decode count")
	}

	return count
}

// SetTokenCount set the total number of token
func (k Keeper) SetTokenCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := types.TokenCountKey
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// IssueToken appends a token in the store with a new id and update the count
func (k Keeper) IssueToken(
	ctx sdk.Context,
	token types.Token,
) uint64 {
	// Create the token count
	count := k.GetTokenCount(ctx)

	// Set the ID of the appended value
	token.Id = count

	store := ctx.KVStore(k.storeKey)
	bytes := k.cdc.MustMarshalBinaryBare(&token)
	store.Set(types.GetTokenKey(token.Id), bytes)

	// Update token count
	k.SetTokenCount(ctx, count+1)

	return count
}

// SetToken set a specific token in the store
func (k Keeper) SetToken(ctx sdk.Context, token types.Token) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&token)
	store.Set(types.GetTokenKey(token.Id), b)
}

// GetToken returns a token from its id
func (k Keeper) GetToken(ctx sdk.Context, id uint64) types.Token {
	store := ctx.KVStore(k.storeKey)
	var token types.Token
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.GetTokenKey(id)), &token)
	return token
}

// HasToken checks if the token exists in the store
func (k Keeper) HasToken(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetTokenKey(id))
}

// RemoveToken removes a token from the store
func (k Keeper) RemoveToken(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetTokenKey(id))
}

// GetAllToken returns all token
func (k Keeper) GetAllToken(ctx sdk.Context) (list []types.Token) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TokenKey)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Token
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Bank keeper functions

func (k Keeper) MintToken(ctx sdk.Context, address string, amount sdk.Coin) error {

	sender, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{amount}); err != nil {
		return sdkerrors.Wrapf(err, "mint vouchers coins: %s", amount)
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{amount}); err != nil {
		return sdkerrors.Wrap(err, "transfer tokens")
	}
	return nil
}


func (k Keeper) SetDenomMetaData(ctx sdk.Context, metadata bank.Metadata) {
	k.bankKeeper.SetDenomMetaData(ctx, metadata)
}

func (k Keeper) GetDenomMetaData(ctx sdk.Context, denom string) bank.Metadata{
	return k.bankKeeper.GetDenomMetaData(ctx,denom)
}