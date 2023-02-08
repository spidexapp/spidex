package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdataChannel = "updata_channel"

var _ sdk.Msg = &MsgUpdateChannel{}

func NewMsgUpdateChannel(creator string, newChannel string) *MsgUpdateChannel {
	return &MsgUpdateChannel{
		Creator:    creator,
		NewChannel: newChannel,
	}
}

func (msg *MsgUpdateChannel) Route() string {
	return RouterKey
}

func (msg *MsgUpdateChannel) Type() string {
	return TypeMsgUpdataChannel
}

func (msg *MsgUpdateChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateChannel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
