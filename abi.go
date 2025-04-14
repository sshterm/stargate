package stargate

import (
	"bytes"
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed stargate.json
var ABI []byte

var IStargateABI abi.ABI

func init() {
	initABI()
}
func initABI() {
	IStargateABI, _ = abi.JSON(bytes.NewReader(ABI))
}
