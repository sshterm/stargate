package stargate

import "github.com/ethereum/go-ethereum/common"

type Chain struct {
	Address       common.Address `json:"address"`
	BridgeAddress common.Address `json:"bridgeAddress"`
	Decimals      int            `json:"decimals"`
	Symbol        string         `json:"symbol"`
	EndpointID    int            `json:"endpointID"`
	Name          string         `json:"name"`
	ChainID       int            `json:"chainID"`
}

var USDT_BSC_TO_ETH Chain = Chain{
	Address:       common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	BridgeAddress: common.HexToAddress("0x138EB30f73BC423c6455C53df6D89CB01d9eBc63"),
	Decimals:      18,
	Symbol:        "USDT",
	EndpointID:    30102,
	Name:          "Tether USD",
	ChainID:       56,
}
var USDC_BSC_TO_ETH Chain = Chain{
	Address:       common.HexToAddress("0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"),
	BridgeAddress: common.HexToAddress("0x962Bd449E630b0d928f308Ce63f1A21F02576057"),
	Decimals:      18,
	Symbol:        "USDC",
	EndpointID:    30102,
	Name:          "USD Coin",
	ChainID:       56,
}
var USDT_ETH_TO_BSC Chain = Chain{
	Address:       common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
	BridgeAddress: common.HexToAddress("0x933597a323Eb81cAe705C5bC29985172fd5A3973"),
	Decimals:      6,
	Symbol:        "USDT",
	EndpointID:    30101,
	Name:          "Tether USD",
	ChainID:       1,
}
var USDC_ETH_TO_BSC Chain = Chain{
	Address:       common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	BridgeAddress: common.HexToAddress("0xc026395860Db2d07ee33e05fE50ed7bD583189C7"),
	Decimals:      6,
	Symbol:        "USDC",
	EndpointID:    30101,
	Name:          "USD Coin",
	ChainID:       1,
}
