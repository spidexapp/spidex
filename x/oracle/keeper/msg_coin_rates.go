package keeper

import (
	"context"
	"time"

	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"

	"github.com/spidexapp/spidex/x/oracle/types"
)

// CoinRatesData creates the CoinRates packet
// data with obi encoded and send it to the channel
func (k msgServer) CoinRatesData(goCtx context.Context, msg *types.MsgCoinRatesData) (*types.MsgCoinRatesDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	sourcePort := types.PortID
	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, params.SourceChannel)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port %s",
			params.SourceChannel,
			sourcePort,
		)
	}
	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, params.SourceChannel)
	if !found {
		return nil, sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, params.SourceChannel)
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, params.SourceChannel))
	if !ok {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound,
			"module does not own channel capability")
	}

	allSymbols := k.GetAllCoinSymbol(ctx)
	var symbols []string
	for _, symbol := range allSymbols {
		symbols = append(symbols, symbol.Symbol)
	}

	coinRate := &types.CoinRatesCallData{
		Symbols:    symbols,
		Multiplier: params.Multiplier,
	}

	encodedCalldata := obi.MustEncode(*coinRate)

	packetData := packet.NewOracleRequestPacketData(
		msg.ClientID,
		params.ScriptId,
		encodedCalldata,
		msg.AskCount,
		msg.MinCount,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)

	err := k.ChannelKeeper.SendPacket(ctx, channelCap, channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		params.SourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.NewHeight(0, 0),
		uint64(ctx.BlockTime().UnixNano()+int64(10*time.Minute)), // Arbitrary timestamp timeout for now
	))
	if err != nil {
		return nil, err
	}
	return &types.MsgCoinRatesDataResponse{}, nil
}

func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.PortKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.ScopedKeeper.ClaimCapability(ctx, cap, name)
}

// ChanCloseInit defines a wrapper function for the channel Keeper's function
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.ScopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.ChannelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.ScopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// GetPort returns the portID for the module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(k.portKey))
}

// SetPort sets the portID for the module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(k.portKey, []byte(portID))
}
