package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetUniquekey set a specific uniquekey in the store from its index
func (k Keeper) SetUniquekey(ctx context.Context, uniquekey types.Uniquekey) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniquekeyKeyPrefix))
	b := k.cdc.MustMarshal(&uniquekey)
	store.Set(types.UniquekeyKey(
		uniquekey.Key,
	), b)
}

// GetUniquekey returns a uniquekey from its index
func (k Keeper) GetUniquekey(
	ctx context.Context,
	key string,

) (val types.Uniquekey, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniquekeyKeyPrefix))

	b := store.Get(types.UniquekeyKey(
		key,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniquekey removes a uniquekey from the store
func (k Keeper) RemoveUniquekey(
	ctx context.Context,
	key string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniquekeyKeyPrefix))
	store.Delete(types.UniquekeyKey(
		key,
	))
}

// GetAllUniquekey returns all uniquekey
func (k Keeper) GetAllUniquekey(ctx context.Context) (list []types.Uniquekey) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniquekeyKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Uniquekey
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
