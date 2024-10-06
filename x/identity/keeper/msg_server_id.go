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

	txtime := ctx.BlockTime().String()

	newdid := k.generateShortDeterministicUserID(msg.Creator+msg.Username+txtime, 40)

	_, isFound := k.GetIdByDidorUsernameorCreator(ctx, newdid)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	_, isUserNameFound := k.GetIdByDidorUsernameorCreator(ctx, msg.Username)
	if isUserNameFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "username already set")
	}

	_, isCreatorFound := k.GetIdByDidorUsernameorCreator(ctx, msg.Creator)
	if isCreatorFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "wallet already set")
	}

	var id = types.Id{
		Creator:  msg.Creator,
		Did:      newdid,
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
	var owner string
	// Check if the value exists
	valFound, isFound := k.GetIdByDidorUsernameorCreator(
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
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "hash invalid")
	}

	_, isCreatorFound := k.GetIdByDidorUsernameorCreator(ctx, msg.Owner)
	if isCreatorFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "wallet already set")
	}

	if msg.Owner == "" {
		owner = valFound.Creator
	}

	if msg.Owner != "" {
		_, Ownererr := sdk.AccAddressFromBech32(msg.Owner)
		if Ownererr != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid new creator address")
		} else {
			owner = msg.Owner
		}
	}

	oldCreator := valFound.Creator
	var id = types.Id{
		Creator:  owner,
		Did:      valFound.Did,
		Hash:     msg.Hash,
		Username: valFound.Username,
	}

	k.UpdateUserId(ctx, id, oldCreator)

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

	// k.RemoveId(
	// 	ctx,
	// 	msg.Did,
	// )

	return &types.MsgDeleteIdResponse{}, nil
}
