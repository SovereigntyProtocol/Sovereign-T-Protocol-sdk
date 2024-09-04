package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetId set a specific id in the store from its index
func (k Keeper) SetId(ctx context.Context, id types.Id) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	b := k.cdc.MustMarshal(&id)
	store.Set(types.IdKey(
		id.Did,
	), b)
}

// GetId returns a id from its index
func (k Keeper) GetId(
	ctx context.Context,
	did string,

) (val types.Id, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))

	b := store.Get(types.IdKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveId removes a id from the store
func (k Keeper) RemoveId(
	ctx context.Context,
	did string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	store.Delete(types.IdKey(
		did,
	))
}

// GetAllId returns all id
func (k Keeper) GetAllId(ctx context.Context) (list []types.Id) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Id
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
