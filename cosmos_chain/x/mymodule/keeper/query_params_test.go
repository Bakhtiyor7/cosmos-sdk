package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/Bakhtiyor7/cosmos_chain/testutil/keeper"
    "github.com/Bakhtiyor7/cosmos_chain/x/mymodule/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.MymoduleKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
