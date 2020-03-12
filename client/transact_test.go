package client

import (
	"fmt"
	"github.com/ok-chain/gosdk/utils"

	"testing"
)

const (
	name   = "alice"
	passWd = "12345678"
	// send's mnemonic
	mnemonic = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	addr     = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	// target address
	addr1 = "okchain1g7c3nvac7mjgn2m9mqllgat8wwd3aptdqket5k"
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

func TestNewOrder(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.NewOrder(fromInfo, passWd, "xxb_okt", "BUY", "11.2", "1.23", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
	fmt.Println("orderId:", GetOrderIdFromResponse(&res))
}

func TestCancelOrder(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.CancelOrder(fromInfo, passWd, "ID0000004307-1", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestDelegate(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Delegate(fromInfo, passWd, "1024.2048okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestUnbond(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Unbond(fromInfo, passWd, "10.24okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}
