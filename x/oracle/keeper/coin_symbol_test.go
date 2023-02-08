package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/spidexapp/spidex/testutil/keeper"
	"github.com/spidexapp/spidex/testutil/nullify"
	"github.com/spidexapp/spidex/x/oracle/keeper"
	"github.com/spidexapp/spidex/x/oracle/types"
)

func createNCoinSymbol(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CoinSymbol {
	items := make([]types.CoinSymbol, n)
	for i := range items {
		items[i].Id = keeper.AppendCoinSymbol(ctx, items[i])
	}
	return items
}

func TestCoinSymbolGet(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNCoinSymbol(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCoinSymbol(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCoinSymbolRemove(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNCoinSymbol(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoinSymbol(ctx, item.Id)
		_, found := keeper.GetCoinSymbol(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCoinSymbolGetAll(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNCoinSymbol(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCoinSymbol(ctx)),
	)
}

func TestCoinSymbolCount(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNCoinSymbol(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCoinSymbolCount(ctx))
}
