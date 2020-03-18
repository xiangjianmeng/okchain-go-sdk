package msg

import (
	"encoding/json"
	"strconv"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/common"
)


type MsgCancelOrders struct {
	Sender       sdk.AccAddress `json:"sender"` // order maker address
	OrderIds     []string       `json:"order_ids"`
}

// NewMsgCancelOrder is a constructor function for MsgCancelOrder
func NewMsgCancelOrders(sender sdk.AccAddress, orderIdItems []string) MsgCancelOrders {
	msgCancelOrder := MsgCancelOrders{
		Sender:       sender,
		OrderIds: orderIdItems,
	}
	return msgCancelOrder
}

// Name Implements Msg.
func (msg MsgCancelOrders) Route() string { return "order" }

// Type Implements Msg.
func (msg MsgCancelOrders) Type() string { return "MsgCancel" }

// ValdateBasic Implements Msg.
func (msg MsgCancelOrders) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}
	if msg.OrderIds == nil || len(msg.OrderIds) == 0 {
		return sdk.ErrUnknownRequest("invalid OrderIds")
	}
	if len(msg.OrderIds) > common.OrderItemLimit {
		return sdk.ErrUnknownRequest("Numbers of CancelOrderItem should not be more than " + strconv.Itoa(common.OrderItemLimit))
	}
	for _, item := range msg.OrderIds {
		if item == "" {
			return sdk.ErrUnauthorized("orderId cannot be empty")
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelOrders) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgCancelOrders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func NewCancelOrder(sender sdk.AccAddress, orderId string) MsgCancelOrders {
	orderIdList := []string{orderId}
	msg := MsgCancelOrders{
		Sender:       sender,
		OrderIds: orderIdList,
	}
	return msg
}
