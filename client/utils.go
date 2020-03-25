package client

import (
	"errors"
	"github.com/okex/okchain-go-sdk/common"
	"github.com/okex/okchain-go-sdk/types"
)

const (
	countDefault = 100
)

func convertToDelegatorResp(delegator types.Delegator, stdUndelegation types.StandardizedUndelegation,
) types.DelegatorResp {
	return types.DelegatorResp{
		DelegatorAddress:     delegator.DelegatorAddress,
		ValidatorAddresses:   delegator.ValidatorAddresses,
		Shares:               delegator.Shares,
		Tokens:               delegator.Tokens.StandardizeToDec(),
		UnbondedTokens:       stdUndelegation.Quantity,
		CompletionTime:       stdUndelegation.CompletionTime,
		// TODO: fit the env of staking pressure test, release the code later
		//IsProxy:              delegator.IsProxy,
		//TotalDelegatedTokens: delegator.TotalDelegatedTokens.StandardizeToDec(),
		//ProxyAddress:         delegator.ProxyAddress,
	}
}

func checkParamsGetTickersInfo(count []int) (countRet int, err error) {
	if len(count) > 1 {
		return 0, errors.New("invalid params input for 'GetTickersInfo'")
	}

	if len(count) == 0 {
		countRet = countDefault
	} else {
		if count[0] < 0 {
			return 0, errors.New("'count' cannot be negative")
		}
		countRet = count[0]
	}
	return
}

func checkParamsGetRecentTxRecord(product string, start, end, page, perPage int) (perPageRet int, err error) {
	if product == "" {
		return 0, errors.New("'product' is empty")
	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return
}

func checkParamsGetOpenClosedOrders(addr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	if !common.IsValidAccaddr(addr) {
		return 0, errors.New("invalid address input")
	}

	if product == "" {
		return 0, errors.New("'product' is empty")
	}

	if !common.IsValidSide(side) {
		return 0, errors.New("'side' can only be 'BUY' or 'SELL'")

	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return

}

func checkParamsGetDealsInfo(addr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	return checkParamsGetOpenClosedOrders(addr, product, side, start, end, page, perPage)
}

func checkParamsGetTransactionsInfo(addr string, type_, start, end, page, perPage int) (perPageRet int, err error) {
	if !common.IsValidAccaddr(addr) {
		return 0, errors.New("invalid address input")
	}

	if type_ < 0 {
		return 0, errors.New("'type_' cannot be negative")

	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return
}

func GetOrderIdFromResponse(result *types.TxResponse) string {
	for i := 0; i < len(result.Events); i++ {
		event := result.Events[i]
		for j := 0; j < len(event.Attributes); j++ {
			attribute := event.Attributes[j]
			if attribute.Key == "orderId" {
				return attribute.Value
			}
		}
	}
	return ""
}
