package updater

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/slices"
	"math/big"
	"sync"
	"time"
)

// listenActions
//	Listens the channels and prioritizes them.
func (p *PairUpdater) listenActions() {
	// Generate the trades channel.
	p.tradeCh = make(chan TradeAction)

	// Generate the sync channel.
	p.syncCh = make(chan SyncAction)

	// Generate the sort channel.
	p.sortCh = make(chan SortAction)

	// The action priority.
	// 1 - Sync
	// 2 - Sort
	// 3 - Trade

	// Read the channels.
	for {
		// The trade actions.
		select {
		case syncAction := <-p.syncCh:
			// Sync action.
			p.onSyncAction(syncAction)
		default:
			select {
			// The sort actions.
			case sortAction := <-p.sortCh:
				// Sort action.
				p.onSortAction(sortAction)
			default:
				select {
				case tradeAction := <-p.tradeCh:
					// Trade action.
					p.onTradeAction(tradeAction)
				default:
					// Silent.
				}
			}
		}
	}
}

// listenHistory
// 	Listens the transaction history channels.
func (p *PairUpdater) listenHistory() {
	// Generate new channels.
	p.txHistoryReset = make(chan bool)
	p.txHistoryAdd = make(chan HistoryAddAction)
	p.txHistorySearch = make(chan HistorySearchAction)

	// Generate new history.
	p.hashToOptionHistory = make(map[common.Hash]*circle.TradeOption)
	p.hashToTxHistory = make(map[common.Hash]*types.Transaction)
	p.hashToTxBlock = make(map[common.Hash]*big.Int)

	// Pair to minimum gas price to frontrun the last
	// transaction that uses the pair.
	p.pairToMinGasPriceMutex = sync.RWMutex{}
	p.pairToMinGasPrice = make(map[common.Address]*big.Int)

	// Priorities.
	// 1 - History reset.
	// 2 - History add.
	// 3 - History search.

	for {
		select {
		// Reset channel is prioritized.
		case _ = <-p.txHistoryReset:
			// Clear account pending history.
			p.accountToPendingTx.Range(func(key interface{}, value interface{}) bool {
				p.accountToPendingTx.Delete(key)
				return true
			})

			// Clear gas price map.
			p.pairToMinGasPriceMutex.Lock()
			logger.Log.WithField("pairs", len(p.pairToMinGasPrice)).Debugln("Pair gas prices are cleared!")
			p.pairToMinGasPrice = make(map[common.Address]*big.Int)
			p.pairToMinGasPriceMutex.Unlock()

			// Reset the history.
			if len(p.hashToTxHistory) == 0 {
				continue
			}

			// Reset maps.
			p.hashToOptionHistory = make(map[common.Hash]*circle.TradeOption)
			p.hashToTxHistory = make(map[common.Hash]*types.Transaction)
			p.hashToTxBlock = make(map[common.Hash]*big.Int)

			logger.Log.WithField("historyLen", len(p.hashToTxHistory)).Debugln("History cleared!")
		default:
			select {
			case txInfo := <-p.txHistoryAdd:
				// Iterate over pair addresses.
				p.hashToOptionHistory[txInfo.Tx.Hash()] = txInfo.Option
				p.hashToTxHistory[txInfo.Tx.Hash()] = txInfo.Tx
				p.hashToTxBlock[txInfo.Tx.Hash()] = txInfo.BlockNumber
				logger.Log.WithField("historyLen", len(p.hashToTxHistory)).Debugln("Added to the history!")
			case searchInfo := <-p.txHistorySearch:
				// Set gas price.
				p.compareAndSwapGasPrice(searchInfo.TargetPairAddr, searchInfo.TargetTx)

				// Skip if history is empty.
				if len(p.hashToOptionHistory) == 0 {
					continue
				}

				// Search history.
				for prevTxHash, prevOption := range p.hashToOptionHistory {
					// Get previous transaction's block number.
					prevTxBlock := p.hashToTxBlock[prevTxHash]
					latestBlock := new(big.Int).SetUint64(p.highestBlockNum.Load().(uint64))
					if latestBlock.Cmp(prevTxBlock) != 0 {
						break
					}

					// Continue if none of the pairs are used in that transaction.
					if !slices.Contains(prevOption.Circle.PairAddresses, searchInfo.TargetPairAddr) {
						continue
					}

					// Get the previous transaction.
					prevTx := p.hashToTxHistory[prevTxHash]
					tradeOption := p.hashToOptionHistory[prevTxHash]

					// Continue if block has already passed.
					if highestBlock, isPassed := p.isBlockPassed(prevTxBlock); isPassed {
						logger.Log.
							WithField("hash", prevTxHash.String()).
							WithField("highestBlock", highestBlock).
							WithField("txBlock", prevTxBlock).
							Infoln("Block has already passed! Skipping this transaction...")
						continue
					}

					// Continue if gas price is lower.
					if prevTx.GasPrice().Cmp(searchInfo.TargetTx.GasPrice()) > 0 {
						continue
					}

					// Calculate the frontrun gas cost. (%0.1 more gas.)
					frontrunGasPrice := new(big.Int).Mul(searchInfo.TargetTx.GasPrice(), big.NewInt(1100))
					frontrunGasPrice.Div(frontrunGasPrice, big.NewInt(1000))

					// Calculate the profit limit of the option.
					newTradeProfitLimit := prevOption.GetTradeCost(frontrunGasPrice)
					tradeProfit, err := prevOption.NormalProfit()
					if err != nil {
						logger.Log.WithError(err).Errorln("Unable to calculate trade profit.")
						utils.PrintTradeOption(prevOption)
						continue
					}

					// Get from.
					prevAccountMsg, err := searchInfo.TargetTx.AsMessage(types.LatestSignerForChainID(variables.ChainId), big.NewInt(1))
					if err != nil {
						logger.Log.WithError(err).Fatalln("Unable to get transaction as message.")
					}

					// Log fields.
					logFields := logrus.Fields{
						"targetTx":        searchInfo.TargetTx.Hash(),
						"account":         prevAccountMsg.From(),
						"nonce":           prevAccountMsg.Nonce(),
						"ourTx":           prevTx.Hash(),
						"pairAddr":        searchInfo.TargetPairAddr,
						"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToGwei(frontrunGasPrice)),
					}

					// Check if we are still in profit.
					if tradeProfit.Cmp(newTradeProfitLimit) >= 0 {
						logger.Log.WithFields(logFields).Infoln("Updating trade transaction to frontrun competitors...")
						logger.Log.Infoln("")

						// Increase the gas price and resend transaction again.
						replacedTx := p.increaseTxGasPrice(prevTx, tradeOption, prevTxBlock, frontrunGasPrice, searchInfo.TargetTx.Hash())

						// Replace.
						if replacedTx != nil {
							delete(p.hashToOptionHistory, prevTxHash)
							delete(p.hashToTxHistory, prevTxHash)
							delete(p.hashToTxBlock, prevTxHash)
							p.hashToOptionHistory[replacedTx.Hash()] = prevOption
							p.hashToTxHistory[replacedTx.Hash()] = replacedTx
							p.hashToTxBlock[replacedTx.Hash()] = prevTxBlock
						}
					} else {
						// Calculate the cancel gas cost. (%11 more gas.)
						cancelGasPrice := new(big.Int).Mul(prevTx.GasPrice(), big.NewInt(111))
						cancelGasPrice.Div(cancelGasPrice, big.NewInt(100))

						// Enable auto-cancel
						if cancelGasPrice.Cmp(variables.CancelThresholdGasPrice) > 0 {
							logger.Log.WithFields(logFields).Infoln("Trade transaction might not be profitable anymore! Cancelling transaction...")

							// Frontrun your own transaction and replace it with blank tx.
							replacedTx := p.cancelTx(prevTx, tradeOption, prevTxBlock, cancelGasPrice, searchInfo.TargetTx.Hash())

							// Delete from history.
							if replacedTx != nil {
								delete(p.hashToOptionHistory, prevTx.Hash())
								delete(p.hashToTxHistory, prevTxHash)
								delete(p.hashToTxBlock, prevTxHash)
							}
						}
					}
				}
			default:
				// Silent
			}
		}
	}
}

// listenHeads
//	Listens the new mined block heads.
func (p *PairUpdater) listenHeads() {
	var err error
	for {
		select {
		case err = <-p.blocksSub.Err():
			// Disconnected, retry.
			close(p.blocksCh)
			logger.Log.WithError(err).Errorln("Disconnected from the new blocks! Reconnecting...")
			p.subscribeToHeads()
			logger.Log.WithError(err).Errorln("Connected back to the new blocks!")
		case header := <-p.blocksCh:
			// Redirect to the listen method.
			if header != nil {
				// Update block number.
				p.txHistoryReset <- true
				p.highestBlockNum.Store(header.Number.Uint64())

				logger.Log.
					WithField("highestBlock", header.Number).
					Debugln("New heads received.")

				// Sync the block.
				go func(toBlock uint64) {
					// Get the last sync block.
					fromBlock := p.GetSyncBlockNumber()
					fromBlock += 1

					// Get logs.
					logs, err := p.backend.FilterLogs(context.Background(), ethereum.FilterQuery{
						Addresses: p.PairAddresses,
						FromBlock: new(big.Int).SetUint64(fromBlock),
						ToBlock:   new(big.Int).SetUint64(toBlock),
					})
					if err != nil {
						logger.Log.
							WithError(err).
							WithField("fromBlock", fromBlock).
							WithField("toBlock", toBlock).
							Errorln("Unable to get logs from block")
						return
					}

					// Iterate over logs.
					for _, log := range logs {
						p.onNewLog(log)
					}
					p.syncBlockNum.Store(toBlock)

					logger.Log.
						WithField("fromBlock", fromBlock).
						WithField("toBlock", toBlock).
						Debugln("Reserves synced with the block!")

					// Trigger sort action.
					p.sortCh <- SortAction{
						BlockNumber: new(big.Int).SetUint64(toBlock),
					}
				}(header.Number.Uint64())
			}
		default:
			// Silent
		}
	}
}

// listenPending
//	Listens the new pending transaction.
func (p *PairUpdater) listenPending() {
	// Sometimes the pending transactions can have same account and nonce
	// so the most profitable one for the miner will get selected.
	p.accountToPendingTx = sync.Map{}

	// Start listening for new transactions.
	go func() {
		// Skip if node not supported.
		if p.pendingSub == nil || p.pendingCh == nil {
			return
		}

		for {
			select {
			case err := <-p.pendingSub.Err():
				// Disconnected, retry.
				close(p.pendingCh)
				logger.Log.WithError(err).Errorln("Disconnected from the new pending transactions! Reconnecting...")
				p.subscribeToPending()
				logger.Log.WithError(err).Errorln("Connected back to the new pending transactions!")
			case hash := <-p.pendingCh:
				if hash != nil {
					go p.onNewPending(hash)
				}
			default:
				// Silent
			}
		}
	}()
}

// onNewLog
//	Gets triggered when new log is found.
func (p *PairUpdater) onNewLog(log types.Log) {
	// Continue if not in addresses.
	if _, ok := p.AddressToPair[log.Address]; !ok || log.Removed {
		return
	}

	// Iterate over topics.
	for _, topic := range log.Topics {
		// Continue if topic is not sync.
		if !bytes.Equal(syncId.Bytes(), topic.Bytes()) {
			continue
		}

		// Decode event.
		syncDetails, err := pairAbi.Unpack("Sync", log.Data)
		if err != nil {
			logger.Log.WithError(err).Errorln("Unable to decode 'sync' events.")
			break
		}

		// Update reserves.
		resA, resB := syncDetails[0].(*big.Int), syncDetails[1].(*big.Int)
		p.syncCh <- SyncAction{
			BlockNumber: new(big.Int).SetUint64(log.BlockNumber),
			TxIndex:     new(big.Int).SetUint64(uint64(log.TxIndex)),
			LogIndex:    new(big.Int).SetUint64(uint64(log.Index)),
			Address:     log.Address,
			Res0:        resA,
			Res1:        resB,
		}
		break
	}
}

// onNewPending
//	Gets triggered when new pending transaction.
func (p *PairUpdater) onNewPending(hash *common.Hash) {
	// Skip if no hash.
	if hash == nil {
		return
	}

	// Get transaction details.
	transaction, isPending, err := variables.EthClient.TransactionByHash(context.Background(), *hash)
	if err != nil || transaction == nil || !isPending {
		return
	}

	// Get from.
	msg, err := transaction.AsMessage(types.LatestSignerForChainID(variables.ChainId), big.NewInt(1))
	if err != nil {
		return
	}

	// Skip our transactions.
	if bytes.EqualFold(msg.From().Bytes(), variables.Wallet.Address().Bytes()) {
		return
	}

	// Check account nonce.
	val, ok := p.accountToPendingTx.Load(msg.From())

	// Check if there are already a pending transaction for that account.
	if ok {
		// Previous account transaction.
		prevAccountTx := val.(*types.Transaction)

		// Check nonce's.
		if transaction.Nonce() == prevAccountTx.Nonce() {
			// Check gas prices.
			if transaction.GasPrice().Cmp(prevAccountTx.GasPrice()) > 0 {
				// New transaction has same nonce but more gas.
				// Replace the transaction.
				p.accountToPendingTx.Store(msg.From(), transaction)
			} else {
				// New transaction has same nonce but less or same gas.
				// Skip.
				return
			}
		} else {
			// Skip if nonce is less.
			return
		}
	} else {
		// Add it to the transactions.
		p.accountToPendingTx.Store(msg.From(), transaction)
	}

	// Filter out transactions.
	if transaction.To() == nil || transaction.Data() == nil {
		return
	}
	if len(transaction.Data()) < 4 {
		return
	}
	if transaction.To().String() == "0x0000000000000000000000000000000000000000" ||
		transaction.To().String() == "0x000000000000000000000000000000000000dEaD" {
		return
	}
	if transaction.GasPrice().Cmp(variables.GasPrice) < 0 || transaction.Gas() < 70000 {
		return
	}

	// Static search.
	if !p.staticSearch(transaction) {
		// Dynamic search.
		p.dynamicSearch(msg, transaction)
	}
}

// onTradeAction
//	Gets triggered when new trade action is received.
func (p *PairUpdater) onTradeAction(action TradeAction) {
	// Check if block has already passed.
	if highestBlock, isPassed := p.isBlockPassed(action.BlockNumber); isPassed {
		logger.Log.WithFields(logrus.Fields{
			"targetBlock":  action.BlockNumber,
			"highestBlock": highestBlock,
		}).Infoln("Unable to send transaction! Block has already passed.")
		return
	}

	// Send the transaction.
	err := p.backend.SendTransaction(context.Background(), action.Transaction)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"targetBlock": action.BlockNumber,
		}).WithError(err).Errorln("Unable to send transaction!")
	}

	// Skip if it's a blank transaction.
	if action.Transaction.Data() == nil || len(action.Transaction.Data()) == 0 || action.Transaction.Gas() == 21000 {
		return
	}

	// Add to history.
	p.txHistoryAdd <- HistoryAddAction{
		BlockNumber: action.BlockNumber,
		Tx:          action.Transaction,
		Option:      action.TradeOption,
	}

	// Log transaction.
	logger.Log.Infoln("")
	logger.Log.Infoln("Arbitrage Transaction Sent!")
	logger.Log.Infoln("===========================")
	logger.Log.Infoln("Hash     :", action.Transaction.Hash().String())

	// Print replacement.
	if action.ReplacedTransaction != nil {
		logger.Log.Infoln("Old Hash :", action.ReplacedTransaction.Hash())
		logger.Log.Infoln("Old Gas P:", fmt.Sprintf("%.2f Gwei", utils.WeiToGwei(action.ReplacedTransaction.GasPrice())))
	}

	profit, _ := action.TradeOption.NormalProfit()
	logger.Log.Infoln("Block    :", action.BlockNumber)
	logger.Log.Infoln("Path     :", action.TradeOption.Circle.SymbolsStr())
	logger.Log.Infoln("Gas Price:", fmt.Sprintf("%.2f Gwei", utils.WeiToGwei(action.Transaction.GasPrice())))
	logger.Log.Infoln("Profit   :", fmt.Sprintf("%.7f WBNB", utils.WeiToEthers(profit)))
	logger.Log.Infoln("Gas Cost :", fmt.Sprintf("%.8f WBNB", utils.WeiToEthers(action.TradeOption.GetTradeCost(action.Transaction.GasPrice()))))

	// Log amount infos.
	logger.Log.Infoln("Amounts  :")
	for _, amountOut := range action.TradeOption.AmountsOut {
		logger.Log.Infoln("	", amountOut.String())
	}

	// Log pair infos.
	logger.Log.Infoln("Pairs    :")
	for i, reserves := range action.TradeOption.Reserves {
		logger.Log.Infoln("	", action.TradeOption.Circle.PairAddresses[i], "->", reserves[0].String(), reserves[1].String())
	}

	logger.Log.Infoln("")
}

// onSortAction
//	Gets triggered when new sort action is received.
func (p *PairUpdater) onSortAction(action SortAction) {
	// Check if block has already passed.
	if highestBlock, isPassed := p.isBlockPassed(action.BlockNumber); isPassed {
		logger.Log.WithFields(logrus.Fields{
			"targetBlock":  action.BlockNumber,
			"highestBlock": highestBlock,
		}).Infoln("Unable to sort circles! Block has already passed.")
		return
	}

	// Sort circles.
	sortStart := time.Now()
	sortedTradeOptions, optionGasPrices := p.quickSortCircles()
	sortElapsed := time.Since(sortStart)

	// Check if block has already passed.
	if highestBlock, isPassed := p.isBlockPassed(action.BlockNumber); isPassed {
		logger.Log.WithFields(logrus.Fields{
			"targetBlock":  action.BlockNumber,
			"highestBlock": highestBlock,
		}).Infoln("Unable to sort circles! Block has already passed.")
		return
	}

	logger.Log.WithFields(logrus.Fields{
		"sortElapsed": sortElapsed,
		"targetBlock": action.BlockNumber,
		"tradesSize":  len(sortedTradeOptions),
	}).Debugln("Trade options are sorted for the block!")

	// Run onSort hook.
	if len(sortedTradeOptions) != 0 {
		// Limit options to best 5.
		if len(sortedTradeOptions) > 5 {
			sortedTradeOptions = sortedTradeOptions[:5]
		}

		go p.OnSort(action.BlockNumber, sortedTradeOptions, optionGasPrices, sortElapsed, p)
	}
}

// onSyncAction
//	Gets triggered when new sync action is received.
func (p *PairUpdater) onSyncAction(action SyncAction) {
	// Check if pair exists.
	if _pair, ok := p.AddressToPair[action.Address]; ok {
		// Update reserves.
		_pair.SetReserves(action.Res0, action.Res1, action.BlockNumber, action.TxIndex, action.LogIndex)
	}
}

// isBlockPassed
//	Returns if the block is passed or not.
func (p *PairUpdater) isBlockPassed(blockNum *big.Int) (uint64, bool) {
	// Get the highest block number.
	highestBlockNum := p.GetHighestBlockNumber()

	// Check if block has already passed.
	return highestBlockNum, highestBlockNum > blockNum.Uint64()
}

// staticSearch
//	Searches the transaction data and check if any of our pairs are used.
func (p *PairUpdater) staticSearch(transaction *types.Transaction) bool {
	// Limit call data to 100 + 4 bytes
	callData := transaction.Data()
	if len(callData) > 104 {
		callData = callData[4:100]
	}

	// Iterate over pairs.
	isFound := false
	for _, pairAddr := range p.PairAddresses {
		// Trigger search if pair address is in the transaction data.
		if bytes.Contains(callData, pairAddr.Bytes()) {
			isFound = true
			p.txHistorySearch <- struct {
				TargetTx       *types.Transaction
				TargetPairAddr common.Address
			}{
				TargetTx:       transaction,
				TargetPairAddr: pairAddr,
			}
		}
	}

	return isFound
}

// dynamicSearch
// 	Simulates a transaction with `debug_traceCall` and checks if any of our pairs are used.
func (p *PairUpdater) dynamicSearch(msg types.Message, transaction *types.Transaction) {
	// Marshall tx.
	simulationTx := make(map[string]string)
	simulationTx["from"] = msg.From().String()
	simulationTx["to"] = transaction.To().String()
	simulationTx["gas"] = fmt.Sprintf("0x%x", transaction.Gas())
	simulationTx["gasPrice"] = fmt.Sprintf("0x%x", transaction.GasPrice())
	simulationTx["value"] = fmt.Sprintf("0x%x", transaction.Value())
	simulationTx["data"] = "0x" + hex.EncodeToString(transaction.Data())
	simulationTx["nonce"] = fmt.Sprintf("0x%x", transaction.Nonce())

	// Tracer options.
	tracerOptions := make(map[string]interface{})
	tracerOptions["onlyTopCall"] = "false"
	tracerOptions["withLog"] = "true"

	// Options.
	options := make(map[string]interface{})
	options["disableStorage"] = true
	options["disableStack"] = false
	options["enableMemory"] = false
	options["timeout"] = "75ms"
	options["tracer"] = "callTracer"
	options["tracerConfig"] = tracerOptions

	// Trace as call.
	var traceCallRes DebugTraceCall
	err := p.pendingBackend.Call(&traceCallRes, "debug_traceCall", simulationTx, "latest", options)
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("hash", transaction.Hash().String()).
			Debugln("Unable to simulate transaction.")
		return
	}

	// Pair addresses channel.
	wg := new(sync.WaitGroup)

	// Recursive checker function.
	var checkCall func(DebugTraceCall)
	checkCall = func(call DebugTraceCall) {
		defer wg.Done()
		if slices.Contains(p.PairAddresses, call.To) {
			p.txHistorySearch <- struct {
				TargetTx       *types.Transaction
				TargetPairAddr common.Address
			}{
				TargetTx:       transaction,
				TargetPairAddr: call.To,
			}
		}

		// Check sub calls.
		if call.Calls != nil && len(call.Calls) > 0 {
			for _, subCall := range call.Calls {
				wg.Add(1)
				go checkCall(subCall)
			}
		}
	}

	// Wait checker.
	wg.Add(1)
	go checkCall(traceCallRes)
	wg.Wait()
}
