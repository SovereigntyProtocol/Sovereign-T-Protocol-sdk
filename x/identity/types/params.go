package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeySuperwallet = []byte("Superwallet")
	// TODO: Determine the default value
	DefaultSuperwallet string = "ssi18evc5mvlvlnmywrxw7avx9dkrj0zrrhmhp7nrj"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	superwallet string,
) Params {
	return Params{
		Superwallet: superwallet,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultSuperwallet,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySuperwallet, &p.Superwallet, validateSuperwallet),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateSuperwallet(p.Superwallet); err != nil {
		return err
	}
	return nil
}

// validateSuperwallet validates the Superwallet param
func validateSuperwallet(v interface{}) error {
	superwallet, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = superwallet

	return nil
}
