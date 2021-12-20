package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateLaunchpad{}

func NewMsgCreateLunchpad(creator string, tokenId uint64,
	denoms []string, softcap, hardcap uint64,
	startTime, endTime time.Time) *MsgCreateLaunchpad {
	return &MsgCreateLaunchpad{
		Creator:         creator,
		TokenId:         tokenId,
		AccepetedDenoms: denoms,
		Softcap:         softcap,
		Hardcap:         hardcap,
		StartTime:       startTime,
		EndTime:         endTime,
		Status:          "CREATED",
	}
}

func (m *MsgCreateLaunchpad) Route() string {
	return RouterKey
}

func (m *MsgCreateLaunchpad) Type() string {
	return "create_lunchpad"
}

func (m *MsgCreateLaunchpad) ValidateBasic() error {
	// TODO: validate msg
	return nil
}

func (m *MsgCreateLaunchpad) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgCreateLaunchpad) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

var _ sdk.Msg = &MsgDepositToLaunchpad{}

func NewMsgDepositToLaunchpad(id uint64, depositor string, amount sdk.Coins) *MsgDepositToLaunchpad {
	return &MsgDepositToLaunchpad{
		Id:        id,
		Depositor: depositor,
		Amount:    amount,
	}
}

func (m *MsgDepositToLaunchpad) Route() string {
	return RouterKey
}

func (m *MsgDepositToLaunchpad) Type() string {
	return "deposit_to_lunchpad"
}

func (m *MsgDepositToLaunchpad) ValidateBasic() error {
	// TODO: validate msg
	return nil
}

func (m *MsgDepositToLaunchpad) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDepositToLaunchpad) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
