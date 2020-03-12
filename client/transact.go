package client

import (
	"fmt"
	"github.com/ok-chain/gosdk/common/transactParams"
	"github.com/ok-chain/gosdk/crypto/keys"
	"github.com/ok-chain/gosdk/types"
	"github.com/ok-chain/gosdk/types/tx"
	"github.com/ok-chain/gosdk/utils"
)

// broadcast mode
const (
	BroadcastBlock = "block"
	BroadcastSync  = "sync"
	BroadcastAsync = "async"
)

func (cli *OKChainClient) Send(fromInfo keys.Info, passWd, toAddr, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transactParams.CheckSendParams(fromInfo, passWd, toAddr); err != nil {
		return types.TxResponse{}, err
	}

	to, err := types.AccAddressFromBech32(toAddr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", toAddr, err)
	}

	coins, err := utils.ParseCoins(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgTokenSend(fromInfo.GetAddress(), to, coins)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

func (cli *OKChainClient) NewOrder(fromInfo keys.Info, passWd, product, side, price, quantity, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transactParams.CheckNewOrderParams(fromInfo, passWd, product, side); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgNewOrder(fromInfo.GetAddress(), product, side, price, quantity)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

func (cli *OKChainClient) CancelOrder(fromInfo keys.Info, passWd, orderID, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transactParams.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgCancelOrder(fromInfo.GetAddress(), orderID)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// design for the pressure test of dev

// Delegate okt for voting
func (cli *OKChainClient) Delegate(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transactParams.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	coin, err := utils.ParseCoin(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgDelegate(fromInfo.GetAddress(), coin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}
