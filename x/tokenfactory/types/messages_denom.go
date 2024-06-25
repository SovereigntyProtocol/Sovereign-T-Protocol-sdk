package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDenom{}

func NewMsgCreateDenom(
	owner string,
	denom string,
	ticker string,
	precision int32,
	maxSupply int32,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Owner:     owner,
		Denom:     denom,
		Ticker:    ticker,
		Precision: precision,
		MaxSupply: maxSupply,
	}
}

func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	tickerLength := len(msg.Ticker)
	if tickerLength < 3 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Ticker length must be at least 3 chars long")
	}
	if tickerLength > 10 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Ticker length must be 10 chars long maximum")
	}
	if msg.MaxSupply <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Max Supply must be greater than 0")
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateDenom{}

func NewMsgUpdateDenom(
	owner string,
	denom string,
	ticker string,
	precision int32,
	maxSupply int32,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Owner:     owner,
		Denom:     denom,
		Ticker:    ticker,
		Precision: precision,
		MaxSupply: maxSupply,
	}
}

func (msg *MsgUpdateDenom) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "opration is not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Owner)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	// }
	// return nil
}

var _ sdk.Msg = &MsgDeleteDenom{}

func NewMsgDeleteDenom(
	owner string,
	denom string,

) *MsgDeleteDenom {
	return &MsgDeleteDenom{
		Owner: owner,
		Denom: denom,
	}
}

func (msg *MsgDeleteDenom) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "opration is not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Owner)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	// }
	// return nil
}
