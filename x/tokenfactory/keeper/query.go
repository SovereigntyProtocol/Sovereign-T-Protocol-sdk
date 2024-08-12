package keeper

import (
	"github.com/cosmos/cosmos-sdk/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
