// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

struct EVMBlockHeader {
    uint256 blockNumber;
    bytes32 receiptsRoot;
}

struct EVMLog {
    address loggerAddress;
    bytes32[] topics;
    bytes data;
}

struct EVMReceipt {
    bool status;
    uint64 cululativeGasUsed;
    bytes bloom;
    EVMLog[] logs;
}

struct EVMEventInfo {
    bytes32 blockchainID;
    uint256 blockNumber;
    bytes32 txHash;
    uint256 logIndex;
    EVMLog log;
}

interface IEventImporter {
    /*
     * @notice 
     * 1. Imports a block hash from another blockchain via Warp.
     * 2. Verifies that the provided blockHeader matches the authenticated block hash.
     * 3. Gets the receipt at the given transaction index by verifying the merkle proof against the block header's receipt root.
     * 4. Decodes and returns the log at the given log index from the receipt.
     */
    function importEvent(bytes calldata blockHeader, uint256 txIndex, bytes calldata receiptProof, uint256 logIndex)
        external
        returns (EVMEventInfo memory);
}
