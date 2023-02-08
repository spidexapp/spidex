package contracts

import (
	_ "embed" // embed compiled smart contracts
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed compiled_contracts/Oracle.json
	OracleJSON []byte // nolint: golint

	// OracleContract is the compiled oracle contracts
	OracleContract evmtypes.CompiledContract

	// OracleAddress is the erc20 module address
	OracleAddress common.Address
)

func init() {
	OracleAddress = common.HexToAddress("0x4200000000000000000000000000000000000042")

	err := json.Unmarshal(OracleJSON, &OracleContract)
	if err != nil {
		panic(err)
	}

	if len(OracleContract.Bin) == 0 {
		panic("load contracts failed")
	}
}
