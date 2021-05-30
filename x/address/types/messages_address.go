package types

import (
	"net/mail"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAddress{}

func NewMsgCreateAddress(creator string, name string, pubKey string) *MsgCreateAddress {
	return &MsgCreateAddress{
		Creator: creator,
		Name:    name,
		PubKey:  pubKey,
	}
}

func (msg *MsgCreateAddress) Route() string {
	return RouterKey
}

func (msg *MsgCreateAddress) Type() string {
	return "CreateAddress"
}

func (msg *MsgCreateAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !IsValidEmailAddress(msg.Name) {
		return sdkerrors.Wrapf(ErrInvalidEmailAddress, "invalid email address (%s)", msg.Name)
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateAddress{}

func NewMsgUpdateAddress(creator string, name string, pubKey string) *MsgUpdateAddress {
	return &MsgUpdateAddress{
		Creator: creator,
		Name:    name,
		PubKey:  pubKey,
	}
}

func (msg *MsgUpdateAddress) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAddress) Type() string {
	return "UpdateAddress"
}

func (msg *MsgUpdateAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !IsValidEmailAddress(msg.Name) {
		return sdkerrors.Wrapf(ErrInvalidEmailAddress, "invalid email address (%s)", msg.Name)
	}

	return nil
}

var _ sdk.Msg = &MsgCreateAddress{}

func IsValidEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
