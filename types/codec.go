package types

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

var MsgCdc = amino.NewCodec()

func init() {
	RegisterMsgCdc(MsgCdc)
}

func RegisterMsgCdc(cdc *amino.Codec) {
	//cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)

	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterConcrete(MsgSend{}, "okchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgNewOrder{}, "okchain/order/MsgNew", nil)
	cdc.RegisterConcrete(MsgCancelOrder{}, "okchain/order/MsgCancel", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "okchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgMint{}, "okchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "okchain/staking/MsgDelegate", nil)
	cdc.RegisterConcrete(MsgUndelegate{}, "okchain/staking/MsgUnDelegate", nil)
	//cdc.RegisterConcrete(msg.MsgVote{}, "okchain/staking/MsgVote", nil)
	//cdc.RegisterConcrete(msg.MsgDestroyValidator{}, "okchain/staking/MsgDestroyValidator", nil)
	//cdc.RegisterConcrete(msg.MsgUnjail{}, "cosmos-sdk/MsgUnjail", nil)

	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}
