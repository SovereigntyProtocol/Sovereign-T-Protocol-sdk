package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UserList:    []User{},
		AddressList: []Address{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in user
	userIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserList {
		index := string(UserKey(elem.Did))
		if _, ok := userIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for user")
		}
		userIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in address
	addressIndexMap := make(map[string]struct{})

	for _, elem := range gs.AddressList {
		index := string(AddressKey(elem.Owner))
		if _, ok := addressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for address")
		}
		addressIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
