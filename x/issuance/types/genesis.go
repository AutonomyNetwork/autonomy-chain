package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TokenList: []*Token{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	tokenIdMap := make(map[uint64]bool)

	for _, elem := range gs.TokenList {
		if _, ok := tokenIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for token")
		}
		tokenIdMap[elem.Id] = true
	}

	return nil
}
