package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://geth.marcelonode.xyz")
	if err != nil {
		log.Fatalln("get client error: ", err.Error())
		return
	}
	// tx, err := SendETH(client)
	tx, err := CallSC(client)
	if err != nil {
		log.Fatalln("sending tx error: ", err.Error())
		return
	}
	log.Println("Transaction sent ", tx.Hash().Hex())
}

func SendETH(client *ethclient.Client) (tx *types.Transaction, err error) {
	value := big.NewInt(100000000000000000)

	privateKey, err := crypto.HexToECDSA("53616f5061756c6f42617263656c6f6e614d616e696c6143616e617269617339")
	if err != nil {
		log.Fatalln("get privateKey error: ", err.Error())
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("get publicKeyECDSA error: ", err.Error())
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln("get PendingNonceAt error: ", err.Error())
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln("get gas price error: ", err.Error())
		return
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalln("get chainID error: ", err.Error())
		return
	}

	to := common.HexToAddress("0x263C3Ab7E4832eDF623fBdD66ACee71c028Ff591")
	txData := &types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      uint64(21000),
		GasPrice: gasPrice,
		Data:     nil,
	}
	tx = types.NewTx(txData)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalln("signedTx error: ", err.Error())
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalln("signedTx error: ", err.Error())
		return
	}

	tx = signedTx
	return
}

func CallSC(client *ethclient.Client) (tx *types.Transaction, err error) {
	err = nil
	contract, err := NewRentalRecords(common.HexToAddress("0x12FF95443a1f08CF5c1c01BD3EcA789818a0ae12"), client)
	if err != nil {
		log.Fatalln("get contract error: ", err.Error())
		return
	}
	owner, err := contract.Owner(nil)
	if err != nil {
		log.Fatalln("get contract owner error: ", err.Error())
		return
	}
	log.Println("Contract owner is ", owner.String())

	auth, err := GetAuth(client, 921000)
	if err != nil {
		log.Fatalln("get auth error: ", err.Error())
		return
	}

	// callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	log.Println("Submitting tx...")
	tx, err = contract.RegisterRental(auth, "Joselito", "Roger", "Rua do Banjo 123", big.NewInt(1000))
	if err != nil {
		log.Fatalln("RegisterRental error: ", err.Error())
		return
	}
	log.Println("Waiting tx...", tx.Hash().Hex())
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatalln("RegisterRental error: ", receipt.Status)
		return
	}
	return
}

func GetAuth(client *ethclient.Client, gasLimit int) (auth *bind.TransactOpts, err error) {
	err = nil
	privateKey, err := crypto.HexToECDSA("53616f5061756c6f42617263656c6f6e614d616e696c6143616e617269617339")
	if err != nil {
		log.Fatalln("get privateKey error: ", err.Error())
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("get publicKeyECDSA error: ", err.Error())
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln("get PendingNonceAt error: ", err.Error())
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln("get gas price error: ", err.Error())
		return
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalln("get chainID error: ", err.Error())
		return
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalln("get auth error: ", err.Error())
		return
	}
	// auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = gasPrice
	auth.Context = context.Background()

	return
}
