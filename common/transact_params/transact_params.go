package transact_params

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/common/libs/pkg/errors"
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
	"strings"
)

func CheckTransferUnitsParams(fromInfo keys.Info, passWd string, transfers []types.TransferUnit) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	transLen := len(transfers)
	if transLen == 0 {
		return errors.New("no receiver input")
	}
	for i := 0; i < 0; i++ {
		if transfers[i].Coins.IsAllPositive() {
			continue
		} else {
			return errors.New("only positive amount of coins is available")
		}
	}

	return nil
}

func CheckVoteParams(fromInfo keys.Info, passWd string, valAddrs []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	if len(valAddrs) == 0 {
		return errors.New("no validator address input")
	}

	return nil
}

func CheckKeyParams(fromInfo keys.Info, passWd string) error {
	if fromInfo == nil {
		return errors.New("input invalid keys info")
	}
	if len(passWd) == 0 {
		return errors.New("no password input")
	}

	return nil
}

func CheckSendParams(fromInfo keys.Info, passWd, toAddr string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	if len(toAddr) != 46 || !strings.HasPrefix(toAddr, "okchain") {
		return errors.New("input invalid receiver address")
	}

	return nil
}

func CheckNewOrderParams(fromInfo keys.Info, passWd string, products, sides, prices, quantities []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	productsLen := len(products)
	if productsLen == 0 {
		return errors.New("no product input")
	}

	if len(sides) != productsLen {
		return errors.New("invalid param side counts")
	}

	if len(prices) != productsLen {
		return errors.New("invalid param price counts")
	}

	if len(quantities) != productsLen {
		return errors.New("invalid param quantity counts")
	}

	for _, side := range sides {
		if side != "BUY" && side != "SELL" {
			return errors.New(`side must only be "BUY" or "SELL"`)
		}
	}

	return nil
}

func CheckCancelOrderParams(fromInfo keys.Info, passWd string, orderIds []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	// check duplicated
	filter := make(map[string]struct{})
	for _, id := range orderIds {
		if _, ok := filter[id]; ok {
			return fmt.Errorf("orderId: %s is duplicated", id)
		}

		filter[id] = struct{}{}
	}

	return nil
}

func checkAccuracyOfStr(num string, accuracy int) bool {
	num = strings.TrimSpace(num)
	strs := strings.Split(num, ".")
	if len(strs) > 2 || len(strs) == 0 {
		return false
	} else if len(strs) == 2 {
		for i, v := range strs[1] {
			if i > accuracy-1 && v != '0' {
				fmt.Printf("the accuracy can't be larger than %d\n", accuracy)
				return false
			}
		}
	}
	return true
}
