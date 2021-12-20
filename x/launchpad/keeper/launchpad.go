package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"

	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// GetLaunchpadCount get the total number of launchpads
func (k Keeper) GetLaunchpadCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LaunchpadCountKey)

	if bz == nil {
		return 0
	}

	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		panic("cannot decode count")
	}

	return count
}

// SetLaunchpadCount set the total number of Launchpad
func (k Keeper) SetLaunchpadCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := types.LaunchpadCountKey
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) SetLaunchpad(ctx sdk.Context, launchpad types.Launchpad) {
	store := ctx.KVStore(k.storeKey)

	b := types.MustMarshalLaunchpad(k.cdc, &launchpad)

	store.Set(types.GetLaunchpadKey(launchpad.Id), b)
}

func (k Keeper) GetLaunchpad(ctx sdk.Context, id uint64) (launchpad types.Launchpad) {
	store := ctx.KVStore(k.storeKey)
	val := store.Get(types.GetLaunchpadKey(id))
	if val == nil {
		return types.Launchpad{}
	}
	launchpad = types.MustUnmashalLaunchpad(k.cdc, val)
	return launchpad
}

func (k Keeper) HasLaunchpad(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetLaunchpadKey(id))
}

func (k Keeper) RemoveLaunchpad(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLaunchpadKey(id))
}

func (k Keeper) GetAllLaunchpad(ctx sdk.Context) (list []types.Launchpad) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LaunchpadKey)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Launchpad
		k.cdc.Unmarshal(iterator.Value(), &val)
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

func (k Keeper) GetDenomMetaData(ctx sdk.Context, denom string) (bank.Metadata, bool) {
	return k.bankKeeper.GetDenomMetaData(ctx, denom)
}

func (k Keeper) GetAllDenomsMetaData(ctx sdk.Context) []bank.Metadata {
	denomsMetaData := make([]bank.Metadata, 0)

	k.bankKeeper.IterateAllDenomMetaData(ctx, func(metadata bank.Metadata) bool {
		denomsMetaData = append(denomsMetaData, (metadata))
		return false
	})
	return denomsMetaData
}

func (k Keeper) GetDenomHolders(ctx sdk.Context, denom string) []bank.Balance {
	denomHolder := make([]bank.Balance, 0)
	totalBalances := k.bankKeeper.GetAccountsBalances(ctx)
	for _, account := range totalBalances {
		balance := account.Coins.AmountOf(denom)
		if balance.IsPositive() {
			acc := bank.Balance{
				Address: account.Address,
				Coins:   sdk.Coins{sdk.NewCoin(denom, balance)},
			}
			denomHolder = append(denomHolder, acc)
		}
	}
	return denomHolder
}
