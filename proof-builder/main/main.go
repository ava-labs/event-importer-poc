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

func buildProofFromRPC() {
	// Create the ethclient
	ctx := context.Background()
	ethClient, err := ethclient.DialContext(ctx, "https://api.avax-test.network/ext/bc/C/rpc")
	// ethClient, err := ethclient.DialContext(ctx, "https://api.avax.network/ext/bc/C/rpc")
	if err != nil {
		panic(err)
	}

	// Get the block info
	// blockHash := common.HexToHash("0x2d9215bce478eb82bfd35f7e9bdc9d76e1814e8d7b5aa10ab05e2f17d145c0cf")
	blockHash := common.HexToHash("0x6081172c9710380bdd37438af268cc7423a20d6b692c656ff48ed417b05b3d9c")
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
		encodedReceipt, err := rlp.EncodeToBytes(receipt)
		if err != nil {
			log.Error("Failed to encode receipt", "txHash", tx.Hash().String(), "err", err)
			panic(err)
		}
		log.Info("Got encoded receipt", "txHash", tx.Hash().String(), "receipt", hex.EncodeToString(encodedReceipt))
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
	// receipt1Key, err := rlp.EncodeToBytes(uint(6))
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

	// RLP decode the value returned by the proof back into a *type.Receipt
	var receipt types.Receipt
	err = rlp.DecodeBytes(val, &receipt)
	if err != nil {
		panic(err)
	}
	log.Info("Decoded receipt", "txHash", receipt.TxHash.String(), "status", receipt.Status, "cumulativeGasUsed", receipt.CumulativeGasUsed, "bloom", hex.EncodeToString(receipt.Bloom[:]), "logs", len(receipt.Logs))
}

func main() {
	// Set log output to stdout
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))

	// Decode the receipt from a hex string.
	encodedReceiptHex := "b9026902f90265018301d488b9010000000000000000000000000000000020000000001000000000000000000000000008000000000000000200010000000000120000000000000000020000000000000000000000000000000000000000000000001000000000000000000000000000000000080000000000000000000000000000000000000000000000000400000000000000000000000000000000000008000000000000000000000000000000000000000000000000400000000000000000000000000000000002001000000000000000000000000001000000000000000000080000000000000040000000000000000000000000000000000000000000000000000000000000000000000000f9015af838945ff137d4b0fdcd49dca30c7cf57e578a026d2789e1a0bb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97280f9011d945ff137d4b0fdcd49dca30c7cf57e578a026d2789f884a049628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419fa01cbf263a8e79bd77754c7c99d3c4bd6fdc26b91dacc009566af6a3e2bab13235a000000000000000000000000060d058410efa141c72693d0c92dde1aa24440a15a00000000000000000000000008817340e0a3435e06254f2ed411e6418cd070d6fb88000000000000000000000000000000000000000000000000000000000000000a70000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000ba2d33f24e940000000000000000000000000000000000000000000000000000000000001cbcd"
	log.Info("Encoded receipt", "receipt", encodedReceiptHex)
	encodedReceipt, err := hex.DecodeString(encodedReceiptHex)
	if err != nil {
		panic(err)
	}
	var receipt types.Receipt
	err = rlp.DecodeBytes(encodedReceipt, &receipt)
	if err != nil {
		panic(err)
	}

	// Create a trie with the single receipt.
	t, err := trie.New(trie.StateTrieID(common.Hash{}), trie.NewDatabase(nil))
	if err != nil {
		panic(err)
	}
	root := types.DeriveSha(types.Receipts([]*types.Receipt{&receipt}), t)
	log.Info("Computed trie root", "root", root.String())

	// Create a proof for the first (and only) key in the trie.
	memoryDB := memorydb.New()
	key, err := rlp.EncodeToBytes(uint(0))
	if err != nil {
		panic(err)
	}
	err = t.Prove(key, 0, memoryDB)
	if err != nil {
		panic(err)
	}

	// Print the proof.
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

	// Verify the proof, which should return the value of the encoded receipt.
	val, err := trie.VerifyProof(root, key, memoryDB)
	if err != nil {
		panic(err)
	}
	log.Info("Verified proof", "key", hex.EncodeToString(key), "value", hex.EncodeToString(val))

	// RLP decode the value returned by the proof back into a *type.Receipt.
	err = rlp.DecodeBytes(val, &receipt)
	if err != nil {
		panic(err)
	}
	log.Info("Decoded receipt", "txHash", receipt.TxHash.String(), "status", receipt.Status, "cumulativeGasUsed", receipt.CumulativeGasUsed, "bloom", hex.EncodeToString(receipt.Bloom[:]), "logs", len(receipt.Logs))
}
