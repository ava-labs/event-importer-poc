package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/coreth/ethclient"
	proofutils "github.com/ava-labs/receipt-proofs-poc/proofs"
	"github.com/ava-labs/receipt-proofs-poc/tests/utils"
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
	mockPriceFeedAggegratorAddress, mockPriceFeedAggregator := utils.DeployMockPriceFeedAggregator(ctx, fundedKey, cChainInfo)

	// Deploy Price Feed Importer contract on Subnet A
	priceFeedImporterAddress, priceFeedImporter := utils.DeployPriceFeedImporter(ctx, fundedKey, subnetAInfo, cChainInfo.BlockchainID, mockPriceFeedAggegratorAddress)

	// Update the Mock Price Feed contract on C-Chain
	mockRound := big.NewInt(42)
	mockAnswer := big.NewInt(121212121212)
	updateAnswerReceipt := utils.UpdateMockPriceFeedAnswer(ctx, fundedKey, cChainInfo, mockPriceFeedAggregator, mockRound, mockAnswer)

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
	proofDB, err := proofutils.ConstructCorethReceiptProof(ctx, corethClient, updateAnswerReceipt.BlockHash, updateAnswerReceipt.TransactionIndex)
	Expect(err).Should(BeNil())
	encodedProof := proofutils.EncodeMerkleProof(proofDB)

	// Get the log index of the AnswerUpdated event
	answerUpdatedLogIndex, _, err := utils.GetEventFromLogs(updateAnswerReceipt.Logs, mockPriceFeedAggregator.ParseAnswerUpdated)
	Expect(err).Should(BeNil())

	// Import the AnswerUpdated event into the Price Feed Importer contract on Subnet A
	importEventTx := utils.ConstructImportEventTransaction(
		ctx,
		fundedKey,
		fundedAddress,
		subnetAInfo,
		priceFeedImporterAddress,
		signedBlockHashMessage,
		encodedHeader,
		updateAnswerReceipt.TransactionIndex,
		encodedProof,
		answerUpdatedLogIndex,
	)

	// Send the transaction and check that it includes the expected log.
	err = subnetAInfo.RPCClient.SendTransaction(ctx, importEventTx)
	Expect(err).Should(BeNil())
	importEventReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnetAInfo, importEventTx.Hash())
	_, answerUpdatedLog, err := utils.GetEventFromLogs(importEventReceipt.Logs, priceFeedImporter.ParseAnswerUpdated)
	Expect(err).Should(BeNil())
	log.Info("Successfully imported event to update price feed", "txReceipt", importEventReceipt, "log", answerUpdatedLog)
}
