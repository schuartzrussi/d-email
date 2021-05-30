package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateEmail{}

func NewMsgCreateEmail(creator string, from string, to string, senderSignature string, senderAddressVersion uint64, subject string, body string, attachments []string, replyTo string, trackIds []string, sendedAt string, decryptionKeys []string, previousDecryptionKey string) *MsgCreateEmail {
	return &MsgCreateEmail{
		Creator:               creator,
		From:                  from,
		To:                    to,
		SenderSignature:       senderSignature,
		SenderAddressVersion:  senderAddressVersion,
		Subject:               subject,
		Body:                  body,
		Attachments:           attachments,
		ReplyTo:               replyTo,
		TrackIds:              trackIds,
		SendedAt:              sendedAt,
		DecryptionKeys:        decryptionKeys,
		PreviousDecryptionKey: previousDecryptionKey,
	}
}

func (msg *MsgCreateEmail) Route() string {
	return RouterKey
}

func (msg *MsgCreateEmail) Type() string {
	return "CreateEmail"
}

func (msg *MsgCreateEmail) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEmail) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEmail) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.From == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'from' field is required")
	}

	if msg.To == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'to' field is required")
	}

	if msg.Subject == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'subject' field is required")
	}

	if msg.Body == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'body' field is required")
	}

	if msg.SendedAt == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'sendedAt' field is required")
	}

	if msg.SenderSignature == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'senderSignature' field is required")
	}

	if msg.SenderAddressVersion == 0 {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'senderAddressVersion' field is required")
	}

	if len(msg.TrackIds) == 0 {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'trackIds' field is required")
	}

	if len(msg.DecryptionKeys) == 0 {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'decryptionKeys' field is required")
	}

	if msg.ReplyTo != "" && msg.PreviousDecryptionKey == "" {
		return sdkerrors.Wrapf(ErrRequiredEmailParam, "'previousDecryptionKey' field is required")
	}

	return nil
}

var _ sdk.Msg = &MsgCreateEmail{}
