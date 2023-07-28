package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "testchain/testutil/keeper"
	"testchain/x/cache/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CacheKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
