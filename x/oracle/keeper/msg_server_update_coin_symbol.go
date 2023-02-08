package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/spidexapp/spidex/x/oracle/types"
)

func (k msgServer) UpdateCoinSymbol(goCtx context.Context, msg *types.MsgUpdateCoinSymbol) (*types.MsgUpdateCoinSymbolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if msg.Creator != params.Admin {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotAdmin, "'%s' is not admin", msg.Creator)
	}

	_, found := k.GetCoinSymbol(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	var coinSymbol = types.CoinSymbol{
		Id:        msg.Id,
		Symbol:    strings.ToUpper(msg.Symbol),
		IsDeleted: msg.IsDelete,
	}

	k.SetCoinSymbol(ctx, coinSymbol)

	return &types.MsgUpdateCoinSymbolResponse{}, nil
}
