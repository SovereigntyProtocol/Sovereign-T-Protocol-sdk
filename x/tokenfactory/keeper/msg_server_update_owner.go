package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)

	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom does not exist")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var denom = types.Denom{
		Owner:     msg.NewOwner,
		Denom:     msg.Denom,
		MaxSupply: valFound.MaxSupply,
		Precision: valFound.Precision,
		Ticker:    valFound.Ticker,
	}

	k.SetDenom(
		ctx,
		denom,
	)

	return &types.MsgUpdateOwnerResponse{}, nil
}
