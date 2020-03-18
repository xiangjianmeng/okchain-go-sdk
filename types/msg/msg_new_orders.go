package msg

import (
	"strconv"
	"strings"
	"fmt"
	"encoding/json"
	sdk "github.com/okex/okchain-go-sdk/types"
)

const OrderItemLimit = 200

//********************MsgNewOrders*************
type MsgNewOrders struct {
	Sender     sdk.AccAddress `json:"sender"` // order maker address
	OrderItems []OrderItem    `json:"order_items"`
}

type OrderItem struct {
	Product  string  `json:"product"`  // product for trading pair in full name of the tokens
	Side     string  `json:"side"`     // BUY/SELL
	Price    sdk.Dec `json:"price"`    // price of the order
	Quantity sdk.Dec `json:"quantity"` // quantity of the order
}

func NewOrderItem(product string, side string, price string,
	quantity string) OrderItem {
	return OrderItem{
		Product:  product,
		Side:     side,
		Price:    sdk.MustNewDecFromStr(price),
		Quantity: sdk.MustNewDecFromStr(quantity),
	}
}

// NewMsgNewOrders is a constructor function for MsgNewOrder
func NewMsgNewOrders(sender sdk.AccAddress, orderItems []OrderItem) MsgNewOrders {
	return MsgNewOrders{
		Sender:     sender,
		OrderItems: orderItems,
	}
}

// Name Implements Msg.
func (msg MsgNewOrders) Route() string { return "order" }

// Type Implements Msg.
func (msg MsgNewOrders) Type() string { return "MsgNew" }

// ValdateBasic Implements Msg.
func (msg MsgNewOrders) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}
	if msg.OrderItems == nil || len(msg.OrderItems) == 0 {
		return sdk.ErrUnknownRequest("invalid OrderItems")
	}
	if len(msg.OrderItems) > OrderItemLimit {
		return sdk.ErrUnknownRequest("Numbers of NewOrderItem should not be more than " + strconv.Itoa(OrderItemLimit))
	}
	for _, item := range msg.OrderItems {
		if len(item.Product) == 0 {
			return sdk.ErrUnknownRequest("Product cannot be empty")
		}
		symbols := strings.Split(item.Product, "_")
		if len(symbols) != 2 {
			return sdk.ErrUnknownRequest("Product should be in the format of \"base_quote\"")
		}
		if symbols[0] == symbols[1] {
			return sdk.ErrUnknownRequest("invalid product")
		}
		if item.Side != "BUY" && item.Side != "SELL" {
			return sdk.ErrUnknownRequest(
				fmt.Sprintf("Side is expected to be \"BUY\" or \"SELL\", but got \"%s\"", item.Side))
		}
		if !(item.Price.IsPositive() && item.Quantity.IsPositive()) {
			return sdk.ErrUnknownRequest("Price/Quantity must be positive")
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgNewOrders) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgNewOrders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func NewPlaceOrder(sender sdk.AccAddress, product string, side string, price string,
	quantity string) MsgNewOrders {
	orderItem := NewOrderItem(product, side, price, quantity)
	msg := MsgNewOrders{
		Sender:     sender,
		OrderItems: []OrderItem{orderItem},
	}
	return msg
}