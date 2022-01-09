package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
)

func CmdCreateLaunchpad() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Crate a new launchpad",
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenId, _ := cmd.Flags().GetUint64(FlagTokenID)
			supply, _ := cmd.Flags().GetUint64(FlagSupply)
			softcap, _ := cmd.Flags().GetUint64(FlagSoftCap)
			hardcap, _ := cmd.Flags().GetUint64(FlagHardCap)

			if softcap > hardcap {
				return fmt.Errorf("softcap shoud be less then hardcap")
			}

			startTime, _ := cmd.Flags().GetString(FlagStartTime)
			endTime, _ := cmd.Flags().GetString(FlagHardCap)

			// str := "2014-11-12T11:45:26.371Z"

			// startTime := str
			// endTime := str

			sTime, err := time.Parse(time.RFC3339, startTime)
			if err != nil {
				return err
			}

			eTime, err := time.Parse(time.RFC3339, endTime)
			if err != nil {
				return err
			}

			if !sTime.After(eTime) {
				return fmt.Errorf("start time shoud be less then end time")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLunchpad(clientCtx.GetFromAddress().String(),
				tokenId, supply, softcap, hardcap, sTime, eTime)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Uint64(FlagTokenID, 0, "launchpad ending time")
	cmd.Flags().Uint64(FlagSupply, 0, "token supply adding for launchpad")
	cmd.Flags().Uint64(FlagSoftCap, 0, "expected softcap value in number")
	cmd.Flags().Uint64(FlagHardCap, 0, "expected hardcap value in number")
	cmd.Flags().String(FlagStartTime, "", "launchpad starting time in nano seconds")
	cmd.Flags().String(FlagEndTime, "", "launchpad ending time in nano seconds")

	return cmd
}

func CmdDepositToLaunchpad() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [id] [amount]",
		Short: "Deposit to launchpad",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			argsId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argsAmount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositToLaunchpad(argsId, clientCtx.GetFromAddress().String(), argsAmount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
