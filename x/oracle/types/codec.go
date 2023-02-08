package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCoinRatesData{}, "oracle/CoinRatesData", nil)
	cdc.RegisterConcrete(&MsgCreateCoinSymbol{}, "oracle/CreateCoinSymbol", nil)
	cdc.RegisterConcrete(&MsgUpdateCoinSymbol{}, "oracle/UpdateCoinSymbol", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCoinRatesData{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCoinSymbol{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateCoinSymbol{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
