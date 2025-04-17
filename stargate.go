package stargate

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type Stargate struct {
	rpc        string
	privateKey []byte
	to         common.Address
	chain      Chain
}

// NewStargate 创建一个新的 Stargate 实例
// 参数:
//   - rpc: RPC 节点地址
//   - privateKey: 私钥
//   - chain: 链配置信息
//   - to: 目标地址
//
// 返回:
//   - *Stargate: 返回一个初始化的 Stargate 指针
func NewStargate(rpc string, privateKey []byte, chain Chain, to common.Address) *Stargate {
	return &Stargate{rpc: rpc, privateKey: privateKey, chain: chain, to: to}
}

// Approve 方法批准一个地址可以花费不超过指定数量的代币。
// 这个方法主要用于在以太坊网络上批准一个桥接地址，允许它代表用户转移一定数量的代币。
// 参数:
//   - amount (*big.Int): 要批准的最大代币数量。
//
// 返回值:
//   - hash (common.Hash): 批准交易的哈希值。
//   - err (error): 如果有错误发生，返回错误信息。
func (s *Stargate) Approve(amount *big.Int) (hash common.Hash, err error) {
	var client *ethclient.Client
	client, err = ethclient.Dial(s.rpc)
	if err != nil {
		return
	}
	defer client.Close()
	var txData []byte
	txData, err = IERC20ABI.Pack("approve", s.chain.BridgeAddress, amount)
	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.ToECDSA(s.privateKey)
	if err != nil {
		return
	}

	toAddress := s.chain.Address

	value := big.NewInt(0)
	hash, err = Transaction(client, big.NewInt(int64(s.chain.ChainID)), toAddress, value, txData, privateKey)
	return
}

// Bridge 通过 Stargate 协议在不同链之间转移代币
// 参数:
//   - dstEid: 目标链的 LayerZero ID
//   - amount: 要转移的代币数量
//
// 返回:
//   - hash: 交易哈希
//   - err: 错误信息
//
// Bridge 函数会验证代币地址,计算跨链费用,并发送跨链交易
func (s *Stargate) Bridge(dstEid int, amount decimal.Decimal) (hash common.Hash, err error) {
	var client *ethclient.Client
	client, err = ethclient.Dial(s.rpc)
	if err != nil {
		return
	}
	defer client.Close()
	var token common.Address
	token, err = s.quoteToken(client)
	if err != nil {
		return
	}
	if token != s.chain.Address {
		err = errors.New("token address does not match chain address")
		return
	}
	sendParam := StargateSendParam{
		DstEid:       uint32(dstEid),
		To:           AddressToBytes32(s.to),
		AmountLD:     ToWei(amount, s.chain.Decimals),
		MinAmountLD:  ToWei(amount, s.chain.Decimals),
		ExtraOptions: []byte{},
		ComposeMsg:   []byte{},
		OftCmd:       []byte{},
	}
	err = s.quoteOFT(client, &sendParam)
	if err != nil {
		return
	}
	var messageFee *StargateMessagingFee
	messageFee, err = s.quoteSend(client, &sendParam)
	if err != nil {
		return
	}

	var txData []byte
	txData, err = IStargateABI.Pack("send", sendParam, messageFee, s.to)
	if err != nil {
		return
	}
	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.ToECDSA(s.privateKey)
	if err != nil {
		return
	}

	toAddress := s.chain.BridgeAddress

	value := messageFee.NativeFee
	hash, err = Transaction(client, big.NewInt(int64(s.chain.ChainID)), toAddress, value, txData, privateKey)
	return
}

func (s *Stargate) quoteToken(client *ethclient.Client) (token common.Address, err error) {
	var data []byte
	data, err = IStargateABI.Pack("token")
	if err != nil {
		return
	}
	var res []byte
	res, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &s.chain.BridgeAddress,
		Data: data,
	}, nil)
	if err != nil {
		return
	}
	err = IStargateABI.UnpackIntoInterface(&token, "token", res)
	if err != nil {
		return
	}

	return
}
func (s *Stargate) quoteOFT(client *ethclient.Client, sendParam *StargateSendParam) (err error) {
	var data []byte
	data, err = IStargateABI.Pack("quoteOFT", sendParam)
	if err != nil {
		return
	}
	var res []byte
	res, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &s.chain.BridgeAddress,
		Data: data,
	}, nil)
	if err != nil {
		return
	}
	var resOFT StargateResOFT
	err = IStargateABI.UnpackIntoInterface(&resOFT, "quoteOFT", res)
	if err != nil {
		return
	}
	sendParam.MinAmountLD = resOFT.Limit.MinAmountLD
	return
}
func (s *Stargate) quoteSend(client *ethclient.Client, sendParam *StargateSendParam) (fee *StargateMessagingFee, err error) {
	var data []byte
	data, err = IStargateABI.Pack("quoteSend", sendParam, false)
	if err != nil {
		return
	}
	var res []byte
	res, err = client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &s.chain.BridgeAddress,
		Data: data,
	}, nil)
	if err != nil {
		return
	}
	var feeRes StargateQuoteResult
	err = IStargateABI.UnpackIntoInterface(&feeRes, "quoteSend", res)
	return &feeRes.Fee, err
}
