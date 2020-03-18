package types

import (
	"github.com/tendermint/tendermint/crypto"
	"time"
)

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

// Commission defines a commission parameters for a given validator.
type Commission struct {
	CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time `json:"update_time"` // the last time the commission rate was changed
}

type Validator struct {
	OperatorAddress         ValAddress    `json:"operator_address"`    // address of the validator's operator; bech encoded in JSON
	ConsPubKey              crypto.PubKey `json:"consensus_pubkey"`    // the consensus public key of the validator; bech encoded in JSON
	Jailed                  bool          `json:"jailed"`              // has the validator been jailed from bonded status?
	Status                  byte          `json:"status"`              // validator status (bonded/unbonding/unbonded)
	Tokens                  Int           `json:"tokens"`              // delegated tokens (incl. self-delegation)
	DelegatorShares         Dec           `json:"delegator_shares"`    // total shares issued to a validator's delegators
	Description             Description   `json:"description"`         // description terms for the validator
	UnbondingHeight         int64         `json:"unbonding_height" `   // if unbonding, height at which this validator has begun unbonding
	UnbondingCompletionTime time.Time     `json:"unbonding_time"`      // if unbonding, min time for the validator to complete unbonding
	Commission              Commission    `json:"commission"`          // commission parameters
	MinSelfDelegation       Int           `json:"min_self_delegation"` // validator's self declared minimum self delegation
}
