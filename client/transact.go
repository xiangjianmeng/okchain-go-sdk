package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	"github.com/okex/okchain-go-sdk/common/transactParams"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/msg"
	"github.com/okex/okchain-go-sdk/types/tx"
	"github.com/okex/okchain-go-sdk/utils"
)

// broadcast mode
const (
	BroadcastBlock = "block"
	BroadcastSync  = "sync"
	BroadcastAsync = "async"
)

func (cli *OKChainClient) Send(fromInfo keys.Info, passWd, toAddr, coinsStr, memo string, accNum, seqNum uint64) (resp types.TxResponse, err error) {
	if !transactParams.IsValidSendParams(fromInfo, passWd, toAddr) {
		return types.TxResponse{}, errors.New("err : params input to send are invalid")
	}

	to, err := types.AccAddressFromBech32(toAddr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Address [%s] error: %s", toAddr, err)
	}

	coins, err := utils.ParseCoins(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := msg.NewMsgTokenSend(fromInfo.GetAddress(), to, coins)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

func (cli *OKChainClient) NewOrder(fromInfo keys.Info, passWd, product, side, price, quantity, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidNewOrderParams(fromInfo, passWd, product, side) {
		return types.TxResponse{}, errors.New("err : params input to pend a order are invalid")
	}
	msg := msg.NewPlaceOrder(fromInfo.GetAddress(), product, side, price, quantity)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

func (cli *OKChainClient) CancelOrder(fromInfo keys.Info, passWd, orderID, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidCancelOrderParams(fromInfo, passWd) {
		return types.TxResponse{}, errors.New("err : params input to cancel a order are invalid")
	}

	msg := msg.NewCancelOrder(fromInfo.GetAddress(), orderID)
	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

func (cli *OKChainClient) NewOrders(fromInfo keys.Info, orderItems []msg.OrderItem, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	for _, item := range orderItems {
		if !transactParams.IsValidNewOrderParams(fromInfo, passWd, item.Product, item.Side) {
			return types.TxResponse{}, errors.New("err : params input to pend a order are invalid")
		}
	}

	msg := msg.NewMsgNewOrders(fromInfo.GetAddress(), orderItems)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

func (cli *OKChainClient) CancelOrders(fromInfo keys.Info, passWd, memo string, orderIdList []string, accNum, seqNum uint64) (types.TxResponse, error) {
	if !transactParams.IsValidCancelOrderParams(fromInfo, passWd) {
		return types.TxResponse{}, errors.New("err : params input to cancel a order are invalid")
	}

	msg := msg.NewMsgCancelOrders(fromInfo.GetAddress(), orderIdList)
	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}
