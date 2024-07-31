package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetAddress set a specific address in the store from its index
func (k Keeper) SetAddress(ctx context.Context, address types.Address) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressKeyPrefix))
	b := k.cdc.MustMarshal(&address)
	store.Set(types.AddressKey(
		address.Owner,
	), b)
}

// GetAddress returns a address from its index
func (k Keeper) GetAddress(
	ctx context.Context,
	owner string,

) (val types.Address, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressKeyPrefix))

	b := store.Get(types.AddressKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAddress removes a address from the store
func (k Keeper) RemoveAddress(
	ctx context.Context,
	owner string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressKeyPrefix))
	store.Delete(types.AddressKey(
		owner,
	))
}

// GetAllAddress returns all address
func (k Keeper) GetAllAddress(ctx context.Context) (list []types.Address) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Address
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
