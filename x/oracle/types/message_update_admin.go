package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdataAdmin = "updata_admin"

var _ sdk.Msg = &MsgUpdateAdmin{}

func NewMsgUpdateAdmin(creator string, newAdmin string) *MsgUpdateAdmin {
	return &MsgUpdateAdmin{
		Creator:  creator,
		NewAdmin: newAdmin,
	}
}

func (msg *MsgUpdateAdmin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdmin) Type() string {
	return TypeMsgUpdataAdmin
}

func (msg *MsgUpdateAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
