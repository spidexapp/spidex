package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateCoinSymbol = "update_coin_symbol"

var _ sdk.Msg = &MsgUpdateCoinSymbol{}

func NewMsgUpdateCoinSymbol(creator string, id uint64, symbol string) *MsgUpdateCoinSymbol {
	return &MsgUpdateCoinSymbol{
		Creator: creator,
		Id:      id,
		Symbol:  symbol,
	}
}

func (msg *MsgUpdateCoinSymbol) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCoinSymbol) Type() string {
	return TypeMsgUpdateCoinSymbol
}

func (msg *MsgUpdateCoinSymbol) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCoinSymbol) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCoinSymbol) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
