package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatascriptId = "update_script_id"

var _ sdk.Msg = &MsgUpdateScriptId{}

func NewMsgUpdateScriptId(creator string, newScriptId uint64) *MsgUpdateScriptId {
	return &MsgUpdateScriptId{
		Creator:     creator,
		NewScriptId: newScriptId,
	}
}

func (msg *MsgUpdateScriptId) Route() string {
	return RouterKey
}

func (msg *MsgUpdateScriptId) Type() string {
	return TypeMsgUpdatascriptId
}

func (msg *MsgUpdateScriptId) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateScriptId) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateScriptId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
