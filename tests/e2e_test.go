// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package local

import (
	"os"
	"testing"

	"github.com/ava-labs/teleporter/tests/local"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

const (
	warpGenesisFile = "./tests/utils/warp-genesis.json"

	erc20TokenHubLabel    = "ERC20TokenHub"
	erc20TokenSpokeLabel  = "ERC20TokenSpoke"
	nativeTokenHubLabel   = "NativeTokenHub"
	nativeTokenSpokeLabel = "NativeTokenSpoke"
	multiHopLabel         = "MultiHop"
	sendAndCallLabel      = "SendAndCall"
	registrationLabel     = "Registration"
)

var LocalNetworkInstance *local.LocalNetwork

func TestE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}
	format.MaxLength = 10000

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	// Create the local network instance
	LocalNetworkInstance = local.NewLocalNetwork(warpGenesisFile)

})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
})

var _ = ginkgo.Describe("[Teleporter Token Bridge integration tests]", func() {
	ginkgo.It("Bridge an ERC20 token between two Subnets",
		ginkgo.Label(erc20TokenHubLabel, erc20TokenSpokeLabel),
		func() {
			return
		})

})
