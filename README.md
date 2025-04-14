# Stargate Cross-Chain Bridge

## 概述

Stargate 是一个高性能、低费用的跨链桥接协议，允许用户在不同的区块链之间无缝转移代币。本项目实现了基于 Stargate 协议的跨链桥接功能，支持多种区块链网络之间的代币转移。

## 特性

- **高性能**：利用 Stargate 协议的高效性，实现快速的跨链交易。
- **低费用**：相比其他跨链解决方案，Stargate 提供了更具竞争力的费用结构。
- **多链支持**：支持多种主流区块链网络，包括但不限于 Ethereum、BSC、Polygon 等。
- **安全性**：采用先进的加密技术和安全协议，确保用户的资产安全。

## 安装

确保您已经安装了 Go 语言环境。然后，您可以使用以下命令安装项目依赖：

```bash
go get github.com/sshterm/stargate
```

## 使用示例
以下是一个简单的使用示例，展示如何通过 Stargate 进行跨链代币转移：

```go
package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/shopspring/decimal"
    "github.com/sshterm/stargate"
)

func main() {
    rpc := "https://bsc-rpc.publicnode.com"
    privateKey := []byte
    to := common.HexToAddress("0xRecipientAddress")

    sg := stargate.NewStargate(rpc, privateKey, stargate.USDT_BSC_TO_ETH, to)
    amount := decimal.NewFromFloat(1.0)

    hash, err := sg.Bridge(30101, amount) // 30101 是目标链的 以太坊 ID
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Transaction Hash:", hash.Hex())
}
```
## Mainnet Contracts
https://stargateprotocol.gitbook.io/stargate/v2-developer-docs/technical-reference/mainnet-contracts