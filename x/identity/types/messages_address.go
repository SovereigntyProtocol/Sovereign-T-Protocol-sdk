package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAddress{}

func NewMsgCreateAddress(
	creator string,
	owner string,

) *MsgCreateAddress {
	return &MsgCreateAddress{
		Creator: creator,
		Owner:   owner,
	}
}

func (msg *MsgCreateAddress) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
	// return nil
}

var _ sdk.Msg = &MsgUpdateAddress{}

func NewMsgUpdateAddress(
	creator string,
	owner string,

) *MsgUpdateAddress {
	return &MsgUpdateAddress{
		Creator: creator,
		Owner:   owner,
	}
}

func (msg *MsgUpdateAddress) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
	// return nil
}

var _ sdk.Msg = &MsgDeleteAddress{}

func NewMsgDeleteAddress(
	creator string,
	owner string,

) *MsgDeleteAddress {
	return &MsgDeleteAddress{
		Creator: creator,
		Owner:   owner,
	}
}

func (msg *MsgDeleteAddress) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
	// return nil
}
