package main

import (
	"context"
	"fmt"
	"github.com/brahma-adshonor/gohook"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	_ "github.com/gosuri/uilive"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/addresses"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/config"
	"github.com/tarik0/DexEqualizer/hub"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/monitor"
	"github.com/tarik0/DexEqualizer/updater"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"github.com/tarik0/DexEqualizer/wallet"
	"golang.org/x/exp/slices"
	"math/big"
	"net/http"
	"os"
	_ "sync"
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

	// Get the hot tokens from the API.
	for tmp := utils.UpdateHotTokens(); variables.IsDev == false && tmp == 0; {
		logger.Log.WithField("tokenCount", tmp).Infoln("Hot tokens API is not ready yet! Waiting 1 min...")
		time.Sleep(1 * time.Minute)
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
		variables.RpcClient,
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
	variables.Hub = hub.NewHub()
	variables.Hub.SetHandler()
	go variables.Hub.Run()
	go variables.Hub.ClearHistory()

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
func onSort(sortBlockNum *big.Int, options []*circle.TradeOption, gasPrices []*big.Int, sortTime time.Duration, u *updater.PairUpdater) {
	// Check balance.
	go checkBalance()

	// broadcast ranks.
	go func() {
		// Print the best 5 options.
		var tradesJson = make([]circle.TradeOptionJSON, 5)
		for i, opt := range options {
			// Get the best gas price for the option.
			tradesJson[i] = opt.GetJSON(gasPrices[i])

			if i == 4 {
				break
			}
		}

		// broadcast ranks
		err := variables.Hub.BroadcastRanks(tradesJson, sortTime.Milliseconds(), sortBlockNum.Uint64())
		if err != nil {
			logger.Log.WithError(err).Fatalln("Unable to marshal trade.")
		}
	}()

	// The pair addresses that we already took an action.
	alreadyUsedPairs := make([]common.Address, 0)

	// Estimate circles.
	// todo it is time to delete that maybe
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
	for i, swapCircle := range options {
		// Check if profitable.
		profit, _ := swapCircle.NormalProfit()
		triggerLim := swapCircle.GetTradeCost(gasPrices[i])
		if profit.Cmp(triggerLim) < 0 {
			// TODO check here before commit
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
		triggerSwap(swapCircle, triggerLim, profit, sortBlockNum, gasPrices[i], u)
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
	transactor.Value = new(big.Int).Set(common.Big0)

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
			Reserves:              tradeOption.Reserves,
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
func triggerSwap(tradeOption *circle.TradeOption, lim *big.Int, profit *big.Int, number *big.Int, gasPrice *big.Int, u *updater.PairUpdater) {
	// broadcast buy.
	go func() {
		err := variables.Hub.BroadcastMsg(fmt.Sprintf(
			"(%s) circle has passed the trigger limit of %.5f WBNB! (%.5f WBNB)",
			tradeOption.Circle.SymbolsStr(),
			utils.WeiToEthers(lim),
			utils.WeiToEthers(profit),
		))
		if err != nil {
			logger.Log.WithError(err).Errorln("Unable to send message to the hub.")
		}
	}()

	// New transactor.
	transactor, err := variables.Wallet.NewTransactor()
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get transactor.")
	}

	// Set transactor values.
	transactor.GasPrice = gasPrice
	transactor.Value = new(big.Int).Set(common.Big0)
	transactor.GasLimit = tradeOption.NormalGasSpent() + tradeOption.NormalGasTokenAmount()*10000
	transactor.NoSend = true

	// The parameter.
	param := abis.SwapParameters{
		Pairs:                 tradeOption.Circle.PairAddresses,
		Reserves:              tradeOption.Reserves,
		Path:                  tradeOption.Circle.Path,
		AmountsOut:            tradeOption.AmountsOut,
		PairTokens:            tradeOption.Circle.PairTokens,
		GasToken:              config.Parsed.Contracts.GasToken,
		GasTokenAmount:        new(big.Int).SetUint64(tradeOption.NormalGasTokenAmount()),
		RevertOnReserveChange: true,
	}

	// Ready transaction.
	tx, err := variables.SwapExec.ExecuteSwap(transactor, param)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"circle": tradeOption.GetJSON(gasPrice),
		}).WithError(err).Errorln("Unable to estimate gas for the transaction.")
		return
	}

	// Send it to the trade channel.
	err = u.DoTrade(updater.TradeAction{
		BlockNumber: number,
		TradeOption: tradeOption,
		Transaction: tx,
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Updater not initialized yet! How tf you did this ?")
	}
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
