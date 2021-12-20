package keeper

import (
	"github.com/AutonomyNetwork/autonomy-chain/x/launchpad/types"
)

var _ types.QueryServer = Keeper{}
