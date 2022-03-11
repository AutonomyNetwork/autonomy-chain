package types

import (
	
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgIssueToken{}

func NewMsgIssueToken(creator, denom, display_name string, decimal uint64, initial_supply uint64) *MsgIssueToken {
	return &MsgIssueToken{
		Creator:       creator,
		Denom:         denom,
		DisplayName:   display_name,
		Decimals:      decimal,
		InitialSupply: initial_supply,
	}
}

func (m *MsgIssueToken) Route() string {
	return RouterKey
}

func (m *MsgIssueToken) Type() string {
	return "issue_token"
}

func (m *MsgIssueToken) ValidateBasic() error {
	// TODO: validate msg
	if m.Decimals<1 {
		return ErrInvalidDecimals
	}
	
	return nil
}

func (m *MsgIssueToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgIssueToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
