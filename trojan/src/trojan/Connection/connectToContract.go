package connection

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	PrivateKeyHex string
	ContractAddress string
	ContractABI string
	ChainRPC string
)

func ConnectToChain() (*ethclient.Client, error) {
	client, err := ethclient.Dial(ChainRPC)
	if err != nil {
		return nil, err
	}
	return client, nil
}




