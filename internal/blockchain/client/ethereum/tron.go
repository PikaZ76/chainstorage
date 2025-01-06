package ethereum

import (
	"github.com/coinbase/chainstorage/internal/blockchain/client/internal"
	"github.com/coinbase/chainstorage/internal/blockchain/parser/ethereum/types"
)

func NewTronClientFactory(params internal.JsonrpcClientParams) internal.ClientFactory {
	// Tron shares the same data schema as Ethereum since it is an EVM chain, but skip Trace data first.
	return NewEthereumClientFactory(params, WithEthereumNodeType(types.EthereumNodeType_ARCHIVAL), WithEthereumTraceType(types.TraceType_PARITY))

}
