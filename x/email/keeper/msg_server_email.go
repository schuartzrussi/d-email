package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/schrsi/d-email/x/email/types"
)

func (k msgServer) CreateEmail(goCtx context.Context, msg *types.MsgCreateEmail) (*types.MsgCreateEmailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.ReplyTo != "" {
		if !k.HasEmail(ctx, msg.ReplyTo) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("cannot find replied email: %s", msg.ReplyTo))
		}
	}

	email := types.Email{
		Creator:               msg.Creator,
		From:                  msg.From,
		To:                    msg.To,
		SenderSignature:       msg.SenderSignature,
		Subject:               msg.Subject,
		Body:                  msg.Body,
		Attachments:           msg.Attachments,
		ReplyTo:               msg.ReplyTo,
		TrackIds:              msg.TrackIds,
		SendedAt:              msg.SendedAt,
		DecryptionKeys:        msg.DecryptionKeys,
		PreviousDecryptionKey: msg.PreviousDecryptionKey,
		SenderAddressVersion:  msg.SenderAddressVersion,
	}

	email.Id = k.GenerateID(email)

	k.SetEmail(ctx, email)

	events := []sdk.Event{
		sdk.NewEvent("email", sdk.NewAttribute("id", email.Id)),
	}

	for _, trackID := range msg.TrackIds {
		events = append(events, sdk.NewEvent("email", sdk.NewAttribute("trackID", trackID)))
	}

	ctx.EventManager().EmitEvents(events)

	return &types.MsgCreateEmailResponse{
		Id: email.Id,
	}, nil
}
