package keeper

import (
	"github.com/Bakhtiyor7/cosmos_chain/x/cosmoschain/types"
)

var _ types.QueryServer = Keeper{}
