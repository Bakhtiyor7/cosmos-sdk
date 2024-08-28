package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/Bakhtiyor7/cosmos_chain/testutil/keeper"
    "github.com/Bakhtiyor7/cosmos_chain/x/cosmoschain/types"
    "github.com/Bakhtiyor7/cosmos_chain/x/cosmoschain/keeper"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmoschainKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}