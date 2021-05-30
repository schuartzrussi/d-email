package email

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/schrsi/d-email/x/email/keeper"
	"github.com/schrsi/d-email/x/email/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the email
	for _, elem := range genState.EmailList {
		k.SetEmail(ctx, *elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	emailList := k.GetAllEmail(ctx)
	for _, elem := range emailList {
		elem := elem
		genesis.EmailList = append(genesis.EmailList, &elem)
	}

	return genesis
}
