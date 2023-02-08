package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/spidexapp/spidex/x/oracle/types"
)

func (k msgServer) UpdateScriptId(goCtx context.Context, msg *types.MsgUpdateScriptId) (*types.MsgUpdateScriptIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if msg.Creator != params.Admin {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotAdmin, "'%s' is not admin", msg.Creator)
	}
	params.ScriptId = msg.NewScriptId
	k.SetParams(ctx, params)
	return &types.MsgUpdateScriptIdResponse{}, nil
}
