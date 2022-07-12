package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

func CmdIssueToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-token [denom] [decimals] [display_name] [initial-supply] ",
		Short: "Issue a new token",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenom, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsDecimals, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argsDisplayName, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsInitial_Supply, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueToken(clientCtx.GetFromAddress().String(), argsDenom, argsDisplayName, argsDecimals, argsInitial_Supply)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
