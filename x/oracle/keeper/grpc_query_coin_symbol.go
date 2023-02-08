package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/spidexapp/spidex/x/oracle/types"
)

func (k Keeper) CoinSymbolAll(c context.Context, req *types.QueryAllCoinSymbolRequest) (*types.QueryAllCoinSymbolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coinSymbols []types.CoinSymbol
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coinSymbolStore := prefix.NewStore(store, types.KeyPrefix(types.CoinSymbolKey))

	pageRes, err := query.Paginate(coinSymbolStore, req.Pagination, func(key []byte, value []byte) error {
		var coinSymbol types.CoinSymbol
		if err := k.cdc.Unmarshal(value, &coinSymbol); err != nil {
			return err
		}

		coinSymbols = append(coinSymbols, coinSymbol)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinSymbolResponse{CoinSymbol: coinSymbols, Pagination: pageRes}, nil
}

func (k Keeper) CoinSymbol(c context.Context, req *types.QueryGetCoinSymbolRequest) (*types.QueryGetCoinSymbolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	coinSymbol, found := k.GetCoinSymbol(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCoinSymbolResponse{CoinSymbol: coinSymbol}, nil
}
