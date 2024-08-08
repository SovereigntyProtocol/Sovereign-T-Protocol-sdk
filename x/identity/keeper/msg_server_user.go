package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUser(
		ctx,
		msg.Did,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	msg.Owner = msg.Creator
	_, isAddressFound := k.GetAddress(
		ctx,
		msg.Owner,
	)

	if isAddressFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "owner index already set")
	}

	var address = types.Address{
		Creator: msg.Creator,
		Owner:   msg.Owner,
	}

	k.SetAddress(ctx, address)

	var user = types.User{
		Creator: msg.Creator,
		Did:     msg.Did,
		Hash:    msg.Hash,
		Owner:   msg.Owner,
	}

	k.SetUser(
		ctx,
		user,
	)
	return &types.MsgCreateUserResponse{}, nil
}

func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUser(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if msg.Hash == valFound.Hash {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "same hash")
	}

	msg.Creator = msg.Owner
	var user = types.User{
		Creator: msg.Creator,
		Did:     msg.Did,
		Hash:    msg.Hash,
		Owner:   msg.Owner,
	}

	k.SetUser(ctx, user)

	var address = types.Address{
		Creator: msg.Creator,
		Owner:   msg.Owner,
	}

	k.SetAddress(ctx, address)

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUser(
		ctx,
		msg.Did,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUser(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteUserResponse{}, nil
}
