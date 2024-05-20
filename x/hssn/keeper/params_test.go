package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hssn/testutil/keeper"
	"hssn/x/hssn/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HssnKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
