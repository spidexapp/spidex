package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdataMultiplier = "update_multiplier"

var _ sdk.Msg = &MsgUpdateMultiplier{}

func NewMsgUpdateMultiplier(creator string, newMultiplier uint64) *MsgUpdateMultiplier {
	return &MsgUpdateMultiplier{
		Creator:       creator,
		NewMultiplier: newMultiplier,
	}
}

func (msg *MsgUpdateMultiplier) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMultiplier) Type() string {
	return TypeMsgUpdataMultiplier
}

func (msg *MsgUpdateMultiplier) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMultiplier) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMultiplier) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
