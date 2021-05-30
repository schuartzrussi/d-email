package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AddressList: []*Address{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated address
	addressIdMap := make(map[string]bool)

	for _, elem := range gs.AddressList {
		key := elem.Name + string(elem.Version)
		if _, ok := addressIdMap[key]; ok {
			return fmt.Errorf("duplicated id for address")
		}
		addressIdMap[key] = true
	}

	return nil
}
