package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	_ "github.com/gosuri/uilive"
	"github.com/tarik0/DexEqualizer/addresses"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/config"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/monitor"
	"github.com/tarik0/DexEqualizer/updater"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"github.com/tarik0/DexEqualizer/ws"
	"math/big"
	"net/http"
	"os"
	_ "sync"
	"time"
)

// The webserver hub.

var hub *ws.Hub

func main() {
	// The rpc.
	var err error
	RPCUrl := os.Args[1]

	// Connect to the RPC.
	logger.Log.WithField("rpc", RPCUrl).Infoln("Connecting to the RPC...")
	variables.RpcClient, err = rpc.Dial(RPCUrl)
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
	)

	logger.Log.Infoln("")
	logger.Log.Infoln("+ Network Settings")
	logger.Log.Infoln("  Chain ID          :", variables.ChainId.String())
	logger.Log.Infoln("  WETH Addr         :", config.Parsed.Network.WETH.String())
	logger.Log.Infoln("  Multicaller       :", config.Parsed.Contracts.Multicaller)
	logger.Log.Infoln("  Flashloan Executor:", config.Parsed.Contracts.Flashloan)
	logger.Log.Infoln("")

	// Start web server.
	monitor.SetWebHandler()

	// New websocket server.
	hub = ws.NewHub(u)
	hub.SetHandler()
	go hub.Run()

	// Set onSync.
	u.OnSync = onSync

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

// onSync gets triggered on new sync event.
func onSync(updateTime time.Duration, sortTime time.Duration, u *updater.PairUpdater) {
	// Get trade options.
	options := u.GetSortedTrades()
	if options == nil {
		return
	}

	// Skip if no trades.
	if len(options) == 0 {
		return
	}

	// Broadcast ranks.
	go func() {
		// Print the best 10 options.
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
				Circles:    tradesJson,
				SortTime:   sortTime.Milliseconds(),
				UpdateTime: updateTime.Milliseconds(),
			},
		})
		if err != nil {
			logger.Log.WithError(err).Fatalln("Unable to marshal trade.")
		}

		// Broadcast
		hub.Broadcast <- rankBytes
	}()

	// Check if profitable.
	profit, _ := options[0].GetProfit()
	triggerLim := options[0].TriggerLimit()
	if profit.Cmp(triggerLim) < 0 {
		return
	}

	// Broadcast buy.
	go func() {
		// Marshall.
		msg := ws.MessageReq{
			Message: fmt.Sprintf(
				"%s circle has passed the trigger limit of %.5f WBNB! (%.5f WBNB)",
				options[0].Circle.SymbolsStr(),
				utils.WeiToEthers(triggerLim),
				utils.WeiToEthers(profit),
			),
		}
		messageBytes, err := json.Marshal(ws.WebsocketReq{
			Type: "Message",
			Data: msg,
		})
		if err != nil {
			logger.Log.WithError(err).Fatalln("Unable to marshal message trade.")
		}

		// Add to history.
		hub.AddToHistory(msg)

		// Broadcast
		hub.Broadcast <- messageBytes
	}()
}
