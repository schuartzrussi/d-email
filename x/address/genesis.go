package address

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/schrsi/d-email/x/address/keeper"
	"github.com/schrsi/d-email/x/address/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the address
	for _, elem := range genState.AddressList {
		k.GenerateTrackID(ctx)
		k.SetAddress(ctx, *elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all address
	addressList := k.GetAllAddress(ctx)
	for _, elem := range addressList {
		elem := elem
		genesis.AddressList = append(genesis.AddressList, &elem)
	}

	return genesis
}
