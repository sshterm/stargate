package stargate

import (
	"bytes"
	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed stargate.json
var ABI []byte

//go:embed IERC20.json
var IERC20 []byte

var IStargateABI abi.ABI
var IERC20ABI abi.ABI

func init() {
	initABI()
}
func initABI() {
	IStargateABI, _ = abi.JSON(bytes.NewReader(ABI))
	IERC20ABI, _ = abi.JSON(bytes.NewReader(IERC20))
}
