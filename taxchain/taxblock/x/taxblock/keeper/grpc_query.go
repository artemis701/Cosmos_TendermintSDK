package keeper

import (
	"taxblock/x/taxblock/types"
)

var _ types.QueryServer = Keeper{}
