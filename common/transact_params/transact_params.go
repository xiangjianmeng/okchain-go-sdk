package transact_params

import (
	"fmt"
	"github.com/ok-chain/gosdk/common/libs/pkg/errors"
	"github.com/ok-chain/gosdk/crypto/keys"
	"strings"
)

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

func CheckNewOrderParams(fromInfo keys.Info, passWd, product, side string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	if len(product) == 0 {
		return errors.New("no product input")
	}
	if side != "BUY" && side != "SELL" {
		return errors.New(`side can only be "BUY" or "SELL"`)
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
