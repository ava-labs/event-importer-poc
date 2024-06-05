package flows

import (
	"context"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ava-labs/coreth/ethclient"
	pricefeedimporter "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/PriceFeedImporter"
	mockpricefeedaggregator "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/mocks/MockPriceFeedAggregator"
	proofutils "github.com/ava-labs/receipt-proofs-poc/proofs"
	"github.com/ava-labs/receipt-proofs-poc/tests/utils"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	. "github.com/onsi/gomega"
)

/**
 * Deploys an ERC20TokenHub contract on the C-Chain
 * Deploys an ERC20TokenSpoke contract on Subnet A
 * Check sending to  unregistered spoke fails
 * Register the ERC20TokenSpoke to hub contract
 * Check sending to non-collateralized spoke fails
 * Collateralize the spoke
 * Check sending to collateralized spoke succeeds and withdraws with correct scale.
 */
func ImportPriceFeed(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy Mock Price Feed contract on C-Chain
	cChainTransactorOpts, err := bind.NewKeyedTransactorWithChainID(fundedKey, cChainInfo.EVMChainID)
	Expect(err).Should(BeNil())
	mockPriceFeedAggregatorAddress, deployMockPriceFeedTx, mockPriceFeedAggregator, err := mockpricefeedaggregator.DeployMockPriceFeedAggregator(
		cChainTransactorOpts,
		cChainInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, deployMockPriceFeedTx.Hash())
	log.Info("Created Mock Price Feed contract", "address", mockPriceFeedAggregatorAddress.Hex())

	// Deploy Price Feed Importer contract on Subnet A
	subnetATransactorOpts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	priceFeedImporterAddress, deployPriceFeedImporterTx, priceFeedImporter, err := pricefeedimporter.DeployPriceFeedImporter(subnetATransactorOpts, subnetAInfo.RPCClient, cChainInfo.BlockchainID, mockPriceFeedAggregatorAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnetAInfo, deployPriceFeedImporterTx.Hash())
	log.Info("Created Price Feed Importer contract", "address", priceFeedImporterAddress.Hex())

	// Update the Mock Price Feed contract on C-Chain
	mockRound := big.NewInt(42)
	mockAnswer := big.NewInt(121212121212)
	updateAnswerTx, err := mockPriceFeedAggregator.UpdateAnswer(cChainTransactorOpts, mockAnswer, mockRound, big.NewInt(time.Now().Unix()))
	Expect(err).Should(BeNil())
	updateAnswerReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, updateAnswerTx.Hash())
	log.Info("Updated Mock Price Feed contract", "txHash", updateAnswerTx.Hash().Hex())

	// Get the block header
	// Create a custom coreth client so that the block hash from the C-Chain is correct
	chainRPCURI := teleporterUtils.HttpToRPCURI(cChainInfo.NodeURIs[0], cChainInfo.BlockchainID.String())
	corethClient, err := ethclient.Dial(chainRPCURI)
	Expect(err).Should(BeNil())
	blockHeader, err := corethClient.HeaderByHash(ctx, updateAnswerReceipt.BlockHash)
	Expect(err).Should(BeNil())
	encodedHeader, err := rlp.EncodeToBytes(blockHeader)
	Expect(err).Should(BeNil())

	// Get a Warp signature of the block hash containing the AnswerUpdated event
	signedBlockHashMessage := utils.GetSignedBlockHashMessage(ctx, cChainInfo, subnetAInfo.SubnetID, updateAnswerReceipt.BlockHash)

	// Construct a Merkle proof of the AnswerUpdated event agains the block's receipts root
	proofDB, err := proofutils.ConstructReceiptProof(ctx, corethClient, updateAnswerReceipt.BlockHash, updateAnswerReceipt.TransactionIndex)
	Expect(err).Should(BeNil())
	encodedProof := make([][]byte, 0)
	it := proofDB.NewIterator(nil, nil)
	for it.Next() {
		encodedProof = append(encodedProof, it.Value())
	}
	for _, proofElem := range encodedProof {
		log.Info("Encoded proof elem", "proofElem", hex.EncodeToString(proofElem))
	}

	// Get the log index of the AnswerUpdated event
	answerUpdatedLogIndex, _, err := utils.GetEventFromLogs(updateAnswerReceipt.Logs, mockPriceFeedAggregator.ParseAnswerUpdated)
	Expect(err).Should(BeNil())

	// Import the AnswerUpdated event into the Price Feed Importer contract on Subnet A
	priceFeedImporterABI, err := pricefeedimporter.PriceFeedImporterMetaData.GetAbi()
	Expect(err).Should(BeNil())
	importEventData, err := priceFeedImporterABI.Pack("importEvent", encodedHeader, big.NewInt(int64(updateAnswerReceipt.TransactionIndex)), encodedProof, big.NewInt(int64(answerUpdatedLogIndex)))
	Expect(err).Should(BeNil())
	gasFeeCap, gasTipCap, nonce := teleporterUtils.CalculateTxParams(ctx, subnetAInfo, fundedAddress)
	unsignedImportEventTx := predicateutils.NewPredicateTx(
		subnetAInfo.EVMChainID,
		nonce,
		&priceFeedImporterAddress,
		3_000_000, // TODO: How much gas is needed?
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		importEventData,
		types.AccessList{},
		warp.ContractAddress,
		signedBlockHashMessage,
	)
	signedImportEventTx := teleporterUtils.SignTransaction(unsignedImportEventTx, fundedKey, subnetAInfo.EVMChainID)
	encodedTx, err := rlp.EncodeToBytes(signedImportEventTx)
	Expect(err).Should(BeNil())
	log.Info("Created signed tx to import event", "txHash", signedImportEventTx.Hash().Hex(), "rawTx", hex.EncodeToString(encodedTx))
	err = subnetAInfo.RPCClient.SendTransaction(ctx, signedImportEventTx)
	Expect(err).Should(BeNil())
	importEventReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnetAInfo, signedImportEventTx.Hash())
	_, answerUpdatedLog, err := utils.GetEventFromLogs(importEventReceipt.Logs, priceFeedImporter.ParseAnswerUpdated)
	Expect(err).Should(BeNil())
	log.Info("Successfully imported event to update price feed", "txReceipt", importEventReceipt, "log", answerUpdatedLog)
}
