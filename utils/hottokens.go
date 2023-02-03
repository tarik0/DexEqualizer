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
	Count      int   `json:"count"`
	LastUpdate int64 `json:"lastUpdate"`
	Tokens     []struct {
		Address common.Address `json:"address"`
		Symbol  string         `json:"symbol"`
		Usage   *big.Int       `json:"usage"`
	} `json:"tokens"`
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
func UpdateHotTokens() int {
	// Get hot tokens.
	hotTokens, err := GetHotTokens()
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to get hot tokens from the API. You sure it is running ?")
		return 0
	}

	// Add hot tokens to all tokens.
	updatedCount := 0
	for _, hotToken := range hotTokens.Tokens {
		if !slices.Contains(variables.TargetTokens, hotToken.Address) {
			variables.TargetTokens = append(variables.TargetTokens, hotToken.Address)
			variables.TokenNames[hotToken.Address] = hotToken.Symbol
			variables.TokenFees[hotToken.Address] = new(big.Int).Set(variables.Big10000)
			updatedCount += 1
		}
	}
	return updatedCount
}
