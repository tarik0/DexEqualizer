package utils

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/slices"
	"math/big"
	"net/http"
	"time"
)

// HotTokensResponse is the JSON struct from hot tokens response.
type HotTokensResponse struct {
	Count  int `json:"Count"`
	Tokens []struct {
		Symbol  string         `json:"Symbol"`
		Address common.Address `json:"Address"`
		BuyFee  *big.Int       `json:"BuyFee"`
		SellFee *big.Int       `json:"SellFee"`
		SwapGas *big.Int       `json:"SwapGas"`
		Usage   *big.Int       `json:"Usage"`
		Pairs   []string       `json:"Pairs"`
	} `json:"Tokens"`
}

// hotTokensURL is the endpoint for hot tokens.
const hotTokensURL = "http://127.0.0.1:8081/"

// httpClient is a basic HTTP client with 10 sec timeout.
var httpClient = &http.Client{Timeout: 10 * time.Second}

// GetHotTokens returns the hot tokens from the API.
func GetHotTokens() (*HotTokensResponse, error) {
	// Send GET request.
	r, err := httpClient.Get(hotTokensURL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// Unmarshall json.
	var response HotTokensResponse
	err = json.NewDecoder(r.Body).Decode(&response)
	return &response, err
}

// UpdateHotTokens updates the tokens with API and returns how many tokens updated.
func UpdateHotTokens() (int, error) {
	// Get hot tokens.
	hotTokens, err := GetHotTokens()
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to get hot tokens from the API. You sure it is running ?")
		return 0, err
	}

	// Check response.
	if hotTokens.Count < 15 {
		return 0, variables.APINotReady
	}

	// Check duplicates.
	nonDupCount := 0
	for _, hotToken := range hotTokens.Tokens {
		if !slices.Contains(variables.TargetTokenAddresses, hotToken.Address) {
			nonDupCount += 1
		}
	}
	if nonDupCount < 15 {
		return 0, variables.APINotReady
	}

	// Add hot tokens to all tokens.
	for _, hotToken := range hotTokens.Tokens {
		if !slices.Contains(variables.TargetTokenAddresses, hotToken.Address) {
			variables.TargetTokenAddresses = append(variables.TargetTokenAddresses, hotToken.Address)
		}
		variables.TargetTokens[hotToken.Address] = &variables.Token{
			Address: hotToken.Address,
			Symbol:  hotToken.Symbol,
			BuyFee:  hotToken.BuyFee,
			SellFee: hotToken.SellFee,
			SwapGas: hotToken.SwapGas,
		}
	}
	return nonDupCount, nil
}
