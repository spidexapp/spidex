package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"

	// Version defines the current version the IBC module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to
	PortID = "oracle"
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address
var AccAddress sdk.AccAddress

func init() {
	AccAddress = authtypes.NewModuleAddress(ModuleName)
	ModuleAddress = common.BytesToAddress(AccAddress.Bytes())
}

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("oracle-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CoinSymbolKey      = "CoinSymbol-value-"
	CoinSymbolCountKey = "CoinSymbol-count-"
)
