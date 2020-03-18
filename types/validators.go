package types

import (
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
)

type ResultValidatorsOutput struct {
	BlockHeight int64             `json:"block_height"`
	Validators  []ValidatorOutput `json:"validators"`
}

func NewResultValidatorsOutput(rv *ctypes.ResultValidators) (rvo ResultValidatorsOutput, err error) {
	rvo.BlockHeight = rv.BlockHeight
	valNum := len(rv.Validators)
	rvo.Validators = make([]ValidatorOutput, valNum)
	for i := 0; i < valNum; i++ {
		if rvo.Validators[i], err = bech32ValidatorOutput(rv.Validators[i]); err != nil {
			return ResultValidatorsOutput{}, err
		}
	}
	return
}

type ValidatorOutput struct {
	Address          ConsAddress `json:"address"`
	PubKey           string      `json:"pub_key"`
	ProposerPriority int64       `json:"proposer_priority"`
	VotingPower      int64       `json:"voting_power"`
}

type ConsAddress []byte

func bech32ValidatorOutput(validator *types.Validator) (ValidatorOutput, error) {
	bechValPubkey, err := Bech32ifyConsPub(validator.PubKey)
	if err != nil {
		return ValidatorOutput{}, err
	}

	return ValidatorOutput{
		Address:          ConsAddress(validator.Address),
		PubKey:           bechValPubkey,
		ProposerPriority: validator.ProposerPriority,
		VotingPower:      validator.VotingPower,
	}, nil
}

type Description struct {
	Moniker  string `json:"moniker"`  // name
	Identity string `json:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website"`  // optional website link
	Details  string `json:"details"`  // optional details
}

func NewDescription(moniker, identity, website, details string) Description {
	return Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

type CommissionRates struct {
	Rate          Dec `json:"rate"`            // the commission rate charged to delegators, as a fraction
	MaxRate       Dec `json:"max_rate"`        // maximum commission rate which validator can ever charge, as a fraction
	MaxChangeRate Dec `json:"max_change_rate"` // maximum daily increase of the validator commission, as a fraction
}

func NewCommissionRates(rate, maxRate, maxChangeRate Dec) CommissionRates {
	return CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxChangeRate,
	}
}
