package ethereum

import (
	"github.com/coinbase/chainstorage/internal/blockchain/parser/ethereum/types"
	"github.com/coinbase/chainstorage/internal/blockchain/parser/internal"
)

func NewTronNativeParser(params internal.ParserParams, opts ...internal.ParserFactoryOption) (internal.NativeParser, error) {
	// Tron shares the same data schema as Ethereum since its an EVM chain except skip trace data
	opts = append(opts, WithEthereumNodeType(types.EthereumNodeType_ALCHEMY_TRON))
	return NewEthereumNativeParser(params, opts...)
}
