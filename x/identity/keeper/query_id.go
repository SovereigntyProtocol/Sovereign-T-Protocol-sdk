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

func (k Keeper) IdAll(ctx context.Context, req *types.QueryAllIdRequest) (*types.QueryAllIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ids []types.Id

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	idStore := prefix.NewStore(store, types.KeyPrefix(types.IdKeyPrefix))

	pageRes, err := query.Paginate(idStore, req.Pagination, func(key []byte, value []byte) error {
		var id types.Id
		if err := k.cdc.Unmarshal(value, &id); err != nil {
			k.Logger().Error("Failed to unmarshal user", "key", key, "error", err)
			return nil
		}
		ids = append(ids, id)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIdResponse{Id: ids, Pagination: pageRes}, nil
}

func (k Keeper) Id(ctx context.Context, req *types.QueryGetIdRequest) (*types.QueryGetIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetIdByDidorUsernameorCreator(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetIdResponse{Id: val}, nil
}
