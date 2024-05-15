package main

import (
	"context"
	"encoding/hex"
	"os"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

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
		log.Info("Got receipt", "txHash", tx.Hash().String(), "status", receipt.Status, "cumulativeGasUsed", receipt.CumulativeGasUsed, "bloom", hex.EncodeToString(receipt.Bloom[:]), "logs", len(receipt.Logs))
	}
	log.Info("Got receipts", "numReceipts", len(receipts))

	// Create a trie of the receipts
	t, err := trie.New(trie.StateTrieID(common.Hash{}), trie.NewDatabase(nil))
	if err != nil {
		panic(err)
	}
	root := types.DeriveSha(types.Receipts(receipts), t)
	log.Info("Computed trie root", "root", root.String())

	// Verify that the root of the trie matches the receipts root of the block
	if root != blockInfo.ReceiptHash() {
		log.Crit("Roots do not match", "root", root.String(), "blockInfo.ReceiptsRoot", blockInfo.ReceiptHash().String())
		panic("Roots do not match")
	}
	log.Info("Receipt root matches block receipt root")

	memoryDB := memorydb.New()
	receipt1Key, err := rlp.EncodeToBytes(uint(0))
	if err != nil {
		panic(err)
	}
	err = t.Prove(receipt1Key, memoryDB)
	if err != nil {
		panic(err)
	}

	// Print the proof
	log.Info("Created proof")
	it := memoryDB.NewIterator(nil, nil)
	encodedProof := make([]string, 0)
	for it.Next() {
		key := it.Key()
		value := it.Value()
		log.Info("Proof element", "key", hex.EncodeToString(key), "value", hex.EncodeToString(value))
		encodedProof = append(encodedProof, hex.EncodeToString(value))
	}
	log.Info("Encoded proof", "proof", encodedProof)

	// Verify the proof
	val, err := trie.VerifyProof(root, receipt1Key, memoryDB)
	if err != nil {
		panic(err)
	}
	log.Info("Verified proof", "key", hex.EncodeToString(receipt1Key), "value", hex.EncodeToString(val))
}
