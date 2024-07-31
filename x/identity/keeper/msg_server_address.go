package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAddress(goCtx context.Context, msg *types.MsgCreateAddress) (*types.MsgCreateAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetAddress(
		ctx,
		msg.Owner,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var address = types.Address{
		Creator: msg.Creator,
		Owner:   msg.Owner,
	}

	k.SetAddress(
		ctx,
		address,
	)
	return &types.MsgCreateAddressResponse{}, nil
}

func (k msgServer) UpdateAddress(goCtx context.Context, msg *types.MsgUpdateAddress) (*types.MsgUpdateAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAddress(
		ctx,
		msg.Owner,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var address = types.Address{
		Creator: msg.Creator,
		Owner:   msg.Owner,
	}

	k.SetAddress(ctx, address)

	return &types.MsgUpdateAddressResponse{}, nil
}

func (k msgServer) DeleteAddress(goCtx context.Context, msg *types.MsgDeleteAddress) (*types.MsgDeleteAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAddress(
		ctx,
		msg.Owner,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAddress(
		ctx,
		msg.Owner,
	)

	return &types.MsgDeleteAddressResponse{}, nil
}
