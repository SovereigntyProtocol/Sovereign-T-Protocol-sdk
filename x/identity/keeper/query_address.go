package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressAll(ctx context.Context, req *types.QueryAllAddressRequest) (*types.QueryAllAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addresss []types.Address

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	addressStore := prefix.NewStore(store, types.KeyPrefix(types.AddressKeyPrefix))

	pageRes, err := query.Paginate(addressStore, req.Pagination, func(key []byte, value []byte) error {
		var address types.Address
		if err := k.cdc.Unmarshal(value, &address); err != nil {
			return err
		}

		addresss = append(addresss, address)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressResponse{Address: addresss, Pagination: pageRes}, nil
}

func (k Keeper) Address(ctx context.Context, req *types.QueryGetAddressRequest) (*types.QueryGetAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetAddress(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAddressResponse{Address: val}, nil
}
