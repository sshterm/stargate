package stargate

import "math/big"

type StargateSendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}
type StargateMessagingFee struct {
	NativeFee  *big.Int `json:"nativeFee"`
	LzTokenFee *big.Int `json:"lzTokenFee"`
}

type StargateQuoteResult struct {
	Fee StargateMessagingFee
}

type StargateReceipt struct {
	AmountSentLD     *big.Int `json:"amountSentLD"`
	AmountReceivedLD *big.Int `json:"amountReceivedLD"`
}
type StargateOftFeeDetails struct {
	FeeAmountLD *big.Int `json:"feeAmountLD"`
	Description string   `json:"description"`
}
type StargateLimit struct {
	MinAmountLD *big.Int `json:"minAmountLD"`
	MaxAmountLD *big.Int `json:"maxAmountLD"`
}
type StargateResOFT struct {
	Limit         StargateLimit           `json:"limit"`
	OftFeeDetails []StargateOftFeeDetails `json:"oftFeeDetails"`
	Receipt       StargateReceipt         `json:"receipt"`
}
