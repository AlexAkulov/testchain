package keeper

import (
	"context"
	"strconv"

	"testchain/x/cache/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetState(goCtx context.Context, msg *types.MsgSetState) (*types.MsgSetStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var state = types.MsgSetState{
		Id:    msg.Id,
		Value: msg.Value,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateKeyPrefix))
	b := k.cdc.MustMarshal(&state)

	store.Set(types.StateKey(
		strconv.FormatUint(state.Id, 10),
	), b)

	return &types.MsgSetStateResponse{}, nil
}
