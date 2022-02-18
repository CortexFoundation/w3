// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*rpcTransactionMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r RPCTransaction) MarshalJSON() ([]byte, error) {
	type RPCTransaction struct {
		BlockHash        *common.Hash      `json:"blockHash"`
		BlockNumber      *hexutil.Big      `json:"blockNumber"`
		From             common.Address    `json:"from"`
		Gas              hexutil.Uint64    `json:"gas"`
		GasPrice         *hexutil.Big      `json:"gasPrice"`
		GasFeeCap        *hexutil.Big      `json:"maxFeePerGas,omitempty"`
		GasTipCap        *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
		Hash             common.Hash       `json:"hash"`
		Input            hexutil.Bytes     `json:"input"`
		Nonce            hexutil.Uint64    `json:"nonce"`
		To               *common.Address   `json:"to"`
		TransactionIndex *hexutil.Uint64   `json:"transactionIndex"`
		Value            *hexutil.Big      `json:"value"`
		Type             hexutil.Uint64    `json:"type"`
		Accesses         *types.AccessList `json:"accessList,omitempty"`
		ChainID          *hexutil.Big      `json:"chainId,omitempty"`
		V                *hexutil.Big      `json:"v"`
		R                *hexutil.Big      `json:"r"`
		S                *hexutil.Big      `json:"s"`
	}
	var enc RPCTransaction
	enc.BlockHash = r.BlockHash
	enc.BlockNumber = (*hexutil.Big)(r.BlockNumber)
	enc.From = r.From
	enc.Gas = hexutil.Uint64(r.Gas)
	enc.GasPrice = (*hexutil.Big)(r.GasPrice)
	enc.GasFeeCap = (*hexutil.Big)(r.GasFeeCap)
	enc.GasTipCap = (*hexutil.Big)(r.GasTipCap)
	enc.Hash = r.Hash
	enc.Input = r.Input
	enc.Nonce = hexutil.Uint64(r.Nonce)
	enc.To = r.To
	enc.TransactionIndex = (*hexutil.Uint64)(r.TransactionIndex)
	enc.Value = (*hexutil.Big)(r.Value)
	enc.Type = hexutil.Uint64(r.Type)
	enc.Accesses = r.Accesses
	enc.ChainID = (*hexutil.Big)(r.ChainID)
	enc.V = (*hexutil.Big)(r.V)
	enc.R = (*hexutil.Big)(r.R)
	enc.S = (*hexutil.Big)(r.S)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *RPCTransaction) UnmarshalJSON(input []byte) error {
	type RPCTransaction struct {
		BlockHash        *common.Hash      `json:"blockHash"`
		BlockNumber      *hexutil.Big      `json:"blockNumber"`
		From             *common.Address   `json:"from"`
		Gas              *hexutil.Uint64   `json:"gas"`
		GasPrice         *hexutil.Big      `json:"gasPrice"`
		GasFeeCap        *hexutil.Big      `json:"maxFeePerGas,omitempty"`
		GasTipCap        *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
		Hash             *common.Hash      `json:"hash"`
		Input            *hexutil.Bytes    `json:"input"`
		Nonce            *hexutil.Uint64   `json:"nonce"`
		To               *common.Address   `json:"to"`
		TransactionIndex *hexutil.Uint64   `json:"transactionIndex"`
		Value            *hexutil.Big      `json:"value"`
		Type             *hexutil.Uint64   `json:"type"`
		Accesses         *types.AccessList `json:"accessList,omitempty"`
		ChainID          *hexutil.Big      `json:"chainId,omitempty"`
		V                *hexutil.Big      `json:"v"`
		R                *hexutil.Big      `json:"r"`
		S                *hexutil.Big      `json:"s"`
	}
	var dec RPCTransaction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.BlockHash != nil {
		r.BlockHash = dec.BlockHash
	}
	if dec.BlockNumber != nil {
		r.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.From != nil {
		r.From = *dec.From
	}
	if dec.Gas != nil {
		r.Gas = uint64(*dec.Gas)
	}
	if dec.GasPrice != nil {
		r.GasPrice = (*big.Int)(dec.GasPrice)
	}
	if dec.GasFeeCap != nil {
		r.GasFeeCap = (*big.Int)(dec.GasFeeCap)
	}
	if dec.GasTipCap != nil {
		r.GasTipCap = (*big.Int)(dec.GasTipCap)
	}
	if dec.Hash != nil {
		r.Hash = *dec.Hash
	}
	if dec.Input != nil {
		r.Input = *dec.Input
	}
	if dec.Nonce != nil {
		r.Nonce = uint64(*dec.Nonce)
	}
	if dec.To != nil {
		r.To = dec.To
	}
	if dec.TransactionIndex != nil {
		r.TransactionIndex = (*uint64)(dec.TransactionIndex)
	}
	if dec.Value != nil {
		r.Value = (*big.Int)(dec.Value)
	}
	if dec.Type != nil {
		r.Type = uint64(*dec.Type)
	}
	if dec.Accesses != nil {
		r.Accesses = dec.Accesses
	}
	if dec.ChainID != nil {
		r.ChainID = (*big.Int)(dec.ChainID)
	}
	if dec.V != nil {
		r.V = (*big.Int)(dec.V)
	}
	if dec.R != nil {
		r.R = (*big.Int)(dec.R)
	}
	if dec.S != nil {
		r.S = (*big.Int)(dec.S)
	}
	return nil
}