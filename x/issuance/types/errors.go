package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/issuance module sentinel errors
var (
	ErrDenomAlreadyExist = sdkerrors.Register(ModuleName, 2,"Denom already exist")
	ErrInvalidDecimals = sdkerrors.Register(ModuleName,3, "Invalid decimals")
	ErrInvalidDenom = sdkerrors.Register(ModuleName,4,"Invalid denom")
)
