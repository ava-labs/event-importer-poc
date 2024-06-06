// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	corethethclient "github.com/ava-labs/coreth/ethclient"
	proofutils "github.com/ava-labs/receipt-proofs-poc/proofs"
	subnetevmethclient "github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/log"
)

/**
 * Utility tool for building a Merkle proof for a transaction receipt.
 * Required parameters are the RPC URL, block hash, and transaction index.
 * Example usage:
 * go run proofs/tools/proof_builder.go --rpc-url https://api.avax.network/ext/bc/C/rpc \
 * --blockhash 0x27bfa26de1022c4d9cbbc67d6819037ce8a4de99832f00105d7be9d2b4d25369 --txindex 1
 */
func main() {
	// Set log output to stdout
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))

	// Parse the RPC URL, block hash, and transaction index from the command line flags
	var rpcEndpoint string
	var blockHash string
	var txIndex int
	flag.StringVar(&rpcEndpoint, "rpc-url", "", "RPC URL")
	flag.StringVar(&blockHash, "block-hash", "", "Block hash")
	flag.IntVar(&txIndex, "tx-index", -1, "Transaction index")
	flag.Parse()
	if len(rpcEndpoint) == 0 {
		panic("RPC URL is required")
	}
	if len(blockHash) == 0 {
		panic("Block hash is required")
	}
	if txIndex < 0 {
		panic("Transaction index is required")
	}

	buildProofFromRPC(rpcEndpoint, blockHash, uint(txIndex))
}

func buildProofFromRPC(rpcURL string, blockHashStr string, txIndex uint) {
	// If the rpcURL contains "/ext/bc/C/", use the coreth client, otherwise use the subnet-evm client
	ctx := context.Background()
	var memoryDB *memorydb.Database
	if strings.Contains(rpcURL, "/ext/bc/C/") {
		// Create the ethclient
		ethClient, err := corethethclient.DialContext(ctx, rpcURL)
		if err != nil {
			panic(err)
		}

		// Create the proof.
		memoryDB, err = proofutils.ConstructCorethReceiptProof(ctx, ethClient, common.HexToHash(blockHashStr), txIndex)
		if err != nil {
			panic(err)
		}
	} else {
		// Create the ethclient
		ethClient, err := subnetevmethclient.DialContext(ctx, rpcURL)
		if err != nil {
			panic(err)
		}

		// Create the proof.
		memoryDB, err = proofutils.ConstructSubnetEVMReceiptProof(ctx, ethClient, common.HexToHash(blockHashStr), txIndex)
		if err != nil {
			panic(err)
		}
	}

	// Print the proof
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
}
