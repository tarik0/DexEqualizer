package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/addresses"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/slices"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	_ "unsafe"
)

// wsMessageSizeLimit is 500 MB
const wsMessageSizeLimit = 500 * 1024 * 1024

//go:linkname newWebsocketCodec github.com/ethereum/go-ethereum/rpc.newWebsocketCodec
func newWebsocketCodec(*websocket.Conn, string, http.Header) rpc.ServerCodec

// newWebsocketCodecHook is a hook for the newWebsocketCodec.
func newWebsocketCodecHook(conn *websocket.Conn, host string, req http.Header) rpc.ServerCodec {
	codec := newWebsocketCodecTramp(conn, host, req)
	conn.SetReadLimit(wsMessageSizeLimit)
	return codec
}

// newWebsocketCodecTramp is a tramp for the newWebsocketCodec.
func newWebsocketCodecTramp(*websocket.Conn, string, http.Header) rpc.ServerCodec {
	for {
		panic("hooking failed")
	}
}

// Max fee amount.

var MaxBuyFee = big.NewInt(11000)   // Max %10
var MaxSellFee = big.NewInt(11000)  // Max %10
var MaxBuyGas = big.NewInt(220000)  // 220,000
var MaxSellGas = big.NewInt(220000) // 220,000
var JunkGas = big.NewInt(40000)
var MaxTokenCount = 120

// The subscriptions and channels

var blocksCh chan *types.Header
var blocksSub ethereum.Subscription

// subscribeToHeads
//	Subscribes to the new blocks.
func subscribeToHeads() {
	// Make new channel.
	blocksCh = make(chan *types.Header)

	// Subscribe to the new blocks.
	var err error
	blocksSub, err = variables.EthClient.SubscribeNewHead(context.Background(), blocksCh)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to subscribe to the new blocks.")
	}
}

func main() {
	// The rpc.
	var err error
	RPCUrl := os.Args[1]

	// Hook the websocket.
	err = gohook.HookByIndirectJmp(newWebsocketCodec, newWebsocketCodecHook, newWebsocketCodecTramp)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to hook newWebsocketCodec.")
	}

	// Is development.
	variables.IsDev = os.Getenv("IS_DEV") == "true"
	if variables.IsDev {
		logger.Log.Level = logrus.DebugLevel
	}

	// Connect to the RPC.
	logger.Log.WithField("rpc", RPCUrl).Infoln("Connecting to the RPC...")
	variables.RpcClient, err = rpc.DialWebsocket(context.Background(), os.Args[1], "")
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to connect to the RPC.")
	}
	variables.EthClient = ethclient.NewClient(variables.RpcClient)
	logger.Log.Infoln("Connected to the RPC!")

	// Get chain id.
	variables.ChainId, err = variables.EthClient.ChainID(context.Background())
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get chain id.")
	}

	//////////////////////////////////////////////

	// Load routers.
	variables.TargetRouters, variables.RouterNames, variables.RouterFees, err = addresses.LoadRouters(variables.ChainId)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to import routers.")
	}

	//////////////////////////////////////////////

	// The sort ticker.
	sortTicker := time.NewTicker(3 * time.Second)

	// Subscribe to new blocks.
	subscribeToHeads()

	// Listen on 8081
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// Lock the mutex.
		rw.RLock()
		rankBytes, err := json.Marshal(TokenInfos{
			Count:  uint64(len(latestSortedTokens)),
			Tokens: latestSortedTokens,
		})
		rw.RUnlock()

		// Check error.
		if err != nil {
			http.Error(writer, "Internal", http.StatusInternalServerError)
			logger.Log.WithError(err).Errorln("Unable to encode ranks.")
			return
		}

		// Write bytes.
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(rankBytes)
	})
	go func() {
		logger.Log.Fatalln(http.ListenAndServe(":8081", nil))
	}()

	// Listen new blocks.
	for {
		select {
		case err = <-blocksSub.Err():
			// Disconnected, retry.
			close(blocksCh)
			logger.Log.WithError(err).Errorln("Disconnected from the new blocks! Reconnecting...")
			subscribeToHeads()
			logger.Log.WithError(err).Errorln("Connected back to the new blocks!")
		default:
			select {
			case <-sortTicker.C:
				onSort()
			default:
				select {
				case header := <-blocksCh:
					// Redirect to the listen method.
					if header != nil {
						// Get block logs.
						logs, err := variables.EthClient.FilterLogs(context.Background(), ethereum.FilterQuery{
							FromBlock: header.Number,
							ToBlock:   header.Number,
						})
						if err != nil {
							logger.Log.
								WithError(err).
								WithField("block", header.Number).
								Errorln("Unable to get logs from block")
							return
						}

						latestBlockNum = header.Number.Uint64()

						// Filter logs.
						for _, log := range logs {
							filterLogs(log)
						}
					}
				default:
					// Silent
				}
			}
		}
	}
}

// pairAbi is the Uniswap pair abi.
var pairAbi, _ = abi.JSON(strings.NewReader(abis.PairMetaData.ABI))

// tokenUsages the token "sync" counter.
var tokenUsages = make(map[common.Address]uint64)

// tokenSymbols is the symbol map
var tokenSymbols = make(map[common.Address]string)

// tokenPairs is the collection of token's pairs.
var tokenPairs = make(map[common.Address][]common.Address)

// latestSortedTokens is the latest sorted tokens.
var rw = &sync.RWMutex{}
var latestSortedTokens = make([]TokenInfo, 0)

// latestBlockNUm is the latest block number.
var latestBlockNum uint64 = 0

// TokenInfos is the token infos JSON struct.
type TokenInfos struct {
	Count  uint64      `json:"Count"`
	Tokens []TokenInfo `json:"Tokens"`
}

// TokenInfo is the token info JSON struct.
type TokenInfo struct {
	Symbol  string           `json:"Symbol"`
	Address common.Address   `json:"Address"`
	BuyFee  *big.Int         `json:"BuyFee" `
	SellFee *big.Int         `json:"SellFee"`
	SwapGas *big.Int         `json:"SwapGas"`
	Usage   uint64           `json:"Usage"`
	Pairs   []common.Address `json:"Pairs"`
}

// onSort gets triggered on new sort event.
func onSort() {
	// Generate token infos.
	var tokenInfos = make([]TokenInfo, 0)
	for key, val := range tokenUsages {
		// Check if pair count is more than one.
		if val, ok := tokenPairs[key]; !ok || len(val) < 2 {
			continue
		}

		// Check token again.
		isTokenOk, tokenBuyFee, tokenSellFee, tokenBuyGas, tokenSellGas, err := filterToken(key)
		if err != nil || !isTokenOk {
			// Remove if previously it was not a honeypot.
			delete(tokenUsages, key)
			delete(tokenPairs, key)
			continue
		}

		// Calculate swap gas.
		swapGas := new(big.Int).Add(tokenBuyGas, tokenSellGas)
		swapGas.Div(swapGas, common.Big2)

		// Append to the list.
		tokenInfos = append(tokenInfos, TokenInfo{
			Address: key,
			Symbol:  tokenSymbols[key],
			BuyFee:  tokenBuyFee,
			SellFee: tokenSellFee,
			SwapGas: swapGas,
			Pairs:   tokenPairs[key],
			Usage:   val,
		})
	}

	// Sort the usages.
	sort.Slice(tokenInfos, func(i, j int) bool {
		return tokenInfos[i].Usage > tokenInfos[j].Usage
	})

	// Copy to the latest.
	rw.Lock()
	latestSortedTokens = tokenInfos
	if len(latestSortedTokens) > MaxTokenCount {
		latestSortedTokens = latestSortedTokens[:MaxTokenCount-1]
	}

	// Return if not ready.
	if len(latestSortedTokens) < 10 {
		rw.Unlock()
		return
	}

	// Print the first 10 tokens.
	logger.Log.Infoln("")
	logger.Log.Infoln("Top 10 Tokens By Trade Amounts - #", latestBlockNum)
	logger.Log.Infoln()
	for i := 0; i < 10 && i < len(latestSortedTokens); i++ {
		logger.Log.Infoln(
			fmt.Sprintf(
				"%d ) %s - %s, | (Buy F/G: %d/%d) (Sell F/G: %d/%d) (Pairs: %d)",
				i+1,
				latestSortedTokens[i].Symbol,
				latestSortedTokens[i].Address,
				latestSortedTokens[i].BuyFee,
				latestSortedTokens[i].SwapGas,
				latestSortedTokens[i].SellFee,
				latestSortedTokens[i].SwapGas,
				len(latestSortedTokens[i].Pairs),
			),
		)
	}
	rw.Unlock()
}

// emptyCallOpts is a blank call options.
var emptyCallOpts = &bind.CallOpts{}

// processToken
//	Process tokens.
func processToken(address common.Address, pair common.Address) bool {
	// Get token contracts.
	tokenContract, err := abis.NewERC20(address, variables.EthClient)
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to cast new token contract for", address.String())
		return false
	}

	// Get symbol if not found.
	if _, ok := tokenSymbols[address]; !ok {
		symbol, err := tokenContract.Symbol(emptyCallOpts)
		if err != nil {
			logger.Log.WithError(err).Errorln("Unable to get symbols for", address.String())
			return true
		}

		tokenSymbols[address] = symbol
	}

	// Check honeypot.
	isTokenOk, _, _, _, _, err := filterToken(address)
	if err != nil || !isTokenOk {
		// Remove if previously it was not a honeypot.
		if _, ok := tokenUsages[address]; ok {
			delete(tokenUsages, address)
			delete(tokenPairs, address)
		}
		return false
	}

	// New array for token pairs.
	if _, ok := tokenPairs[address]; !ok {
		tokenPairs[address] = make([]common.Address, 0)
	}

	// Append to the pairs.
	if !slices.Contains(tokenPairs[address], pair) {
		tokenPairs[address] = append(tokenPairs[address], pair)
	}

	// Increase the counters.
	tokenUsages[address] += 1
	return true
}

// filterLogs
//	Filters the logs.
func filterLogs(log types.Log) {
	// Check if log is a "sync" log.
	isSyncLog := false
	for _, topic := range log.Topics {
		// Continue if topic is not sync.
		if bytes.Equal(pairAbi.Events["Sync"].ID.Bytes(), topic.Bytes()) {
			isSyncLog = true
			break
		}
	}

	// Return if it's not a "Sync" log.
	if !isSyncLog {
		return
	}

	// Unpack the log details.
	_, err := pairAbi.Unpack("Sync", log.Data)
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to decode 'sync' events.")
		return
	}

	// Get pair contract.
	pairContract, err := abis.NewPair(log.Address, variables.EthClient)
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to cast new pair contract for", log.Address.String())
		return
	}

	// Get tokens.
	token0, err := pairContract.Token0(emptyCallOpts)
	token1, err := pairContract.Token1(emptyCallOpts)
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to get tokens for the pair", log.Address)
		return
	}

	// Process tokens.
	processToken(token0, log.Address)
	processToken(token1, log.Address)
}

// filterToken
//	Filters token according to fee info.
func filterToken(address common.Address) (
	isOk bool,
	buyFee *big.Int,
	sellFee *big.Int,
	buyGas *big.Int,
	sellGas *big.Int,
	err error,
) {
	// Check honeypot.
	buyExpectedOut, buyActualOut, sellExpectedOut, sellActualOut, buyGas, sellGas, err := callHoneypotIs(address)
	if err != nil {
		isOk = false
		return
	}

	// Calculate fees.
	buyFee, err = calcFee(buyExpectedOut, buyActualOut)
	sellFee, err = calcFee(sellExpectedOut, sellActualOut)
	if err != nil {
		isOk = false
		return
	}

	// Expect fees to be lower than max.
	if buyFee.Cmp(MaxBuyFee) > 0 || sellFee.Cmp(MaxSellFee) > 0 {
		isOk = false
		return
	}

	// Expect they are the same.
	if buyFee.Cmp(sellFee) != 0 {
		isOk = false
		return
	}

	// Subtract unnecessary gas.
	buyGas.Sub(buyGas, big.NewInt(35000)) // deposit
	buyGas.Sub(buyGas, JunkGas)
	sellGas.Sub(sellGas, JunkGas)

	// Gas should be too much.
	if buyGas.Cmp(MaxBuyGas) > 0 || sellGas.Cmp(MaxSellGas) > 0 {
		isOk = false
		return
	}

	isOk = true
	return
}

// calcFee calculates the token fee.
func calcFee(expected, actual *big.Int) (*big.Int, error) {
	if actual.Cmp(common.Big0) <= 0 {
		return nil, variables.InvalidInput
	}

	num := new(big.Int).Set(actual)
	num.Mul(num, variables.Big10000)
	den := new(big.Int).Set(expected)
	fee := new(big.Int).Div(num, den)

	return fee, nil
}

// callHoneypotIs
//	Calls honeypot.is contract.
func callHoneypotIs(address common.Address) (
	buyExpectedOut *big.Int,
	buyActualOut *big.Int,
	sellExpectedOut *big.Int,
	sellActualOut *big.Int,
	buyGasUsed *big.Int,
	sellGasUsed *big.Int,
	err error,
) {

	// The method.
	methodId, _ := hex.DecodeString("d66383cb")

	// Generate new call data.
	callData := new(bytes.Buffer)
	callData.Write(methodId)
	callData.Write(common.LeftPadBytes(address.Bytes(), 32))

	// Generate new ethereum message.
	toAddr := common.HexToAddress("0x2bf75fd2fab5fc635a4c6073864c708dfc8396fc")
	msg := ethereum.CallMsg{
		To:    &toAddr,
		From:  common.HexToAddress("0x8894e0a0c962cb723c1976a4421c95949be2d4e3"),
		Data:  callData.Bytes(),
		Value: big.NewInt(1e16),
	}

	// Call.
	var res []byte
	res, err = variables.EthClient.PendingCallContract(context.Background(), msg)
	if err != nil {
		return
	}

	// Decode output.
	buyExpectedOut = new(big.Int).SetBytes(res[32*0 : 32*1])
	buyActualOut = new(big.Int).SetBytes(res[32*1 : 32*2])
	sellExpectedOut = new(big.Int).SetBytes(res[32*2 : 32*3])
	sellActualOut = new(big.Int).SetBytes(res[32*3 : 32*4])
	buyGasUsed = new(big.Int).SetBytes(res[32*4 : 32*5])
	sellGasUsed = new(big.Int).SetBytes(res[32*5 : 32*6])
	return
}
