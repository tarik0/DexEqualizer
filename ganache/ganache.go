package ganache

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
	"os/exec"
	"strings"
	"time"
)

// StartGanache starts the Ganache as daemon.
func StartGanache(rpcUrl string, port int) (string, error) {
	// Execute.
	cmd := exec.Command(
		"ganache",
		"--fork.url", rpcUrl,
		"--fork.blockNumber", "latest",
		//"--fork.preLatestConfirmations", "0",
		"--server.ws",
		"--server.port", fmt.Sprintf("%d", port),
		"--detach",
	)

	// Read output.
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return strings.Trim(string(output), "\n"), err
}

// StopGanache stops the Ganache daemon.
func StopGanache(name string) (string, error) {
	// Execute.
	cmd := exec.Command("ganache", "instances", "stop", name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), err
}

// SimulateTransaction simulates transaction in Ganache.
func SimulateTransaction(transaction *types.Transaction) (*types.Receipt, time.Duration, error) {
	// Measure time.
	startTime := time.Now()

	// Get from.
	msg, err := transaction.AsMessage(types.LatestSignerForChainID(variables.ChainId), big.NewInt(1))
	if err != nil {
		return nil, 0, err
	}

	// Add account.
	var result bool
	err = variables.GanacheRpcClient.Call(&result, "evm_addAccount", msg.From().String(), "password")
	if err != nil {
		return nil, 0, err
	}

	// Unlock account.
	err = variables.GanacheRpcClient.Call(&result, "personal_unlockAccount", msg.From().String(), "password", 300)
	if err != nil {
		return nil, 0, err
	}
	if !result {
		return nil, 0, variables.UnableToUnlock
	}

	// Marshall tx.
	simulationTx := make(map[string]string)
	simulationTx["from"] = msg.From().String()
	simulationTx["to"] = transaction.To().String()
	simulationTx["gas"] = fmt.Sprintf("0x%x", transaction.Gas())
	simulationTx["gasPrice"] = fmt.Sprintf("0x%x", transaction.GasPrice())
	simulationTx["value"] = fmt.Sprintf("0x%x", transaction.Value())
	simulationTx["data"] = "0x" + hex.EncodeToString(transaction.Data())
	simulationTx["nonce"] = fmt.Sprintf("0x%x", transaction.Nonce())
	simulationTx["type"] = fmt.Sprintf("0x%x", transaction.Type())

	// Send transaction.
	var resultStr common.Hash
	err = variables.GanacheRpcClient.Call(&resultStr, "eth_sendTransaction", simulationTx)
	if err != nil {
		return nil, 0, err
	}

	// Get receipt.
	receipt, err := variables.GanacheClient.TransactionReceipt(context.Background(), resultStr)
	if err != nil {
		return nil, 0, err
	}

	return receipt, time.Since(startTime), err
}
