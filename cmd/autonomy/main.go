package main

import (
	"os"
	
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/server"
	
	
	"github.com/AutonomyNetwork/autonomy-chain/app"
	"github.com/AutonomyNetwork/autonomy-chain/cmd/autonomy/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		
		default:
			os.Exit(1)
		}
	}
}
