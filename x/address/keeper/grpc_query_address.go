package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/schrsi/d-email/x/address/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressAll(c context.Context, req *types.QueryAllAddressRequest) (*types.QueryAllAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addresss []*types.Address
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	addressStore := prefix.NewStore(store, types.KeyPrefix(types.AddressKey))

	pageRes, err := query.Paginate(addressStore, req.Pagination, func(key []byte, value []byte) error {
		var address types.Address
		if err := k.cdc.UnmarshalBinaryBare(value, &address); err != nil {
			return err
		}

		addresss = append(addresss, &address)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressResponse{Address: addresss, Pagination: pageRes}, nil
}

func (k Keeper) AddressByName(c context.Context, req *types.QueryGetAddressByNameRequest) (*types.QueryGetAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	address := k.GetAddressByName(ctx, req.Name)
	if address == (types.Address{}) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAddressResponse{Address: &address}, nil
}

func (k Keeper) AddressByVersion(c context.Context, req *types.QueryGetAddressByVersionRequest) (*types.QueryGetAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var address types.Address
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAddress(ctx, req.Name, req.Version) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressStoreKeyInBytes(req.Name, req.Version)), &address)

	return &types.QueryGetAddressResponse{Address: &address}, nil
}
