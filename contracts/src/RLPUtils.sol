// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {EVMLog, EVMReceipt} from "./IEventImporter.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RLPReader.sol";

/**
 * THIS IS AN EXAMPLE LIBRARY THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

library RLPUtils {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for RLPReader.Iterator;

    function decodeReceipt(RLPReader.RLPItem memory encodedReceipt) internal pure returns (EVMReceipt memory) {
        // If the encoded receipt is not a list, then the first byte is the transaction type,
        // followed by the RLP encoding of the receipt afterwards. If encoded receipt is already
        // a list itself, then the transaction type is 0 (legacy tx).
        uint8 txType;
        if (!encodedReceipt.isList()) {
            uint256 memptr = encodedReceipt.memPtr;
            // solhint-disable-next-line no-inline-assembly
            assembly {
                txType := byte(0, mload(memptr))
            }
            require(txType == 1 || txType == 2, "Invalid tx type for non-legacy tx");
            encodedReceipt = RLPReader.RLPItem({len: encodedReceipt.len - 1, memPtr: encodedReceipt.memPtr + 1});
        }

        // Four items in every receipt are:
        // 1. Post-state or status
        // 2. Cumulative gas used
        // 3. Bloom filter
        // 4. Logs
        RLPReader.RLPItem[] memory receiptItems = encodedReceipt.toList();
        require(receiptItems.length == 4, "Invalid number of RLP elements in receipt");
        EVMReceipt memory result;
        result.txType = txType;
        result.postStateOrStatus = receiptItems[0].toBytes();
        result.cululativeGasUsed = uint64(receiptItems[1].toUint());
        result.bloom = receiptItems[2].toBytes();
        RLPReader.RLPItem[] memory logs = receiptItems[3].toList();
        result.logs = new EVMLog[](logs.length);
        for (uint256 i = 0; i < logs.length; i++) {
            result.logs[i] = decodeLog(logs[i]);
        }

        return result;
    }

    function decodeLog(RLPReader.RLPItem memory encodedLog) internal pure returns (EVMLog memory) {
        RLPReader.RLPItem[] memory log = encodedLog.toList();
        // Three items in every receipt are:
        // 1. Address of the logger
        // 2. Log topics, of which there is always at least one (the event signature)
        // 3. Log data (arbitrary bytes)
        require(log.length == 3, "Invalid number of RLP elements in log");
        EVMLog memory evmLog;
        evmLog.loggerAddress = log[0].toAddress();
        RLPReader.RLPItem[] memory topics = log[1].toList();
        evmLog.topics = new bytes32[](topics.length);
        for (uint256 i = 0; i < topics.length; i++) {
            evmLog.topics[i] = bytes32(topics[i].toBytes());
        }
        evmLog.data = log[2].toBytes();
        return evmLog;
    }

    function decodeBlockNumberAndReceiptsRoot(bytes memory encodedBlockHeader)
        internal
        pure
        returns (uint256, bytes32)
    {
        // RLP decode the block header.
        RLPReader.RLPItem[] memory blockHeader = encodedBlockHeader.toRlpItem().toList();
        require(blockHeader.length >= 15, "Invalid number of RLP elements in block header");

        // Extract the block number and the receipts root from the RLP encoding.
        uint256 blockNumber = blockHeader[8].toUint();
        bytes32 receiptsRoot = bytes32(blockHeader[5].toBytes());

        return (blockNumber, receiptsRoot);
    }

    function encodeUint256(uint256 index) internal pure returns (bytes memory) {
        if (index == 0) {
            return hex"80";
        }
        // TODO: Support RLP encoding of larger tx indexes.
        require(index <= 127, "Invalid tx index");
        return abi.encodePacked(uint8(index));
    }
}
