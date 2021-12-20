package cli

import (
	"context"
	"strconv"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLaunchpad() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-launchpads",
		Short: "list all launchpad",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLaunchpadRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LaunchpadAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowLaunchpad() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-launchpad [id]",
		Short: "shows a launchpad",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetLaunchpadRequest{
				Id: id,
			}

			res, err := queryClient.Launchpad(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
