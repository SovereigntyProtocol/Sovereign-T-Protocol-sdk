package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/testutil/nullify"
	"github.com/cosmos/cosmos-sdk/x/identity/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUniquekeyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	msgs := createNUniquekey(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetUniquekeyRequest
		response *types.QueryGetUniquekeyResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUniquekeyRequest{
				Key: msgs[0].Key,
			},
			response: &types.QueryGetUniquekeyResponse{Uniquekey: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUniquekeyRequest{
				Key: msgs[1].Key,
			},
			response: &types.QueryGetUniquekeyResponse{Uniquekey: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUniquekeyRequest{
				Key: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Uniquekey(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestUniquekeyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	msgs := createNUniquekey(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUniquekeyRequest {
		return &types.QueryAllUniquekeyRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UniquekeyAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Uniquekey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Uniquekey),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UniquekeyAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Uniquekey), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Uniquekey),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UniquekeyAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Uniquekey),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UniquekeyAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
