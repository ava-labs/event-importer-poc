// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {EVMLog, EVMEventInfo, IEventImporter} from "./IEventImporter.sol";
import {WarpBlockHash, IWarpMessenger} from "@subnet-evm/contracts/interfaces/IWarpMessenger.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RLPReader.sol";
import {RLPUtils} from "./RLPUtils.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Abstract contract for importing events from another blockchain.
 * Uses the Warp precompile to authenticate the block hash of the block including the events to be imported.
 * Inheriting contracts must implement the _onEventImport function to handle event imports.
 */
abstract contract EventImporter is IEventImporter {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;

    struct TxLogIndex {
        uint256 txIndex;
        uint256 logIndex;
    }

    IWarpMessenger public warpMessenger;

    constructor() {
        warpMessenger = IWarpMessenger(0x0200000000000000000000000000000000000005);
    }

    /*
     * @notice Imports an event log from another blockchain.
     * 1. Imports a block hash from another blockchain via Warp.
     * 2. Verifies that the provided blockHeader matches the authenticated block hash.
     * 3. Gets the receipt at the given transaction index by verifying the merkle proof against the block header's receipt root.
     * 4. Decodes and returns the log at the given log index from the receipt.
     * Note: The first argument (sourceBlockchainID) is not used in this example because it is fetched from
     * WarpBlockHash provided by the Warp precompile.
     */
    function importEvent(
        bytes32,
        bytes calldata blockHeader,
        uint256 txIndex,
        bytes[] calldata receiptProof,
        uint256 logIndex
    ) external {
        // TODO it is more expensive to do that because of the memory allocation
        // but Solidity function inlining is not easy to predict.
        TxLogIndex[] memory txLogIndexes = new TxLogIndex[](1);
        txLogIndexes[0] = TxLogIndex({txIndex: txIndex, logIndex: logIndex});
        importEvents(bytes32(0), blockHeader, receiptProof, txLogIndexes);
    }

    // @notice import multiple events from the same receipt
    function importEvents(
        bytes32,
        bytes calldata blockHeader,
        bytes[] calldata receiptProof,
        TxLogIndex[] memory txLogIndexes
    ) public {
        // Get the verified block has via the Warp precompile.
        (WarpBlockHash memory warpBlockHash, bool valid) = warpMessenger.getVerifiedWarpBlockHash(0);
        require(valid, "Invalid WarpBlockHash");

        // Check that the blockHeader matches the authenticated block hash.
        require(keccak256(blockHeader) == warpBlockHash.blockHash, "Invalid blockHeader");

        // RLP decode the required values from the block header.
        (uint256 blockNumber, bytes32 receiptsRoot) = RLPUtils.decodeBlockNumberAndReceiptsRoot(blockHeader);

        // Construct the key of the trie receipt proof.
        bytes[] memory receiptKeys = new bytes[](txLogIndexes.length);
        for (uint256 i; i < txLogIndexes.length; i++) {
            receiptKeys[i] = RLPUtils.encodeUint256(txLogIndexes[i].txIndex);
        }

        // Verify the trie proof against the receipts root.
        StorageValue[] memory results = MerklePatricia.VerifyEthereumProof(receiptsRoot, receiptProof, receiptKeys);
        require(results.length == receiptKeys.length, "Invalid number of results in receipt proof");
        for (uint256 i; i < results.length; i++) {
            require(results[i].value.length > 0, "Invalid receipt proof");

            TxLogIndex memory txLogIndex = txLogIndexes[i];

            EVMLog memory log = RLPUtils.decodeLogFast(results[i].value.toRlpItem(), txLogIndex.logIndex);

            _onEventImport(
                EVMEventInfo({
                    blockchainID: warpBlockHash.sourceChainID,
                    blockNumber: blockNumber,
                    txIndex: txLogIndex.txIndex,
                    logIndex: txLogIndex.logIndex,
                    log: log
                })
            );

            emit EventImported(
                warpBlockHash.sourceChainID,
                warpBlockHash.blockHash,
                log.loggerAddress,
                txLogIndex.txIndex,
                txLogIndex.logIndex
            );
        }
    }

    function _onEventImport(EVMEventInfo memory eventInfo) internal virtual;
}
