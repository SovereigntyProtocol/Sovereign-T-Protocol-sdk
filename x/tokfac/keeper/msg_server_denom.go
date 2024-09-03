package keeper

import (
	"context"
	"math/big"

	"github.com/cosmos/cosmos-sdk/x/tokfac/types"

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

	minttokenpointer, ok := new(big.Int).SetString(msg.Maxsupply, 10)

	if !ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "big supply error")
	}

	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewIntFromBigInt(minttokenpointer)))

	sdkError := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)

	if sdkError != nil {
		return nil, errorsmod.Wrap(sdkError, "Mint coin error")
	}

	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	sdkError2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, mintCoins)

	if sdkError2 != nil {
		return nil, errorsmod.Wrap(sdkError, "error while sending coins")
	}

	// bankkeeper.BaseKeeper(bankkeeper.BaseKeeper{}).SetDenomMetaData(ctx, newMetadata)

	var denom = types.Denom{
		Creator:     msg.Creator,
		Denom:       msg.Denom,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Description: msg.Description,
		Maxsupply:   msg.Maxsupply,
		Denomunit:   msg.Denomunit,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var denom = types.Denom{
		Creator:     msg.Description,
		Denom:       valFound.Denom,
		Name:        valFound.Name,
		Symbol:      valFound.Symbol,
		Description: valFound.Description,
		Maxsupply:   valFound.Maxsupply,
		Denomunit:   valFound.Denomunit,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}

func (k msgServer) DeleteDenom(goCtx context.Context, msg *types.MsgDeleteDenom) (*types.MsgDeleteDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if msg.Creator == valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "not allowed")
	}

	k.RemoveDenom(
		ctx,
		msg.Denom,
	)

	return &types.MsgDeleteDenomResponse{}, nil
}
