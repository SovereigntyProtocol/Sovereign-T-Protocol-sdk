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

func (k Keeper) UniquekeyAll(ctx context.Context, req *types.QueryAllUniquekeyRequest) (*types.QueryAllUniquekeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var uniquekeys []types.Uniquekey

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	uniquekeyStore := prefix.NewStore(store, types.KeyPrefix(types.UniquekeyKeyPrefix))

	pageRes, err := query.Paginate(uniquekeyStore, req.Pagination, func(key []byte, value []byte) error {
		var uniquekey types.Uniquekey
		if err := k.cdc.Unmarshal(value, &uniquekey); err != nil {
			return err
		}

		uniquekeys = append(uniquekeys, uniquekey)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUniquekeyResponse{Uniquekey: uniquekeys, Pagination: pageRes}, nil
}

func (k Keeper) Uniquekey(ctx context.Context, req *types.QueryGetUniquekeyRequest) (*types.QueryGetUniquekeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetUniquekey(
		ctx,
		req.Key,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUniquekeyResponse{Uniquekey: val}, nil
}
