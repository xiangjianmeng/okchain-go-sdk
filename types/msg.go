package types

import (
	"encoding/json"
	"github.com/tendermint/tendermint/crypto"
)

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
	Amount      DecCoins   `json:"amount"`
}

// NewMsgTokenSend is a constructor function for MsgSend
func NewMsgTokenSend(from, to AccAddress, coins DecCoins) MsgSend {
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

// nolint
func (MsgSend) Route() string            { return "" }
func (MsgSend) Type() string             { return "" }
func (MsgSend) ValidateBasic() Error     { return nil }
func (MsgSend) GetSigners() []AccAddress { return nil }

type MsgNewOrders struct {
	Sender     AccAddress  `json:"sender"`
	OrderItems []OrderItem `json:"order_items"`
}

// NewMsgNewOrders is a constructor function for MsgNewOrders
func NewMsgNewOrders(sender AccAddress, orderItems []OrderItem) MsgNewOrders {
	return MsgNewOrders{
		Sender:     sender,
		OrderItems: orderItems,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgNewOrders) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

// nolint
func (MsgNewOrders) Route() string            { return "" }
func (MsgNewOrders) Type() string             { return "" }
func (MsgNewOrders) ValidateBasic() Error     { return nil }
func (MsgNewOrders) GetSigners() []AccAddress { return nil }

type MsgCancelOrders struct {
	Sender   AccAddress `json:"sender"`
	OrderIds []string   `json:"order_ids"`
}

// NewMsgCancelOrders is a constructor function for MsgCancelOrders
func NewMsgCancelOrders(sender AccAddress, orderIdItems []string) MsgCancelOrders {
	return MsgCancelOrders{
		Sender:   sender,
		OrderIds: orderIdItems,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelOrders) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return MustSortJSON(b)
}

// nolint
func (MsgCancelOrders) Route() string            { return "" }
func (MsgCancelOrders) Type() string             { return "" }
func (MsgCancelOrders) ValidateBasic() Error     { return nil }
func (MsgCancelOrders) GetSigners() []AccAddress { return nil }

type MsgMultiSend struct {
	From      AccAddress     `json:"from"`
	Transfers []TransferUnit `json:"transfers"`
}

// NewMsgMultiSend is a constructor function for MsgMultiSend
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

// nolint
func (MsgMultiSend) Route() string            { return "" }
func (MsgMultiSend) Type() string             { return "" }
func (MsgMultiSend) ValidateBasic() Error     { return nil }
func (MsgMultiSend) GetSigners() []AccAddress { return nil }

type MsgMint struct {
	Symbol string     `json:"symbol"`
	Amount int64      `json:"amount"`
	Owner  AccAddress `json:"owner"`
}

// NewMsgMint is a constructor function for MsgMint
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

// nolint
func (MsgMint) Route() string            { return "" }
func (MsgMint) Type() string             { return "" }
func (MsgMint) ValidateBasic() Error     { return nil }
func (MsgMint) GetSigners() []AccAddress { return nil }

type MsgDelegate struct {
	DelegatorAddress AccAddress `json:"delegator_address"`
	Amount           DecCoin    `json:"quantity"`
}

// NewMsgDelegate creates a msg of delegating
func NewMsgDelegate(delAddr AccAddress, amount DecCoin) MsgDelegate {
	return MsgDelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDelegate) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDelegate) Route() string            { return "" }
func (MsgDelegate) Type() string             { return "" }
func (MsgDelegate) ValidateBasic() Error     { return nil }
func (MsgDelegate) GetSigners() []AccAddress { return nil }

type MsgUndelegate struct {
	DelegatorAddress AccAddress `json:"delegator_address" `
	Amount           DecCoin    `json:"quantity"`
}

// NewMsgUndelegate creates a msg of undelegating
func NewMsgUndelegate(delAddr AccAddress, amount DecCoin) MsgUndelegate {
	return MsgUndelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUndelegate) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUndelegate) Route() string            { return "" }
func (MsgUndelegate) Type() string             { return "" }
func (MsgUndelegate) ValidateBasic() Error     { return nil }
func (MsgUndelegate) GetSigners() []AccAddress { return nil }

type MsgVote struct {
	DelAddr  AccAddress   `json:"delegator_address"`
	ValAddrs []ValAddress `json:"validator_addresses"`
}

// NewMsgVote creates a msg of multi voting
func NewMsgVote(delAddr AccAddress, valAddrs []ValAddress) MsgVote {
	return MsgVote{
		DelAddr:  delAddr,
		ValAddrs: valAddrs,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgVote) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgVote) Route() string            { return "" }
func (MsgVote) Type() string             { return "" }
func (MsgVote) ValidateBasic() Error     { return nil }
func (MsgVote) GetSigners() []AccAddress { return nil }

type MsgDestroyValidator struct {
	DelAddr AccAddress `json:"delegator_address"`
}

// NewMsgDestroyValidator creates a msg of destroy-validator
func NewMsgDestroyValidator(delAddr AccAddress) MsgDestroyValidator {
	return MsgDestroyValidator{
		DelAddr: delAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDestroyValidator) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDestroyValidator) Route() string            { return "" }
func (MsgDestroyValidator) Type() string             { return "" }
func (MsgDestroyValidator) ValidateBasic() Error     { return nil }
func (MsgDestroyValidator) GetSigners() []AccAddress { return nil }

type MsgUnjail struct {
	ValidatorAddr ValAddress `json:"address"`
}

// NewMsgUnjail creates a msg of unjailing
func NewMsgUnjail(validatorAddr ValAddress) MsgUnjail {
	return MsgUnjail{
		ValidatorAddr: validatorAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnjail) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnjail) Route() string            { return "" }
func (MsgUnjail) Type() string             { return "" }
func (MsgUnjail) ValidateBasic() Error     { return nil }
func (MsgUnjail) GetSigners() []AccAddress { return nil }

type MsgCreateValidator struct {
	Description       Description   `json:"description"`
	MinSelfDelegation DecCoin       `json:"min_self_delegation"`
	DelegatorAddress  AccAddress    `json:"delegator_address"`
	ValidatorAddress  ValAddress    `json:"validator_address"`
	PubKey            crypto.PubKey `json:"pubkey"`
}

type msgCreateValidatorJSON struct {
	Description       Description `json:"description"`
	MinSelfDelegation DecCoin     `json:"min_self_delegation"`
	DelegatorAddress  AccAddress  `json:"delegator_address"`
	ValidatorAddress  ValAddress  `json:"validator_address"`
	PubKey            string      `json:"pubkey"`
}

// NewMsgCreateValidator creates a msg of create-validator
// Delegator address and validator address are the same
func NewMsgCreateValidator(valAddr ValAddress, pubKey crypto.PubKey, description Description, minSelfDelegation DecCoin,
) MsgCreateValidator {
	return MsgCreateValidator{
		Description:       description,
		DelegatorAddress:  AccAddress(valAddr),
		ValidatorAddress:  valAddr,
		PubKey:            pubKey,
		MinSelfDelegation: minSelfDelegation,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateValidator) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// MarshalJSON is useful for the signing of msg MsgCreateValidator
func (msg MsgCreateValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Description:       msg.Description,
		DelegatorAddress:  msg.DelegatorAddress,
		ValidatorAddress:  msg.ValidatorAddress,
		PubKey:            MustBech32ifyConsPub(msg.PubKey),
		MinSelfDelegation: msg.MinSelfDelegation,
	})
}

// nolint
func (MsgCreateValidator) Route() string            { return "" }
func (MsgCreateValidator) Type() string             { return "" }
func (MsgCreateValidator) ValidateBasic() Error     { return nil }
func (MsgCreateValidator) GetSigners() []AccAddress { return nil }

type MsgEditValidator struct {
	Description
	ValidatorAddress ValAddress `json:"address"`
}

// NewMsgEditValidator creates a msg of edit-validator
func NewMsgEditValidator(valAddr ValAddress, description Description) MsgEditValidator {
	return MsgEditValidator{
		Description:      description,
		ValidatorAddress: valAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgEditValidator) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgEditValidator) Route() string            { return "" }
func (MsgEditValidator) Type() string             { return "" }
func (MsgEditValidator) ValidateBasic() Error     { return nil }
func (MsgEditValidator) GetSigners() []AccAddress { return nil }

type MsgRegProxy struct {
	ProxyAddress AccAddress `json:"proxy_address"`
	Reg          bool       `json:"reg"`
}

// NewMsgRegProxy creates a msg of registering or unregistering proxy
func NewMsgRegProxy(proxyAddress AccAddress, reg bool) MsgRegProxy {
	return MsgRegProxy{
		ProxyAddress: proxyAddress,
		Reg:          reg,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgRegProxy) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgRegProxy) Route() string            { return "" }
func (MsgRegProxy) Type() string             { return "" }
func (MsgRegProxy) ValidateBasic() Error     { return nil }
func (MsgRegProxy) GetSigners() []AccAddress { return nil }

type MsgBindProxy struct {
	DelAddr      AccAddress `json:"delegator_address"`
	ProxyAddress AccAddress `json:"proxy_address"`
}

// NewMsgBindProxy creates a msg of binding proxy
func NewMsgBindProxy(delAddr, proxyAddr AccAddress) MsgBindProxy {
	return MsgBindProxy{
		DelAddr:      delAddr,
		ProxyAddress: proxyAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgBindProxy) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgBindProxy) Route() string            { return "" }
func (MsgBindProxy) Type() string             { return "" }
func (MsgBindProxy) ValidateBasic() Error     { return nil }
func (MsgBindProxy) GetSigners() []AccAddress { return nil }

type MsgUnbindProxy struct {
	DelAddr AccAddress `json:"delegator_address"`
}

// NewMsgDestroyValidator creates a msg of unbinding proxy
func NewMsgUnbindProxy(delAddr AccAddress) MsgUnbindProxy {
	return MsgUnbindProxy{
		DelAddr: delAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnbindProxy) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnbindProxy) Route() string            { return "" }
func (MsgUnbindProxy) Type() string             { return "" }
func (MsgUnbindProxy) ValidateBasic() Error     { return nil }
func (MsgUnbindProxy) GetSigners() []AccAddress { return nil }
