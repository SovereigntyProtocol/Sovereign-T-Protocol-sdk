package identity_test

import (
	"testing"

	keepertest "identity/testutil/keeper"
	"identity/testutil/nullify"
	identity "identity/x/identity/module"
	"identity/x/identity/types"

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
