package main

import (
	"context"
	"encoding/hex"
	"os"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

type slicePutter struct {
	slice [][]byte
}

func newSlicePutter() *slicePutter {
	return &slicePutter{
		slice: make([][]byte, 0),
	}
}

func (sp *slicePutter) Put(key []byte, value []byte) error {
	log.Info("Slice putter put", "key", hex.EncodeToString(key), "value", hex.EncodeToString(value))
	sp.slice = append(sp.slice, value)
	return nil
}

func (sp *slicePutter) Delete(key []byte) error {
	// No-op
	return nil
}

func main() {
	// Set log output to stdout
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))

	// Create the ethclient
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, "https://api.avax-test.network/ext/bc/C/rpc")
	if err != nil {
		panic(err)
	}

	// Get the block info
	blockHash := common.HexToHash("0x5504425badb74aa81d4f6a028a8c3b27cc364ad8d251c123a5c48b7e479e4d1f")
	blockInfo, err := ethClient.BlockByHash(ctx, blockHash)
	if err != nil {
		panic(err)
	}
	log.Info("Got block", "blockHash", blockHash.String())

	// Get the receipts for each transaction in the block
	receipts := make([]*types.Receipt, blockInfo.Transactions().Len())
	for i, tx := range blockInfo.Transactions() {
		receipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			panic(err)
		}
		receipts[i] = receipt
	}

	receipts100 := make([]*types.Receipt, 100*len(receipts))
	for i := 0; i < 100; i++ {
		for j, r := range receipts {
			receipts100[(i*len(receipts))+j] = r
		}
	}
	log.Info("Got receipts", "numReceipts", len(receipts100))

	// Create a trie of the receipts
	t, err := trie.New(trie.StateTrieID(common.Hash{}), trie.NewDatabase(nil))
	if err != nil {
		panic(err)
	}
	root := types.DeriveSha(types.Receipts(receipts100), t)
	log.Info("Computed trie root", "root", root.String())

	// Verify that the root of the trie matches the receipts root of the block
	// if root != blockInfo.ReceiptHash() {
	// 	log.Crit("Roots do not match", "root", root.String(), "blockInfo.ReceiptsRoot", blockInfo.ReceiptHash().String())
	// 	panic("Roots do not match")
	// }
	// log.Info("Receipt root matches block receipt root")

	sp := newSlicePutter()
	receipt1Key, err := rlp.EncodeToBytes(uint(0))
	if err != nil {
		panic(err)
	}
	err = t.Prove(receipt1Key, sp)
	if err != nil {
		panic(err)
	}

	log.Info("Created proof")
	for _, proofElement := range sp.slice {
		log.Info("Proof element", "element", hex.EncodeToString(proofElement))
	}

}
