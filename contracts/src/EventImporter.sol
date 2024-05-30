// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {EVMEventInfo, EVMLog, EVMReceipt, EVMBlockHeader, IEventImporter} from "./IEventImporter.sol";
import {WarpBlockHash, IWarpMessenger} from "@subnet-evm/contracts/interfaces/IWarpMessenger.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPUtils} from "./RLPUtils.sol";

/**
 * @notice Abstract contract for importing events from another blockchain.
 * Uses the Warp precompile to authenticate the block hash of the block including the events to be imported.
 * Inheriting contracts must implement the _onEventImport function to handle event imports.
 */
abstract contract EventImporter is IEventImporter {
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
     */
    function importEvent(bytes calldata blockHeader, uint256 txIndex, bytes[] calldata receiptProof, uint256 logIndex)
        external
    {
        // Get the verified block has via the Warp precompile.
        (WarpBlockHash memory warpBlockHash, bool valid) = warpMessenger.getVerifiedWarpBlockHash(0);
        require(valid, "Invalid WarpBlockHash");

        // Check that the blockHeader matches the authenticated block hash.
        require(keccak256(blockHeader) == warpBlockHash.blockHash, "Invalid blockHeader");

        // RLP decode the required values from the block header.
        (uint256 blockNumber, bytes32 receiptsRoot) = RLPUtils.decodeBlockNumberAndReceiptsRoot(blockHeader);

        // Construct the key of the trie receipt proof.
        bytes[] memory receiptKeys = new bytes[](1);
        receiptKeys[0] = RLPUtils.encodeUint256(txIndex);

        // Verify the trie proof against the receipts root.
        StorageValue[] memory results = MerklePatricia.VerifyEthereumProof(receiptsRoot, receiptProof, receiptKeys);
        require(results.length == 1, "Invalid number of results in receipt proof");
        require(results[0].value.length > 0, "Invalid receipt proof");

        EVMReceipt memory receipt = RLPUtils.decodeReceipt(results[0].value);
        require(logIndex < receipt.logs.length, "Invalid log index");

        _onEventImport(
            EVMEventInfo({
                blockchainID: warpBlockHash.sourceChainID,
                blockNumber: blockNumber,
                txIndex: txIndex,
                logIndex: logIndex,
                log: receipt.logs[logIndex]
            })
        );
    }

    function _onEventImport(EVMEventInfo memory eventInfo) internal virtual;
}
