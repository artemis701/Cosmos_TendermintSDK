package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "pontechain/testutil/keeper"
	"pontechain/x/pontechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PontechainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
