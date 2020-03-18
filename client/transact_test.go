package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"

	"testing"
	"github.com/okex/okchain-go-sdk/types/msg"
)

const (
	name   = "alice"
	passWd = "12345678"
	// send's mnemonic
	mnemonic = "total lottery arena when pudding best candy until army spoil drill pool"
	// target address
	addr1 = "okchain1g7c3nvac7mjgn2m9mqllgat8wwd3aptdqket5k"

	testCoin = "xxb-127"
	baseCoin = "okt"
	testProduct = testCoin + "_" + baseCoin
)

func TestSend(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.Send(fromInfo, passWd, addr1, "10.24okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestNewOrderAndCancelOrder(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.NewOrder(fromInfo, passWd, testProduct, "BUY", "11.2", "1.23", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)

	orderIdList := GetOrderIdListFromResponse(&res)

	assertNotEqual(t, len(orderIdList), 1)
	fmt.Println("orderId:", orderIdList)

	res, err = cli.CancelOrder(fromInfo, passWd, orderIdList[0], "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence() + 1)
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestNewOrdersAndCancelOrders(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	items := []msg.OrderItem{
		msg.NewOrderItem(testProduct, "BUY", "11.2", "1.23"),
		msg.NewOrderItem(testProduct, "BUY", "11.2", "1.23"),
	}

	res, err := cli.NewOrders(fromInfo, items, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
	orderIdList := GetOrderIdListFromResponse(&res)
	fmt.Println("orderId:", orderIdList)

	res, err = cli.CancelOrders(fromInfo, passWd, "my memo", orderIdList, accInfo.GetAccountNumber(), accInfo.GetSequence() + 1)
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}
