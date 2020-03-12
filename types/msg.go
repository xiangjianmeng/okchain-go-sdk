package types

import "encoding/json"

// Transactions messages must fulfill the Msg
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Route() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}

//__________________________________________________________

// Transactions objects must fulfill the Tx
type Tx interface {
	// Gets the all the transaction's messages.
	GetMsgs() []Msg

	// ValidateBasic does a simple and lightweight validation check that doesn't
	// require access to any other information.
	ValidateBasic() Error
}

//__________________________________________________________

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, Error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)

//__________________________________________________________
// msgs of okchain's tx

type MsgSend struct {
	FromAddress AccAddress `json:"from_address"`
	ToAddress   AccAddress `json:"to_address"`
	Amount      Coins      `json:"amount"`
}

func NewMsgTokenSend(from, to AccAddress, coins Coins) MsgSend {
	return MsgSend{
		FromAddress: from,
		ToAddress:   to,
		Amount:      coins,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgSend) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

func (MsgSend) Route() string            { return "" }
func (MsgSend) Type() string             { return "" }
func (MsgSend) ValidateBasic() Error     { return nil }
func (MsgSend) GetSigners() []AccAddress { return nil }

type MsgNewOrder struct {
	Sender   AccAddress `json:"sender"`
	Product  string     `json:"product"`
	Side     string     `json:"side"`
	Price    Dec        `json:"price"`
	Quantity Dec        `json:"quantity"`
}

func NewMsgNewOrder(sender AccAddress, product string, side string, price string, quantity string) MsgNewOrder {
	return MsgNewOrder{
		Sender:   sender,
		Product:  product,
		Side:     side,
		Price:    MustNewDecFromStr(price),
		Quantity: MustNewDecFromStr(quantity),
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgNewOrder) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

func (MsgNewOrder) Route() string            { return "" }
func (MsgNewOrder) Type() string             { return "" }
func (MsgNewOrder) ValidateBasic() Error     { return nil }
func (MsgNewOrder) GetSigners() []AccAddress { return nil }

type MsgCancelOrder struct {
	Sender  AccAddress `json:"sender"`
	OrderId string     `json:"order_id"`
}

func NewMsgCancelOrder(sender AccAddress, orderId string) MsgCancelOrder {
	msgCancelOrder := MsgCancelOrder{
		Sender:  sender,
		OrderId: orderId,
	}
	return msgCancelOrder
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelOrder) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

func (MsgCancelOrder) Route() string            { return "" }
func (MsgCancelOrder) Type() string             { return "" }
func (MsgCancelOrder) ValidateBasic() Error     { return nil }
func (MsgCancelOrder) GetSigners() []AccAddress { return nil }

type MsgMultiSend struct {
	From      AccAddress     `json:"from"`
	Transfers []TransferUnit `json:"transfers"`
}

func NewMsgMultiSend(from AccAddress, transfers []TransferUnit) MsgMultiSend {
	return MsgMultiSend{
		From:      from,
		Transfers: transfers,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgMultiSend) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

func (MsgMultiSend) Route() string            { return "" }
func (MsgMultiSend) Type() string             { return "" }
func (MsgMultiSend) ValidateBasic() Error     { return nil }
func (MsgMultiSend) GetSigners() []AccAddress { return nil }

type MsgMint struct {
	Symbol string     `json:"symbol"`
	Amount int64      `json:"amount"`
	Owner  AccAddress `json:"owner"`
}

func NewMsgMint(symbol string, amount int64, owner AccAddress) MsgMint {
	return MsgMint{
		Symbol: symbol,
		Amount: amount,
		Owner:  owner,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgMint) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

func (MsgMint) Route() string            { return "" }
func (MsgMint) Type() string             { return "" }
func (MsgMint) ValidateBasic() Error     { return nil }
func (MsgMint) GetSigners() []AccAddress { return nil }

type MsgDelegate struct {
	DelegatorAddress AccAddress `json:"delegator_address"`
	Amount           Coin       `json:"quantity"`
}

func NewMsgDelegate(delAddr AccAddress, amount Coin) MsgDelegate {
	return MsgDelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDelegate) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

func (MsgDelegate) Route() string            { return "" }
func (MsgDelegate) Type() string             { return "" }
func (MsgDelegate) ValidateBasic() Error     { return nil }
func (MsgDelegate) GetSigners() []AccAddress { return nil }
