package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	ParamStoreKeyAdmin         = []byte("Admin")
	ParamStoreKeySourceChannel = []byte("EnableEVMHook")
	ParamStoreKeyMultiplier    = []byte("Multiplier")
	ParamStoreKeyScriptId      = []byte("ScriptId")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(admin, sourceChannel string, multiplier uint64) Params {
	return Params{
		Admin:         admin,
		SourceChannel: sourceChannel,
		Multiplier:    multiplier,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", "", 0)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyAdmin, &p.Admin, validateString),
		paramtypes.NewParamSetPair(ParamStoreKeySourceChannel, &p.SourceChannel, validateString),
		paramtypes.NewParamSetPair(ParamStoreKeyMultiplier, &p.Multiplier, validateUint64),
		paramtypes.NewParamSetPair(ParamStoreKeyScriptId, &p.ScriptId, validateUint64),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

func validateUint64(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateString(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
