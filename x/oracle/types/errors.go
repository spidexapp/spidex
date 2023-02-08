package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")

	ErrInvalidVersion = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrUnexpectedEvent  = sdkerrors.Register(ModuleName, 1, "unexpected event")
	ErrABIPack          = sdkerrors.Register(ModuleName, 2, "contracts ABI pack failed")
	ErrCoinSymbolExists = sdkerrors.Register(ModuleName, 3, "coin symbol already exists")
	ErrCreatorNotAdmin  = sdkerrors.Register(ModuleName, 4, "creator is not admin account")
	ErrNotTypeOfAddress = sdkerrors.Register(ModuleName, 5, "address not a bech32 address")
)
