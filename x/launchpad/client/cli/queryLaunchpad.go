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

func CmdCreatedLaunchpads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-created-launchpads",
		Short: "list all created launchpad",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCreatedLaunchpadRequest{}

			res, err := queryClient.CreatedLaunchpads(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdActiveLaunchpads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-active-launchpads",
		Short: "list all active launchpad",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryActiveLaunchpadRequest{}

			res, err := queryClient.ActiveLaunchpads(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdSuccessLaunchpads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-success-launchpads",
		Short: "list all success launchpad",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySuccessLaunchpadRequest{}

			res, err := queryClient.SuccessLaunchpads(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdFailLaunchpads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-fail-launchpads",
		Short: "list all fail launchpad",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFailLaunchpadRequest{}

			res, err := queryClient.FailLaunchpads(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

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
