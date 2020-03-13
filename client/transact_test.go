package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"

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
	// validator address
	valAddr1 = "okchainvaloper10q0rk5qnyag7wfvvt7rtphlw589m7frs863s3m"
	valAddr2 = "okchainvaloper1g7znsf24w4jc3xfca88pq9kmlyjdare6mph5rx"
	valAddr3 = "okchainvaloper1alq9na49n9yycysh889rl90g9nhe58lcs50wu5"
	valAddr4 = "okchainvaloper1svzxp4ts5le2s4zugx34ajt6shz2hg42a3gl7g"
	// validator mnemonic
	valMnemonic = "race imitate stay curtain puppy suggest spend toe old bridge sunset pride"
	valName     = "validator"
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

	res, err := cli.CancelOrder(fromInfo, passWd, "ID1-0000000244-1", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
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

func TestVote(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	// delegate first
	sequence := accInfo.GetSequence()
	_, err = cli.Delegate(fromInfo, passWd, "1024.2048okt", "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)

	// vote then
	sequence++
	valsToVoted := []string{valAddr1, valAddr2, valAddr3, valAddr4}
	res, err := cli.Vote(fromInfo, passWd, valsToVoted, "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestDestroyValidator(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(valMnemonic, valName, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.DestroyValidator(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestUnjail(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(valMnemonic, valName, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Unjail(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestMultiSend(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	transStr := `okchain1g7c3nvac7mjgn2m9mqllgat8wwd3aptdqket5k 1.024okt
okchain1aac2la53t933t265nhat9pexf9sde8kjnagh9m 2.048okt`
	transfers, err := utils.ParseTransfersStr(transStr)
	assertNotEqual(t, err, nil)

	res, err := cli.MultiSend(fromInfo, passWd, transfers, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}
