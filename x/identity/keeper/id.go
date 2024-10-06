package keeper

import (
	"context"
	"fmt"

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

	uniquekey := types.IdKey(id.Did + ":" + id.Username)
	store.Set(uniquekey, b)

	uniquestore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueIdKeyPrefix))
	uniquestore.Set(types.IdKey(id.Did), uniquekey)
	uniquestore.Set(types.IdKey(id.Creator), uniquekey)
	uniquestore.Set(types.IdKey(id.Username), uniquekey)
}

// update the current user and delete the olduser uniquekey
func (k Keeper) UpdateUserId(ctx context.Context, id types.Id, oldCreator string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	uniquestore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueIdKeyPrefix))

	b := k.cdc.MustMarshal(&id)

	uniquekey := types.IdKey(id.Did + ":" + id.Username)
	uniquestore.Delete(types.IdKey(oldCreator))

	store.Set(uniquekey, b)

	uniquestore.Set(types.IdKey(id.Did), uniquekey)

	uniquestore.Set(types.IdKey(id.Creator), uniquekey)

	uniquestore.Set(types.IdKey(id.Username), uniquekey)
}

// GetId returns a id from its index
func (k Keeper) GetId(
	ctx context.Context,
	did string,

) (val types.Id, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	uniquestore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueIdKeyPrefix))

	uniquekey := uniquestore.Get(types.IdKey(did))
	if uniquekey == nil {
		return val, false
	}

	data := store.Get(uniquekey)

	k.cdc.MustUnmarshal(data, &val)
	return val, true
}

func (k Keeper) GetIdByDidorUsernameorCreator(
	ctx context.Context,
	key string,

) (val types.Id, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	uniquestore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueIdKeyPrefix))

	uniquekey := uniquestore.Get(types.IdKey(key))
	if uniquekey == nil {
		return val, false
	}

	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	data := store.Get(uniquekey)

	k.cdc.MustUnmarshal(data, &val)
	return val, true
}

// RemoveId removes a id from the store
func (k Keeper) RemoveId(
	ctx context.Context,
	did string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IdKeyPrefix))
	fmt.Println(store)
	// store.Delete(types.IdKey(
	// 	did,
	// ))
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
