package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateId{}

func NewMsgCreateId(
	creator string,
	did string,
	hash string,
	owner string,
	username string,

) *MsgCreateId {
	return &MsgCreateId{
		Creator:  creator,
		Did:      did,
		Hash:     hash,
		Owner:    owner,
		Username: username,
	}
}

func (msg *MsgCreateId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateId{}

func NewMsgUpdateId(
	creator string,
	did string,
	hash string,
	owner string,
	username string,

) *MsgUpdateId {
	return &MsgUpdateId{
		Creator:  creator,
		Did:      did,
		Hash:     hash,
		Owner:    owner,
		Username: username,
	}
}

func (msg *MsgUpdateId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteId{}

func NewMsgDeleteId(
	creator string,
	did string,

) *MsgDeleteId {
	return &MsgDeleteId{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgDeleteId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
