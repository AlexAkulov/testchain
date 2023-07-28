package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"testchain/x/cache/types"
)

func (k msgServer) SetState(goCtx context.Context, msg *types.MsgSetState) (*types.MsgSetStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetStateResponse{}, nil
}
