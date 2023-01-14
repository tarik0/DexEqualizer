package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/big"
)

// The config.

var Parsed *ArbZeroConfig

// ArbZeroConfig
//	The config template.
type ArbZeroConfig struct {
	Network struct {
		PrivateKey string         `yaml:"Private Key"`
		WETH       common.Address `yaml:"WETH"`
	} `yaml:"Network"`
	Contracts struct {
		Executor    common.Address `yaml:"Executor"`
		GasToken    common.Address `yaml:"Gas Token"`
		Multicaller common.Address `yaml:"Multicaller"`
	} `yaml:"Contracts"`
	ArbitrageOptions struct {
		Limiters struct {
			MaxAmountIn *big.Float `yaml:"Max Amount In"`
			StopBalance *big.Float `yaml:"Stop Balance"`
		} `yaml:"Limiters"`
		GasOptions struct {
			ExtraGasPercent *big.Int `yaml:"Extra Gas Percent"`
		} `yaml:"Gas Options"`
	} `yaml:"Arbitrage Options"`
}

// LoadConfig
// 	Loads config.
func LoadConfig(chainId *big.Int) (*ArbZeroConfig, error) {
	// Read the bytes.
	rawFile, err := ioutil.ReadFile(fmt.Sprintf("./chains/%d/config.yml", chainId))
	if err != nil {
		return nil, err
	}

	// Parse bytes.
	var parsed ArbZeroConfig
	err = yaml.Unmarshal(rawFile, &parsed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
