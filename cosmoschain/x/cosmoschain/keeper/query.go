package keeper

import (
	"github.com/Bakhtiyor7/cosmoschain/x/cosmoschain/types"
)

var _ types.QueryServer = Keeper{}
