package tx

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
)

func buildTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (types.StdTx, error) {
	signMsg := StdSignMsg{
		ChainID:       "okchain",
		AccountNumber: accNumber,
		Sequence:      seqNumber,
		Memo:          memo,
		Msgs:          msgs,
		Fee:           types.NewStdFee(0, nil),
	}

	sig, err := makeSignature(utils.Kb, fromName, passphrase, signMsg)
	if err != nil {
		return types.StdTx{}, err
	}

	return types.NewStdTx(signMsg.Msgs, signMsg.Fee, []types.StdSignature{sig}, signMsg.Memo), nil
}

func BuildAndSignAndEncodeStdTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) ([]byte, error) {
	stdTx, err := buildTx(fromName, passphrase, memo, msgs, accNumber, seqNumber)
	if err != nil {
		return nil, fmt.Errorf("build stdTx error: %s", err)
	}

	// amino encoded
	txBytes, err := types.MsgCdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, fmt.Errorf("amino encoded stdTx error: %s", err)
	}
	return txBytes, nil
}
