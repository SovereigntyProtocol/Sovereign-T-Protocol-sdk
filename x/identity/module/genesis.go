package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/identity/keeper"
	"github.com/cosmos/cosmos-sdk/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, elem)
	}
	// Set all the address
	for _, elem := range genState.AddressList {
		k.SetAddress(ctx, elem)
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

	genesis.UserList = k.GetAllUser(ctx)
	genesis.AddressList = k.GetAllAddress(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
