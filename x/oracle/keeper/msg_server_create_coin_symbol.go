package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/spidexapp/spidex/x/oracle/types"
)

func (k msgServer) CreateCoinSymbol(goCtx context.Context, msg *types.MsgCreateCoinSymbol) (*types.MsgCreateCoinSymbolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if msg.Creator != params.Admin {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotAdmin, "'%s' is not admin", msg.Creator)
	}

	var coinSymbol = types.CoinSymbol{
		Symbol:    strings.ToUpper(msg.Symbol),
		IsDeleted: false,
	}

	list := k.GetAllCoinSymbol(ctx)
	for _, v := range list {
		if v.Symbol == coinSymbol.Symbol {
			return nil, sdkerrors.Wrapf(types.ErrCoinSymbolExists, "coin symbol already exists")
		}
	}

	id := k.AppendCoinSymbol(
		ctx,
		coinSymbol,
	)

	return &types.MsgCreateCoinSymbolResponse{
		Id: id,
	}, nil
}
