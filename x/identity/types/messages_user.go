package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgCreateUser(
	creator string,
	did string,
	hash string,
	owner string,

) *MsgCreateUser {
	return &MsgCreateUser{
		Creator: creator,
		Did:     did,
		Hash:    hash,
		Owner:   owner,
	}
}

func (msg *MsgCreateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, diderr := VerifyDidFormat(msg.Did)
	if diderr != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid did (%s)", diderr)
	}

	if msg.Hash == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid user's hash")
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateUser{}

func NewMsgUpdateUser(
	creator string,
	did string,
	hash string,
	owner string,

) *MsgUpdateUser {
	return &MsgUpdateUser{
		Creator: creator,
		Did:     did,
		Hash:    hash,
		Owner:   owner,
	}
}

func (msg *MsgUpdateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Hash == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid user's hash")
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUser{}

func NewMsgDeleteUser(
	creator string,
	did string,

) *MsgDeleteUser {
	return &MsgDeleteUser{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgDeleteUser) ValidateBasic() error {
	return errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "not allowed")
	// _, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	// }
	// return nil
}
