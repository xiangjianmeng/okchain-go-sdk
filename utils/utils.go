package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/ok-chain/gosdk/common/libs/pkg/errors"
	"github.com/ok-chain/gosdk/crypto/go-bip39"
	"github.com/ok-chain/gosdk/crypto/keys/hd"
	"github.com/ok-chain/gosdk/types"
)

var AddressStoreKeyPrefix = []byte{0x01}

// parse validator address string to types.ValAddress
func ParseValAddresses(valAddrsStr []string) ([]types.ValAddress, error) {
	valLen := len(valAddrsStr)
	valAddrs := make([]types.ValAddress, valLen)
	var err error
	for i := 0; i < valLen; i++ {
		valAddrs[i], err = types.ValAddressFromBech32(valAddrsStr[i])
		if err != nil {
			return nil, fmt.Errorf("invalid validator address: %s", valAddrsStr[i])
		}
	}
	return valAddrs, nil
}

func AddressStoreKey(addr types.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}

func GeneratePrivateKeyFromMnemo(mnemo string) (string, error) {
	hdPath := hd.NewFundraiserParams(0, 0)
	seed, err := bip39.NewSeedWithErrorChecking(mnemo, "")
	if err != nil {
		return "", err
	}
	masterPrivateKey, ch := hd.ComputeMastersFromSeed(seed)
	derivedPrivateKey, err := hd.DerivePrivateKeyForPath(masterPrivateKey, ch, hdPath.String())
	return hex.EncodeToString(derivedPrivateKey[:]), nil
}

func slice2Array(s []byte) (byteArray [32]byte, err error) {
	if len(s) != 32 {
		return byteArray, errors.New("byte slice's length is not 32")
	}
	for i := 0; i < 32; i++ {
		byteArray[i] = s[i]
	}
	return
}
