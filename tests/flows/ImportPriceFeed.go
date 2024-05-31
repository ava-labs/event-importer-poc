package flows

import (
	"context"

	"github.com/ava-labs/coreth/accounts/abi/bind"
	"github.com/ava-labs/receipt-proof-poc/abi-bindings/go/mocks/mockpricefeedaggregator"
	mockpricefeedaggregator "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/mocks/MockPriceFeedAggregator"
	mockpricefeedaggregator "github.com/ava-labs/receipt-proofs-poc/abi-bindings/go/MockPriceFeedAggregator"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
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
	mockPriceFeedAggregatorAddress, tx, mockPriceFeedAggregator, err := mockpricefeedaggregator.DeployMockPriceFeedAggregator(
		cChainTransactorOpts,
		cChainInfo.RPCClient,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, cChainInfo, tx.Hash())

	// Deploy Price Feed Importer contract on Subnet A
	subnetATransactorOpts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	priceFeedImporterAddress, tx, priceFeedImporter, err := 

	// Update the Mock Price Feed contract on C-Chain

	// Get a Warp signature of the block hash containing the AnswerUpdated event

	// Construct a Merkle proof of the AnswerUpdated event agains the block's receipts root

	// Import the AnswerUpdated event into the Price Feed Importer contract on Subnet A
}
