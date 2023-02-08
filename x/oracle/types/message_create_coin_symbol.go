package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateCoinSymbol = "create_coin_symbol"

var _ sdk.Msg = &MsgCreateCoinSymbol{}

func NewMsgCreateCoinSymbol(creator string, symbol string) *MsgCreateCoinSymbol {
	return &MsgCreateCoinSymbol{
		Creator: creator,
		Symbol:  symbol,
	}
}

func (msg *MsgCreateCoinSymbol) Route() string {
	return RouterKey
}

func (msg *MsgCreateCoinSymbol) Type() string {
	return TypeMsgCreateCoinSymbol
}

func (msg *MsgCreateCoinSymbol) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCoinSymbol) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCoinSymbol) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
