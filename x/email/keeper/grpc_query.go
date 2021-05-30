package keeper

import (
	"github.com/schrsi/d-email/x/email/types"
)

var _ types.QueryServer = Keeper{}
