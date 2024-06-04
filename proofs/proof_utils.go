package proofs

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/ethclient"
	"github.com/ava-labs/coreth/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

func ConstructReceiptProof(
	ctx context.Context,
	ethClient ethclient.Client,
	blockHash common.Hash,
	txIndex uint,
) (*memorydb.Database, error) {
	// Get the block info
	blockInfo, err := ethClient.BlockByHash(ctx, blockHash)
	if err != nil || blockInfo == nil {
		log.Error("Failed to get block info", "blockHash", blockHash.String(), "err", err)
		return nil, err
	}
	if blockInfo.Hash() != blockHash {
		log.Error("Block hash does not match", "blockHash", blockHash.String())
		return nil, fmt.Errorf("block hash does not match")
	}

	// Get the receipts for each transaction in the block
	receipts := make([]*types.Receipt, blockInfo.Transactions().Len())
	for i, tx := range blockInfo.Transactions() {
		receipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			log.Error("Failed to get transaction receipt", "txHash", tx.Hash().String(), "err", err)
			return nil, err
		}
		receipts[i] = receipt
		encodedReceipt, err := rlp.EncodeToBytes(receipt)
		if err != nil {
			log.Error("Failed to encode receipt", "txHash", tx.Hash().String(), "err", err)
			return nil, err
		}
		log.Info("Got encoded receipt", "txHash", tx.Hash().String(), "receipt", hex.EncodeToString(encodedReceipt))
	}

	// Create a trie of the receipts
	receiptTrie, err := trie.New(trie.StateTrieID(common.Hash{}), trie.NewDatabase(nil))
	if err != nil {
		log.Error("Failed to create receipt trie", "err", err)
		return nil, err
	}

	// Defensive check that the receipts root matches the block header.
	// This should always be the case.
	receiptsRoot := types.DeriveSha(types.Receipts(receipts), receiptTrie)
	if receiptsRoot != blockInfo.Header().ReceiptHash {
		log.Error("Receipts root does not match", "blockHash", blockHash.String())
		return nil, err
	}

	// Construct the proof of the request receipt against the trie.
	key, err := rlp.EncodeToBytes(txIndex)
	if err != nil {
		log.Error("Failed to encode tx index", "err", err)
		return nil, err
	}
	memoryDB := memorydb.New()
	err = receiptTrie.Prove(key, 0, memoryDB)
	if err != nil {
		log.Error("Failed to prove receipt", "err", err)
		return nil, err
	}

	// Double check that the proof is valid.
	verifiedValue, err := trie.VerifyProof(receiptsRoot, key, memoryDB)
	if err != nil {
		log.Error("Failed to verify proof", "err", err)
		return nil, err
	}
	log.Info("Verified proof", "value", hex.EncodeToString(verifiedValue))

	return memoryDB, nil
}
