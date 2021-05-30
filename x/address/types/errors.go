package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/address module sentinel errors
var (
	ErrAddressAlreadyExists = sdkerrors.Register(ModuleName, 1100, "address already exists")
	ErrInvalidEmailAddress  = sdkerrors.Register(ModuleName, 1101, "invalid email address")
)
