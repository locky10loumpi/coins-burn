package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "mychain/testutil/keeper"
	"mychain/x/coinactions/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CoinactionsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
