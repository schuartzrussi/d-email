package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/email module sentinel errors
var (
	ErrRequiredEmailParam = sdkerrors.Register(ModuleName, 1100, "required email param")
)
