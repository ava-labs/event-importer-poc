// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * THIS IS AN EXAMPLE INTERFACE THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/*
 * @notice Struct representing an EVM event log.
 */
struct EVMLog {
    address loggerAddress;
    bytes32[] topics;
    bytes data;
}

/*
 * @notice Struct representing an EVM transaction receipt.
 */
struct EVMReceipt {
    uint8 txType;
    bytes postStateOrStatus;
    uint64 cululativeGasUsed;
    bytes bloom;
    EVMLog[] logs;
}

/*
 * @notice Struct representing an EVM event log and its associated metadata.
 */
struct EVMEventInfo {
    bytes32 blockchainID;
    uint256 blockNumber;
    uint256 txIndex;
    uint256 logIndex;
    EVMLog log;
}

interface IEventImporter {
    /*
     * @notice Event emitted when an event is imported from another blockchain.
     */
    event EventImported(
        bytes32 indexed sourceBlockchainID,
        bytes32 indexed sourceBlockHash,
        address indexed loggerAddress,
        uint256 txIndex,
        uint256 logIndex
    );

    /*
     * @notice Imports an event log from another blockchain.
     * 1. Imports a block hash from another blockchain via Warp.
     * 2. Verifies that the provided blockHeader matches the authenticated block hash.
     * 3. Gets the receipt at the given transaction index by verifying the merkle proof against 
     * the block header's receipt root.
     * 4. Decodes and returns the log at the given log index from the receipt.
     */
    function importEvent(
        bytes32 sourceBlockchainID,
        bytes calldata blockHeader,
        uint256 txIndex,
        bytes[] calldata receiptProof,
        uint256 logIndex
    ) external;
}
