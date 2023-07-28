package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetState = "set_state"

var _ sdk.Msg = &MsgSetState{}

func NewMsgSetState(creator string, id uint64, value string) *MsgSetState {
	return &MsgSetState{
		Creator: creator,
		Id:      id,
		Value:   value,
	}
}

func (msg *MsgSetState) Route() string {
	return RouterKey
}

func (msg *MsgSetState) Type() string {
	return TypeMsgSetState
}

func (msg *MsgSetState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
