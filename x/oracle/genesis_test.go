package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/spidexapp/spidex/testutil/keeper"
	"github.com/spidexapp/spidex/testutil/nullify"
	"github.com/spidexapp/spidex/x/oracle"
	"github.com/spidexapp/spidex/x/oracle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)
	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
