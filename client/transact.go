package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/transact_params"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	"github.com/okex/okchain-go-sdk/utils"
)

// broadcast mode
const (
	BroadcastBlock = "block"
	BroadcastSync  = "sync"
	BroadcastAsync = "async"
)

func (cli *OKChainClient) Send(fromInfo keys.Info, passWd, toAddr, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckSendParams(fromInfo, passWd, toAddr); err != nil {
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
	if err := transact_params.CheckNewOrderParams(fromInfo, passWd, product, side); err != nil {
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
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
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
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
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

// unbond the delegation on okchain
func (cli *OKChainClient) Unbond(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	coin, err := utils.ParseCoin(coinsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : parse Coins [%s] error: %s", coinsStr, err)
	}

	msg := types.NewMsgUndelegate(fromInfo.GetAddress(), coin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// vote to the validators
func (cli *OKChainClient) Vote(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckVoteParams(fromInfo, passWd, valAddrsStr); err != nil {
		return types.TxResponse{}, err
	}

	valAddrs, err := utils.ParseValAddresses(valAddrsStr)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : validator address parsed error: %s", err.Error())
	}

	msg := types.NewMsgVote(fromInfo.GetAddress(), valAddrs)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// destroy the validator and unbond the min-self-delegation
func (cli *OKChainClient) DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgDestroyValidator(fromInfo.GetAddress())

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// unjail the own validator which was jailed by slashing module
func (cli *OKChainClient) Unjail(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgUnjail(types.ValAddress(fromInfo.GetAddress()))

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// multi-send coins to several receivers
func (cli *OKChainClient) MultiSend(fromInfo keys.Info, passWd string, transfers []types.TransferUnit, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckTransferUnitsParams(fromInfo, passWd, transfers); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgMultiSend(fromInfo.GetAddress(), transfers)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// create a new validator
func (cli *OKChainClient) CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details, minSelfDelegationStr, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	pubkey, err := types.GetConsPubKeyBech32(pubkeyStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	description := types.NewDescription(moniker, identity, website, details)

	minSelfDelegationCoin, err := utils.ParseCoin(minSelfDelegationStr)
	if err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgCreateValidator(types.ValAddress(fromInfo.GetAddress()), pubkey, description, minSelfDelegationCoin)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)
}

// EditValidator edits the description on a validator by the owner
func (cli *OKChainClient) EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	description := types.NewDescription(moniker, identity, website, details)

	msg := types.NewMsgEditValidator(types.ValAddress(fromInfo.GetAddress()), description)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}

// RegisterProxy registers the identity of proxy
func (cli *OKChainClient) RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error) {
	if err := transact_params.CheckKeyParams(fromInfo, passWd); err != nil {
		return types.TxResponse{}, err
	}

	msg := types.NewMsgRegProxy(fromInfo.GetAddress(), true)

	stdBytes, err := tx.BuildAndSignAndEncodeStdTx(fromInfo.GetName(), passWd, memo, []types.Msg{msg}, accNum, seqNum)
	if err != nil {
		return types.TxResponse{}, fmt.Errorf("err : build and sign stdTx error: %s", err.Error())
	}

	return cli.broadcast(stdBytes, BroadcastBlock)

}
