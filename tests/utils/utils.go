package utils

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

func GetSignedBlockHashMessage(ctx context.Context, subnet interfaces.SubnetTestInfo, importingSubnetID ids.ID, blockHash common.Hash) []byte {
	warpClient, err := warpBackend.NewClient(subnet.NodeURIs[0], subnet.BlockchainID.String())
	Expect(err).Should(BeNil())

	signedWarpMessageBytes, err := warpClient.GetBlockAggregateSignature(
		ctx,
		ids.ID(blockHash),
		warp.WarpDefaultQuorumNumerator,
		importingSubnetID.String(),
	)
	Expect(err).Should(BeNil())
	log.Info("Got signed block hash message", "blockHash", blockHash.String(), "signedWarpMessageBytes", hex.EncodeToString(signedWarpMessageBytes))

	return signedWarpMessageBytes
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
