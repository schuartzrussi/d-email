package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/schrsi/d-email/x/address/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) CreateAddress(goCtx context.Context, msg *types.MsgCreateAddress) (*types.MsgCreateAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	current := k.GetAddressByName(ctx, msg.Name)
	if current != (types.Address{}) {
		return nil, sdkerrors.Wrap(types.ErrAddressAlreadyExists, fmt.Sprintf("address %s already exists", msg.Name))
	}

	err := k.sendCoinsToModule(ctx, msg.Creator, "1token")
	if err != nil {
		return nil, err
	}

	trackID := k.GenerateTrackID(ctx)

	address := types.Address{
		Creator: msg.Creator,
		Name:    msg.Name,
		PubKey:  msg.PubKey,
		TrackID: trackID,
		Version: 1,
	}

	k.SetAddress(ctx, address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("address", sdk.NewAttribute("name", msg.Name)),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("address", sdk.NewAttribute("trackID", strconv.FormatUint(trackID, 10))),
	)

	return &types.MsgCreateAddressResponse{
		TrackID: trackID,
	}, nil
}

func (k msgServer) UpdateAddress(goCtx context.Context, msg *types.MsgUpdateAddress) (*types.MsgUpdateAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	current := k.GetAddressByName(ctx, msg.Name)
	if current == (types.Address{}) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("address %s doesn't exist", msg.Name))
	}

	if msg.Creator != current.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	err := k.sendCoinsToModule(ctx, msg.Creator, "1token")
	if err != nil {
		return nil, err
	}

	address := types.Address{
		Creator: msg.Creator,
		Name:    current.Name,
		PubKey:  msg.PubKey,
		TrackID: current.TrackID,
		Version: current.Version + 1,
	}

	k.SetAddress(ctx, address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("address", sdk.NewAttribute("name", msg.Name)),
	)

	return &types.MsgUpdateAddressResponse{
		Version: current.Version + 1,
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
