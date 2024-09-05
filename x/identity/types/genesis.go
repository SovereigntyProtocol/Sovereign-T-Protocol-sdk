package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		IdList:        []Id{},
		UniquekeyList: []Uniquekey{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in id
	idIndexMap := make(map[string]struct{})

	for _, elem := range gs.IdList {
		index := string(IdKey(elem.Did))
		if _, ok := idIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for id")
		}
		idIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in uniquekey
	uniquekeyIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniquekeyList {
		index := string(UniquekeyKey(elem.Key))
		if _, ok := uniquekeyIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for uniquekey")
		}
		uniquekeyIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
