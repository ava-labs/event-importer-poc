package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/trie"
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
	ethClient, err := ethclient.DialContext(ctx, "https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		panic(err)
	}

	// Get the block info
	blockHash := common.HexToHash("0x2d9215bce478eb82bfd35f7e9bdc9d76e1814e8d7b5aa10ab05e2f17d145c0cf")
	blockInfo, err := ethClient.BlockByHash(ctx, blockHash)
	if err != nil || blockInfo == nil {
		panic(err)
	}
	if blockInfo.Hash() != blockHash {
		panic("Block hash does not match")
	}

	encodedHeader, err := rlp.EncodeToBytes(blockInfo.Header())
	if err != nil {
		panic(err)
	}
	log.Info("Got block", "blockHash", blockHash.String())
	log.Info("RLP Encoded block header", "header", hex.EncodeToString(encodedHeader))
	log.Info("Actual header hash", "hash", blockInfo.Header().Hash().String())
	log.Info("Actual block hash", "hash", blockInfo.Hash().String())

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
	receipt1Key, err := rlp.EncodeToBytes(uint(8))
	if err != nil {
		panic(err)
	}
	err = t.Prove(receipt1Key, 0, memoryDB)
	if err != nil {
		panic(err)
	}

	// Print the proof
	log.Info("Created proof")
	it := memoryDB.NewIterator(nil, nil)
	encodedProof := make([]string, 0)
	i := 0
	for it.Next() {
		value := it.Value()
		encodedProof = append(encodedProof, hex.EncodeToString(value))
		formattedProofElem := fmt.Sprintf("proof[%d] = hex\"%s\";", i, hex.EncodeToString(value))
		log.Info("Encoded proof element", "key", hex.EncodeToString(it.Key()), "elem", formattedProofElem)
		i++
	}
	log.Info("Encoded proof", "proof", encodedProof)

	// Verify the proof
	val, err := trie.VerifyProof(root, receipt1Key, memoryDB)
	if err != nil {
		panic(err)
	}
	log.Info("Verified proof", "key", hex.EncodeToString(receipt1Key), "value", hex.EncodeToString(val))
}
