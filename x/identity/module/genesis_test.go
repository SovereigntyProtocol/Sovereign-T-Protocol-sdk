package identity_test

import (
	"testing"

	keepertest "github.com/cosmos/cosmos-sdk/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/testutil/nullify"
	identity "github.com/cosmos/cosmos-sdk/x/identity/module"
	"github.com/cosmos/cosmos-sdk/x/identity/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IdList: []types.Id{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, k, genesisState)
	got := identity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdList, got.IdList)
	// this line is used by starport scaffolding # genesis/test/assert
}
