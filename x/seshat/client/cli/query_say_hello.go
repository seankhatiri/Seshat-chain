package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"seshat/x/seshat/types"
)

var _ = strconv.Itoa(0)

func CmdSayHello() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "say_hello [name]",
		Short: "Query say-hello",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqName := args[0]
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err == nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QuerySayHelloRequest{
				Name: reqName,
			}
			res, err := queryClient.SayHello(cmd.Context(), params)

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
