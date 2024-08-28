package types

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDenom{}

func NewMsgCreateDenom(
	creator string,
	denom string,
	name string,
	symbol string,
	description string,
	maxsupply string,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Creator:     creator,
		Denom:       denom,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Maxsupply:   maxsupply,
	}
}

func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Denom == sdk.DefaultBondDenom {
		return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "not allowed")
	}

	minttokenpointer, ok := new(big.Int).SetString(msg.Maxsupply, 10)

	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "big supply error")
	}

	if math.NewIntFromBigInt(minttokenpointer).IsNegative() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "supply is negative")
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateDenom{}

func NewMsgUpdateDenom(
	creator string,
	denom string,
	name string,
	symbol string,
	description string,
	maxsupply string,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Creator:     creator,
		Denom:       denom,
		Name:        name,
		Symbol:      symbol,
		Description: description,
		Maxsupply:   maxsupply,
	}
}

func (msg *MsgUpdateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, updatecreatorerr := sdk.AccAddressFromBech32(msg.Description)

	if updatecreatorerr != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid  address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteDenom{}

func NewMsgDeleteDenom(
	creator string,
	denom string,

) *MsgDeleteDenom {
	return &MsgDeleteDenom{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgDeleteDenom) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
}
