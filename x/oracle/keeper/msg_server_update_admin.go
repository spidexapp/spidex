package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/spidexapp/spidex/x/oracle/types"
)

func (k msgServer) UpdateAdmin(goCtx context.Context, msg *types.MsgUpdateAdmin) (*types.MsgUpdateAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if msg.Creator != params.Admin {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotAdmin, "'%s' is not admin", msg.Creator)
	}
	_, err := sdk.AccAddressFromBech32(msg.NewAdmin)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrNotTypeOfAddress, "'%s' to type of address", msg.NewAdmin)
	}
	params.Admin = msg.NewAdmin
	k.SetParams(ctx, params)
	return &types.MsgUpdateAdminResponse{}, nil
}
