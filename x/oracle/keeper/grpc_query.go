package keeper

import (
	"github.com/spidexapp/spidex/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
