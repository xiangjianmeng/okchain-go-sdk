package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"

	"testing"
)

const (
	name     = "alice"
	passWd   = "12345678"
	// send's mnemonic
	mnemonic = "sustain hole urban away boy core lazy brick wait drive tiger tell"
	// target address
	addr1    = "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"
)

func TestSend(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.Send(fromInfo, passWd, addr1, "10.24tokt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestNewOrder(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.NewOrder(fromInfo, passWd, "tokt_tusdk", "BUY", "11.2", "1.23", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
	fmt.Println("orderId:", res.Tags[1].Value)
}

func TestCancelOrder(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)
	res, err := cli.CancelOrder(fromInfo, passWd, "ID0000000122-1", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}
