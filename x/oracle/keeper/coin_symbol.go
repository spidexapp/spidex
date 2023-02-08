package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spidexapp/spidex/x/oracle/types"
)

// GetCoinSymbolCount get the total number of coinSymbol
func (k Keeper) GetCoinSymbolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoinSymbolCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCoinSymbolCount set the total number of coinSymbol
func (k Keeper) SetCoinSymbolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoinSymbolCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCoinSymbol appends a coinSymbol in the store with a new id and update the count
func (k Keeper) AppendCoinSymbol(
	ctx sdk.Context,
	coinSymbol types.CoinSymbol,
) uint64 {
	// Create the coinSymbol
	count := k.GetCoinSymbolCount(ctx)

	// Set the ID of the appended value
	coinSymbol.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinSymbolKey))
	appendedValue := k.cdc.MustMarshal(&coinSymbol)
	store.Set(GetCoinSymbolIDBytes(coinSymbol.Id), appendedValue)

	// Update coinSymbol count
	k.SetCoinSymbolCount(ctx, count+1)

	return count
}

// SetCoinSymbol set a specific coinSymbol in the store
func (k Keeper) SetCoinSymbol(ctx sdk.Context, coinSymbol types.CoinSymbol) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinSymbolKey))
	b := k.cdc.MustMarshal(&coinSymbol)
	store.Set(GetCoinSymbolIDBytes(coinSymbol.Id), b)
}

// GetCoinSymbol returns a coinSymbol from its id
func (k Keeper) GetCoinSymbol(ctx sdk.Context, id uint64) (val types.CoinSymbol, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinSymbolKey))
	b := store.Get(GetCoinSymbolIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCoinSymbol removes a coinSymbol from the store
func (k Keeper) RemoveCoinSymbol(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinSymbolKey))
	store.Delete(GetCoinSymbolIDBytes(id))
}

// GetAllCoinSymbol returns all coinSymbol
func (k Keeper) GetAllCoinSymbol(ctx sdk.Context) (list []types.CoinSymbol) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoinSymbolKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoinSymbol
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinSymbolIDBytes returns the byte representation of the ID
func GetCoinSymbolIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoinSymbolIDFromBytes returns ID in uint64 format from a byte array
func GetCoinSymbolIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
