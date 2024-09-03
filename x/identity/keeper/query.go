package keeper

import (
	"github.com/cosmos/cosmos-sdk/x/identity/types"
)

var _ types.QueryServer = Keeper{}
