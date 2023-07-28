package keeper

import (
	"testchain/x/cache/types"
)

var _ types.QueryServer = Keeper{}
