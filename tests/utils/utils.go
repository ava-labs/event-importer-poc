// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	ieventimporter "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/IEventImporter"
	pricefeedimporter "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/PriceFeedImporter"
	mockpricefeedaggregator "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/mocks/MockPriceFeedAggregator"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

func GetSignedBlockHashMessage(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	importingSubnetID ids.ID,
	blockHash common.Hash,
) []byte {
	warpClient, err := warpBackend.NewClient(subnet.NodeURIs[0], subnet.BlockchainID.String())
	Expect(err).Should(BeNil())

	signedWarpMessageBytes, err := warpClient.GetBlockAggregateSignature(
		ctx,
		ids.ID(blockHash),
		warp.WarpDefaultQuorumNumerator,
		importingSubnetID.String(),
	)
	Expect(err).Should(BeNil())
	log.Info(
		"Got signed block hash message",
		"blockHash", blockHash.String(),
		"signedWarpMessageBytes", hex.EncodeToString(signedWarpMessageBytes),
	)

	return signedWarpMessageBytes
}

func DeployMockPriceFeedAggregator(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *mockpricefeedaggregator.MockPriceFeedAggregator) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, aggregator, err := mockpricefeedaggregator.DeployMockPriceFeedAggregator(
		opts,
		subnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	log.Info("Created Mock Price Feed contract", "address", address.Hex())
	return address, aggregator
}

func DeployPriceFeedImporter(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	aggregatorBlockchainID ids.ID,
	aggregatorAddress common.Address,
) (common.Address, *pricefeedimporter.PriceFeedImporter) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, importer, err := pricefeedimporter.DeployPriceFeedImporter(
		opts,
		subnet.RPCClient,
		aggregatorBlockchainID,
		aggregatorAddress,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	log.Info("Created Price Feed Importer contract", "address", address.Hex())
	return address, importer
}

func UpdateMockPriceFeedAnswer(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	mockPriceFeedAggregator *mockpricefeedaggregator.MockPriceFeedAggregator,
	round *big.Int,
	answer *big.Int,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	updateAnswerTx, err := mockPriceFeedAggregator.UpdateAnswer(opts, answer, round, big.NewInt(time.Now().Unix()))
	Expect(err).Should(BeNil())
	updateAnswerReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, updateAnswerTx.Hash())
	log.Info("Updated Mock Price Feed contract", "txHash", updateAnswerTx.Hash().Hex())
	return updateAnswerReceipt
}

func ConstructImportEventTransaction(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	senderAddress common.Address,
	subnet interfaces.SubnetTestInfo,
	eventImporterAddress common.Address,
	signedBlockHashMessage []byte,
	encodedBlockHeader []byte,
	txIndex uint,
	encodedReceiptProof [][]byte,
	logIndex int,
) *types.Transaction {
	importerABI, err := ieventimporter.IEventImporterMetaData.GetAbi()
	Expect(err).Should(BeNil())
	importEventData, err := importerABI.Pack(
		"importEvent",
		encodedBlockHeader,
		big.NewInt(int64(txIndex)),
		encodedReceiptProof,
		big.NewInt(int64(logIndex)),
	)
	Expect(err).Should(BeNil())
	gasFeeCap, gasTipCap, nonce := teleporterUtils.CalculateTxParams(ctx, subnet, senderAddress)
	unsignedImportEventTx := predicateutils.NewPredicateTx(
		subnet.EVMChainID,
		nonce,
		&eventImporterAddress,
		3_000_000, // TODO: How much gas is needed?
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		importEventData,
		types.AccessList{},
		warp.ContractAddress,
		signedBlockHashMessage,
	)
	return teleporterUtils.SignTransaction(unsignedImportEventTx, senderKey, subnet.EVMChainID)
}

// Returns the first log in 'logs' that is successfully parsed by 'parser'
func GetEventFromLogs[T any](logs []*types.Log, parser func(log types.Log) (T, error)) (int, T, error) {
	for i, log := range logs {
		event, err := parser(*log)
		if err == nil {
			return i, event, nil
		}
	}
	return -1, *new(T), fmt.Errorf("failed to find %T event in receipt logs", *new(T))
}
