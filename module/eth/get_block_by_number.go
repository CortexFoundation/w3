package eth

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// BlockByNumber requests the block with full transactions with the given
// number.
func BlockByNumber(number *big.Int) *BlockByNumberFactory {
	return &BlockByNumberFactory{number: number}
}

type BlockByNumberFactory struct {
	// args
	number *big.Int

	// returns
	result  json.RawMessage
	returns *types.Block
}

func (f *BlockByNumberFactory) Returns(block *types.Block) *BlockByNumberFactory {
	f.returns = block
	return f
}

// CreateRequest implements the core.RequestCreator interface.
func (f *BlockByNumberFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_getBlockByNumber",
		Args:   []interface{}{toBlockNumberArg(f.number), true},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the core.ResponseHandler interface.
func (f *BlockByNumberFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	if len(f.result) <= 4 { // empty or "null"
		return errNotFound
	}

	var head *types.Header
	var body rpcBlock
	if err := json.Unmarshal(f.result, &head); err != nil {
		return err
	}
	if err := json.Unmarshal(f.result, &body); err != nil {
		return err
	}

	block := types.NewBlockWithHeader(head).WithBody(body.Transactions, nil)
	*f.returns = *block
	return nil
}

// HeaderByNumber requests the header with the given number.
func HeaderByNumber(number *big.Int) *HeaderByNumberFactory {
	return &HeaderByNumberFactory{number: number}
}

type HeaderByNumberFactory struct {
	// args
	number *big.Int

	// returns
	result  *types.Header
	returns *types.Header
}

func (f *HeaderByNumberFactory) Returns(header *types.Header) *HeaderByNumberFactory {
	f.returns = header
	return f
}

// CreateRequest implements the core.RequestCreator interface.
func (f *HeaderByNumberFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_getBlockByNumber",
		Args:   []interface{}{toBlockNumberArg(f.number), false},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the core.ResponseHandler interface.
func (f *HeaderByNumberFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	*f.returns = *f.result
	return nil
}
