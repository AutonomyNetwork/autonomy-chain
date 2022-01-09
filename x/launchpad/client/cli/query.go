package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListLaunchpad())
	cmd.AddCommand(CmdCreatedLaunchpads())
	cmd.AddCommand(CmdActiveLaunchpads())
	cmd.AddCommand(CmdSuccessLaunchpads())
	cmd.AddCommand(CmdFailLaunchpads())
	cmd.AddCommand(CmdShowLaunchpad())

	return cmd
}
