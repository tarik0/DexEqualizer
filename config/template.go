package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/big"
)

// The config.

var Parsed *DexEqConfig

// DexEqConfig
//	The config template.
type DexEqConfig struct {
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
			MaxHops     int        `yaml:"Max Hops"`
			MinHops     int        `yaml:"Min Hops"`
			MaxCircles  uint64     `yaml:"Max Circles"`
		} `yaml:"Limiters"`
		GasOptions struct {
			ExtraGasPercent *big.Int `yaml:"Extra Gas Percent"`
		} `yaml:"Gas Options"`
	} `yaml:"Arbitrage Options"`
}

// LoadConfig
// 	Loads config.
func LoadConfig(chainId *big.Int) (*DexEqConfig, error) {
	// Read the bytes.
	rawFile, err := ioutil.ReadFile(fmt.Sprintf("./chains/%d/config.yml", chainId))
	if err != nil {
		return nil, err
	}

	// Parse bytes.
	var parsed DexEqConfig
	err = yaml.Unmarshal(rawFile, &parsed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
