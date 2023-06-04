package seshat_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "seshat/testutil/keeper"
	"seshat/testutil/nullify"
	"seshat/x/seshat"
	"seshat/x/seshat/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SeshatKeeper(t)
	seshat.InitGenesis(ctx, *k, genesisState)
	got := seshat.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
