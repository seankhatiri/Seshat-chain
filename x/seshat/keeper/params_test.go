package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "seshat/testutil/keeper"
	"seshat/x/seshat/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SeshatKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
