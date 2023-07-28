package keeper

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testchain/x/cache/types"
)

func (k Keeper) GetState(goCtx context.Context, req *types.QueryGetStateRequest) (*types.QueryGetStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateKeyPrefix))

	b := store.Get(types.StateKey(
		strconv.FormatUint(req.Id, 10),
	))
	if b == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}
	val := types.MsgSetState{}
	k.cdc.MustUnmarshal(b, &val)

	return &types.QueryGetStateResponse{Value: val.Value}, nil
}
