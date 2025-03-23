package stargate

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Transaction(client *ethclient.Client, chainID *big.Int, toAddress common.Address, value *big.Int, txData []byte, privateKey *ecdsa.PrivateKey) (hash common.Hash, err error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	var gasPrice *big.Int
	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	var gasLimit uint64
	gasLimit, err = client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &toAddress,
		Value:    value,
		Data:     txData,
		GasPrice: gasPrice,
	})
	if err != nil {
		return
	}
	var nonce uint64
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, txData)

	var signedTx *types.Transaction
	signedTx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	hash = signedTx.Hash()

	return
}
