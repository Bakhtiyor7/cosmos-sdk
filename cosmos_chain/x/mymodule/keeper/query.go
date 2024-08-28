package keeper

import (
	"github.com/Bakhtiyor7/cosmos_chain/x/mymodule/types"
)

var _ types.QueryServer = Keeper{}
