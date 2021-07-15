package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/schrsi/d-email/x/email/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) CreateEmail(goCtx context.Context, msg *types.MsgCreateEmail) (*types.MsgCreateEmailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.ReplyTo != "" {
		if !k.HasEmail(ctx, msg.ReplyTo) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("cannot find replied email: %s", msg.ReplyTo))
		}
	}

	if k.HasEmail(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(types.ErrInvalidEmailId, fmt.Sprintf("invalid id: %s", msg.Id))
	}

	err := k.sendCoinsToModule(ctx, msg.Creator, "20token")
	if err != nil {
		return nil, err
	}

	email := types.Email{
		Id:                    msg.Id,
		Creator:               msg.Creator,
		From:                  msg.From,
		To:                    msg.To,
		SenderSignature:       msg.SenderSignature,
		Subject:               msg.Subject,
		Body:                  msg.Body,
		ReplyTo:               msg.ReplyTo,
		TrackIds:              msg.TrackIds,
		SendedAt:              msg.SendedAt,
		DecryptionKeys:        msg.DecryptionKeys,
		PreviousDecryptionKey: msg.PreviousDecryptionKey,
		SenderAddressVersion:  msg.SenderAddressVersion,
	}

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

func (k msgServer) sendCoinsToModule(ctx sdk.Context, creator string, quantity string) error {
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	feeCoins, err := sdk.ParseCoinsNormalized(quantity)
	if err != nil {
		return err
	}

	creatorAddress, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoins(ctx, creatorAddress, moduleAcct, feeCoins); err != nil {
		return err
	}

	return nil
}
