package coinactions_test

import (
	"testing"

	keepertest "mychain/testutil/keeper"
	"mychain/testutil/nullify"
	coinactions "mychain/x/coinactions/module"
	"mychain/x/coinactions/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoinactionsKeeper(t)
	coinactions.InitGenesis(ctx, k, genesisState)
	got := coinactions.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
