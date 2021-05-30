package keeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/schrsi/d-email/x/address/types"
)

const (
	maxTrackIDQuantity = 2
)

func (k Keeper) SetAddress(ctx sdk.Context, address types.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	b := k.cdc.MustMarshalBinaryBare(&address)
	store.Set(GetAddressStoreKeyInBytes(address.Name, address.Version), b)
}

func (k Keeper) GetAddressByName(ctx sdk.Context, name string) types.Address {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	var address types.Address

	iterator := sdk.KVStorePrefixIterator(store, []byte(name+"-"))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Address
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)

		if val.Version > address.Version {
			address = val
		}
	}

	return address
}

func (k Keeper) HasAddress(ctx sdk.Context, name string, version uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	return store.Has(GetAddressStoreKeyInBytes(name, version))
}

func (k Keeper) GetAddressByVersion(ctx sdk.Context, name string, version uint64) types.Address {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	var address types.Address
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddressStoreKeyInBytes(name, version)), &address)
	return address
}

func (k Keeper) GetAllAddress(ctx sdk.Context) (list []types.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Address
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func GetAddressStoreKeyInBytes(name string, version uint64) []byte {
	key := fmt.Sprintf("%s-%d", name, version)
	return []byte(key)
}

func (k Keeper) GenerateTrackID(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressTrackIDCountKey))

	currentByteValue := store.Get([]byte("value"))
	currentQuantityByteValue := store.Get([]byte("quantity"))

	if currentByteValue == nil || currentQuantityByteValue == nil {
		store.Set([]byte("quantity"), []byte(strconv.FormatUint(1, 10)))
		store.Set([]byte("value"), []byte(strconv.FormatUint(1, 10)))

		return 1
	}

	currentValue, err := strconv.ParseUint(string(currentByteValue), 10, 64)
	if err != nil {
		panic("cannot decode currentValue")
	}

	currentQuantity, err := strconv.ParseUint(string(currentQuantityByteValue), 10, 64)
	if err != nil {
		panic("cannot decode currentQuantity")
	}

	if currentQuantity >= maxTrackIDQuantity {
		store.Set([]byte("quantity"), []byte(strconv.FormatUint(1, 10)))
		store.Set([]byte("value"), []byte(strconv.FormatUint(currentValue+1, 10)))

		return currentValue + 1
	}

	store.Set([]byte("quantity"), []byte(strconv.FormatUint(currentQuantity+1, 10)))
	return currentValue
}
