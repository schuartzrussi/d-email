package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/schrsi/d-email/x/email/types"
)

func (k Keeper) SetEmail(ctx sdk.Context, email types.Email) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmailKey))
	b := k.cdc.MustMarshalBinaryBare(&email)
	store.Set(GetEmailIDBytes(email.Id), b)
}

func (k Keeper) GetEmail(ctx sdk.Context, id string) types.Email {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmailKey))
	var email types.Email
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetEmailIDBytes(id)), &email)
	return email
}

func (k Keeper) HasEmail(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmailKey))
	return store.Has(GetEmailIDBytes(id))
}

func (k Keeper) GetEmailOwner(ctx sdk.Context, id string) string {
	return k.GetEmail(ctx, id).Creator
}

func (k Keeper) GetAllEmail(ctx sdk.Context) (list []types.Email) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmailKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Email
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetEmailIDBytes(id string) []byte {
	return []byte(id)
}
