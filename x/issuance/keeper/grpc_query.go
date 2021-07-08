package keeper

import (
	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
)

var _ types.QueryServer = Keeper{}
