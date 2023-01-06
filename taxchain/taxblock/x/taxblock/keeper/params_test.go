package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "taxblock/testutil/keeper"
	"taxblock/x/taxblock/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TaxblockKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
