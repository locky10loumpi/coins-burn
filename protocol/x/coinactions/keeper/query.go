package keeper

import (
	"mychain/x/coinactions/types"
)

var _ types.QueryServer = Keeper{}
