package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LaunchpadList: []*Launchpad{},
	}
}

func (gs GenesisState) Validate() error {
	launchpadIdMap := make(map[uint64]bool)

	for _, elem := range gs.LaunchpadList {
		if _, ok := launchpadIdMap[elem.TokenId]; ok {
			return fmt.Errorf("duplicated id for launchpad")
		}
		launchpadIdMap[elem.TokenId] = true
	}

	return nil
}
