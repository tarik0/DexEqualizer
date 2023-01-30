package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	_ "github.com/gosuri/uilive"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/addresses"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/config"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/monitor"
	"github.com/tarik0/DexEqualizer/updater"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"github.com/tarik0/DexEqualizer/wallet"
	"github.com/tarik0/DexEqualizer/ws"
	"golang.org/x/exp/slices"
	"math/big"
	"net/http"
	"os"
	_ "sync"
	"time"
	_ "unsafe"
)

// The webserver hub.

var hub *ws.Hub

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

	// Load tokens.
	variables.TargetTokens, variables.TokenNames, variables.TokenFees, err = addresses.LoadTokens(variables.ChainId)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to import tokens.")
	}

	// Load config.
	config.Parsed, err = config.LoadConfig(variables.ChainId)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to import config.")
	}

	//////////////////////////////////////////////

	// Load executor.
	variables.SwapExec, err = abis.NewSwapExecutorV2(config.Parsed.Contracts.Executor, variables.EthClient)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to load executor.")
	}

	// Check wallet.
	variables.Wallet, err = wallet.InitWallet(config.Parsed.Network.PrivateKey, variables.ChainId)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to import private key as wallet.")
	}

	//////////////////////////////////////////////

	// Connect RPC client agan. (to increase read limit.)
	simulationClient, err := rpc.Dial(os.Args[1])
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to connect to the RPC.")
	}

	// The updater.
	u := updater.NewPairUpdater(
		&updater.PairUpdaterParams{
			Routers: struct {
				Addresses []common.Address
				Names     map[common.Address]string
				Fees      map[common.Address]*big.Int
			}{
				Addresses: variables.TargetRouters,
				Names:     variables.RouterNames,
				Fees:      variables.RouterFees,
			},
			Tokens: struct {
				MainAddress common.Address
				Addresses   []common.Address
				Symbols     map[common.Address]string
				Fees        map[common.Address]*big.Int
			}{
				MainAddress: config.Parsed.Network.WETH,
				Addresses:   variables.TargetTokens,
				Symbols:     variables.TokenNames,
				Fees:        variables.TokenFees,
			},
			Multicaller: struct {
				Address common.Address
			}{
				Address: config.Parsed.Contracts.Multicaller,
			},
		},
		variables.EthClient,
		simulationClient,
	)

	logger.Log.Infoln("")
	logger.Log.Infoln("+ Network Settings")
	logger.Log.Infoln("  Chain ID          :", variables.ChainId.String())
	logger.Log.Infoln("  WETH Addr         :", config.Parsed.Network.WETH.String())
	logger.Log.Infoln("  Multicaller       :", config.Parsed.Contracts.Multicaller)
	logger.Log.Infoln("  Executor          :", config.Parsed.Contracts.Executor)
	logger.Log.Infoln("")

	if variables.IsDev {
		logger.Log.Debugln("Development mode activated!")
		logger.Log.Debugln("")
	}

	// Start web server.
	monitor.SetWebHandler()

	// New websocket server.
	hub = ws.NewHub(u)
	hub.SetHandler()
	go hub.Run()
	go hub.ClearHistory()

	// Set onSort.
	u.OnSort = onSort

	logger.Log.WithField("tokenCount", len(variables.TokenNames)).Infoln("Loading pair information...")

	// Start listening.
	err = u.Start()
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to generate new updater.")
	}

	logger.Log.Infoln("")
	logger.Log.Infoln("Monitor server started at \"http://0.0.0.0:8080\"!")

	// Start web server.
	logger.Log.Fatalln(http.ListenAndServe(":8080", nil))
}

// onSort gets triggered on new sort event.
func onSort(header *types.Header, options []*circle.TradeOption, updateTime time.Duration, u *updater.PairUpdater) {
	// Check balance.
	go checkBalance()

	// Skip if no trades.
	if len(options) == 0 {
		return
	}
	options = options[:5]

	// Broadcast ranks.
	go func() {
		// Print the best 5 options.
		var tradesJson = make([]circle.TradeOptionJSON, 5)
		for i, opt := range options {
			tradesJson[i] = opt.GetJSON()

			if i == 4 {
				break
			}
		}

		// Marshall.
		rankBytes, err := json.Marshal(ws.WebsocketReq{
			Type: "Rank",
			Data: ws.RankReq{
				Circles:     tradesJson,
				SortTime:    updateTime.Milliseconds(),
				BlockNumber: header.Number.Uint64(),
			},
		})
		if err != nil {
			logger.Log.WithError(err).Fatalln("Unable to marshal trade.")
		}

		// Broadcast
		hub.Broadcast <- rankBytes
	}()

	// The pair addresses that we already took an action.
	alreadyUsedPairs := make([]common.Address, 0)

	// Estimate circles.
	if variables.IsDev {
		go func() {
			circleGases, _, errs := estimateCircles(options)
			for i, _ := range circleGases {
				if errs[i] != nil &&
					fmt.Sprint(errs[i]) != "execution reverted: SE2" &&
					fmt.Sprint(errs[i]) != "execution reverted: FE2" &&
					fmt.Sprint(errs[i]) != "execution reverted: FE6" {
					utils.PrintTradeOption(options[i])
					logger.Log.WithError(errs[i]).Fatalln("Circle failed the simulation.")
				}
			}
		}()
	}

	// Check circles.
	for _, swapCircle := range options {
		// Check if profitable.
		profit, _ := swapCircle.NormalProfit()
		triggerLim := swapCircle.NormalTriggerProfit(variables.GasPrice)
		if profit.Cmp(triggerLim) < 0 {
			return
		}

		// Check if we took action for one of the pairs.
		alreadyUsed := false
		for _, pairAddr := range swapCircle.Circle.PairAddresses {
			if slices.Contains(alreadyUsedPairs, pairAddr) {
				alreadyUsed = true
				break
			}

			// Add pair to the used pairs.
			alreadyUsedPairs = append(alreadyUsedPairs, pairAddr)
		}
		if alreadyUsed {
			continue
		}

		// Trigger the best swap.
		triggerSwap(swapCircle, triggerLim, profit, header.Number, u)
		break
	}
}

// estimateCircles estimates the gas limit for circles.
func estimateCircles(swapCircles []*circle.TradeOption) ([]uint64, uint64, []error) {
	// Arguments.
	transactor, err := variables.Wallet.NewTransactor()
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get transactor.")
	}
	transactor.NoSend = true
	transactor.GasPrice = variables.GasPrice
	transactor.Value = common.Big0

	// Wait group and channel.
	ch := make(chan struct {
		Id       int
		GasLimit uint64
		Err      error
	})

	// Iterate over circles.
	for i, tradeOption := range swapCircles {
		// The parameters.
		param := abis.SwapParameters{
			Pairs:                 tradeOption.Circle.PairAddresses,
			Reserves:              tradeOption.Circle.PairReserves,
			Path:                  tradeOption.Circle.Path,
			AmountsOut:            tradeOption.AmountsOut,
			PairTokens:            tradeOption.Circle.PairTokens,
			GasToken:              config.Parsed.Contracts.GasToken,
			GasTokenAmount:        new(big.Int).SetUint64(tradeOption.NormalGasTokenAmount()),
			RevertOnReserveChange: true,
		}

		// Estimate gas.
		i := i
		go func() {
			// Execute.
			tx, err := variables.SwapExec.ExecuteSwap(transactor, param)
			if err != nil {
				ch <- struct {
					Id       int
					GasLimit uint64
					Err      error
				}{GasLimit: 0, Err: err, Id: i}
				return
			}

			ch <- struct {
				Id       int
				GasLimit uint64
				Err      error
			}{GasLimit: tx.Gas(), Err: nil, Id: i}
		}()
	}

	// Arrays.
	gasLimits := make([]uint64, len(swapCircles))
	estimateErrors := make([]error, len(swapCircles))

	// Take average hop gas.
	var allGasTotal uint64 = 0
	var txCount uint64 = 0

	// Iterate over channels.
	for i := 0; i < len(swapCircles); i++ {
		select {
		case estimated := <-ch:
			gasLimits[estimated.Id] = estimated.GasLimit
			estimateErrors[estimated.Id] = estimated.Err

			if estimated.Err == nil {
				allGasTotal += estimated.GasLimit / uint64(len(swapCircles[estimated.Id].Circle.Pairs))
				txCount += 1
			}
		}
	}

	if txCount == 0 {
		txCount += 1
	}

	return gasLimits, allGasTotal / txCount, estimateErrors
}

// triggerSwap triggers a new swap with circle.
func triggerSwap(tradeOption *circle.TradeOption, lim *big.Int, profit *big.Int, number *big.Int, u *updater.PairUpdater) {
	// Broadcast buy.
	go func() {
		// Encoder.
		var buff = new(bytes.Buffer)
		e := json.NewEncoder(buff)
		e.SetEscapeHTML(true)

		// Marshall.
		msg := ws.MessageReq{
			Timestamp: time.Now().UnixMilli(),
			Message: fmt.Sprintf(
				"%s circle has passed the trigger limit of %.5f WBNB! (%.5f WBNB)",
				tradeOption.Circle.SymbolsStr(),
				utils.WeiToEthers(lim),
				utils.WeiToEthers(profit),
			),
		}

		// Encode.
		err := e.Encode(ws.WebsocketReq{
			Type: "Message",
			Data: msg,
		})
		if err != nil {
			logger.Log.WithError(err).Fatalln("Unable to marshal message.")
		}

		// Add to history.
		hub.AddToHistory(msg)

		// Broadcast
		hub.Broadcast <- buff.Bytes()
	}()

	// New transactor.
	transactor, err := variables.Wallet.NewTransactor()
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get transactor.")
	}

	// Set transactor values.
	transactor.GasPrice = variables.GasPrice
	transactor.Value = common.Big0
	transactor.GasLimit = tradeOption.NormalGasSpent() + tradeOption.NormalGasTokenAmount()*10000

	// The parameter.
	param := abis.SwapParameters{
		Pairs:                 tradeOption.Circle.PairAddresses,
		Reserves:              tradeOption.Circle.PairReserves,
		Path:                  tradeOption.Circle.Path,
		AmountsOut:            tradeOption.AmountsOut,
		PairTokens:            tradeOption.Circle.PairTokens,
		GasToken:              config.Parsed.Contracts.GasToken,
		GasTokenAmount:        new(big.Int).SetUint64(tradeOption.NormalGasTokenAmount()),
		RevertOnReserveChange: true,
	}

	// Send transaction.
	tx, err := variables.SwapExec.ExecuteSwap(transactor, param)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"circle": tradeOption.GetJSON(),
		}).WithError(err).Errorln("Unable to estimate gas for the transaction.")
		return
	}

	// Add to the history.
	u.TxHistoryAdd <- struct {
		Tx          *types.Transaction
		Option      *circle.TradeOption
		BlockNumber *big.Int
	}{Tx: tx, Option: tradeOption, BlockNumber: number}

	// Log transaction.
	logger.Log.Infoln("")
	logger.Log.WithFields(logrus.Fields{
		"hash":            tx.Hash().String(),
		"circle":          tradeOption.GetJSON(),
		"profitLimit":     fmt.Sprintf("%.18f BNB", utils.WeiToEthers(tradeOption.NormalTriggerProfit(transactor.GasPrice))),
		"gasSpent":        tradeOption.NormalGasSpent(),
		"chiAmount":       tradeOption.NormalGasTokenAmount(),
		"gasSpentWithChi": (tradeOption.NormalGasSpent() + tradeOption.NormalGasTokenAmount()*10000) - tradeOption.NormalChiRefund(),
	}).Infoln("Arbitrage transaction sent!")
	logger.Log.Infoln("")
}

// checkBalance checks the wallet balance and stops when too low.
func checkBalance() {
	// Return if dev.
	if variables.IsDev {
		return
	}

	// Get WETH balance.
	balance, err := variables.EthClient.PendingBalanceAt(context.Background(), variables.Wallet.Address())
	if balance.Cmp(utils.EthersToWei(config.Parsed.ArbitrageOptions.Limiters.StopBalance)) <= 0 {
		logger.Log.WithField("balance", utils.WeiToEthers(balance)).Fatalln("Low balance!")
	}
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get balance!")
	}
}
