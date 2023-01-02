package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strings"
)

// WeiToEthers parses the amounts to ethers.
func WeiToEthers(wei *big.Int) *big.Float {
	return WeiToUnit(wei, big.NewInt(18))
}

// WeiToUnit parses the amounts to some decimals.
func WeiToUnit(wei *big.Int, decimals *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), decimals, nil)))
}

// EthersToWei parses ethers to amounts.
func EthersToWei(ethers *big.Float) *big.Int {
	truncInt, _ := ethers.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", ethers), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

// TokensToWei comment tokens to wei.
func TokensToWei(amounts *big.Float, decimals *big.Int) *big.Int {
	truncInt, _ := amounts.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, new(big.Int).Exp(big.NewInt(10), decimals, nil))
	fracStr := strings.Split(fmt.Sprintf("%.18f", amounts), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}
