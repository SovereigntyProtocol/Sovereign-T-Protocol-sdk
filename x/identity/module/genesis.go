package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/identity/keeper"
	"github.com/cosmos/cosmos-sdk/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the id
	for _, elem := range genState.IdList {
		k.SetId(ctx, elem)
	}
	// Set all the uniquekey
	for _, elem := range genState.UniquekeyList {
		k.SetUniquekey(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.IdList = k.GetAllId(ctx)
	genesis.UniquekeyList = k.GetAllUniquekey(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
