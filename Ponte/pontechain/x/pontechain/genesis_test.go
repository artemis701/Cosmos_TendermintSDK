package pontechain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "pontechain/testutil/keeper"
	"pontechain/testutil/nullify"
	"pontechain/x/pontechain"
	"pontechain/x/pontechain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PontechainKeeper(t)
	pontechain.InitGenesis(ctx, *k, genesisState)
	got := pontechain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
