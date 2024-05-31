// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

/*
 * @notice Struct representing an EVM block header.
 * Note: The block header format may vary between chains with different VM implementations (i.e. subnet-evm vs coreth).
*/
struct EVMBlockHeader {
    bytes32 parentHash;
    bytes32 sha3Uncles;
    address miner;
    bytes32 stateRoot;
    bytes32 transactionsRoot;
    bytes32 receiptsRoot;
    bytes logsBloom;
    uint256 difficulty;
    uint256 number;
    uint256 gasLimit;
    uint256 gasUsed;
    uint256 timestamp;
    bytes extraData;
    bytes32 mixHash;
    uint256 nonce;
    uint256 baseFeePerGas;
    uint256 blockGasCost;
}

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
    bool status;
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
     * @notice Imports an event log from another blockchain.
     * 1. Imports a block hash from another blockchain via Warp.
     * 2. Verifies that the provided blockHeader matches the authenticated block hash.
     * 3. Gets the receipt at the given transaction index by verifying the merkle proof against 
     *    the block header's receipt root.
     * 4. Decodes and returns the log at the given log index from the receipt.
     */
    function importEvent(bytes calldata blockHeader, uint256 txIndex, bytes[] calldata receiptProof, uint256 logIndex)
        external;
}
