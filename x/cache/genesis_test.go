package cache_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testchain/testutil/keeper"
	"testchain/testutil/nullify"
	"testchain/x/cache"
	"testchain/x/cache/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CacheKeeper(t)
	cache.InitGenesis(ctx, *k, genesisState)
	got := cache.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
