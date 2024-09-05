package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateId(goCtx context.Context, msg *types.MsgCreateId) (*types.MsgCreateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newDid, diderr := types.CreateNewDid()

	if diderr != nil {
		return nil, errorsmod.Wrap(diderr, "did error, please try again")
	}

	_, isFound := k.GetIdByUniqueKey(ctx, newDid)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// _, isUserNameFound := k.GetIdByUniqueKey(ctx, msg.Username)
	// if isUserNameFound {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "username already set")
	// }

	// _, isCreatorFound := k.GetIdByUniqueKey(ctx, msg.Creator)
	// if isCreatorFound {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "wallet already set")
	// }

	var id = types.Id{
		Creator:  msg.Creator,
		Did:      newDid,
		Hash:     msg.Hash,
		Username: msg.Username,
	}

	k.SetId(
		ctx,
		id,
	)
	return &types.MsgCreateIdResponse{}, nil
}

func (k msgServer) UpdateId(goCtx context.Context, msg *types.MsgUpdateId) (*types.MsgUpdateIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetId(
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

	var id = types.Id{
		Creator:  msg.Creator,
		Did:      msg.Did,
		Hash:     msg.Hash,
		Owner:    msg.Owner,
		Username: msg.Username,
	}

	k.SetId(ctx, id)

	return &types.MsgUpdateIdResponse{}, nil
}

func (k msgServer) DeleteId(goCtx context.Context, msg *types.MsgDeleteId) (*types.MsgDeleteIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetId(
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

	k.RemoveId(
		ctx,
		msg.Did,
	)

	return &types.MsgDeleteIdResponse{}, nil
}
