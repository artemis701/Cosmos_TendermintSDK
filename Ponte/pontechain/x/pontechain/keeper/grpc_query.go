package keeper

import (
	"pontechain/x/pontechain/types"
)

var _ types.QueryServer = Keeper{}
