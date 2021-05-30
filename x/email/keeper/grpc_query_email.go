package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/schrsi/d-email/x/email/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EmailAll(c context.Context, req *types.QueryAllEmailRequest) (*types.QueryAllEmailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var emails []*types.Email
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	emailStore := prefix.NewStore(store, types.KeyPrefix(types.EmailKey))

	pageRes, err := query.Paginate(emailStore, req.Pagination, func(key []byte, value []byte) error {
		var email types.Email
		if err := k.cdc.UnmarshalBinaryBare(value, &email); err != nil {
			return err
		}

		emails = append(emails, &email)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEmailResponse{Email: emails, Pagination: pageRes}, nil
}

func (k Keeper) EmailById(c context.Context, req *types.QueryGetEmailByIdRequest) (*types.QueryGetEmailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var email types.Email
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasEmail(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmailKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetEmailIDBytes(req.Id)), &email)

	return &types.QueryGetEmailResponse{Email: &email}, nil
}
