package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}
	var mintCoins sdk.Coins

	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.MaxSupply))))

	sdkError := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)

	if sdkError != nil {
		return nil, errorsmod.Wrap(sdkError, "Mint coin error")
	}

	lender, _ := sdk.AccAddressFromBech32(msg.Owner)
	sdkError2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, mintCoins)

	if sdkError2 != nil {
		return nil, errorsmod.Wrap(sdkError, "error while sending coins")
	}

	var denom = types.Denom{
		Owner:     msg.Owner,
		Denom:     msg.Denom,
		Ticker:    msg.Ticker,
		Precision: msg.Precision,
		MaxSupply: msg.MaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)

	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "this operation is not allowed")
	// Check if the value exists
	// valFound, isFound := k.GetDenom(
	// 	ctx,
	// 	msg.Denom,
	// )
	// if !isFound {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	// }

	// Checks if the msg owner is the same as the current owner
	// if msg.Owner != valFound.Owner {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// var denom = types.Denom{
	// 	Owner:     msg.Owner,
	// 	Denom:     msg.Denom,
	// 	Ticker:    msg.Ticker,
	// 	Precision: msg.Precision,
	// 	MaxSupply: msg.MaxSupply,
	// }

	// k.SetDenom(ctx, denom)

	// return &types.MsgUpdateDenomResponse{}, nil
}

func (k msgServer) DeleteDenom(goCtx context.Context, msg *types.MsgDeleteDenom) (*types.MsgDeleteDenomResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "this operation is not allowed")
	// Check if the value exists
	// valFound, isFound := k.GetDenom(
	// 	ctx,
	// 	msg.Denom,
	// )
	// if !isFound {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	// }

	// // Checks if the msg owner is the same as the current owner
	// if msg.Owner != valFound.Owner {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// k.RemoveDenom(
	// 	ctx,
	// 	msg.Denom,
	// )

	// return &types.MsgDeleteDenomResponse{}, nil
}
