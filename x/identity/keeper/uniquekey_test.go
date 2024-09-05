package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/testutil/nullify"
	"github.com/cosmos/cosmos-sdk/x/identity/keeper"
	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUniquekey(keeper keeper.Keeper, ctx context.Context, n int) []types.Uniquekey {
	items := make([]types.Uniquekey, n)
	for i := range items {
		items[i].Key = strconv.Itoa(i)

		keeper.SetUniquekey(ctx, items[i])
	}
	return items
}

func TestUniquekeyGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNUniquekey(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniquekey(ctx,
			item.Key,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUniquekeyRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNUniquekey(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniquekey(ctx,
			item.Key,
		)
		_, found := keeper.GetUniquekey(ctx,
			item.Key,
		)
		require.False(t, found)
	}
}

func TestUniquekeyGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNUniquekey(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniquekey(ctx)),
	)
}
