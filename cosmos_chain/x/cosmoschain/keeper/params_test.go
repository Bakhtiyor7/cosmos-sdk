package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/Bakhtiyor7/cosmos_chain/testutil/keeper"
    "github.com/Bakhtiyor7/cosmos_chain/x/cosmoschain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CosmoschainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
