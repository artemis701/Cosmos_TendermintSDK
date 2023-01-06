package taxblock_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "taxblock/testutil/keeper"
	"taxblock/testutil/nullify"
	"taxblock/x/taxblock"
	"taxblock/x/taxblock/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TaxblockKeeper(t)
	taxblock.InitGenesis(ctx, *k, genesisState)
	got := taxblock.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
