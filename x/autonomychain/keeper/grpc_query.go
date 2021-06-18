package keeper

import (
	"github.com/AutonomyNetwork/autonomy-chain/x/autonomychain/types"
)

var _ types.QueryServer = Keeper{}
