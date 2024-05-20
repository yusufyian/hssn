package hssn_test

import (
	"testing"

	keepertest "hssn/testutil/keeper"
	"hssn/testutil/nullify"
	hssn "hssn/x/hssn/module"
	"hssn/x/hssn/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HssnKeeper(t)
	hssn.InitGenesis(ctx, k, genesisState)
	got := hssn.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
