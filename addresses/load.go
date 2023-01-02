package addresses

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"strings"
)

// LoadTokens
//	Loads tokens from the tokens.txt file.
func LoadTokens(chainId *big.Int) ([]common.Address, map[common.Address]string, map[common.Address]*big.Int, error) {
	t1, t2, t3, err := loadAddresses(fmt.Sprintf("./chains/%d/tokens.txt", chainId), true)
	return t1, t2, t3, err
}

// LoadRouters
//	Loads routers from the routers.txt file.
func LoadRouters(chainId *big.Int) ([]common.Address, map[common.Address]string, map[common.Address]*big.Int, error) {
	return loadAddresses(fmt.Sprintf("./chains/%d/routers.txt", chainId), true)
}

// loadAddresses
//	Loads addresses from path.
func loadAddresses(path string, isRouter bool) (
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
		addresses = append(addresses, address)
		if !isRouter {
			symbols[address] = tmp[1]
		} else {
			fees[address], _ = new(big.Int).SetString(tmp[1], 10)
			symbols[address] = tmp[2]
		}
	}

	if err := newScanner.Err(); err != nil {
		return nil, nil, nil, err
	}

	return addresses, symbols, fees, nil
}
