package flows

import (
	"context"

	"github.com/ava-labs/receipt-proof-poc/abi-bindings/go/mocks/mockpricefeedaggregator"
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
	mockPriceFeedAggregator, err := mockpricefeedaggregator.NewMockPriceFeedAggregator(fundedAddress, network.GetCChainClient())

	// Deploy Price Feed Importer contract on Subnet A

	// Update the Mock Price Feed contract on C-Chain

	// Get a Warp signature of the block hash containing the AnswerUpdated event

	// Construct a Merkle proof of the AnswerUpdated event agains the block's receipts root

	// Import the AnswerUpdated event into the Price Feed Importer contract on Subnet A
}
