package addresses

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/slices"
	"math/big"
	"os"
	"strings"
)

// LoadTokens
//	Loads tokens from the tokens.txt file.
func LoadTokens(chainId *big.Int) (
	tokenAddresses []common.Address,
	tokens map[common.Address]*variables.Token,
	err error,
) {
	// Open handle.
	readFile, err := os.Open(fmt.Sprintf("./chains/%d/tokens.txt", chainId))
	if err != nil {
		return nil, nil, err
	}

	// New scanner.
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Read lines.
	tokenAddresses = make([]common.Address, 0)
	tokens = make(map[common.Address]*variables.Token)
	for fileScanner.Scan() {
		// Get the line.
		line := fileScanner.Text()

		// Skip comment lines.
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line.
		lineChunks := strings.Split(line, " ")

		// Check length.
		if len(lineChunks) < 5 {
			logger.Log.Infoln("Invalid token line:", line)
			continue
		}

		// Check values.
		isAddrOk := common.IsHexAddress(lineChunks[0])
		buyFee, ok0 := new(big.Int).SetString(lineChunks[1], 10)
		sellFee, ok1 := new(big.Int).SetString(lineChunks[2], 10)
		swapGas, ok2 := new(big.Int).SetString(lineChunks[3], 10)
		if !isAddrOk || !ok0 || !ok1 || !ok2 {
			logger.Log.Infoln("Invalid token line:", line)
			continue
		}

		// Add to the map.
		addr := common.HexToAddress(lineChunks[0])
		tokenAddresses = append(tokenAddresses, addr)
		tokens[addr] = &variables.Token{
			Address: addr,
			Symbol:  lineChunks[4],
			BuyFee:  buyFee,
			SellFee: sellFee,
			SwapGas: swapGas,
		}
	}

	return tokenAddresses, tokens, nil
}

// LoadRouters
//	Loads routers from the routers.txt file.
func LoadRouters(chainId *big.Int) ([]common.Address, map[common.Address]string, map[common.Address]*big.Int, error) {
	return loadRouterAddresses(fmt.Sprintf("./chains/%d/routers.txt", chainId), true)
}

// loadRouterAddresses
//	Loads addresses from path.
func loadRouterAddresses(path string, isRouter bool) (
	[]common.Address,
	map[common.Address]string,
	map[common.Address]*big.Int,
	error,
) {
	// Open file.
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	var addresses []common.Address
	symbols := make(map[common.Address]string)
	fees := make(map[common.Address]*big.Int)

	// Scan the files.
	newScanner := bufio.NewScanner(file)
	for newScanner.Scan() {
		text := newScanner.Text()

		// Skip if line starts with "#"
		if len(text) == 0 || string(text[0]) == "#" {
			continue
		}

		tmp := strings.Split(text, " ")

		if len(tmp) < 2 {
			continue
		}

		address := common.HexToAddress(tmp[0])

		// Check if already added.
		if slices.Contains(addresses, address) {
			logger.Log.WithField("addr", address.String()).Warningln("Duplicate address detected! Skipping...")
			continue
		}

		if !isRouter {
			symbols[address] = tmp[1]
		} else {
			fees[address], _ = new(big.Int).SetString(tmp[1], 10)
			symbols[address] = tmp[2]
		}
		addresses = append(addresses, address)
	}

	if err := newScanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return addresses, symbols, fees, nil
}
