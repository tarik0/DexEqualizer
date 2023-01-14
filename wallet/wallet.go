package wallet

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// Wallet is the helper function to sign transactions.
type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	chainId    *big.Int
	address    common.Address
}

// PrivateKey returns the private key of the wallet.
func (wallet *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return wallet.privateKey
}

// PublicKey returns the public key of the wallet.
func (wallet *Wallet) PublicKey() *ecdsa.PublicKey {
	return wallet.publicKey
}

// Address returns the address of the wallet.
func (wallet *Wallet) Address() common.Address {
	return wallet.address
}

func (wallet *Wallet) NewTransactor() (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(wallet.PrivateKey(), wallet.chainId)
}

// InitWallet initializes a new wallet.
func InitWallet(privateKeyStr string, chainId *big.Int) (*Wallet, error) {
	// Parse the private key.
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	// Get the public key.
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("couldn't cast public key to ECDSA")
	}

	// Get the address
	return &Wallet{
		privateKey: privateKey,
		publicKey:  publicKeyECDSA,
		chainId:    chainId,
		address:    crypto.PubkeyToAddress(*publicKeyECDSA),
	}, nil
}
