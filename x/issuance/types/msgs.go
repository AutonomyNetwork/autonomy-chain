package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgIssueToken{}

func NewMsgIssueToken(creator, denom string, decimal uint64, initial_supply uint64, display, description string) *MsgIssueToken {
	return &MsgIssueToken{
		Creator:       creator,
		Denom:         denom,
		Decimals:      decimal,
		InitialSupply: initial_supply,
		Display: display,
		Description: description,
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
