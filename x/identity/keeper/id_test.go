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

func createNId(keeper keeper.Keeper, ctx context.Context, n int) []types.Id {
	items := make([]types.Id, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)

		keeper.SetId(ctx, items[i])
	}
	return items
}

func TestIdGet(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetId(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestIdRemove(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNId(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveId(ctx,
			item.Did,
		)
		_, found := keeper.GetId(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdentityKeeper(t)
	items := createNId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllId(ctx)),
	)
}
