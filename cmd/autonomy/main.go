package main

import (
	"os"

	"github.com/AutonomyNetwork/autonomy-chain/app"
	"github.com/AutonomyNetwork/autonomy-chain/cmd/autonomy/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
